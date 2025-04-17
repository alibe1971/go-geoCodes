package geoCodes

import(
    "reflect"
    "encoding/json"
    "crypto/rand"
    "time"
    "github.com/google/uuid"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "log"
    "os"
    "runtime"
    "bytes"
    "strings"
    "strconv"
    "unicode"
    Structs "github.com/alibe1971/go-geoCodes/geoCodes/Structs"
)

func strToLower(str string) string {
    return strings.ToLower(str)
}

func strToUpper(str string) string {
    return strings.ToUpper(str)
}


func ucfirst(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}


func lcfirst(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

func toString(value interface{}) string {
    switch v := value.(type) {
        case string:
            return v
        case int:
            return strconv.Itoa(v)
        case float64:
            return strconv.FormatFloat(v, 'f', -1, 64)
        default:
            logPanicWithStackTrace(fmt.Sprintf("Unsupported type: %T", v))
            return ""
    }
}

func setElementsToStrings(values []interface{}) []string {
    var stringValues []string
    for _, value := range values {
        stringValues = append(stringValues, toString(value))
    }
    return stringValues
}

func setElementsToIntegers(values []interface{}) []int {
    var intValues []int
    for _, value := range values {
        switch v := value.(type) {
            case int:
                intValues = append(intValues, v)
            case string:
                if intVal, err := strconv.Atoi(v); err == nil {
                    intValues = append(intValues, intVal)
                } else {
                    logPanicWithStackTrace(fmt.Sprintf("Unable to convert `%s` to int", v))
                }
            default:
                logPanicWithStackTrace(fmt.Sprintf("Unsupported type for integer conversion: %T", v))
        }
    }
    return intValues
}

func loadData(key string, jsonData string, v interface{}) {
    err := json.Unmarshal([]byte(jsonData), v)
    if err != nil {
        fmt.Printf("Failed to unmarshal JSON for %s: %v\n", key, err)
        return
    }
    dataMap[key] = v
}


func getData(key string) interface{} {
    if data, ok := dataMap[key]; ok {
            return data
        }
    fmt.Printf("Data with key %s not found\n", key)
    return nil
}

func isInSlice(slice []string, value string) bool {
    for _, item := range slice {
        if item == value {
            return true
        }
    }
    return false
}


func isInMap(m map[string]interface{}, value string) bool {
    _, exists := m[value]
    return exists
}

func getStructPropertiesNamesSingleLevel(s interface{}) []string {
    var fieldNames []string
    t := reflect.TypeOf(s)
    if t.Kind() != reflect.Struct {
        return nil
    }
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        fieldNames = append(fieldNames, field.Name)
    }
    return fieldNames
}

func getStructPropertiesNamesMultilevel(s interface{}, prefix string) []string {
	var fieldNames []string
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Struct {
		return nil
	}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldName := field.Name
		fullName := fieldName
		if prefix != "" {
			fullName = prefix + "." + fieldName
		}
		fieldNames = append(fieldNames, fullName)

		fieldType := field.Type
		if fieldType.Kind() == reflect.Struct {
			nestedStruct := reflect.New(fieldType).Elem().Interface()
			fieldNames = append(fieldNames, getStructPropertiesNamesMultilevel(nestedStruct, fullName)...)
		}

		if fieldType.Kind() == reflect.Ptr && fieldType.Elem().Kind() == reflect.Struct {
			nestedStruct := reflect.New(fieldType.Elem()).Elem().Interface()
			fieldNames = append(fieldNames, getStructPropertiesNamesMultilevel(nestedStruct, fullName)...)
		}
	}
	return fieldNames
}




func generateUniqueString() Structs.GeoCodeReference {
    uuidStr := uuid.New().String()
    timestamp := time.Now().UnixNano()
    randomBytes := make([]byte, 8)
    _, err := rand.Read(randomBytes)
    if err != nil {
        panic(err)
    }
    data := fmt.Sprintf("%s-%d-%x", uuidStr, timestamp, randomBytes)
    hash := sha256.Sum256([]byte(data))
    reference := Structs.GeoCodeReference(hex.EncodeToString(hash[:]))
    return reference
}


func logFatalWithTrace(message string) {
    _, file, line, ok := runtime.Caller(3)
    if ok {
        log.SetOutput(os.Stderr)
        log.Fatalf("Critical: %s\nFile: %s\nLine: %d\n", message, file, line)
    } else {
        log.SetOutput(os.Stderr)
        log.Fatalln("Critical: %s", message)
    }
}

func logPanicWithStackTrace(message string) {
    const depth = 15
    buf := make([]byte, 1024)
    n := runtime.Stack(buf, false)
    stackTrace := string(buf[:n])
    reversedStackTrace := reverseStackTrace(stackTrace)
    log.SetOutput(os.Stderr)
    panic(fmt.Sprintf("Critical Error: %s\nStack Trace:\n%s", message, reversedStackTrace))
}


