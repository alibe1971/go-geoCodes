package geoCodesTest

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/alibe1971/go-geoCodes/geoCodes"
    "github.com/alibe1971/go-geoCodes/geoCodes/Tests/TestLib"
    "fmt"
    "math/rand"
)

const geoSetsTotalCount int = 62
const geoSetsPrimaryKey string = "InternalCode"
var geoSetsGlobalObject = map[string]string{
    "firstElement": "CONV-G20",
    "lastElement":  "ZONE-EZ",
}

func TestGeoSets(t *testing.T) {
    t.Run("TestTheGeoSetsFunctionalities", func(t *testing.T) {
        t.Run("CheckTheGeoSetsObjectIsCorrectlyInstantiated", func(t *testing.T) {
            assert.NotNil(t, geoCodes.GeoSets(), "The object `GeoSets` cannot be `nil`")
            t.Run("CheckTheGeoSetsObjectHasTheCorrectNumberOfElements", func(t *testing.T) {
                // Check with the alias commands
                lengthObj := geoCodes.GeoSets().Length()
                countObj  := geoCodes.GeoSets().Count()
                assert.True(
                    t,
                    countObj == lengthObj && lengthObj == geoSetsTotalCount,
                    "The number of the elements in the object must be " + fmt.Sprint(geoSetsTotalCount),
                )
            })
        })

        t.Run("CheckTheGeoSetsAsListOfElements:`.Get()`", func(t *testing.T) {
            t.Run("CheckTheGeoSetsAsSliceListOfElements", func(t *testing.T) {
                geoSets := geoCodes.GeoSets().Get()
                //** Let's work on the first element **//
                geoSetTypeAssertion := geoSets.Data.([]map[string]interface{})[0]
                assert.Equal(
                    t,
                    geoSetTypeAssertion["Name"],
                    geoSets.Pick("0.Name"),
                    "Wrong match for `Name`",
                )
                assert.Equal(
                    t,
                    geoSetTypeAssertion["InternalCode"],
                    geoSets.Pick("0.InternalCode"),
                    "Wrong match for `InternalCode`",
                )
                assert.Equal(
                    t,
                    geoSetTypeAssertion["UnM49"],
                    geoSets.Pick("0.UnM49"),
                    "Wrong match for `UnM49`",
                )
                assert.Equal(
                    t,
                    geoSetTypeAssertion["Tags"],
                    geoSets.Pick("0.Tags"),
                    "Wrong match for `Tags`",
                )
                assert.Equal(
                    t,
                    geoSetTypeAssertion["Tags"].([]interface{})[0],
                    geoSets.Pick("0.Tags.0"),
                    "Wrong match for `Tags.0`",
                )
                assert.Equal(
                    t,
                    geoSetTypeAssertion["CountryCodes"],
                    geoSets.Pick("0.CountryCodes"),
                    "Wrong match for `CountryCodes`",
                )
                assert.Equal(
                    t,
                    geoSetTypeAssertion["CountryCodes"].([]interface{})[0],
                    geoSets.Pick("0.CountryCodes.0"),
                    "Wrong match for `CountryCodes.0`",
                )
            })

            t.Run("CheckTheGeoSetsAsMapListOfElementsUsing:`.WithIndex()`", func(t *testing.T) {
                geoSets := geoCodes.GeoSets().WithIndex(geoSetsPrimaryKey).Get()
                //** Let's work on the `GEOG-EU` element **//
                geoSetTypeAssertion := geoSets.Data.(map[string]map[string]interface{})["GEOG-EU"]
                assert.Equal(
                    t,
                    geoSetTypeAssertion["Name"],
                    geoSets.Pick("GEOG-EU.Name"),
                    "Wrong match for `Name`",
                )
                assert.Equal(
                    t,
                    geoSetTypeAssertion["InternalCode"],
                    geoSets.Pick("GEOG-EU.InternalCode"),
                    "Wrong match for `InternalCode`",
                )
                assert.Equal(
                    t,
                    geoSetTypeAssertion["UnM49"],
                    geoSets.Pick("GEOG-EU.UnM49"),
                    "Wrong match for `UnM49`",
                )
                assert.Equal(
                    t,
                    geoSetTypeAssertion["Tags"],
                    geoSets.Pick("GEOG-EU.Tags"),
                    "Wrong match for `Tags`",
                )
                assert.Equal(
                    t,
                    geoSetTypeAssertion["Tags"].([]interface{})[0],
                    geoSets.Pick("GEOG-EU.Tags.0"),
                    "Wrong match for `Tags.0`",
                )
                assert.Equal(
                    t,
                    geoSetTypeAssertion["CountryCodes"],
                    geoSets.Pick("GEOG-EU.CountryCodes"),
                    "Wrong match for `CountryCodes`",
                )
                assert.Equal(
                    t,
                    geoSetTypeAssertion["CountryCodes"].([]interface{})[0],
                    geoSets.Pick("GEOG-EU.CountryCodes.0"),
                    "Wrong match for `CountryCodes.0`",
                )
            })
        })

        t.Run("CheckTheGeoSetAsSingleElement:`.First()`", func(t *testing.T) {
            geoSet := geoCodes.GeoSets().First()
            geoSetTypeAssertion := geoSet.Data.(map[string]interface{})
            assert.Equal(
                t,
                geoSetTypeAssertion["Name"],
                geoSet.Pick("Name"),
                "Wrong match for `Name`",
            )
            assert.Equal(
                t,
                geoSetTypeAssertion["InternalCode"],
                geoSet.Pick("InternalCode"),
                "Wrong match for `InternalCode`",
            )
            assert.Equal(
                t,
                geoSetTypeAssertion["UnM49"],
                geoSet.Pick("UnM49"),
                "Wrong match for `UnM49`",
            )
            assert.Equal(
                t,
                geoSetTypeAssertion["Tags"],
                geoSet.Pick("Tags"),
                "Wrong match for `Tags`",
            )
            assert.Equal(
                t,
                geoSetTypeAssertion["Tags"].([]interface{})[0],
                geoSet.Pick("Tags.0"),
                "Wrong match for `Tags.0`",
            )
            assert.Equal(
                t,
                geoSetTypeAssertion["CountryCodes"],
                geoSet.Pick("CountryCodes"),
                "Wrong match for `CountryCodes`",
            )
            assert.Equal(
                t,
                geoSetTypeAssertion["CountryCodes"].([]interface{})[0],
                geoSet.Pick("CountryCodes.0"),
                "Wrong match for `CountryCodes.0`",
            )

            t.Run("TestThatTheUseOf`.WithIndex()`HasNoInfluenceOn`.First()`", func(t *testing.T) {
                _, okWithIndex := geoCodes.GeoSets().
                    WithIndex(geoSetsPrimaryKey).
                    First().
                    Data.(map[string]interface{})
                assert.True(
                    t,
                    okWithIndex,
                    "Wrong Type",
                )
                _, okWithoutIndex := geoCodes.GeoSets().First().Data.(map[string]interface{})
                assert.True(
                    t,
                    okWithoutIndex,
                    "Wrong Type",
                )
                assert.Equal(
                    t,
                    geoCodes.GeoSets().First().Pick(geoSetsPrimaryKey),
                    geoCodes.GeoSets().WithIndex(geoSetsPrimaryKey).First().Pick(geoSetsPrimaryKey),
                    "WithIndex().First() is different from First()",
                )
            })
        })

        t.Run("TestThe`.Pick()`Features", func(t *testing.T) {
            t.Run("TestThe`.Pick()`Aliases", func(t *testing.T) {
                geoSet := geoCodes.GeoSets().First()
                pick := geoSet.Pick(geoSetsPrimaryKey)
                val := geoSet.Val(geoSetsPrimaryKey)
                value := geoSet.Value(geoSetsPrimaryKey)
                lookup := geoSet.Lookup(geoSetsPrimaryKey)
                assert.True(
                    t,
                    pick == val && val == value && value == lookup && lookup == geoSetsGlobalObject["firstElement"],
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
                geoCodes.GeoSets().First().Pick("NotExistentPropertyName")
            })
        })

        t.Run("TestTheStringEndpoints", func(t *testing.T) {
            t.Run("TestThe`.ToJson()`Endpoint", func(t *testing.T) {
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
            t.Run("TestThe`.ToYaml()`Endpoint", func(t *testing.T) {
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
            t.Run("TestThe`.ToXml()`Endpoint", func(t *testing.T) {
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
            t.Run("TestTheExistenceForTheXsdRelatedToTheList(`.GetXsd()`)Endpoint", func(t *testing.T) {
                xsd := geoCodes.GeoSets().GetXsd()
                assert.NotEmpty(t, xsd, "The content of the XSD is empty")
                assert.Contains(
                    t,
                    xsd,
                    "<xs:schema", "The XSD does not contain the <xs:schema> tag, so it may not be valid",
                )
            })
            t.Run("TestTheExistenceForTheXsdRelatedToTheSingleObject(`.GetXsdSingle()`)Endpoint", func(t *testing.T) {
                xsd := geoCodes.GeoSets().GetXsdSingle()
                assert.NotEmpty(t, xsd, "The content of the XSD is empty")
                assert.Contains(
                    t,
                    xsd,
                    "<xs:schema", "The XSD does not contain the <xs:schema> tag, so it may not be valid",
                )
            })

            t.Run("TestThe`.ToFlatten()`Endpoint", func(t *testing.T) {
                list := geoCodes.GeoSets().Get()
                listFlatten := list.ToFlatten(".")
                for i := 0; i < 5; i++ {
                    key := rand.Intn(geoSetsTotalCount)
                    assert.Equal(
                        t,
                        list.Pick(fmt.Sprintf("%d.InternalCode", key)),
                        listFlatten[fmt.Sprintf("%d.InternalCode", key)],
                        fmt.Sprintf(
                            "The flatten structure does not work (issue on `InternalCode` for %v)",
                            list.Pick(fmt.Sprintf("%d.InternalCode", key)),
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
                }
                t.Run("TestTheBehaviorOf`.ToFlatten()`WithWrongProperty", func(t *testing.T) {
                    assert.Nil(
                        t,
                        geoCodes.GeoSets().First().ToFlatten(".")["NotExistentPropertyName"],
                        "The value must be `nil`",
                    )
                })
            })
        })

        t.Run("TestThePackageLanguages", func(t *testing.T) {
            geoCodes.UseLanguage("en")
            assert.Equal(
                t,
                geoCodes.GeoSets().WithIndex(geoSetsPrimaryKey).Get().Pick("ORGS-WTO.Name"),
                "World Trade Organization (WTO)",
                "The chosen language does not seem to work",
            )
            geoCodes.UseLanguage("it")
            assert.Equal(
                t,
                geoCodes.GeoSets().WithIndex(geoSetsPrimaryKey).Get().Pick("ORGS-WTO.Name"),
                "Organizzazione Mondiale del Commercio (OMC)",
                "The chosen language does not seem to work",
            )
        })
    })
}


// func TestElibeGeoSets(t *testing.T) {
//
// //     geoSet0 := geoCodes.GeoSets().First().ToJson()
// //     geoSet0 := geoCodes.GeoSets().Get().ToJson()
// //     geoSet0 := geoCodes.GeoSets().WithIndex(geoSetsPrimaryKey).Get().ToJson()
//
// //     geoSet0 := geoCodes.GeoSets().First().ToYaml()
// //     geoSet0 := geoCodes.GeoSets().Get().ToYaml()
// //     geoSet0 := geoCodes.GeoSets().WithIndex(geoSetsPrimaryKey).Get().ToYaml()
//
//     geoSet0 := geoCodes.GeoSets().First().ToXml()
// //     geoSet0 := geoCodes.GeoSets().Get().ToXml()
// //     geoSet0 := geoCodes.GeoSets().WithIndex(geoSetsPrimaryKey).Get().ToXml()
//
//
//
// //     geoCodes.UseLanguage("it")
//     fmt.Printf("%v\n", geoSet0)
// }
