package geoCodes

import (
    Structs "github.com/alibe1971/go-geoCodes/geoCodes/Structs"
    data "github.com/alibe1971/go-geoCodes/geoCodes/Data"
)


var superDefaultLanguage string
var defaultLanguage string
var currentLanguage string
var availableLanguages []string

var geocodesMap = make(map[Structs.GeoCodeReference]*Structs.GeoCode)
var dataMap = make(map[string]interface{})
var dataDefaultLength = make(map[string]int)

func Initialize() {
    loadData("config", data.Config, &Structs.Config{})
    var languages = getData("config").(*Structs.Config).Settings.Languages
    superDefaultLanguage = languages.Default
    defaultLanguage = superDefaultLanguage
    currentLanguage = superDefaultLanguage

    keys := make([]string, 0, len(languages.InPackage))
    for key := range languages.InPackage {
        keys = append(keys, key)
    }
    availableLanguages = keys
}

func GetLanguages(typeStr string) interface{} {
	switch typeStr {
        case "default":
            return defaultLanguage
        case "current":
            return currentLanguage
        case "available":
            return availableLanguages
        default:
            return nil
	}
}

func SetLanguage(typeStr string, language string) {
    language = strToLower(language)
    valid := isInSlice(availableLanguages, language)
    if !valid {
        logPanicWithStackTrace("not a valid language")
    }
    switch typeStr {
        case "default":
            defaultLanguage = language
        case "current":
            currentLanguage = language
    }
}

func ResetLanguages() {
    defaultLanguage = superDefaultLanguage
    currentLanguage = superDefaultLanguage
}

func InitializeGeoCodeSet(setStr string) Structs.GeoCodeReference {
    geoCode, err := initializeGeoCodeSet(setStr)
    if err != nil {
        logPanicWithStackTrace("PACKAGE NOT INITIALIZED")
    }
    var reference Structs.GeoCodeReference
    reference = generateUniqueString()
    geocodesMap[reference] = geoCode
    return reference
}

func OutPutObject(reference Structs.GeoCodeReference, method string) interface{} {
	switch method {
        case "get":
            return getGeoCodeData(reference, false)
        case "first":
            return getGeoCodeData(reference, true)
	}
	return nil
}

func OutPutString(reference Structs.GeoCodeReference, data interface{}, method string) string {
    toStringData, err := getDataOnString(reference, data, method)
    if err != nil {
        logPanicWithStackTrace("Error occurred: " + err.Error())
        return ""
    }
    return toStringData
}


func Setters(reference Structs.GeoCodeReference, method string, values ...interface{}) {
    settings := Structs.SettingsMap[geocodesMap[reference].SetType].(Structs.SettingsType)
    switch method {
        case "index":
            setIndex(settings, reference, toString(values[0]))
        case "select":
            props := setElementsToStrings(values)
            setSelect(settings, reference, props...)
        case "orderBy":
            props := setElementsToStrings(values)
            setOrderBy(settings, reference, props...)
        case "offset", "limit":
            props := setElementsToIntegers(values)
            setInterval(reference, method, props[0])
    }
}
