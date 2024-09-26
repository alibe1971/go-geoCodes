package geoCodes

import (
    "reflect"
    Structs "github.com/alibe1971/go-geoCodes/geoCodes/Structs"
    data "github.com/alibe1971/go-geoCodes/geoCodes/Data"
    transData "github.com/alibe1971/go-geoCodes/geoCodes/Data/Translations"
)

func initializeGeoCodeSet(setStr string) (*Structs.GeoCode, error) {
    var SetObject map[string]interface{}
    var primaryKey string
    locale := getData("config").(*Structs.Config).Settings.Languages.InPackage[currentLanguage]
    switch setStr {
        case "countries":
            SetObject = initializeGeoCode(setStr, data.Countries, &Structs.Countries{}, transData.Countries, reflect.TypeOf(Structs.TransCountries{}))
            primaryKey = Structs.CountrySettings.PrimaryKey
        case "geoSets":
            SetObject = initializeGeoCode(setStr, data.GeoSets, &Structs.GeoSets{}, transData.GeoSets, reflect.TypeOf(Structs.TransGeneric{}))
            primaryKey = Structs.GeoSetSettings.PrimaryKey
        case "currencies":
            SetObject = initializeGeoCode(setStr, data.Currencies, &Structs.Currencies{}, transData.Currencies, reflect.TypeOf(Structs.TransGeneric{}))
            primaryKey = Structs.CurrencySettings.PrimaryKey
    }

    return &Structs.GeoCode{
        SetType:        setStr,
        SetObject:      SetObject,
        SetLocale:      locale,
        SetEnquiries:   Structs.Enquiries {
            Interval: Structs.IntervalStruct{
                Offset: 0,
                Limit:  dataDefaultLength[setStr],
            },
            OrderBy: Structs.OrderByStruct{
                Property:   primaryKey,
                OrderType:  "ASC",
            },
        },
    }, nil
}

func initializeGeoCode(key string, mainData string, mainStruct interface{}, langData map[string]string, langStruct reflect.Type) map[string]interface{} {
    settings := Structs.SettingsMap[key].(Structs.SettingsType)
    if (!isInMap(dataMap, key)) {
        loadData(key, mainData, mainStruct)
    }
    langMap := loadLanguages(key, langData, langStruct)

    langFunctionalityMap := make(map[string]reflect.Value)
    for langKey, langVal := range langMap {
        langFunctionalityMap[langKey] = reflect.ValueOf(langVal).Elem()
    }
    langProperties := getStructPropertiesNamesMultilevel(Structs.TransSettingsMap[key], "")

    dataStructured := getData(key)
    dataVal := reflect.ValueOf(dataStructured)
    dataVal = dataVal.Elem()
    dataLen := dataVal.Len()
    if _, exists := dataDefaultLength[key]; !exists {
        dataDefaultLength[key] = dataLen
    }

    var dataStructuredMap = make(map[string]interface{})
    for i := 0; i < dataLen; i++ {
        prop := dataVal.Index(i)
        primaryKey := prop.FieldByName(settings.PrimaryKey).String()
        superDefault := langFunctionalityMap[superDefaultLanguage].MapIndex(reflect.ValueOf(primaryKey))
        defDefault := langFunctionalityMap[defaultLanguage].MapIndex(reflect.ValueOf(primaryKey))
        current := langFunctionalityMap[currentLanguage].MapIndex(reflect.ValueOf(primaryKey))
        for j := 0; j < prop.NumField(); j++ {

            property := prop.Type().Field(j).Name
            if isInSlice(langProperties, property) {
                translatedValue := getTheTranslatedProperty(property, current, defDefault, superDefault)
                field := prop.Field(j)
                if field.CanSet() {
                    switch field.Kind() {
                    case reflect.String:
                        if val, ok := translatedValue.(string); ok {
                            field.SetString(val)
                        }
                    case reflect.Slice:
                        if val, ok := translatedValue.([]string); ok {
                            field.Set(reflect.ValueOf(val))
                        }
                    }
                }
            }
        }
        dataStructuredMap[primaryKey] = prop
    }
    return dataStructuredMap
}

func getTheTranslatedProperty(property string, current, defDefault, superDefault reflect.Value) interface{} {
	getValue := func(v reflect.Value) interface{} {
		if v.IsValid() {
			if v.Kind() == reflect.String && v.String() != "" {
				return v.String()
			} else if v.Kind() == reflect.Slice && v.Len() > 0 {
				return v.Interface()
			}
		}
		return nil
	}

	value := getValue(current.FieldByName(property))
	if value != nil {
		return value
	}

	value = getValue(defDefault.FieldByName(property))
	if value != nil {
		return value
	}

	value = getValue(superDefault.FieldByName(property))
	return value
}

func createNewReflectStruct(myStruct reflect.Type) interface{} {
    return reflect.New(myStruct).Interface()
}

func loadLanguages(keyLang string, dataLang map[string]string, structLangType reflect.Type) map[string]interface{} {
    var langMap = make(map[string]interface{})
    for _, language := range []string{superDefaultLanguage, defaultLanguage, currentLanguage} {
        languageRef := keyLang + "_" + language
        if !isInMap(dataMap, languageRef) {
            newStruct := createNewReflectStruct(structLangType)
            loadData(languageRef, dataLang[language], newStruct)
            langMap[language] = newStruct
        } else {
            langMap[language] = getData(languageRef)
        }
    }
    return langMap
}