package geoCodes

type GeoCodeReference string

type GeoCode struct {
	SetType         string
	SetLocale       string
	SetObject       map[string]interface{}
	SetEnquiries    Enquiries
}

type GeoCodeResult interface{}

type SettingsType struct {
    PrimaryKey string
    Indexes    []string
    Public     []string
}

var SettingsMap = map[string]interface{}{
    "countries": CountrySettings,
    "geoSets":  GeoSetSettings,
    "currencies":  CurrencySettings,
}

var TransSettingsMap = map[string]interface{}{
    "countries": TransCountry{},
    "geoSets":  TransGenericItem{},
    "currencies":  TransGenericItem{},
}

type Enquiries struct {
	Index       *string
	Select      []string
    OrderBy     OrderByStruct
    Interval    IntervalStruct
}

type IntervalStruct struct {
    Offset  int
    Limit   int
}

type OrderByStruct struct {
    Property    string
    OrderType   string
}

var SingleItemName = map[string]string{
    "countries": "country",
    "geoSets":  "geoSet",
    "currencies":  "currency",
}

var TypeMap = map[string]func() interface{}{
    "countries":    func() interface{} { return &Countries{} },
    "country":      func() interface{} { return &Country{} },
    "currencies":   func() interface{} { return &Currencies{} },
    "currency":     func() interface{} { return &Currency{} },
    "geoSets":      func() interface{} { return &GeoSets{} },
    "geoSet":       func() interface{} { return &GeoSet{} },
}

var TypeMapXml = map[string]func() interface{}{
    "countries":    func() interface{} { return &CountriesXml{} },
    "currencies":   func() interface{} { return &CurrenciesXml{} },
    "geoSets":   func() interface{} { return &GeoSetsXml{} },
}

var ConverterMapXml = map[string]func(interface{}) interface{}{
    "country": func(data interface{}) interface{} {
        country := data.(*Country)
        return CountryToXML(*country)
    },
    "currency": func(data interface{}) interface{} {
        currency := data.(*Currency)
        return CurrencyToXML(*currency)
    },
    "geoSet": func(data interface{}) interface{} {
        geoSet := data.(*GeoSet)
        return GeoSetToXML(*geoSet)
    },
}

type CDATA struct {
	Value string `xml:",cdata"`
}