func reverseStackTrace(trace string) string {
    var reversed bytes.Buffer
    lines := strings.Split(trace, "\n")
    for i := len(lines) - 1; i >= 0; i-- {
        reversed.WriteString(lines[i] + "\n")
    }
    return reversed.String()
}

/********/
func flattenMap(prefix string, m map[string]interface{}, sep string, out map[string]interface{}) {
    for k, v := range m {
        key := k
        if prefix != "" {
            key = prefix + sep + key
        }
        switch vv := v.(type) {
        case map[string]interface{}:
            flattenMap(key, vv, sep, out)
        case []interface{}:
            flattenSlice(key, vv, sep, out)
        default:
            out[key] = vv
        }
    }
}

func flattenSlice(prefix string, s []interface{}, sep string, out map[string]interface{}) {
    for i, v := range s {
        key := toString(i) //fmt.Sprintf("%d", i)
        if prefix != "" {
            key = prefix + sep + key
        }
        switch vv := v.(type) {
        case map[string]interface{}:
            flattenMap(key, vv, sep, out)
        case []interface{}:
            flattenSlice(key, vv, sep, out)
        default:
            out[key] = vv
        }
    }
}

/********/

func LowerCamelCaseKeys(data interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		newMap := make(map[string]interface{})
		for key, val := range v {
			newKey := lcfirst(key)
			newMap[newKey] = LowerCamelCaseKeys(val)
		}
		return newMap
	case []interface{}:
		for i, item := range v {
			v[i] = LowerCamelCaseKeys(item)
		}
		return v
	default:
		return data
	}
}


func structToMap(v reflect.Value) map[string]interface{} {
	result := make(map[string]interface{})
	vType := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := vType.Field(i).Name
		result[fieldName] = field.Interface()
	}
	return result
}

// Tipo per indicare che il campo non è presente
type notFoundInData struct{}

var notFound = notFoundInData{}

func mapListToXML(
    data interface{},
    containerName string, // Il nome del container (es: "currencies")
    dataType string,      // Una stringa che descrive il tipo, ad es. "[]map[string]interface {}"
    xmlMapping map[string]Structs.XmlFieldMapping,
) (string, error) {

    if data == nil {
        return "", nil
    }

    var listTag string
    var sb strings.Builder
    var constructor map[string]Structs.XmlFieldMapping

    for _, mapping := range xmlMapping {
        listTag = mapping.TagName
        constructor = mapping.Children
        break
    }

    sb.WriteString("<" + containerName + ">\n")
    switch dataType {
        case "map[string]map[string]interface {}":
            dataMap, ok := data.(map[string]interface {})
            if !ok {
                return "", fmt.Errorf("Errore: atteso []map[string]interface{} ma ottenuto %T", data)
            }
            for index, singleItem := range dataMap {
                xmlString, err := mapToXML(singleItem, listTag, constructor, index, 1)
                if err != nil {
                    return "", err
                }
                sb.WriteString(xmlString)
            }
        case "[]map[string]interface {}":
            dataSlice, ok := data.([]interface{})
            if !ok {
                return "", fmt.Errorf("Errore: atteso []map[string]interface{} ma ottenuto %T", data)
            }
            for _, singleItem := range dataSlice {
                xmlString, err := mapToXML(singleItem, listTag, constructor, "", 1)
                if err != nil {
                    return "", err
                }
                sb.WriteString(xmlString)
            }
        default:
            return "", fmt.Errorf("Tipo di dato non gestito: %s", dataType)
    }

    sb.WriteString("</" + containerName + ">\n")
    return sb.String(), nil
}



// mapToXML è la funzione "principale" che genera l'XML per un singolo oggetto.
func mapToXML(
    data interface{},
    itemTag string,
    xmlMapping map[string]Structs.XmlFieldMapping,
    rootAttribute string,
    indentLevel int,
) (string, error) {
    if data == nil {
        return "", nil
    }

    var sb strings.Builder

    // Apertura tag radice
    sb.WriteString(indentString(indentLevel))
    if rootAttribute == "" {
        sb.WriteString("<" + itemTag + ">\n")
    } else {
        sb.WriteString("<" + itemTag + " index=\"" + rootAttribute +"\">\n")
    }

    // Iterazione sulle chiavi del mapping
    for fieldName, fieldMap := range xmlMapping {
        val := extractFieldValue(data, fieldName)
        if val == notFound {
            continue
        }
        fieldXML, err := buildXMLForField(val, fieldMap, indentLevel+1)
        if err != nil {
            return "", err
        }
        sb.WriteString(fieldXML)
    }

    // Chiusura tag radice
    sb.WriteString(indentString(indentLevel))
    sb.WriteString("</" + itemTag + ">\n")

    return sb.String(), nil
}

