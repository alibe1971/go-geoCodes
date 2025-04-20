package geoCodesTest

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/alibe1971/go-geoCodes/geoCodes"
    "github.com/alibe1971/go-geoCodes/geoCodes/Tests/TestLib"
    "fmt"
    "math/rand"
)

var geoSetsTotalCount int = 62

var geoSetsPrimaryKey string = "InternalCode"
var geoSetsFirstElementOfTheObjectPrimKey string = "CONV-G20"
var geoSetsLastElementOfTheObjectPrimKey  string = "ZONE-EZ"

func TestGeoSets(t *testing.T) {
    t.Run("TestTheGeoSetsFunctionality", func(t *testing.T) {
        t.Run("CheckTheGeoSetsObjectIsCorrectlyInstantiated", func(t *testing.T) {
            assert.NotNil(t, geoCodes.GeoSets(), "The object `GeoSets` cannot be `nil`")
            t.Run("CheckTheGeoSetsObjectHasTheCorrectNumberOfElements", func(t *testing.T) {
                // Check with the alias commands
                lengthObj := geoCodes.GeoSets().Length()
                countObj  := geoCodes.GeoSets().Count()
                assert.Equal(t, countObj, lengthObj, "The methods Length() e Count() do not seem to be alias")
                assert.Equal(
                    t,
                    countObj,
                    geoSetsTotalCount,
                    "The number of the elements in the object must be " + fmt.Sprint(geoSetsTotalCount),
                )
            })
        })

        t.Run("CheckTheGeoSetsAsListOfElements:Get()", func(t *testing.T) {
            t.Run("CheckTheGeoSetsAsSliceListOfElements", func(t *testing.T) {
                t.Run("CheckTheGeoSetsListWithStructureDeclaration", func(t *testing.T) {
                    geoSet := geoCodes.GeoSets().Get().Data.([]map[string]interface{})[0]
                    assert.Equal(
                        t,
                        geoSet[geoSetsPrimaryKey],
                        geoSetsFirstElementOfTheObjectPrimKey,
                        fmt.Sprintf("The first element of the object (%v) does not match the expected one (%v)",
                            geoSet[geoSetsPrimaryKey],
                            geoSetsFirstElementOfTheObjectPrimKey,
                        ),
                    )
                })

                t.Run("TestTheGeoSetsListAsSliceWithDirectCommand:AsSlice()", func(t *testing.T) {
                    geoSet := geoCodes.GeoSets().Get().AsSlice()[0]
                    assert.Equal(
                        t,
                        geoSet[geoSetsPrimaryKey],
                        geoSetsFirstElementOfTheObjectPrimKey,
                        fmt.Sprintf("The first element of the object (%v) does not match the expected one (%v)",
                            geoSet[geoSetsPrimaryKey],
                            geoSetsFirstElementOfTheObjectPrimKey,
                        ),
                    )
                })
            })

            t.Run("CheckTheGeoSetsAsMapListOfElements", func(t *testing.T) {
                t.Run("CheckTheGeoSetsListWithStructureDeclaration", func(t *testing.T) {
                    geoSet := geoCodes.GeoSets().WithIndex(geoSetsPrimaryKey).Get().
                        Data.(map[string]map[string]interface{})["ORGS-EU"]
                    assert.Equal(
                        t,
                        geoSet[geoSetsPrimaryKey],
                        "ORGS-EU",
                        fmt.Sprintf("The selected element of the object (%v) does not match the expected one (%v)",
                            geoSet[geoSetsPrimaryKey],
                            "ORGS-EU",
                        ),
                    )
                })

                t.Run("TestTheGeoSetsListAsMapWithDirectCommand:AsMap()", func(t *testing.T) {
                    geoSet := geoCodes.GeoSets().WithIndex(geoSetsPrimaryKey).Get().AsMap()["ORGS-EU"]
                    assert.Equal(
                        t,
                        geoSet[geoSetsPrimaryKey],
                        "ORGS-EU",
                        fmt.Sprintf("The selected element of the object (%v) does not match the expected one (%v)",
                            geoSet[geoSetsPrimaryKey],
                            "ORGS-EU",
                        ),
                    )
                })
            })

        })

        t.Run("CheckTheGeoSetAsSingleElement:First()", func(t *testing.T) {
            t.Run("CheckTheGeoSetWithStructureDeclaration", func(t *testing.T) {
                geoSet := geoCodes.GeoSets().First().Data.(map[string]interface{})
                assert.Equal(
                    t,
                    geoSet[geoSetsPrimaryKey],
                    geoSetsFirstElementOfTheObjectPrimKey,
                    fmt.Sprintf("The first element of the object (%v) does not match the expected one (%v)",
                        geoSet[geoSetsPrimaryKey],
                        geoSetsFirstElementOfTheObjectPrimKey,
                    ),
                )
            })

            t.Run("TestTheGeoSetAsObjectWithDirectCommand:AsObj()", func(t *testing.T) {
                geoSet := geoCodes.GeoSets().First().AsObj()
                assert.Equal(
                    t,
                    geoSet[geoSetsPrimaryKey],
                    geoSetsFirstElementOfTheObjectPrimKey,
                    fmt.Sprintf("The first element of the object (%v) does not match the expected one (%v)",
                        geoSet[geoSetsPrimaryKey],
                        geoSetsFirstElementOfTheObjectPrimKey,
                    ),
                )
            })

            t.Run("TestThatThe:WithIndex():HasNotInfluenceOnDirectCommand:AsObj()", func(t *testing.T) {
                geoSet := geoCodes.GeoSets().WithIndex(geoSetsPrimaryKey).First().AsObj()
                assert.Equal(
                    t,
                    geoSet[geoSetsPrimaryKey],
                    geoSetsFirstElementOfTheObjectPrimKey,
                    fmt.Sprintf("The first element of the object (%v) does not match the expected one (%v)",
                        geoSet[geoSetsPrimaryKey],
                        geoSetsFirstElementOfTheObjectPrimKey,
                    ),
                )
            })

        })

        t.Run("TestTheWrongUseOf:AsMap():AsSlice():AsObj()", func(t *testing.T) {
            t.Run("TestTheWrongUseInPresenceOfSliceList", func(t *testing.T) {
                geoSet := geoCodes.GeoSets().Get()
                t.Run("WrongUseOf:AsMap()", func(t *testing.T) {
                    assert.Empty(t, geoSet.AsMap(), "The map should be empty")
                })
                t.Run("WrongUseOf:AsObj()", func(t *testing.T) {
                    assert.Empty(t, geoSet.AsObj(), "The object should be empty")
                })
            })
            t.Run("TestTheWrongUseInPresenceOfMapList", func(t *testing.T) {
                geoSet := geoCodes.GeoSets().WithIndex(geoSetsPrimaryKey).Get()
                t.Run("WrongUseOf:AsSlice()", func(t *testing.T) {
                    assert.Empty(t, geoSet.AsSlice(), "The slice should be empty")
                })
                t.Run("WrongUseOf:AsObj()", func(t *testing.T) {
                    assert.Empty(t, geoSet.AsObj(), "The object should be empty")
                })
            })
            t.Run("TestTheWrongUseInPresenceOfSingleObject", func(t *testing.T) {
                geoSet := geoCodes.GeoSets().First()
                t.Run("WrongUseOf:AsSlice()", func(t *testing.T) {
                    assert.Empty(t, geoSet.AsSlice(), "The slice should be empty")
                })
                t.Run("WrongUseOf:AsMap()", func(t *testing.T) {
                    assert.Empty(t, geoSet.AsMap(), "The map should be empty")
                })
            })
        })

        t.Run("TestTheStringEndpoints", func(t *testing.T) {
            t.Run("TestThe:ToJson():Endpoint", func(t *testing.T) {
                assert.Nil(
                    t,
                    TestLib.ValidateJSON([]byte(geoCodes.GeoSets().Get().ToJson())),
                    "Not a valid Json",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateJSON([]byte(geoCodes.GeoSets().WithIndex(geoSetsPrimaryKey).Get().ToJson())),
                    "Not a valid Json",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateJSON([]byte(geoCodes.GeoSets().First().ToJson())),
                    "Not a valid Json",
                )
            })
            t.Run("TestThe:ToYaml():Endpoint", func(t *testing.T) {
                assert.Nil(
                    t,
                    TestLib.ValidateYAML([]byte(geoCodes.GeoSets().Get().ToYaml())),
                    "Not a valid Yaml",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateYAML([]byte(geoCodes.GeoSets().WithIndex(geoSetsPrimaryKey).Get().ToYaml())),
                    "Not a valid Yaml",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateYAML([]byte(geoCodes.GeoSets().First().ToYaml())),
                    "Not a valid Yaml",
                )
            })
            t.Run("TestThe:ToXml():Endpoint", func(t *testing.T) {
                assert.Nil(
                    t,
                    TestLib.ValidateXML([]byte(geoCodes.GeoSets().Get().ToXml())),
                    "Not a valid Xml",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateXML([]byte(geoCodes.GeoSets().WithIndex(geoSetsPrimaryKey).Get().ToXml())),
                    "Not a valid Xml",
                )
                assert.Nil(
                    t,
                    TestLib.ValidateXML([]byte(geoCodes.GeoSets().First().ToXml())),
                    "Not a valid Xml",
                )
            })
            t.Run("TestTheExistenceForTheXsdRelatedToTheList:GetXsd():Endpoint", func(t *testing.T) {
                xsd := geoCodes.GeoSets().GetXsd()
                assert.NotEmpty(t, xsd, "The content of the XSD is empty")
                assert.Contains(
                    t,
                    xsd,
                    "<xs:schema", "The XSD does not contain the <xs:schema> tag, so it may not be valid",
                )
            })
            t.Run("TestTheExistenceForTheXsdRelatedToTheSingleObject:GetXsdSingle():Endpoint", func(t *testing.T) {
                xsd := geoCodes.GeoSets().GetXsdSingle()
                assert.NotEmpty(t, xsd, "The content of the XSD is empty")
                assert.Contains(
                    t,
                    xsd,
                    "<xs:schema", "The XSD does not contain the <xs:schema> tag, so it may not be valid",
                )
            })

            t.Run("TestThe:ToFlatten():Endpoint", func(t *testing.T) {
                list := geoCodes.GeoSets().Get()
                listSlice := list.AsSlice()
                listFlatten := list.ToFlatten(".")
                for i := 0; i < 5; i++ {
                    key := rand.Intn(geoSetsTotalCount)

                    assert.Equal(
                        t,
                        listSlice[key]["InternalCode"],
                        listFlatten[fmt.Sprintf("%d.InternalCode", key)],
                        fmt.Sprintf(
                            "The flatten structure does not work (issue on `Alpha2` for %v)",
                            listSlice[key]["InternalCode"],
                        ),
                    )
                    assert.Equal(
                        t,
                        listSlice[key]["UnM49"],
                        listFlatten[fmt.Sprintf("%d.UnM49", key)],
                        fmt.Sprintf(
                            "The flatten structure does not work (issue on `UnM49` for %v)",
                            listSlice[key]["InternalCode"],
                        ),
                    )
                    assert.Equal(
                        t,
                        listSlice[key]["Name"],
                        listFlatten[fmt.Sprintf("%d.Name", key)],
                        fmt.Sprintf(
                            "The flatten structure does not work (issue on `Name` for %v)",
                            listSlice[key]["InternalCode"],
                        ),
                    )
                }

            })
        })

        t.Run("TestThePackageLanguages", func(t *testing.T) {
            geoCodes.UseLanguage("en")
            assert.Equal(
                t,
                geoCodes.GeoSets().WithIndex(geoSetsPrimaryKey).Get().AsMap()["ORGS-WTO"]["Name"],
                "World Trade Organization (WTO)",
                "The chosen language does not seem to work",
            )
            geoCodes.UseLanguage("it")
            assert.Equal(
                t,
                geoCodes.GeoSets().WithIndex(geoSetsPrimaryKey).Get().AsMap()["ORGS-WTO"]["Name"],
                "Organizzazione Mondiale del Commercio (OMC)",
                "The chosen language does not seem to work",
            )
        })
    })
}


