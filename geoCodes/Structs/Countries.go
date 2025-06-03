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
    LocalesIcu      []string                `json:"localesIcu" yaml:"localesIcu" xml:"localesIcu>locale"`
    OtherAppsIds    OtherAppsIds            `json:"otherAppsIds" yaml:"otherAppsIds" xml:"otherAppsIds"`
    Keywords        []string                `json:"keywords" yaml:"keywords" xml:"keywords"`
}

type Flags struct {
	Emoji   string `json:"emoji" yaml:"emoji" xml:"emoji"`
	Svg     string `json:"svg" yaml:"svg" xml:"svg"`
}

type Mottos struct {
	Official map[string]string       `json:"official" yaml:"official" xml:"official"`
	Popular map[string]string       `json:"popular" yaml:"popular" xml:"popular"`
	Royal map[string]string         `json:"royal" yaml:"royal" xml:"royal"`
	Presidential map[string]string  `json:"presidential" yaml:"presidential" xml:"presidential"`
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
        "Flags.Emoji",
        "Flags.Svg",
        "Dependency",
        "Mottos",
        "Mottos.Official",
        "Mottos.Popular",
        "Mottos.Royal",
        "Mottos.Presidential",
        "Currencies",
        "Currencies.LegalTenders",
        "Currencies.WidelyAccepted",
        "DialCodes",
        "DialCodes.Main",
        "DialCodes.Exceptions",
        "CcTld",
        "TimeZones",
        "Languages",
        "LocalesIcu",
        "OtherAppsIds",
        "OtherAppsIds.GeoNamesOrg",
    },
}

var MapBuildXmlCountries = map[string]XmlFieldMapping{
    "country": {
        TagName: "country",
        Children: MapBuildXmlCountry,
    },
}

var MapBuildXmlCountry = map[string]XmlFieldMapping{
    "alpha2": {
        TagName: "alpha2",
    },
    "alpha3": {
        TagName: "alpha3",
    },
    "unM49": {
        TagName: "unM49",
    },
    "name": {
        TagName: "name",
    },
    "fullName": {
        TagName: "fullName",
    },
    "officialName": {
        TagName: "officialName",
        AsAttributes: true,
        Children: map[string]XmlFieldMapping{
            "name": {
                TagName: "name",
                AttrName: "lang", // L'attributo "lang"
            },
        },
    },
    "flags": {
        TagName: "flags",
        Children: map[string]XmlFieldMapping{
            "emoji": {
                TagName: "emoji",
            },
            "svg": {
                TagName: "svg",
                CDATA:   true,
            },
        },
    },
    "dependency": {
        TagName: "dependency",
    },
    "mottos": {
        TagName: "mottos",
        Children: map[string]XmlFieldMapping{
            "official": {
                TagName: "official",
                Children: map[string]XmlFieldMapping{
                    "motto": {
                        TagName: "motto",
                        AttrName: "lang", // L'attributo "lang" per il motto
                    },
                },
            },
            "popular": {
                TagName: "popular",
                Children: map[string]XmlFieldMapping{
                    "motto": {
                        TagName: "motto",
                        AttrName: "lang", // L'attributo "lang" per il motto
                    },
                },
            },
            "royal": {
                TagName: "royal",
                Children: map[string]XmlFieldMapping{
                    "motto": {
                        TagName: "motto",
                        AttrName: "lang", // L'attributo "lang" per il motto
                    },
                },
            },
            "presidential": {
                TagName: "presidential",
                Children: map[string]XmlFieldMapping{
                    "motto": {
                        TagName: "motto",
                        AttrName: "lang", // L'attributo "lang" per il motto
                    },
                },
            }
        },
    },
    "currencies": {
        TagName: "currencies",
        Children: map[string]XmlFieldMapping{
            "legalTenders": {
                TagName: "legalTenders",
                Children: map[string]XmlFieldMapping{
                    "currency": {
                        TagName: "currency",
                        Children: map[string]XmlFieldMapping{
                            "isoAlpha": {
                                TagName: "isoAlpha",
                            },
                            "isoNumber": {
                                TagName: "isoNumber",
                            },
                            "name": {
                                TagName: "name",
                            },
                            "symbol": {
                                TagName: "symbol",
                            },
                            "decimal": {
                                TagName: "decimal",
                            },
                        },
                    },
                },
            },
            "widelyAccepted": {
                TagName: "widelyAccepted",
                Children: map[string]XmlFieldMapping{
                    "currency": {
                        TagName: "currency",
                        Children: map[string]XmlFieldMapping{
                            "isoAlpha": {
                                TagName: "isoAlpha",
                            },
                            "isoNumber": {
                                TagName: "isoNumber",
                            },
                            "name": {
                                TagName: "name",
                            },
                            "symbol": {
                                TagName: "symbol",
                            },
                            "decimal": {
                                TagName: "decimal",
                            },
                        },
                    },
                },
            },
        },
    },
    "dialCodes": {
        TagName: "dialCodes",
        Children: map[string]XmlFieldMapping{
            "main": {
                TagName: "main",
                Children: map[string]XmlFieldMapping{
                    "dial": {
                        TagName: "dial",
                    },
                },
            },
            "exceptions": {
                TagName: "exceptions",
                Children: map[string]XmlFieldMapping{
                    "dial": {
                        TagName: "dial",
                    },
                },
            },
        },
    },
    "ccTld": {
        TagName: "ccTld",
    },
    "timeZones": {
        TagName: "timeZones",
        Children: map[string]XmlFieldMapping{
            "tz": {
                TagName: "tz",
            },
        },
    },
    "localesIcu": {
        TagName: "localesIcu",
        Children: map[string]XmlFieldMapping{
            "locale": {
                TagName: "locale",
            },
        },
    },
    "demonyms": {
        TagName: "demonyms",
        Children: map[string]XmlFieldMapping{
            "demonym": {
                TagName: "demonym",
            },
        },
    },
    "otherAppsIds": {
        TagName: "otherAppsIds",
        Children: map[string]XmlFieldMapping{
            "geoNamesOrg": {
                TagName: "geoNamesOrg",
                AsInt: true,
            },
        },
    },
}



/******/

type CountriesXml struct {
    XMLName  xml.Name    `xml:"countries"`
    Countries []CountryXml `xml:"country"`
}

type CountryXml struct {
    XMLName      xml.Name        `xml:"country"`
    Index        string          `xml:"index,attr,omitempty"`
    Country
    Dependency   string          `xml:"dependency"`
    CcTld        string          `xml:"ccTld"`
    OfficialName []LangStructXml `xml:"officialName>name"`
    Mottos       MottosXml       `xml:"mottos"`
    Flags        FlagsXml        `xml:"flags"`
}

type LangStructXml struct {
    Lang  string `xml:"lang,attr"`
    Value string `xml:",chardata"`
}

type MottosXml struct {
	Official []LangStructXml  `xml:"official>motto"`
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

    var dep string
    if country.Dependency != nil {
        dep = *country.Dependency
    } else {
        dep = ""
    }

    var tld string
    if country.CcTld != nil {
        tld = *country.CcTld
    } else {
        tld = ""
    }


    return CountryXml{
        Country:      country,
        Dependency:   dep,
        CcTld:        tld,
        OfficialName: officialName,
        Mottos:       MottosXml{Official: mottos},
        Flags:        FlagsXml{Svg: CDATA{Value: country.Flags.Svg}},
    }
}

