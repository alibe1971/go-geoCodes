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
    "regexp"
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

	case nil:
		sb.WriteString(indentString(indentLevel))
		sb.WriteString("<" + tagName + "></" + tagName + ">\n")
		return sb.String(), nil

	case float64:
		var strVal string
		if mapping.AsInt {
			strVal = fmt.Sprintf("%d", int64(actual))
		} else {
			strVal = fmt.Sprintf("%v", actual)
		}
		sb.WriteString(indentString(indentLevel))
		sb.WriteString("<" + tagName + ">")
		if mapping.CDATA && strVal != "" {
			sb.WriteString("<![CDATA[" + strVal + "]]>")
		} else {
			sb.WriteString(strVal)
		}
		sb.WriteString("</" + tagName + ">\n")
		return sb.String(), nil

	case string, int, int64, bool:
		strVal := fmt.Sprintf("%v", actual)
		sb.WriteString(indentString(indentLevel))
		sb.WriteString("<" + tagName + ">")
		if mapping.CDATA && strVal != "" {
			sb.WriteString("<![CDATA[" + strVal + "]]>")
		} else {
			sb.WriteString(strVal)
		}
		sb.WriteString("</" + tagName + ">\n")
		return sb.String(), nil

	case []interface{}:
		sb.WriteString(indentString(indentLevel))
		sb.WriteString("<" + tagName + ">\n")

		for _, elem := range actual {
			if len(mapping.Children) == 1 {
				var theChild Structs.XmlFieldMapping
				for _, cMap := range mapping.Children {
					theChild = cMap
					break
				}
				subXML, err := buildXMLForField(elem, theChild, indentLevel+1)
				if err != nil {
					return "", err
				}
				sb.WriteString(subXML)
			} else {
				sb.WriteString(indentString(indentLevel+1))
				sb.WriteString("<item>" + fmt.Sprintf("%v", elem) + "</item>\n")
			}
		}

		sb.WriteString(indentString(indentLevel))
		sb.WriteString("</" + tagName + ">\n")
		return sb.String(), nil

	case map[string]interface{}:
		if allStrings(actual) && shouldConvertToMapStringString(mapping) {
			conv := make(map[string]string)
			for k, rawVal := range actual {
				conv[k] = rawVal.(string)
			}
			return buildXMLForField(conv, mapping, indentLevel)
		}
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
					sb.WriteString("<![CDATA[" + v + "]]>")
				} else {
					sb.WriteString(v)
				}
				sb.WriteString("</" + childTag + ">\n")
			}
		}
		sb.WriteString(indentString(indentLevel))
		sb.WriteString("</" + tagName + ">\n")
		return sb.String(), nil

	case []string:
		sb.WriteString(indentString(indentLevel))
		sb.WriteString("<" + tagName + ">\n")
		for _, s := range actual {
			sb.WriteString(indentString(indentLevel+1))
			sb.WriteString("<item>" + s + "</item>\n")
		}
		sb.WriteString(indentString(indentLevel))
		sb.WriteString("</" + tagName + ">\n")
		return sb.String(), nil

	default:
		rv := reflect.ValueOf(val)
		if rv.Kind() == reflect.Slice {
			sb.WriteString(indentString(indentLevel))
			sb.WriteString("<" + tagName + ">\n")

			for i := 0; i < rv.Len(); i++ {
				elem := rv.Index(i).Interface()
				if mapping.Children != nil {
					for _, childMap := range mapping.Children {
						subXML, err := buildXMLForField(elem, childMap, indentLevel+1)
						if err != nil {
							return "", err
						}
						sb.WriteString(subXML)
					}
				} else {
					sb.WriteString(indentString(indentLevel+1))
					sb.WriteString("<item>" + fmt.Sprintf("%v", elem) + "</item>\n")
				}
			}

			sb.WriteString(indentString(indentLevel))
			sb.WriteString("</" + tagName + ">\n")
			return sb.String(), nil
		}
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


/******/

