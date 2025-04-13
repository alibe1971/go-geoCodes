package geoCodesTest

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/alibe1971/go-geoCodes/geoCodes"
    "github.com/alibe1971/go-geoCodes/geoCodes/Tests/TestLib"
    "fmt"
)

var currenciesTotalCount int = 180

var currenciesPrimaryKey string = "IsoAlpha"
var currenciesFirstElementOfTheObjectPrimKey string = "AED"
var currenciesLastElementOfTheObjectPrimKey  string = "ZWL"

func TestCurrencies(t *testing.T) {
    t.Run("TestTheGeoSetsFunctionality", func(t *testing.T) {
        t.Run("CheckTheGeoSetsObjectIsCorrectlyInstantiated", func(t *testing.T) {
            assert.NotNil(t, geoCodes.Currencies(), "The object `GeoSets` cannot be `nil`")
            t.Run("CheckTheGeoSetsObjectHasTheCorrectNumberOfElements", func(t *testing.T) {
                // Check with the alias commands
                lengthObj := geoCodes.Currencies().Length()
                countObj  := geoCodes.Currencies().Count()
                assert.Equal(t, countObj, lengthObj, "The methods Length() e Count() do not seem to be alias")
                assert.Equal(
                    t,
                    countObj,
                    currenciesTotalCount,
                    "The number of the elements in the object must be " + fmt.Sprint(currenciesTotalCount),
                )
            })
        })

        t.Run("CheckTheGeoSetsAsListOfElements:Get()", func(t *testing.T) {
            t.Run("CheckTheGeoSetsAsSliceListOfElements", func(t *testing.T) {
                t.Run("CheckTheGeoSetsListWithStructureDeclaration", func(t *testing.T) {
                    currency := geoCodes.Currencies().Get().Data.([]map[string]interface{})[0]
                    assert.Equal(
                        t,
                        currency[currenciesPrimaryKey],
                        currenciesFirstElementOfTheObjectPrimKey,
                        fmt.Sprintf("The first element of the object (%v) does not match the expected one (%v)",
                            currency[currenciesPrimaryKey],
                            currenciesFirstElementOfTheObjectPrimKey,
                        ),
                    )
                })

                t.Run("TestTheGeoSetsListAsSliceWithDirectCommand:AsSlice()", func(t *testing.T) {
                    currency := geoCodes.Currencies().Get().AsSlice()[0]
                    assert.Equal(
                        t,
                        currency[currenciesPrimaryKey],
                        currenciesFirstElementOfTheObjectPrimKey,
                        fmt.Sprintf("The first element of the object (%v) does not match the expected one (%v)",
                            currency[currenciesPrimaryKey],
                            currenciesFirstElementOfTheObjectPrimKey,
                        ),
                    )
                })
            })

            t.Run("CheckTheGeoSetsAsMapListOfElements", func(t *testing.T) {
                t.Run("CheckTheGeoSetsListWithStructureDeclaration", func(t *testing.T) {
                    currency := geoCodes.Currencies().WithIndex(currenciesPrimaryKey).Get().
                        Data.(map[string]map[string]interface{})["EUR"]
                    assert.Equal(
                        t,
                        currency[currenciesPrimaryKey],
                        "EUR",
                        fmt.Sprintf("The selected element of the object (%v) does not match the expected one (%v)",
                            currency[currenciesPrimaryKey],
                            "EUR",
                        ),
                    )
                })

                t.Run("TestTheGeoSetsListAsMapWithDirectCommand:AsMap()", func(t *testing.T) {
                    currency := geoCodes.Currencies().WithIndex(currenciesPrimaryKey).Get().AsMap()["EUR"]
                    assert.Equal(
                        t,
                        currency[currenciesPrimaryKey],
                        "EUR",
                        fmt.Sprintf("The selected element of the object (%v) does not match the expected one (%v)",
                            currency[currenciesPrimaryKey],
                            "EUR",
                        ),
                    )
                })
            })

        })

        t.Run("CheckTheGeoSetAsSingleElement:First()", func(t *testing.T) {
            t.Run("CheckTheGeoSetWithStructureDeclaration", func(t *testing.T) {
                currency := geoCodes.Currencies().First().Data.(map[string]interface{})
                assert.Equal(
                    t,
                    currency[currenciesPrimaryKey],
                    currenciesFirstElementOfTheObjectPrimKey,
                    fmt.Sprintf("The first element of the object (%v) does not match the expected one (%v)",
                        currency[currenciesPrimaryKey],
                        currenciesFirstElementOfTheObjectPrimKey,
                    ),
                )
            })

            t.Run("TestTheGeoSetAsObjectWithDirectCommand:AsObj()", func(t *testing.T) {
                currency := geoCodes.Currencies().First().AsObj()
                assert.Equal(
                    t,
                    currency[currenciesPrimaryKey],
                    currenciesFirstElementOfTheObjectPrimKey,
                    fmt.Sprintf("The first element of the object (%v) does not match the expected one (%v)",
                        currency[currenciesPrimaryKey],
                        currenciesFirstElementOfTheObjectPrimKey,
                    ),
                )
            })

            t.Run("TestThatThe:WithIndex():HasNotInfluenceOnDirectCommand:AsObj()", func(t *testing.T) {
                currency := geoCodes.Currencies().WithIndex(currenciesPrimaryKey).First().AsObj()
                assert.Equal(
                    t,
                    currency[currenciesPrimaryKey],
                    currenciesFirstElementOfTheObjectPrimKey,
                    fmt.Sprintf("The first element of the object (%v) does not match the expected one (%v)",
                        currency[currenciesPrimaryKey],
                        currenciesFirstElementOfTheObjectPrimKey,
                    ),
                )
            })

        })

        t.Run("TestTheWrongUseOf:AsMap():AsSlice():AsObj()", func(t *testing.T) {
            t.Run("TestTheWrongUseInPresenceOfSliceList", func(t *testing.T) {
                currency := geoCodes.Currencies().Get()
                t.Run("WrongUseOf:AsMap()", func(t *testing.T) {
                    assert.Empty(t, currency.AsMap(), "The map should be empty")
                })
                t.Run("WrongUseOf:AsObj()", func(t *testing.T) {
                    assert.Empty(t, currency.AsObj(), "The object should be empty")
                })
            })
            t.Run("TestTheWrongUseInPresenceOfMapList", func(t *testing.T) {
                currency := geoCodes.Currencies().WithIndex(currenciesPrimaryKey).Get()
                t.Run("WrongUseOf:AsSlice()", func(t *testing.T) {
                    assert.Empty(t, currency.AsSlice(), "The slice should be empty")
                })
                t.Run("WrongUseOf:AsObj()", func(t *testing.T) {
                    assert.Empty(t, currency.AsObj(), "The object should be empty")
                })
            })
            t.Run("TestTheWrongUseInPresenceOfSingleObject", func(t *testing.T) {
                currency := geoCodes.Currencies().First()
                t.Run("WrongUseOf:AsSlice()", func(t *testing.T) {
                    assert.Empty(t, currency.AsSlice(), "The slice should be empty")
                })
                t.Run("WrongUseOf:AsMap()", func(t *testing.T) {
                    assert.Empty(t, currency.AsMap(), "The map should be empty")
                })
            })
        })

        t.Run("TestTheStringEndpoints", func(t *testing.T) {
            t.Run("TestThe:ToJson():Endpoint", func(t *testing.T) {
                assert.Nil(
                    t,
                    TestLib.ValidateJSON([]byte(geoCodes.Currencies().Get().ToJson())),
                    "Not a valid Json",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateJSON([]byte(geoCodes.Currencies().WithIndex(currenciesPrimaryKey).Get().ToJson())),
                    "Not a valid Json",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateJSON([]byte(geoCodes.Currencies().First().ToJson())),
                    "Not a valid Json",
                )
            })
            t.Run("TestThe:ToYaml():Endpoint", func(t *testing.T) {
                assert.Nil(
                    t,
                    TestLib.ValidateYAML([]byte(geoCodes.Currencies().Get().ToYaml())),
                    "Not a valid Yaml",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateYAML([]byte(geoCodes.Currencies().WithIndex(currenciesPrimaryKey).Get().ToYaml())),
                    "Not a valid Yaml",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateYAML([]byte(geoCodes.Currencies().First().ToYaml())),
                    "Not a valid Yaml",
                )
            })
            t.Run("TestThe:ToXml():Endpoint", func(t *testing.T) {
                assert.Nil(
                    t,
                    TestLib.ValidateXML([]byte(geoCodes.Currencies().Get().ToXml())),
                    "Not a valid Xml",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateXML([]byte(geoCodes.Currencies().WithIndex(currenciesPrimaryKey).Get().ToXml())),
                    "Not a valid Xml",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateXML([]byte(geoCodes.Currencies().First().ToXml())),
                    "Not a valid Xml",
                )
            })
            t.Run("TestTheExistenceForTheXsdRelatedToTheList:GetXsd():Endpoint", func(t *testing.T) {
                xsd := geoCodes.Currencies().GetXsd()
                assert.NotEmpty(t, xsd, "The content of the XSD is empty")
                assert.Contains(
                    t,
                    xsd,
                    "<xs:schema", "The XSD does not contain the <xs:schema> tag, so it may not be valid",
                )
            })
            t.Run("TestTheExistenceForTheXsdRelatedToTheSingleObject:GetXsdSingle():Endpoint", func(t *testing.T) {
                xsd := geoCodes.Currencies().GetXsdSingle()
                assert.NotEmpty(t, xsd, "The content of the XSD is empty")
                assert.Contains(
                    t,
                    xsd,
                    "<xs:schema", "The XSD does not contain the <xs:schema> tag, so it may not be valid",
                )
            })
        })

        t.Run("TestThePackageLanguages", func(t *testing.T) {
            geoCodes.UseLanguage("en")
            assert.Equal(
                t,
                geoCodes.Currencies().WithIndex(currenciesPrimaryKey).Get().AsMap()["AED"]["Name"],
                "UAE Dirham",
                "The chosen language does not seem to work",
            )
            geoCodes.UseLanguage("it")
            assert.Equal(
                t,
                geoCodes.Currencies().WithIndex(currenciesPrimaryKey).Get().AsMap()["AED"]["Name"],
                "Dirham degli Emirati Arabi Uniti",
                "The chosen language does not seem to work",
            )
        })
    })
}


