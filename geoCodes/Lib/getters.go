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

func getDataAsFlattenMap(reference Structs.GeoCodeReference, data interface{}, sep string) (map[string]interface{}, error) {
    if data == nil {
        return nil, nil
    }
    // 1) JSON.Marshal di QUALUNQUE data (struct, map[string]interface{}, slice, ecc.)
    b, err := json.Marshal(data)
    if err != nil {
        return nil, fmt.Errorf("json.Marshal fallita: %w", err)
    }
    // 2) JSON.Unmarshal in interface{} per catturare sia oggetti che array
    var intermediate interface{}
    if err := json.Unmarshal(b, &intermediate); err != nil {
        return nil, fmt.Errorf("json.Unmarshal fallita: %w", err)
    }
    // 3) a seconda di cosa ottengo, appiattisco la mappa o il slice
    flat := make(map[string]interface{})
    switch root := intermediate.(type) {
    case map[string]interface{}:
        flattenMap("", root, sep, flat)
    case []interface{}:
        flattenSlice("", root, sep, flat)
    default:
        return nil, fmt.Errorf("tipo root non supportato: %T", root)
    }
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
            toStringData, err = json.MarshalIndent(data, "", "  ")
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

func getGeoCodeData(reference Structs.GeoCodeReference, onlyFirst bool) Structs.GeoCodeResult {
    var result interface{}
    orderBy := geocodesMap[reference].SetEnquiries.OrderBy.Property
    orderDir := geocodesMap[reference].SetEnquiries.OrderBy.OrderType
    object := geocodesMap[reference].SetObject
    selectedFields := getSelectedFields(reference)

    if geocodesMap[reference].SetEnquiries.Index != nil {
        result = make(map[string]map[string]interface{})
    } else {
        result = make([]map[string]interface{}, 0)
    }

    processItem, _ := processMap[geocodesMap[reference].SetType]

    offsetNum := geocodesMap[reference].SetEnquiries.Interval.Offset
    limitNum := geocodesMap[reference].SetEnquiries.Interval.Limit
    if onlyFirst {
        offsetNum = 0
        limitNum = 1
    }

    kIn, kOut := 0, 0

    items := make([]interface{}, 0)
    for _, value := range object {
        items = append(items, value.(reflect.Value).Interface())
    }

    collator := collate.New(language.Make(getData("config").(*Structs.Config).Settings.Languages.InPackage[currentLanguage]))

    sort.Slice(items, func(i, j int) bool {
        return compareItems(items[i], items[j], orderBy, orderDir, collator)
    })

    for _, item := range items {
        if kIn < offsetNum {
            kIn++
            continue
        }
        kOut++
        if kOut > limitNum {
            return limitNum
        }

        parsedItem := processItem(item, selectedFields)

        if onlyFirst {
            return parsedItem
        }

        switch res := result.(type) {
        case map[string]map[string]interface{}:
            key := reflect.ValueOf(item).FieldByName(*geocodesMap[reference].SetEnquiries.Index).String()
            res[key] = parsedItem
        case []map[string]interface{}:
            result = append(res, parsedItem)
        }

        kIn++
    }

    return result
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
