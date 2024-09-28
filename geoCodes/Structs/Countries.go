package geoCodes

import (
    "encoding/xml"
)

type Countries []Country

type Country struct {
    Name            string                  `json:"name" yaml:"name" xml:"name"`
    FullName        string                  `json:"fullName" yaml:"fullName" xml:"fullName"`
    OfficialName    map[string]string       `json:"officialName" yaml:"officialName" xml:"officialName>name"`
    Alpha2          string                  `json:"alpha2" yaml:"alpha2" xml:"alpha2"`
    Alpha3          string                  `json:"alpha3" yaml:"alpha3" xml:"alpha3"`
    UnM49           string                  `json:"unM49" yaml:"unM49" xml:"unM49"`
    Flags           Flags                   `json:"flags" yaml:"flags" xml:"flags"`
    Dependency      *string                 `json:"dependency" yaml:"dependency" xml:"dependency"`
    Mottos          Mottos                  `json:"mottos" yaml:"mottos" xml:"mottos"`
    Currencies      CcCurrencies            `json:"currencies" yaml:"currencies" xml:"currencies"`
    DialCodes       DialCodes               `json:"dialCodes" yaml:"dialCodes" xml:"dialCodes"`
    CcTld           *string                 `json:"ccTld" yaml:"ccTld" xml:"ccTld"`
    TimeZones       []string                `json:"timeZones" yaml:"timeZones" xml:"timeZones>tz"`
    Languages       []string                `json:"languages" yaml:"languages" xml:"languages>lang"`
    Locales         []string                `json:"locales" yaml:"locales" xml:"locales>locale"`
    OtherAppsIds    OtherAppsIds            `json:"otherAppsIds" yaml:"otherAppsIds" xml:"otherAppsIds"`
    Keywords        []string                `json:"keywords" yaml:"keywords" xml:"keywords"`
}

type Flags struct {
	Svg string `json:"svg" yaml:"svg" xml:"svg"`
}

type Mottos struct {
	Official map[string]string  `json:"official" yaml:"official" xml:"official"`
}

type CcCurrencies struct {
	LegalTenders   []string      `json:"legalTenders" yaml:"legalTenders" xml:"legalTenders>currency"`
	WidelyAccepted []string      `json:"widelyAccepted" yaml:"widelyAccepted" xml:"widelyAccepted>currency"`
}

type DialCodes struct {
	Main       []string      `json:"main" yaml:"main" xml:"main>code"`
	Exceptions []string      `json:"exceptions" yaml:"exceptions" xml:"exceptions>code"`
}

type OtherAppsIds struct {
	GeoNamesOrg *int64 `json:"geoNamesOrg" yaml:"geoNamesOrg" xml:"geoNamesOrg"`
}

var CountrySettings = SettingsType {
    PrimaryKey: "Alpha2",
    Indexes: []string{
        "Name",
        "FullName",
        "Alpha2",
        "Alpha3",
        "UnM49",
    },
    Public: []string{
        "Name",
        "FullName",
        "OfficialName",
        "Alpha2",
        "Alpha3",
        "UnM49",
        "Flags",
        "Flags.Svg",
        "Dependency",
        "Mottos",
        "Mottos.Official",
        "Currencies",
        "Currencies.LegalTenders",
        "Currencies.WidelyAccepted",
        "DialCodes",
        "DialCodes.Main",
        "DialCodes.Exceptions",
        "CcTld",
        "TimeZones",
        "Languages",
        "Locales",
        "OtherAppsIds",
        "OtherAppsIds.GeoNamesOrg",
    },
}

type CountriesXml struct {
    XMLName  xml.Name    `xml:"countries"`
    Countries []CountryXml `xml:"country"`
}

type CountryXml struct {
    XMLName      xml.Name        `xml:"country"`
    Index        string          `xml:"index,attr,omitempty"`
    Country
    OfficialName []LangStructXml `xml:"officialName>name"`
    Mottos       MottosXml       `xml:"mottos"`
    Flags        FlagsXml        `xml:"flags"`
}

type LangStructXml struct {
    Lang  string `xml:"lang,attr"`
    Value string `xml:",chardata"`
}

type MottosXml struct {
	Official []LangStructXml  `json:"official" yaml:"official" xml:"official>motto"`
}

type FlagsXml struct {
	Svg CDATA `xml:"svg"`
}

func CountryToXML(country Country) CountryXml {
    officialName := []LangStructXml{}
    for key, value := range country.OfficialName {
        officialName = append(officialName, LangStructXml{Lang: key, Value: value})
    }

    mottos := []LangStructXml{}
    for key, value := range country.Mottos.Official {
        mottos = append(mottos, LangStructXml{Lang: key, Value: value})
    }

    return CountryXml{
        Country:      country,
        OfficialName: officialName,
        Mottos:       MottosXml{Official: mottos},
        Flags:        FlagsXml{Svg: CDATA{Value: country.Flags.Svg}},
    }
}






