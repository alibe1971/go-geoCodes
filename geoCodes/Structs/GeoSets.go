package geoCodes

import (
    "encoding/xml"
)

type GeoSets []GeoSet

type GeoSet struct {
	Name            string   `json:"name" yaml:"name" xml:"name"`
	InternalCode    string   `json:"internalCode" yaml:"internalCode" xml:"internalCode"`
	UnM49           *string  `json:"unM49" yaml:"unM49" xml:"unM49"`
	Tags            []string `json:"tags" yaml:"tags" xml:"tags>tag"`
	CountryCodes    []string `json:"countryCodes" yaml:"countryCodes" xml:"countryCodes>cc"`
}


var GeoSetSettings = SettingsType {
    PrimaryKey: "InternalCode",
    Indexes: []string{
        "Name",
        "InternalCode",
    },
    Public: []string{
        "Name",
        "InternalCode",
        "UnM49",
        "Tags",
        "CountryCodes",
    },
}

type GeoSetsXml struct {
    XMLName  xml.Name     `xml:"geoSets"`
    GeoSets []GeoSetXml    `xml:"geoSet"`
}

type GeoSetXml struct {
    XMLName   xml.Name `xml:"geoSet"`
    Index     string   `xml:"index,attr,omitempty"`
    GeoSet
}

func GeoSetToXML(geoSet GeoSet) GeoSetXml {
    return GeoSetXml{
        GeoSet:     geoSet,
    }
}