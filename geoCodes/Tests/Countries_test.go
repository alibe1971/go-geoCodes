package geoCodesTest

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/alibe1971/go-geoCodes/geoCodes"
    "github.com/alibe1971/go-geoCodes/geoCodes/Tests/TestLib"
    "fmt"
    "math/rand"
)

// var Countries = geoCodes.Countries()

const countriesTotalCount int = 250
const countriesPrimaryKey string = "Alpha2"
var countriesGlobalObject = map[string]string{
    "firstElement": "AD",
    "lastElement":  "ZW",
}


func TestCountries(t *testing.T) {
    t.Run("TestTheCountriesFunctionalities", func(t *testing.T) {
        t.Run("CheckTheCountriesObjectIsCorrectlyInstantiated", func(t *testing.T) {
            assert.NotNil(t, geoCodes.Countries(), "The object `Countries` cannot be `nil`")
            t.Run("CheckTheCountriesObjectHasTheCorrectNumberOfElements", func(t *testing.T) {
                // Check with the alias commands
                lengthObj := geoCodes.Countries().Length()
                countObj  := geoCodes.Countries().Count()
                assert.True(
                    t,
                    countObj == lengthObj && lengthObj == countriesTotalCount,
                    "The number of the elements in the object must be " + fmt.Sprint(countriesTotalCount),
                )
            })
        })

        t.Run("CheckTheCountriesAsListOfElements:`.Get()`", func(t *testing.T) {
            t.Run("CheckTheCountriesAsSliceListOfElements", func(t *testing.T) {
                countries := geoCodes.Countries().Get()
                //** Let's work on the first element **//
                countryTypeAssertion := countries.Data.([]map[string]interface{})[0]
                assert.Equal(
                    t,
                    countryTypeAssertion["Alpha2"],
                    countries.Pick("0.Alpha2"),
                    "Wrong match for `Alpha2`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["Alpha3"],
                    countries.Pick("0.Alpha3"),
                    "Wrong match for `Alpha3`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["UnM49"],
                    countries.Pick("0.UnM49"),
                    "Wrong match for `UnM49`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["Name"],
                    countries.Pick("0.Name"),
                    "Wrong match for `Name`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["FullName"],
                    countries.Pick("0.FullName"),
                    "Wrong match for `FullName`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["Dependency"],
                    countries.Pick("0.Dependency"),
                    "Wrong match for `Dependency`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["CcTld"],
                    countries.Pick("0.CcTld"),
                    "Wrong match for `CcTld`",
                )

                assert.Equal(
                    t,
                    countryTypeAssertion["OfficialName"],
                    countries.Pick("0.OfficialName"),
                    "Wrong match for `OfficialName`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["OfficialName"].(map[string]interface{})["ca"],
                    countries.Pick("0.OfficialName.ca"),
                    "Wrong match for `OfficialName.ca`",
                )

                assert.Equal(
                    t,
                    countryTypeAssertion["Flags"],
                    countries.Pick("0.Flags"),
                    "Wrong match for `Flags`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["Flags"].(map[string]interface{})["Svg"],
                    countries.Pick("0.Flags.Svg"),
                    "Wrong match for `Flags.Svg`",
                )

                assert.Equal(
                    t,
                    countryTypeAssertion["Mottos"],
                    countries.Pick("0.Mottos"),
                    "Wrong match for `Mottos`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["Mottos"].(map[string]interface{})["Official"],
                    countries.Pick("0.Mottos.Official"),
                    "Wrong match for `Mottos.Official`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["Mottos"].(map[string]interface{})["Official"].(map[string]interface{})["la"],
                    countries.Pick("0.Mottos.Official.la"),
                    "Wrong match for `Mottos.Official.la`",
                )

                assert.Equal(
                    t,
                    countryTypeAssertion["Currencies"],
                    countries.Pick("0.Currencies"),
                    "Wrong match for `Currencies`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["Currencies"].(map[string]interface{})["LegalTenders"],
                    countries.Pick("0.Currencies.LegalTenders"),
                    "Wrong match for `Currencies.LegalTenders`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["Currencies"].(map[string]interface{})["LegalTenders"].([]interface{})[0],
                    countries.Pick("0.Currencies.LegalTenders.0"),
                    "Wrong match for `Currencies.LegalTenders.0`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["Currencies"].(map[string]interface{})["WidelyAccepted"],
                    countries.Pick("0.Currencies.WidelyAccepted"),
                    "Wrong match for `Currencies.WidelyAccepted`",
                )

                assert.Equal(
                    t,
                    countryTypeAssertion["DialCodes"],
                    countries.Pick("0.DialCodes"),
                    "Wrong match for `DialCodes`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["DialCodes"].(map[string]interface{})["Main"],
                    countries.Pick("0.DialCodes.Main"),
                    "Wrong match for `DialCodes.Main`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["DialCodes"].(map[string]interface{})["Main"].([]interface{})[0],
                    countries.Pick("0.DialCodes.Main.0"),
                    "Wrong match for `DialCodes.Main.0`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["DialCodes"].(map[string]interface{})["Exceptions"],
                    countries.Pick("0.DialCodes.Exceptions"),
                    "Wrong match for `DialCodes.Exceptions`",
                )

                assert.Equal(
                    t,
                    countryTypeAssertion["TimeZones"],
                    countries.Pick("0.TimeZones"),
                    "Wrong match for `TimeZones`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["TimeZones"].([]interface{})[0],
                    countries.Pick("0.TimeZones.0"),
                    "Wrong match for `TimeZones.0`",
                )

                assert.Equal(
                    t,
                    countryTypeAssertion["Locales"],
                    countries.Pick("0.Locales"),
                    "Wrong match for `Locales`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["Locales"].([]interface{})[0],
                    countries.Pick("0.Locales.0"),
                    "Wrong match for `Locales.0`",
                )

                assert.Equal(
                    t,
                    countryTypeAssertion["OtherAppsIds"],
                    countries.Pick("0.OtherAppsIds"),
                    "Wrong match for `OtherAppsIds`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["OtherAppsIds"].(map[string]interface{})["GeoNamesOrg"],
                    countries.Pick("0.OtherAppsIds.GeoNamesOrg"),
                    "Wrong match for `OtherAppsIds.GeoNamesOrg`",
                )
            })

            t.Run("CheckTheCountriesAsMapListOfElementsUsing:`.WithIndex()`", func(t *testing.T) {
                countries := geoCodes.Countries().WithIndex(countriesPrimaryKey).Get()
                //** Let's work on the `AE` element **//
                countryTypeAssertion := countries.Data.(map[string]map[string]interface{})["AE"]
                assert.Equal(
                    t,
                    countryTypeAssertion["Alpha2"],
                    countries.Pick("AE.Alpha2"),
                    "Wrong match for `Alpha2`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["Alpha3"],
                    countries.Pick("AE.Alpha3"),
                    "Wrong match for `Alpha3`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["UnM49"],
                    countries.Pick("AE.UnM49"),
                    "Wrong match for `UnM49`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["Name"],
                    countries.Pick("AE.Name"),
                    "Wrong match for `Name`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["FullName"],
                    countries.Pick("AE.FullName"),
                    "Wrong match for `FullName`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["Dependency"],
                    countries.Pick("AE.Dependency"),
                    "Wrong match for `Dependency`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["CcTld"],
                    countries.Pick("AE.CcTld"),
                    "Wrong match for `CcTld`",
                )

                assert.Equal(
                    t,
                    countryTypeAssertion["OfficialName"],
                    countries.Pick("AE.OfficialName"),
                    "Wrong match for `OfficialName`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["OfficialName"].(map[string]interface{})["ar"],
                    countries.Pick("AE.OfficialName.ar"),
                    "Wrong match for `OfficialName.ar`",
                )

                assert.Equal(
                    t,
                    countryTypeAssertion["Flags"],
                    countries.Pick("AE.Flags"),
                    "Wrong match for `Flags`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["Flags"].(map[string]interface{})["Svg"],
                    countries.Pick("AE.Flags.Svg"),
                    "Wrong match for `Flags.Svg`",
                )

                assert.Equal(
                    t,
                    countryTypeAssertion["Mottos"],
                    countries.Pick("AE.Mottos"),
                    "Wrong match for `Mottos`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["Mottos"].(map[string]interface{})["Official"],
                    countries.Pick("AE.Mottos.Official"),
                    "Wrong match for `Mottos.Official`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["Mottos"].(map[string]interface{})["Official"].(map[string]interface{})["ar"],
                    countries.Pick("AE.Mottos.Official.ar"),
                    "Wrong match for `Mottos.Official.ar`",
                )

                assert.Equal(
                    t,
                    countryTypeAssertion["Currencies"],
                    countries.Pick("AE.Currencies"),
                    "Wrong match for `Currencies`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["Currencies"].(map[string]interface{})["LegalTenders"],
                    countries.Pick("AE.Currencies.LegalTenders"),
                    "Wrong match for `Currencies.LegalTenders`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["Currencies"].(map[string]interface{})["LegalTenders"].([]interface{})[0],
                    countries.Pick("AE.Currencies.LegalTenders.0"),
                    "Wrong match for `Currencies.LegalTenders.0`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["Currencies"].(map[string]interface{})["WidelyAccepted"],
                    countries.Pick("AE.Currencies.WidelyAccepted"),
                    "Wrong match for `Currencies.WidelyAccepted`",
                )

                assert.Equal(
                    t,
                    countryTypeAssertion["DialCodes"],
                    countries.Pick("AE.DialCodes"),
                    "Wrong match for `DialCodes`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["DialCodes"].(map[string]interface{})["Main"],
                    countries.Pick("AE.DialCodes.Main"),
                    "Wrong match for `DialCodes.Main`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["DialCodes"].(map[string]interface{})["Main"].([]interface{})[0],
                    countries.Pick("AE.DialCodes.Main.0"),
                    "Wrong match for `DialCodes.Main.0`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["DialCodes"].(map[string]interface{})["Exceptions"],
                    countries.Pick("AE.DialCodes.Exceptions"),
                    "Wrong match for `DialCodes.Exceptions`",
                )

                assert.Equal(
                    t,
                    countryTypeAssertion["TimeZones"],
                    countries.Pick("AE.TimeZones"),
                    "Wrong match for `TimeZones`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["TimeZones"].([]interface{})[0],
                    countries.Pick("AE.TimeZones.0"),
                    "Wrong match for `TimeZones.0`",
                )

                assert.Equal(
                    t,
                    countryTypeAssertion["Locales"],
                    countries.Pick("AE.Locales"),
                    "Wrong match for `Locales`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["Locales"].([]interface{})[0],
                    countries.Pick("AE.Locales.0"),
                    "Wrong match for `Locales.0`",
                )

                assert.Equal(
                    t,
                    countryTypeAssertion["OtherAppsIds"],
                    countries.Pick("AE.OtherAppsIds"),
                    "Wrong match for `OtherAppsIds`",
                )
                assert.Equal(
                    t,
                    countryTypeAssertion["OtherAppsIds"].(map[string]interface{})["GeoNamesOrg"],
                    countries.Pick("AE.OtherAppsIds.GeoNamesOrg"),
                    "Wrong match for `OtherAppsIds.GeoNamesOrg`",
                )
            })
        })

        t.Run("CheckTheCountryAsSingleElement:`.First()`", func(t *testing.T) {
            country := geoCodes.Countries().First()
            countryTypeAssertion := country.Data.(map[string]interface{})
            assert.Equal(
                t,
                countryTypeAssertion["Alpha2"],
                country.Pick("Alpha2"),
                "Wrong match for `Alpha2`",
            )
            assert.Equal(
                t,
                countryTypeAssertion["Alpha3"],
                country.Pick("Alpha3"),
                "Wrong match for `Alpha3`",
            )
            assert.Equal(
                t,
                countryTypeAssertion["UnM49"],
                country.Pick("UnM49"),
                "Wrong match for `UnM49`",
            )
            assert.Equal(
                t,
                countryTypeAssertion["Name"],
                country.Pick("Name"),
                "Wrong match for `Name`",
            )
            assert.Equal(
                t,
                countryTypeAssertion["FullName"],
                country.Pick("FullName"),
                "Wrong match for `FullName`",
            )
            assert.Equal(
                t,
                countryTypeAssertion["Dependency"],
                country.Pick("Dependency"),
                "Wrong match for `Dependency`",
            )
            assert.Equal(
                t,
                countryTypeAssertion["CcTld"],
                country.Pick("CcTld"),
                "Wrong match for `CcTld`",
            )

            assert.Equal(
                t,
                countryTypeAssertion["OfficialName"],
                country.Pick("OfficialName"),
                "Wrong match for `OfficialName`",
            )
            assert.Equal(
                t,
                countryTypeAssertion["OfficialName"].(map[string]interface{})["ca"],
                country.Pick("OfficialName.ca"),
                "Wrong match for `OfficialName.ca`",
            )

            assert.Equal(
                t,
                countryTypeAssertion["Flags"],
                country.Pick("Flags"),
                "Wrong match for `Flags`",
            )
            assert.Equal(
                t,
                countryTypeAssertion["Flags"].(map[string]interface{})["Svg"],
                country.Pick("Flags.Svg"),
                "Wrong match for `Flags.Svg`",
            )

            assert.Equal(
                t,
                countryTypeAssertion["Mottos"],
                country.Pick("Mottos"),
                "Wrong match for `Mottos`",
            )
            assert.Equal(
                t,
                countryTypeAssertion["Mottos"].(map[string]interface{})["Official"],
                country.Pick("Mottos.Official"),
                "Wrong match for `Mottos.Official`",
            )
            assert.Equal(
                t,
                countryTypeAssertion["Mottos"].(map[string]interface{})["Official"].(map[string]interface{})["la"],
                country.Pick("Mottos.Official.la"),
                "Wrong match for `Mottos.Official.la`",
            )

            assert.Equal(
                t,
                countryTypeAssertion["Currencies"],
                country.Pick("Currencies"),
                "Wrong match for `Currencies`",
            )
            assert.Equal(
                t,
                countryTypeAssertion["Currencies"].(map[string]interface{})["LegalTenders"],
                country.Pick("Currencies.LegalTenders"),
                "Wrong match for `Currencies.LegalTenders`",
            )
            assert.Equal(
                t,
                countryTypeAssertion["Currencies"].(map[string]interface{})["LegalTenders"].([]interface{})[0],
                country.Pick("Currencies.LegalTenders.0"),
                "Wrong match for `Currencies.LegalTenders.0`",
            )
            assert.Equal(
                t,
                countryTypeAssertion["Currencies"].(map[string]interface{})["WidelyAccepted"],
                country.Pick("Currencies.WidelyAccepted"),
                "Wrong match for `Currencies.WidelyAccepted`",
            )

            assert.Equal(
                t,
                countryTypeAssertion["DialCodes"],
                country.Pick("DialCodes"),
                "Wrong match for `DialCodes`",
            )
            assert.Equal(
                t,
                countryTypeAssertion["DialCodes"].(map[string]interface{})["Main"],
                country.Pick("DialCodes.Main"),
                "Wrong match for `DialCodes.Main`",
            )
            assert.Equal(
                t,
                countryTypeAssertion["DialCodes"].(map[string]interface{})["Main"].([]interface{})[0],
                country.Pick("DialCodes.Main.0"),
                "Wrong match for `DialCodes.Main.0`",
            )
            assert.Equal(
                t,
                countryTypeAssertion["DialCodes"].(map[string]interface{})["Exceptions"],
                country.Pick("DialCodes.Exceptions"),
                "Wrong match for `DialCodes.Exceptions`",
            )

            assert.Equal(
                t,
                countryTypeAssertion["TimeZones"],
                country.Pick("TimeZones"),
                "Wrong match for `TimeZones`",
            )
            assert.Equal(
                t,
                countryTypeAssertion["TimeZones"].([]interface{})[0],
                country.Pick("TimeZones.0"),
                "Wrong match for `TimeZones.0`",
            )

            assert.Equal(
                t,
                countryTypeAssertion["Locales"],
                country.Pick("Locales"),
                "Wrong match for `Locales`",
            )
            assert.Equal(
                t,
                countryTypeAssertion["Locales"].([]interface{})[0],
                country.Pick("Locales.0"),
                "Wrong match for `Locales.0`",
            )

            assert.Equal(
                t,
                countryTypeAssertion["OtherAppsIds"],
                country.Pick("OtherAppsIds"),
                "Wrong match for `OtherAppsIds`",
            )
            assert.Equal(
                t,
                countryTypeAssertion["OtherAppsIds"].(map[string]interface{})["GeoNamesOrg"],
                country.Pick("OtherAppsIds.GeoNamesOrg"),
                "Wrong match for `OtherAppsIds.GeoNamesOrg`",
            )

            t.Run("TestThatTheUseOf`.WithIndex()`HasNoInfluenceOn`.First()`", func(t *testing.T) {
                _, okWithIndex := geoCodes.Countries().
                    WithIndex(countriesPrimaryKey).
                    First().
                    Data.(map[string]interface{})
                assert.True(
                    t,
                    okWithIndex,
                    "Wrong Type",
                )
                _, okWithoutIndex := geoCodes.Countries().First().Data.(map[string]interface{})
                assert.True(
                    t,
                    okWithoutIndex,
                    "Wrong Type",
                )
                assert.Equal(
                    t,
                    geoCodes.Countries().First().Pick(countriesPrimaryKey),
                    geoCodes.Countries().WithIndex(countriesPrimaryKey).First().Pick(countriesPrimaryKey),
                    "WithIndex().First() is different from First()",
                )
            })
        })

        t.Run("TestThe`.Pick()`Features", func(t *testing.T) {
            t.Run("TestThe`.Pick()`Aliases", func(t *testing.T) {
                country := geoCodes.Countries().First()
                pick := country.Pick(countriesPrimaryKey)
                val := country.Val(countriesPrimaryKey)
                value := country.Value(countriesPrimaryKey)
                lookup := country.Lookup(countriesPrimaryKey)
                assert.True(
                    t,
                    pick == val && val == value && value == lookup && lookup == countriesGlobalObject["firstElement"],
                    "Wrong Type",
                )
            })

            t.Run("TestTheBehaviorOf`.Pick()`WithWrongProperty", func(t *testing.T) {
                defer func() {
                    if r := recover(); r != nil {
                      assert.Contains(t, r.(string), "not found")
                      return
                    }
                    t.Error("Expected panic, but no panic occurred")
                }()
                geoCodes.Countries().First().Pick("NotExistentPropertyName")
            })
        })

        t.Run("TestTheStringEndpoints", func(t *testing.T) {
            t.Run("TestThe`.ToJson()`Endpoint", func(t *testing.T) {
                assert.Nil(
                    t,
                    TestLib.ValidateJSON([]byte(geoCodes.Countries().Get().ToJson())),
                    "Not a valid Json",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateJSON([]byte(geoCodes.Countries().WithIndex(countriesPrimaryKey).Get().ToJson())),
                    "Not a valid Json",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateJSON([]byte(geoCodes.Countries().First().ToJson())),
                    "Not a valid Json",
                )
            })
            t.Run("TestThe`.ToYaml()`Endpoint", func(t *testing.T) {
                assert.Nil(
                    t,
                    TestLib.ValidateYAML([]byte(geoCodes.Countries().Get().ToYaml())),
                    "Not a valid Yaml",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateYAML([]byte(geoCodes.Countries().WithIndex(countriesPrimaryKey).Get().ToYaml())),
                    "Not a valid Yaml",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateYAML([]byte(geoCodes.Countries().First().ToYaml())),
                    "Not a valid Yaml",
                )
            })
            t.Run("TestThe`.ToXml()`Endpoint", func(t *testing.T) {
                assert.Nil(
                    t,
                    TestLib.ValidateXML([]byte(geoCodes.Countries().Get().ToXml())),
                    "Not a valid Xml",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateXML([]byte(geoCodes.Countries().WithIndex(countriesPrimaryKey).Get().ToXml())),
                    "Not a valid Xml",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateXML([]byte(geoCodes.Countries().First().ToXml())),
                    "Not a valid Xml",
                )
            })
            t.Run("TestTheExistenceForTheXsdRelatedToTheList(`.GetXsd()`)Endpoint", func(t *testing.T) {
                xsd := geoCodes.Countries().GetXsd()
                assert.NotEmpty(t, xsd, "The content of the XSD is empty")
                assert.Contains(
                    t,
                    xsd,
                    "<xs:schema", "The XSD does not contain the <xs:schema> tag, so it may not be valid",
                )
            })
            t.Run("TestTheExistenceForTheXsdRelatedToTheSingleObject(`.GetXsdSingle()`)Endpoint", func(t *testing.T) {
                xsd := geoCodes.Countries().GetXsdSingle()
                assert.NotEmpty(t, xsd, "The content of the XSD is empty")
                assert.Contains(
                    t,
                    xsd,
                    "<xs:schema", "The XSD does not contain the <xs:schema> tag, so it may not be valid",
                )
            })

            t.Run("TestThe`.ToFlatten()`Endpoint", func(t *testing.T) {
                list := geoCodes.Countries().Get()
                listFlatten := list.ToFlatten(".")
                for i := 0; i < 5; i++ {
                    key := rand.Intn(countriesTotalCount)
                    assert.Equal(
                        t,
                        list.Pick(fmt.Sprintf("%d.Alpha2", key)),
                        listFlatten[fmt.Sprintf("%d.Alpha2", key)],
                        fmt.Sprintf(
                            "The flatten structure does not work (issue on `Alpha2` for %v)",
                            list.Pick(fmt.Sprintf("%d.Alpha2", key)),
                        ),
                    )
                    assert.Equal(
                        t,
                        list.Pick(fmt.Sprintf("%d.Alpha3", key)),
                        listFlatten[fmt.Sprintf("%d.Alpha3", key)],
                        fmt.Sprintf(
                            "The flatten structure does not work (issue on `Alpha3` for %v)",
                            list.Pick(fmt.Sprintf("%d.Alpha3", key)),
                        ),
                    )
                    assert.Equal(
                        t,
                        list.Pick(fmt.Sprintf("%d.UnM49", key)),
                        listFlatten[fmt.Sprintf("%d.UnM49", key)],
                        fmt.Sprintf(
                            "The flatten structure does not work (issue on `UnM49` for %v)",
                            list.Pick(fmt.Sprintf("%d.UnM49", key)),
                        ),
                    )
                    assert.Equal(
                        t,
                        list.Pick(fmt.Sprintf("%d.Name", key)),
                        listFlatten[fmt.Sprintf("%d.Name", key)],
                        fmt.Sprintf(
                            "The flatten structure does not work (issue on `Name` for %v)",
                            list.Pick(fmt.Sprintf("%d.Name", key)),
                        ),
                    )
                    assert.Equal(
                        t,
                        list.Pick(fmt.Sprintf("%d.Dependency", key)),
                        listFlatten[fmt.Sprintf("%d.Dependency", key)],
                        fmt.Sprintf(
                            "The flatten structure does not work (issue on `Dependency` for %v)",
                            list.Pick(fmt.Sprintf("%d.Dependency", key)),
                        ),
                    )
                }

                t.Run("TestTheBehaviorOf`.ToFlatten()`WithWrongProperty", func(t *testing.T) {
                    assert.Nil(
                        t,
                        geoCodes.Countries().First().ToFlatten(".")["NotExistentPropertyName"],
                        "The value must be `nil`",
                    )
                })
            })
        })

        t.Run("TestThePackageLanguages", func(t *testing.T) {
            geoCodes.UseLanguage("en")
            assert.Equal(
                t,
                geoCodes.Countries().WithIndex(countriesPrimaryKey).Get().Pick("IE.FullName"),
                "Republic of Ireland",
                "The chosen language does not seem to work",
            )
            geoCodes.UseLanguage("it")
            assert.Equal(
                t,
                geoCodes.Countries().WithIndex(countriesPrimaryKey).Get().Pick("IE.FullName"),
                "Repubblica d'Irlanda",
                "The chosen language does not seem to work",
            )
            geoCodes.UseDefaultLanguage()
            assert.Equal(
                t,
                geoCodes.Countries().WithIndex(countriesPrimaryKey).Get().Pick("IE.FullName"),
                "Republic of Ireland",
                "The chosen language does not seem to work",
            )
        })
    })
}


// func TestElibeCountries(t *testing.T) {
//     country0 := geoCodes.Countries().WithIndex(countriesPrimaryKey).Get().Pick("IE.FullName")
// //     country0 := geoCodes.Countries().First().ToFlatten(".")["FullNameWrong"]
//     fmt.Printf("%v\n", country0)
//     TestLib.WriteDataToFile("\n")
//
// }
