package geoCodes


import (
    "encoding/json"
    "encoding/xml"
    "bytes"
    "gopkg.in/yaml.v2"
    "strings"
    "reflect"
    Structs "github.com/alibe1971/go-geoCodes/geoCodes/Structs"
    "golang.org/x/text/collate"
    "golang.org/x/text/language"
    "fmt"
    "errors"
    "sort"
    "io/ioutil"
    "path/filepath"
)

func filterFields(src interface{}, fieldsToKeep []string) map[string]interface{} {
	srcValue := reflect.ValueOf(src)
	filtered := make(map[string]interface{})
	for _, fieldPath := range fieldsToKeep {
		fieldParts := strings.Split(fieldPath, ".")
		currentValue := srcValue
		var currentField reflect.StructField
		var found bool
		for _, part := range fieldParts {
			if currentValue.Kind() == reflect.Struct {
				currentField, found = currentValue.Type().FieldByName(part)
				if found {
					currentValue = currentValue.FieldByName(part)
				} else {
					break
				}
			}
		}
		if found {
			if len(fieldParts) == 1 {
				filtered[currentField.Name] = currentValue.Interface()
			} else {
				addNestedField(filtered, fieldParts, currentValue.Interface())
			}
		}
	}
	return filtered
}


func addNestedField(m map[string]interface{}, fieldParts []string, value interface{}) {
	if len(fieldParts) == 1 {
		m[fieldParts[0]] = value
		return
	}
	if _, ok := m[fieldParts[0]]; !ok {
		m[fieldParts[0]] = make(map[string]interface{})
	}
	if nestedMap, ok := m[fieldParts[0]].(map[string]interface{}); ok {
		addNestedField(nestedMap, fieldParts[1:], value)
	}
}

func getXsd(name string) ([]byte, error) {
    moduleDir, err := filepath.Abs(filepath.Dir("."))
    if err != nil {
        return nil, err
    }
    schemaPath := filepath.Join(moduleDir, "..", "Xsd", name+".xsd")
    xsd, err := ioutil.ReadFile(schemaPath)
    if err != nil {
        return nil, errors.New("invalid XSD")
    }
    return xsd, nil
}


func getDataAsFlattenMap(data interface{}, sep string) (map[string]interface{}, error) {
    if data == nil {
        return nil, nil
    }
    flat := make(map[string]interface{})
    v := reflect.ValueOf(data)
    flattenReflect("", v, sep, flat)
    return flat, nil
}


func getDataOnString(reference Structs.GeoCodeReference, data interface{}, method string) (string, error) {

    var toStringData []byte
    var err error
    var rootTag string = geocodesMap[reference].SetType
    var itemTag string = Structs.SingleItemName[rootTag]
    var instanceTag string = rootTag
    dataType := "nil"
    if data != nil {
        dataType = reflect.TypeOf(data).String()
    }

    switch dataType {
        case "map[string]interface {}":
            data = LowerCamelCaseKeys(data)
            instanceTag = itemTag
        case "map[string]map[string]interface {}":
            dataMap := data.(map[string]map[string]interface{})
            var result = make(map[string]interface{})
            for key, value := range dataMap {
                result[key] = LowerCamelCaseKeys(value)
            }
            data = result
        case "[]map[string]interface {}":
            dataSlice := data.([]map[string]interface{})
            var result []interface{}
            for _, value := range dataSlice {
                result = append(result, LowerCamelCaseKeys(value))
            }
            data = result
    }

    switch method {
        case "xsd":
            toStringData, err = getXsd(rootTag)
        case "xsdSingle":
            toStringData, err = getXsd(itemTag)
        case "json":
//             toStringData, err = json.MarshalIndent(data, "", "  ")
            var buf bytes.Buffer
            enc := json.NewEncoder(&buf)
            enc.SetEscapeHTML(false)       // <<--- disabilita l’escaping di <, >, &
            enc.SetIndent("", "  ")
            if err = enc.Encode(data); err != nil {
                return "", err
            }
            toStringData = buf.Bytes()
        case "yaml":
            outerMap := map[string]interface{}{
                instanceTag: data,
            }
            toStringData, err = yaml.Marshal(outerMap)
        case "xml":
            jsonBytes, err := json.Marshal(data)
            if err != nil {
                return "", err
            }
            var unstructured interface{}
            err = json.Unmarshal(jsonBytes, &unstructured)
            if err != nil {
                return "", err
            }
            var xmlString string
            var constructor map[string]Structs.XmlFieldMapping
            constructor = Structs.TypeMapBuildXml[instanceTag]

            if instanceTag == rootTag {
                xmlString, err = mapListToXML(unstructured, instanceTag, dataType, constructor)
            } else {
                xmlString, err = mapToXML(unstructured, instanceTag, constructor, "", 0)
            }
            if err != nil {
                fmt.Println("ERRORE XML:", err)
            }
            toStringData = []byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" + xmlString)
    }
    if err != nil {
        return "", err
    }
    return string(toStringData), nil
}


func getSelectedFields(reference Structs.GeoCodeReference) []string {
    if len(geocodesMap[reference].SetEnquiries.Select) == 0 {
        return Structs.SettingsMap[geocodesMap[reference].SetType].(Structs.SettingsType).Public
    } else {
        return geocodesMap[reference].SetEnquiries.Select
    }
}