// func TestElibeGeoSets(t *testing.T) {
// //     geoSet0 := geoCodes.GeoSets().First().AsObj()["Name"]
// //     geoSet0 := geoCodes.GeoSets().Get().AsSlice()[0]
// //     geoSet0 := geoCodes.GeoSets().WithIndex(geoSetsPrimaryKey).Get().AsMap()["ORGS-EU"]
//
// //     geoSet0 := geoCodes.GeoSets().First().ToJson()
// //     geoSet0 := geoCodes.GeoSets().Get().ToJson()
// //     geoSet0 := geoCodes.GeoSets().WithIndex(geoSetsPrimaryKey).Get().ToJson()
//
// //     geoSet0 := geoCodes.GeoSets().First().ToYaml()
// //     geoSet0 := geoCodes.GeoSets().Get().ToYaml()
// //     geoSet0 := geoCodes.GeoSets().WithIndex(geoSetsPrimaryKey).Get().ToYaml()
//
// //     geoSet0 := geoCodes.GeoSets().First().ToXml()
// //     geoSet0 := geoCodes.GeoSets().Get().ToXml()
// //     geoSet0 := geoCodes.GeoSets().WithIndex(geoSetsPrimaryKey).Get().ToXml()
//
//
//
// //     geoCodes.UseLanguage("it")
// //     geoSet0 := geoCodes.GeoSets().WithIndex(geoSetsPrimaryKey).Get().AsMap()["ORGS-EU"]["Name"]
// //     fmt.Printf("%v", geoSet0)
// }