func buildXMLForField(
    val interface{},
    mapping Structs.XmlFieldMapping,
    indentLevel int,
) (string, error) {

    var sb strings.Builder
    tagName := mapping.TagName
    if tagName == "" {
        tagName = mapping.Field
    }

    switch actual := val.(type) {

    // Caso 1: valore nil ⇒ tag vuoto
    case nil:
        sb.WriteString(indentString(indentLevel))
        sb.WriteString("<" + tagName + "></" + tagName + ">\n")
        return sb.String(), nil

    // ---- A) float64 con gestione AsInt
    case float64:
        // Se AsInt = true, lo stampiamo come intero
        var strVal string
        if mapping.AsInt {
            // Converte forzatamente in int64
            intVal := int64(actual)
            strVal = fmt.Sprintf("%d", intVal)
        } else {
            // Altrimenti, stampa in modo "normale"
            // Se preferisci niente notazione scientifica, puoi usare:
            // strVal = strconv.FormatFloat(actual, 'f', -1, 64)
            // Oppure semplice "%v".
            strVal = fmt.Sprintf("%v", actual)
        }

        sb.WriteString(indentString(indentLevel))
        sb.WriteString("<" + tagName + ">")
        if mapping.CDATA && strVal != "" {
            sb.WriteString("<![CDATA[")
            sb.WriteString(strVal)
            sb.WriteString("]]>")
        } else {
            sb.WriteString(strVal)
        }
        sb.WriteString("</" + tagName + ">\n")
        return sb.String(), nil

    // ---- B) Altri tipi "scalari"
    case string, int, int64, bool:
        strVal := fmt.Sprintf("%v", actual)
        sb.WriteString(indentString(indentLevel))
        sb.WriteString("<" + tagName + ">")
        if mapping.CDATA && strVal != "" {
            sb.WriteString("<![CDATA[")
            sb.WriteString(strVal)
            sb.WriteString("]]>")
        } else {
            sb.WriteString(strVal)
        }
        sb.WriteString("</" + tagName + ">\n")
        return sb.String(), nil

    // ---- C) slice/array generico
    case []interface{}:
        sb.WriteString(indentString(indentLevel))
        sb.WriteString("<" + tagName + ">\n")

        for _, elem := range actual {
            // Se esiste un mapping.Children, occorre decidere:
            // 1) Se l’elem è un oggetto (map), allora ha senso fare la mini mappa
            // 2) Se è uno scalare (string, int, ecc.), passiamo il valore diretto

            if len(mapping.Children) == 1 {
                // Di solito abbiamo 1 child field "tz" o simile
                // Recuperiamo l’unico childMap
                var theChildName string
                var theChild Structs.XmlFieldMapping
                for cName, cMap := range mapping.Children {
                    theChildName = cName
                    theChild = cMap
                    break
                }

                // Controlliamo se elem è un map[string]interface{}
                // (in tal caso costruiamo fValMap), altrimenti passiamo l’elem così com’è
                switch elemTyped := elem.(type) {
                case map[string]interface{}:
                    // oggetto => usiamo la logica classica
                    fValMap := map[string]interface{}{theChildName: elemTyped}
                    subXML, err := buildXMLForField(fValMap, theChild, indentLevel+1)
                    if err != nil { return "", err }
                    sb.WriteString(subXML)

                default:
                    // valore scalare => passiamo direttamente
                    subXML, err := buildXMLForField(elemTyped, theChild, indentLevel+1)
                    if err != nil { return "", err }
                    sb.WriteString(subXML)
                }

            } else if len(mapping.Children) > 1 {
                // Caso più complesso: se ci sono piú children,
                // magari gli elem dovrebbero essere mappe con piú campi
                // ... la logica esistente ...
                for childField, childMap := range mapping.Children {
                    fValMap := map[string]interface{}{childField: elem}
                    subXML, err := buildXMLForField(fValMap, childMap, indentLevel+1)
                    if err != nil { return "", err }
                    sb.WriteString(subXML)
                }

            } else {
                // Se non ci sono children, stampiamo <tagName>valore</tagName> come fallback
                sb.WriteString(indentString(indentLevel+1))
                sb.WriteString("<" + tagName + ">")
                sb.WriteString(fmt.Sprintf("%v", elem))
                sb.WriteString("</" + tagName + ">\n")
            }
        }

        sb.WriteString(indentString(indentLevel))
        sb.WriteString("</" + tagName + ">\n")
        return sb.String(), nil


    // ---- D) map[string]interface{}
    case map[string]interface{}:
        // --- PROVA A CONVERTIRE A map[string]string? ---
        // (1) se la mappa ha TUTTI valori di tipo string,
        // (2) e se ci aspettiamo di dover usare Attributi (o c'è un solo child che ha AttrName, ecc.),
        // allora deviamo verso la logica "map[string]string".

        if allStrings(actual) && shouldConvertToMapStringString(mapping) {
            // costruiamo una map[string]string
            conv := make(map[string]string)
            for k, rawVal := range actual {
                conv[k] = rawVal.(string) // safe: allStrings() ha già verificato
            }
            // e richiamiamo buildXMLForField sul “conv”
            return buildXMLForField(conv, mapping, indentLevel)
        }

        // --- Se non rientra nei criteri di "mappa di sole stringhe", prosegui con D) standard ---
        sb.WriteString(indentString(indentLevel))
        sb.WriteString("<" + tagName + ">\n")
        for childName, childMap := range mapping.Children {
            childVal := extractFieldValue(actual, childName)
            if childVal == notFound {
                continue
            }
            subXML, err := buildXMLForField(childVal, childMap, indentLevel+1)
            if err != nil {
                return "", err
            }
            sb.WriteString(subXML)
        }
        sb.WriteString(indentString(indentLevel))
        sb.WriteString("</" + tagName + ">\n")
        return sb.String(), nil


    // ---- E) map[string]string (per officialName, mottos, ecc.)
    case map[string]string:
        sb.WriteString(indentString(indentLevel))
        sb.WriteString("<" + tagName + ">\n")
        if len(mapping.Children) == 1 {
            var soleChild Structs.XmlFieldMapping
            for _, c := range mapping.Children {
                soleChild = c
                break
            }
            childTag := soleChild.TagName
            if childTag == "" {
                childTag = soleChild.Field
            }
            attr := soleChild.AttrName

            for k, v := range actual {
                sb.WriteString(indentString(indentLevel+1))
                sb.WriteString("<" + childTag)
                if attr != "" {
                    sb.WriteString(" " + attr + "=\"" + k + "\"")
                }
                sb.WriteString(">")
                if soleChild.CDATA && v != "" {
                    sb.WriteString("<![CDATA[")
                    sb.WriteString(v)
                    sb.WriteString("]]>")
                } else {
                    sb.WriteString(v)
                }
                sb.WriteString("</" + childTag + ">\n")
            }
        } else {
            // fallback generico
            for k, v := range actual {
                sb.WriteString(indentString(indentLevel+1))
                sb.WriteString("<child key=\"" + k + "\">")
                sb.WriteString(v)
                sb.WriteString("</child>\n")
            }
        }
        sb.WriteString(indentString(indentLevel))
        sb.WriteString("</" + tagName + ">\n")
        return sb.String(), nil

    // ---- F) slice di string
    case []string:
        sb.WriteString(indentString(indentLevel))
        sb.WriteString("<" + tagName + ">\n")
        for _, s := range actual {
            sb.WriteString(indentString(indentLevel+1))
            sb.WriteString("<item>")
            sb.WriteString(s)
            sb.WriteString("</item>\n")
        }
        sb.WriteString(indentString(indentLevel))
        sb.WriteString("</" + tagName + ">\n")
        return sb.String(), nil

    default:
        // Tipo sconosciuto
        sb.WriteString(indentString(indentLevel))
        sb.WriteString("<" + tagName + ">UNKNOWN_TYPE</" + tagName + ">\n")
        return sb.String(), nil
    }
}


// Funzione di supporto: recupera data[fieldName], o notFound se assente
func extractFieldValue(data interface{}, fieldName string) interface{} {
    if m, ok := data.(map[string]interface{}); ok {
        val, exists := m[fieldName]
        if !exists {
            return notFound
        }
        return val
    }
    return notFound
}

// Semplice indentazione
func indentString(level int) string {
    return strings.Repeat("  ", level)
}


// Controlla se tutti i valori di m sono string
func allStrings(m map[string]interface{}) bool {
    for _, val := range m {
        if _, ok := val.(string); !ok {
            return false
        }
    }
    return true
}

// Decide se dobbiamo/devono convertire a map[string]string
// (per esempio: se mapping.AsAttributes == true, OPPURE
//  se c'è un solo child e quell'unico child ha un AttrName, ecc.)
func shouldConvertToMapStringString(mapping Structs.XmlFieldMapping) bool {
    if mapping.AsAttributes {
        return true
    }
    // Oppure: se c'è un solo figlio e quell’unico figlio ha un .AttrName
    if len(mapping.Children) == 1 {
        for _, c := range mapping.Children {
            if c.AttrName != "" {
                return true
            }
        }
    }
    return false
}