var processMap = map[string]func(interface{}, []string) map[string]interface{}{
    "countries": func(item interface{}, selectedFields []string) map[string]interface{} {
        return filterFields(item.(Structs.Country), selectedFields)
    },
    "geoSets": func(item interface{}, selectedFields []string) map[string]interface{} {
        return filterFields(item.(Structs.GeoSet), selectedFields)
    },
    "currencies": func(item interface{}, selectedFields []string) map[string]interface{} {
        return filterFields(item.(Structs.Currency), selectedFields)
    },
}

func compareItems(a, b interface{}, orderBy string, direction string, collator *collate.Collator) bool {
    aVal := reflect.ValueOf(a).FieldByName(orderBy)
    bVal := reflect.ValueOf(b).FieldByName(orderBy)
    if aVal.Kind() == reflect.String && bVal.Kind() == reflect.String {
        comparison := collator.CompareString(aVal.String(), bVal.String())
        if direction == "ASC" {
            return comparison < 0
        } else {
            return comparison > 0
        }
    }
    return false
}


// Funzione per codificare i dati in XML
func encodeToXML(data interface{}) (string, error) {
	// Crea un buffer per l'output XML
	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	enc.Indent("", "    ")

	// Codifica in XML
	if err := enc.Encode(data); err != nil {
		return "", fmt.Errorf("error encoding XML: %v", err)
	}

	// Restituisci i dati XML come stringa
	return buf.String(), nil
}



/******/

func getGeoCodeData(
    reference Structs.GeoCodeReference,
    onlyFirst bool,
) Structs.GeoCodeResult {
    // --- 1) Init result (map o slice)
    var result interface{}
    rawObject := geocodesMap[reference].SetObject
    if geocodesMap[reference].SetEnquiries.Index != nil {
        result = make(map[string]map[string]interface{}, len(rawObject))
    } else {
        result = make([]map[string]interface{}, 0, len(rawObject))
    }

    // --- 2) Estraggo e “unwrapDeep” profondo
    type entry struct {
        raw   reflect.Value
        clean interface{}
    }
    entries := make([]entry, 0, len(rawObject))
    for _, rv := range rawObject {
        v, ok := rv.(reflect.Value)
        if !ok {
            continue
        }
        entries = append(entries, entry{
            raw:   v,
            clean: unwrapDeepValue(v),
        })
    }

    // --- 3) (facoltativo) WHERE su entries[i].clean
    // ...

    // --- 4) Ordina col collator
    orderBy  := geocodesMap[reference].SetEnquiries.OrderBy.Property
    orderDir := geocodesMap[reference].SetEnquiries.OrderBy.OrderType
    langTag  := language.Make(
        getData("config").(*Structs.Config).
            Settings.Languages.InPackage[currentLanguage],
    )
    collator := collate.New(langTag)
    sort.Slice(entries, func(i, j int) bool {
        return compareItems(
            entries[i].raw.Interface(),
            entries[j].raw.Interface(),
            orderBy, orderDir, collator,
        )
    })

    // --- 5) Offset / Limit
    offset := geocodesMap[reference].SetEnquiries.Interval.Offset
    limit  := geocodesMap[reference].SetEnquiries.Interval.Limit
    if onlyFirst {
        limit = 1
    }

    // --- 6) Costruisco result da “clean”
    in, out := 0, 0
    for _, e := range entries {
        if in < offset {
            in++
            continue
        }
        in++; out++
        if out > limit {
            break
        }

        if onlyFirst {
            return e.clean
        }

        switch r := result.(type) {
        case map[string]map[string]interface{}:
            m, ok := e.clean.(map[string]interface{})
            if !ok {
                continue
            }
            idxName := *geocodesMap[reference].SetEnquiries.Index
            key     := e.raw.FieldByName(idxName).String()
            r[key]  = m

        case []map[string]interface{}:
            m, ok := e.clean.(map[string]interface{})
            if !ok {
                continue
            }
            result = append(r, m)
        }
    }

    return result
}

func unwrapDeepValue(v reflect.Value) interface{} {
    if !v.IsValid() {
        return nil
    }
    // apri Ptr/Interface
    for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
        if v.IsNil() {
            return nil
        }
        v = v.Elem()
    }

    switch v.Kind() {
    case reflect.Struct:
        out := make(map[string]interface{}, v.NumField())
        t := v.Type()
        for i := 0; i < v.NumField(); i++ {
            f := t.Field(i)
            fv := v.Field(i)
            if fv.CanInterface() {
                out[f.Name] = unwrapDeepValue(fv)
            }
        }
        return out

    case reflect.Slice, reflect.Array:
        n := v.Len()
        arr := make([]interface{}, n)
        for i := 0; i < n; i++ {
            arr[i] = unwrapDeepValue(v.Index(i))
        }
        return arr

    case reflect.Map:
        if v.Type().Key().Kind() == reflect.String {
            out := make(map[string]interface{}, v.Len())
            for _, key := range v.MapKeys() {
                out[key.String()] = unwrapDeepValue(v.MapIndex(key))
            }
            return out
        }
        out2 := make(map[interface{}]interface{}, v.Len())
        for _, key := range v.MapKeys() {
            out2[key.Interface()] = unwrapDeepValue(v.MapIndex(key))
        }
        return out2

    default:
        return v.Interface()
    }
}
