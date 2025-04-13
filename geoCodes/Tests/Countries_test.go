package geoCodesTest

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/alibe1971/go-geoCodes/geoCodes"
    "github.com/alibe1971/go-geoCodes/geoCodes/Tests/TestLib"
    "fmt"
)

// var Countries = geoCodes.Countries()

var countriesTotalCount int = 250

var primaryKey string = "Alpha2"
var firstElementOfTheObjectPrimKey string = "AD"
var lastElementOfTheObjectPrimKey  string = "ZW"

func TestCountries(t *testing.T) {
    t.Run("TestTheCountriesFunctionality", func(t *testing.T) {
        t.Run("CheckTheCountriesObjectIsCorrectlyInstantiated", func(t *testing.T) {
            assert.NotNil(t, geoCodes.Countries(), "The object `Countries` cannot be `nil`")
            t.Run("CheckTheCountriesObjectHasTheCorrectNumberOfElements", func(t *testing.T) {
                // Check with the alias commands
                lengthObj := geoCodes.Countries().Length()
                countObj  := geoCodes.Countries().Count()
                assert.Equal(t, countObj, lengthObj, "The methods Length() e Count() do not seem to be alias")
                assert.Equal(
                    t,
                    countObj,
                    countriesTotalCount,
                    "The number of the elements in the object must be " + fmt.Sprint(countriesTotalCount),
                )
            })
        })

        t.Run("CheckTheCountriesAsListOfElements:Get()", func(t *testing.T) {
            t.Run("CheckTheCountriesAsSliceListOfElements", func(t *testing.T) {
                t.Run("CheckTheCountriesListWithStructureDeclaration", func(t *testing.T) {
                    country := geoCodes.Countries().Get().Data.([]map[string]interface{})[0]
                    assert.Equal(
                        t,
                        country[primaryKey],
                        firstElementOfTheObjectPrimKey,
                        fmt.Sprintf("The first element of the object (%v) does not match the expected one (%v)",
                            country[primaryKey],
                            firstElementOfTheObjectPrimKey,
                        ),
                    )
                })

                t.Run("TestTheCountriesListAsSliceWithDirectCommand:AsSlice()", func(t *testing.T) {
                    country := geoCodes.Countries().Get().AsSlice()[0]
                    assert.Equal(
                        t,
                        country[primaryKey],
                        firstElementOfTheObjectPrimKey,
                        fmt.Sprintf("The first element of the object (%v) does not match the expected one (%v)",
                            country[primaryKey],
                            firstElementOfTheObjectPrimKey,
                        ),
                    )
                })
            })

            t.Run("CheckTheCountriesAsMapListOfElements", func(t *testing.T) {
                t.Run("CheckTheCountriesListWithStructureDeclaration", func(t *testing.T) {
                    country := geoCodes.Countries().WithIndex(primaryKey).Get().
                        Data.(map[string]map[string]interface{})["IT"]
                    assert.Equal(
                        t,
                        country[primaryKey],
                        "IT",
                        fmt.Sprintf("The selected element of the object (%v) does not match the expected one (%v)",
                            country[primaryKey],
                            "IT",
                        ),
                    )
                })

                t.Run("TestTheCountriesListAsMapWithDirectCommand:AsMap()", func(t *testing.T) {
                    country := geoCodes.Countries().WithIndex(primaryKey).Get().AsMap()["IT"]
                    assert.Equal(
                        t,
                        country[primaryKey],
                        "IT",
                        fmt.Sprintf("The selected element of the object (%v) does not match the expected one (%v)",
                            country[primaryKey],
                            "IT",
                        ),
                    )
                })
            })

        })

        t.Run("CheckTheCountryAsSingleElement:First()", func(t *testing.T) {
            t.Run("CheckTheCountryWithStructureDeclaration", func(t *testing.T) {
                country := geoCodes.Countries().First().Data.(map[string]interface{})
                assert.Equal(
                    t,
                    country[primaryKey],
                    firstElementOfTheObjectPrimKey,
                    fmt.Sprintf("The first element of the object (%v) does not match the expected one (%v)",
                        country[primaryKey],
                        firstElementOfTheObjectPrimKey,
                    ),
                )
            })

            t.Run("TestTheCountryAsObjectWithDirectCommand:AsObj()", func(t *testing.T) {
                country := geoCodes.Countries().First().AsObj()
                assert.Equal(
                    t,
                    country[primaryKey],
                    firstElementOfTheObjectPrimKey,
                    fmt.Sprintf("The first element of the object (%v) does not match the expected one (%v)",
                        country[primaryKey],
                        firstElementOfTheObjectPrimKey,
                    ),
                )
            })

            t.Run("TestThatThe:WithIndex():HasNotInfluenceOnDirectCommand:AsObj()", func(t *testing.T) {
                country := geoCodes.Countries().WithIndex(primaryKey).First().AsObj()
                assert.Equal(
                    t,
                    country[primaryKey],
                    firstElementOfTheObjectPrimKey,
                    fmt.Sprintf("The first element of the object (%v) does not match the expected one (%v)",
                        country[primaryKey],
                        firstElementOfTheObjectPrimKey,
                    ),
                )
            })

        })

        t.Run("TestTheWrongUseOf:AsMap():AsSlice():AsObj()", func(t *testing.T) {
            t.Run("TestTheWrongUseInPresenceOfSliceList", func(t *testing.T) {
                country := geoCodes.Countries().Get()
                t.Run("WrongUseOf:AsMap()", func(t *testing.T) {
                    assert.Empty(t, country.AsMap(), "The map should be empty")
                })
                t.Run("WrongUseOf:AsObj()", func(t *testing.T) {
                    assert.Empty(t, country.AsObj(), "The object should be empty")
                })
            })
            t.Run("TestTheWrongUseInPresenceOfMapList", func(t *testing.T) {
                country := geoCodes.Countries().WithIndex(primaryKey).Get()
                t.Run("WrongUseOf:AsSlice()", func(t *testing.T) {
                    assert.Empty(t, country.AsSlice(), "The slice should be empty")
                })
                t.Run("WrongUseOf:AsObj()", func(t *testing.T) {
                    assert.Empty(t, country.AsObj(), "The object should be empty")
                })
            })
            t.Run("TestTheWrongUseInPresenceOfSingleObject", func(t *testing.T) {
                country := geoCodes.Countries().First()
                t.Run("WrongUseOf:AsSlice()", func(t *testing.T) {
                    assert.Empty(t, country.AsSlice(), "The slice should be empty")
                })
                t.Run("WrongUseOf:AsMap()", func(t *testing.T) {
                    assert.Empty(t, country.AsMap(), "The map should be empty")
                })
            })
        })

        t.Run("TestTheStringEndpoints", func(t *testing.T) {
            t.Run("TestThe:ToJson():Endpoint", func(t *testing.T) {
                assert.Nil(
                    t,
                    TestLib.ValidateJSON([]byte(geoCodes.Countries().Get().ToJson())),
                    "Not a valid Json",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateJSON([]byte(geoCodes.Countries().WithIndex(primaryKey).Get().ToJson())),
                    "Not a valid Json",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateJSON([]byte(geoCodes.Countries().First().ToJson())),
                    "Not a valid Json",
                )
            })
            t.Run("TestThe:ToYaml():Endpoint", func(t *testing.T) {
                assert.Nil(
                    t,
                    TestLib.ValidateYAML([]byte(geoCodes.Countries().Get().ToYaml())),
                    "Not a valid Yaml",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateYAML([]byte(geoCodes.Countries().WithIndex(primaryKey).Get().ToYaml())),
                    "Not a valid Yaml",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateYAML([]byte(geoCodes.Countries().First().ToYaml())),
                    "Not a valid Yaml",
                )
            })
            t.Run("TestThe:ToXml():Endpoint", func(t *testing.T) {
                assert.Nil(
                    t,
                    TestLib.ValidateXML([]byte(geoCodes.Countries().Get().ToXml())),
                    "Not a valid Xml",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateXML([]byte(geoCodes.Countries().WithIndex(primaryKey).Get().ToXml())),
                    "Not a valid Xml",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateXML([]byte(geoCodes.Countries().First().ToXml())),
                    "Not a valid Xml",
                )
            })
            t.Run("TestTheExistenceForTheXsdRelatedToTheList:GetXsd():Endpoint", func(t *testing.T) {
                xsd := geoCodes.Countries().GetXsd()
                assert.NotEmpty(t, xsd, "The content of the XSD is empty")
                assert.Contains(
                    t,
                    xsd,
                    "<xs:schema", "The XSD does not contain the <xs:schema> tag, so it may not be valid",
                )
            })
            t.Run("TestTheExistenceForTheXsdRelatedToTheSingleObject:GetXsdSingle():Endpoint", func(t *testing.T) {
                xsd := geoCodes.Countries().GetXsdSingle()
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
                geoCodes.Countries().WithIndex(primaryKey).Get().AsMap()["IE"]["FullName"],
                "Republic of Ireland",
                "The chosen language does not seem to work",
            )
            geoCodes.UseLanguage("it")
            assert.Equal(
                t,
                geoCodes.Countries().WithIndex(primaryKey).Get().AsMap()["IE"]["FullName"],
                "Repubblica d'Irlanda",
                "The chosen language does not seem to work",
            )
        })
    })
}


func TestElibeCountries(t *testing.T) {
     geoCodes.UseLanguage("it")
     country0 := geoCodes.Countries().WithIndex(primaryKey).Get().AsMap()["IE"]["FullName"]

//     fmt.Printf("Lingua: %v\n", geoCodes.GetAvailableLanguages())
//     fmt.Println("   \n")
//      geoCodes.UseLanguage("itss")
//     stica := "stica"


// TEST GLOBALI
    //country0 := geoCodes.Countries().First().ToJson()
//     country0 := geoCodes.Countries().Select("Currencies.LegalTenders", "Alpha3", "Name", "OfficialName").First().ToXml()
//     country0 := geoCodes.Countries().Get().ToJson()
//     country0 := geoCodes.Countries().WithIndex(primaryKey).Get().ToJson()
//     fmt.Printf("%v", country0)

//     country0 := geoCodes.Countries().WithIndex("FullName").Select("Alpha2", "Alpha3", "Name", "OfficialName").Get().ToXml()
//     fmt.Printf("%v", country0)


//     country0 := geoCodes.Countries().GetXsd()
//     fmt.Printf("%v\n", country0)
//     country0 := geoCodes.Countries().GetXsdSingle()



    fmt.Printf("%v", country0)



//     TestLib.WriteDataToFile("\n")
//     geoset0 := geoCodes.GeoSets()
//     TestLib.WriteDataToFile(geoset0)
//
//     currency0 := geoCodes.Currencies()
//     TestLib.WriteDataToFile(currency0)

}