// flattenReflect è la funzione ricorsiva:
// – “apre” pointer e interface
// – per Struct va sui campi esportati
// – per Map/Slice/Array itera
// – altrimenti scrive il valore terminale
func flattenReflect(prefix string, v reflect.Value, sep string, out map[string]interface{}) {
    // srotola Ptr e Interface
    for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
        if v.IsNil() {
            // nil pointer/interface → nil terminale
            out[prefix] = nil
            return
        }
        v = v.Elem()
    }

    switch v.Kind() {
    case reflect.Struct:
        t := v.Type()
        for i := 0; i < v.NumField(); i++ {
            field := t.Field(i)
            if !field.IsExported() {
                continue
            }
            fv := v.Field(i)
            key := field.Name
            if prefix != "" {
                key = prefix + sep + key
            }
            flattenReflect(key, fv, sep, out)
        }

    case reflect.Map:
        // solo chiavi stringhe; altrimenti le ignoro
        if v.Type().Key().Kind() != reflect.String {
            return
        }
        for _, k := range v.MapKeys() {
            strKey := k.String()
            val := v.MapIndex(k)
            key := strKey
            if prefix != "" {
                key = prefix + sep + strKey
            }
            flattenReflect(key, val, sep, out)
        }

    case reflect.Slice, reflect.Array:
        for i := 0; i < v.Len(); i++ {
            elem := v.Index(i)
            idx := strconv.Itoa(i)
            key := idx
            if prefix != "" {
                key = prefix + sep + idx
            }
            flattenReflect(key, elem, sep, out)
        }

    default:
        // tipi base: mantengono perfettamente int64, float64, string, bool, ecc.
        out[prefix] = v.Interface()
    }
}


/******/

// normalizeForPick trasforma tutte le possibili varianti di Data
// in map[string]interface{} o []interface{} pronte per l’indexing.
func normalizeDataForPick(data interface{}) interface{} {
    switch t := data.(type) {
    case map[string]interface{}:
        return t
    case map[string]map[string]interface{}:
        m := make(map[string]interface{}, len(t))
        for k, v := range t {
            m[k] = v
        }
        return m
    case []map[string]interface{}:
        s := make([]interface{}, len(t))
        for i, v := range t {
            s[i] = v
        }
        return s
    default:
        // già []interface{}, primitive, o strutture personalizzate
        return data
    }
}

// getPathValue scende ricorsivamente lungo la slice di chiavi/index
// e gestisce map[string]interface{} e []interface{}.
func getPathValue(current interface{}, parts []string) (interface{}, bool) {
    if len(parts) == 0 {
        return current, true
    }
    head, tail := parts[0], parts[1:]
    switch node := current.(type) {
    case map[string]interface{}:
        v, ok := node[head]
        if !ok {
            return nil, false
        }
        return getPathValue(v, tail)

    case []interface{}:
        idx, err := strconv.Atoi(head)
        if err != nil || idx < 0 || idx >= len(node) {
            return nil, false
        }
        return getPathValue(node[idx], tail)

    default:
        // se non è un container, solo tail vuota ci permette di tornare il valore
        if len(tail) == 0 {
            return current, true
        }
        return nil, false
    }
}


func fixEmojiField4Yaml(yamlData []byte) ([]byte, error) {
	re := regexp.MustCompile(`(?m)^(\s*emoji:\s+)"((?:\\U[0-9A-Fa-f]{8})+)"`)

	return re.ReplaceAllFunc(yamlData, func(line []byte) []byte {
		matches := re.FindSubmatch(line)
		if len(matches) != 3 {
			return line
		}

		escaped := string(matches[2]) // es: \U0001F1E6\U0001F1E9

		reSeq := regexp.MustCompile(`\\U([0-9A-Fa-f]{8})`)
		decoded := reSeq.ReplaceAllStringFunc(escaped, func(seq string) string {
			hex := seq[2:] // salta "\U"
			codepoint, err := strconv.ParseInt(hex, 16, 32)
			if err != nil {
				return seq
			}
			return string(rune(codepoint))
		})

		// ricostruisce la riga: stessa indentazione + emoji decodificata
		return []byte(string(matches[1]) + `"` + decoded + `"`)
	}), nil
}