func TestElibeCurrencies(t *testing.T) {
//     currency0 := geoCodes.Currencies().First().AsObj()["Name"]
//     currency0 := geoCodes.Currencies().Get().AsSlice()[0]
//     currency0 := geoCodes.Currencies().WithIndex(currenciesPrimaryKey).Get().AsMap()["EUR"]

//     currency0 := geoCodes.Currencies().First().ToJson()
//     currency0 := geoCodes.Currencies().Get().ToJson()
//     currency0 := geoCodes.Currencies().WithIndex(currenciesPrimaryKey).Get().ToJson()

//     currency0 := geoCodes.Currencies().First().ToYaml()
//     currency0 := geoCodes.Currencies().Get().ToYaml()
//     currency0 := geoCodes.Currencies().WithIndex(currenciesPrimaryKey).Get().ToYaml()

    currency0 := geoCodes.Currencies().First().ToXml()
//     currency0 := geoCodes.Currencies().Get().ToXml()
//     currency0 := geoCodes.Currencies().WithIndex(currenciesPrimaryKey).Get().ToXml()



//     geoCodes.UseLanguage("it")
//     currency0 := geoCodes.Currencies().WithIndex(currenciesPrimaryKey).Get().AsMap()["EUR"]["Name"]
    fmt.Printf("%v", currency0)
}
