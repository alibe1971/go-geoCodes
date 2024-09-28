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
    "os"
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


func getDataOnString(reference Structs.GeoCodeReference, data interface{}, method string) (string, error) {
    var toStringData []byte
    var xsd []byte
    var err error
    var rootTag string = geocodesMap[reference].SetType
    var itemTag string = Structs.SingleItemName[rootTag]
    var constructor func() interface{}

    dataType := "nil"
    if data != nil {
        dataType = reflect.TypeOf(data).String()
    }

    instanceTag := rootTag
    if dataType == "map[string]interface {}" || dataType == "map[string]map[string]interface {}" {
        instanceTag = itemTag
    }
    constructor = Structs.TypeMap[instanceTag]

    switch method {
        case "json","yaml":
            switch dataType {
                case "map[string]interface {}", "[]map[string]interface {}":
                    instance := constructor()
                    tmp, _ := json.Marshal(data)
                    if err := json.Unmarshal(tmp, instance); err != nil {
                        return "", err
                    }
                    data = instance
                case "map[string]map[string]interface {}":
                    dataMap := data.(map[string]map[string]interface{})
                    var result = make(map[string]interface{})
                    for key, value := range dataMap {
                        instance := constructor()
                        tmp, _ := json.Marshal(value)
                        if err := json.Unmarshal(tmp, instance); err != nil {
                            return "", err
                        }
                        result[key] = instance
                    }
                    data = result
            }
        case "xml", "xmlValidate":
            switch dataType {
                case "map[string]interface {}":
                    instance := constructor()
                    tmp, _ := json.Marshal(data)
                    if err := json.Unmarshal(tmp, instance); err != nil {
                        return "", err
                    }
                    converter, _ := Structs.ConverterMapXml[instanceTag]
                    data = converter(instance)
                case "[]map[string]interface {}":
                    constructor := Structs.TypeMap[itemTag]
                    dataMap := data.([]map[string]interface{})
                    var result interface{}
                    switch rootTag {
                        case "countries":
                            result = Structs.TypeMapXml[rootTag]().(*Structs.CountriesXml)
                        case "currencies":
                            result = Structs.TypeMapXml[rootTag]().(*Structs.CurrenciesXml)
                        case "geoSets":
                            result = Structs.TypeMapXml[rootTag]().(*Structs.GeoSetsXml)
                    }
                    for _, value := range dataMap {
                        instance := constructor()
                        tmp, _ := json.Marshal(value)
                        json.Unmarshal(tmp, instance)
                        converter := Structs.ConverterMapXml[itemTag]
                        switch rootTag {
                            case "countries":
                                countriesResult := result.(*Structs.CountriesXml)
                                countriesResult.Countries = append(countriesResult.Countries, converter(instance).(Structs.CountryXml))
                            case "currencies":
                                currenciesResult := result.(*Structs.CurrenciesXml)
                                currenciesResult.Currencies = append(currenciesResult.Currencies, converter(instance).(Structs.CurrencyXml))
                            case "geoSets":
                                geoSetsResult := result.(*Structs.GeoSetsXml)
                                geoSetsResult.GeoSets = append(geoSetsResult.GeoSets, converter(instance).(Structs.GeoSetXml))
                        }
                    }
                    data = result

                case "map[string]map[string]interface {}":
                    constructor := Structs.TypeMap[itemTag]
                    dataMap := data.(map[string]map[string]interface{})
                    var result interface{}

                    switch rootTag {
                        case "countries":
                            result = Structs.TypeMapXml[rootTag]().(*Structs.CountriesXml)
                        case "currencies":
                            result = Structs.TypeMapXml[rootTag]().(*Structs.CurrenciesXml)
                        case "geoSets":
                            result = Structs.TypeMapXml[rootTag]().(*Structs.GeoSetsXml)
                    }
                    for key, value := range dataMap {
                        instance := constructor()
                        tmp, _ := json.Marshal(value)
                        json.Unmarshal(tmp, instance)
                        converter := Structs.ConverterMapXml[itemTag]
                        switch rootTag {
                            case "countries":
                                countryInstance := converter(instance).(Structs.CountryXml)
                                countryInstance.Index = key
                                countriesResult := result.(*Structs.CountriesXml)
                                countriesResult.Countries = append(countriesResult.Countries, countryInstance)
                            case "currencies":
                                currencyInstance := converter(instance).(Structs.CurrencyXml)
                                currencyInstance.Index = key
                                currenciesResult := result.(*Structs.CurrenciesXml)
                                currenciesResult.Currencies = append(currenciesResult.Currencies, currencyInstance)
                            case "geoSets":
                                geoSetInstance := converter(instance).(Structs.GeoSetXml)
                                geoSetInstance.Index = key
                                geoSetsResult := result.(*Structs.GeoSetsXml)
                                geoSetsResult.GeoSets = append(geoSetsResult.GeoSets, geoSetInstance)
                        }
                    }
                    data = result

            }
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
        case "xml", "xmlValidate":
            toStringData, err = xml.MarshalIndent(data, "", "  ")
            if method == "xmlValidate" {
                xsd, err = getXsd(instanceTag)
                validateXMLAgainstXSD(toStringData, xsd)
            }
    }

    if err != nil {
        return "", err
    }
    return string(toStringData), nil
}

func validateXMLAgainstXSD(xmlData []byte, xsdSchema []byte) {
//     [todo]  https://chatgpt.com/c/66e16d6d-1ec8-8004-aacb-878ba7bd24fc
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

//     processMap := map[string]func(interface{}) map[string]interface{}{
//         "countries": func(item interface{}) map[string]interface{} {
//             return filterFields(item.(Structs.Country), selectedFields)
//         },
//         "geoSets": func(item interface{}) map[string]interface{} {
//             return filterFields(item.(Structs.GeoSet), selectedFields)
//         },
//         "currencies": func(item interface{}) map[string]interface{} {
//             return filterFields(item.(Structs.Currency), selectedFields)
//         },
//     }

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


func STICA() {
    fmt.Println("STICA\n")
    os.Exit(1)
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



