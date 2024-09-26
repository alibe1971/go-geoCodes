package geoCodes

import (
    lib "github.com/alibe1971/go-geoCodes/geoCodes/Lib"
    Structs "github.com/alibe1971/go-geoCodes/geoCodes/Structs"
)


/******************
*  Initialization *
*******************/
func init() {
    lib.Initialize();
}

/***********************
*  Languages Functions *
************************/
func GetAvailableLanguages() []string {
    return lib.GetLanguages("available").([]string)
}

func GetCurrentLanguage() string {
	return lib.GetLanguages("current").(string)
}

func GetDefaultLanguage() string {
    return lib.GetLanguages("default").(string)
}

func SetDefaultLanguage(lang string) {
    lib.SetLanguage("default", lang)
}

func UseLanguage(lang string) {
    lib.SetLanguage("current", lang)
}

func ResetLanguages() {
    lib.ResetLanguages()
}

/********************
*  GeoCodes Objects *
*********************/
type geoCode struct {
    Reference Structs.GeoCodeReference
}

func Countries() *geoCode {
    reference := lib.InitializeGeoCodeSet("countries")
    return &geoCode{Reference: reference}
}

func GeoSets() *geoCode {
    reference := lib.InitializeGeoCodeSet("geoSets")
    return &geoCode{Reference: reference}
}

func Currencies() *geoCode {
    reference := lib.InitializeGeoCodeSet("currencies")
    return &geoCode{Reference: reference}
}

/******************
*  Getters
*****************/
type geoCodeResult struct {
    Reference Structs.GeoCodeReference
    Data Structs.GeoCodeResult
}

func (gc *geoCode) Get() *geoCodeResult {
    data := lib.OutPutObject(gc.Reference, "get")
    return &geoCodeResult{
        Reference: gc.Reference,
        Data: data,
    }
}

func (gc *geoCode) First() *geoCodeResult {
    data := lib.OutPutObject(gc.Reference, "first")
    return &geoCodeResult{
        Reference: gc.Reference,
        Data: data,
    }
}

func (gc *geoCode) Count() int {
    data := lib.OutPutObject(gc.Reference, "get")
    switch v := data.(type) {
        case []map[string]interface{}:
            return len(v)
        case map[string]map[string]interface{}:
            return len(v)
        default:
            return 0
    }
}

func (gc *geoCode) GetXsd() string {
    return lib.OutPutString(gc.Reference, nil, "xsd")
}

func (gc *geoCode) GetXsdSingle() string {
    return lib.OutPutString(gc.Reference, nil, "xsdSingle")
}


func (gcr *geoCodeResult) AsMap() (map[string]map[string]interface{}) {
    if data, ok := gcr.Data.(map[string]map[string]interface{}); ok {
        return data
    }
    return nil
}
func (gcr *geoCodeResult) AsSlice() ([]map[string]interface{}) {
    if data, ok := gcr.Data.([]map[string]interface{}); ok {
        return data
    }
    return nil
}


func (gcr *geoCodeResult) ToJson() (string) {
    return lib.OutPutString(gcr.Reference, gcr.Data, "json")
}

func (gcr *geoCodeResult) ToXml() (string) {
    return lib.OutPutString(gcr.Reference, gcr.Data, "xml")
}

func (gcr *geoCodeResult) ToXmlAndValidate() (string) {
    return lib.OutPutString(gcr.Reference, gcr.Data, "xmlValidate")
}

func (gcr *geoCodeResult) ToYaml() (string) {
    return lib.OutPutString(gcr.Reference, gcr.Data, "yaml")
}

/******************
*  Setters
*****************/
func (gc *geoCode) WithIndex(property string) *geoCode {
    lib.Setters(gc.Reference, "index", property)
    return gc
}

func (gc *geoCode) Select(properties ...string) *geoCode {
    propsInterface := make([]interface{}, len(properties))
    for i, v := range properties {
        propsInterface[i] = v
    }
    lib.Setters(gc.Reference, "select", propsInterface...)
    return gc
}

func (gc *geoCode) OrderBy(property string, orderType string) *geoCode {
    lib.Setters(gc.Reference, "orderBy", property, orderType)
    return gc
}
// func (gc *geoCode) OrderBy(property string, orderType ...string) *geoCode {
//     var orderTypeV string
//     if len(orderType) == 0 {
//         orderTypeV = ""
//     } else {
//         orderTypeV = orderType[0]
//     }
//     lib.Setters(gc.Reference, "orderBy", property, orderTypeV)
//     return gc
// }

func (gc *geoCode) Offset(offset int) *geoCode {
    lib.Setters(gc.Reference, "offset", offset)
    return gc
}
func (gc *geoCode) Skip(offset int) *geoCode {
    return gc.Offset(offset)
}

func (gc *geoCode) Limit(limit int) *geoCode {
    lib.Setters(gc.Reference, "limit", limit)
    return gc
}
func (gc *geoCode) Take(limit int) *geoCode {
    return gc.Limit(limit)
}

