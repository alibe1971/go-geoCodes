package geoCodes

import (
    "encoding/xml"
)

type Currencies []Currency

type Currency struct {
    Name      string    `json:"name" yaml:"name" xml:"name"`
	IsoAlpha  string    `json:"isoAlpha" yaml:"isoAlpha" xml:"isoAlpha"`
	IsoNumber string    `json:"isoNumber" yaml:"isoNumber" xml:"isoNumber"`
	Symbol    *string   `json:"symbol" yaml:"symbol" xml:"symbol,omitempty"`
	Decimal   *int64    `json:"decimal" yaml:"decimal" xml:"decimal,omitempty"`
}

var CurrencySettings = SettingsType {
    PrimaryKey: "IsoAlpha",
    Indexes: []string{
        "Name",
        "IsoAlpha",
        "IsoNumber",
    },
    Public: []string{
        "Name",
        "IsoAlpha",
        "IsoNumber",
        "Symbol",
        "Decimal",
    },
}

type CurrenciesXml struct {
    XMLName  xml.Name           `xml:"currencies"`
    Currencies []CurrencyXml    `xml:"currency"`
}

type CurrencyXml struct {
    XMLName   xml.Name `xml:"currency"`
    Index     string   `xml:"index,attr,omitempty"`
    Currency
}

func CurrencyToXML(currency Currency) CurrencyXml {
    return CurrencyXml{
        Currency:     currency,
    }
}
