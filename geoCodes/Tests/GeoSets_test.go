package geoCodesTest

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/alibe1971/go-geoCodes/geoCodes"
    "github.com/alibe1971/go-geoCodes/geoCodes/Tests/TestLib"
    "fmt"
    "strings"
    "math/rand"
)

const geoSetsTotalCount int = 62
var geoSetsPrimaryKey string = geoCodes.GeoSets().GetPrimaryKey()
var geoSetsIndexes []string = geoCodes.GeoSets().GetIndexes()
var geoSetsFields []string = geoCodes.GeoSets().GetFields()
var geoSetsExpectedOrderBy = map[string]map[string]string{
    "InternalCode": {
        "ASC":  "CONV-G20",
        "DESC": "ZONE-EZ",
    },
    "Name": {
        "ASC":  "Africa",
        "DESC": "World Trade Organization (WTO)",
    },
}
var geoSetsExpectedLimit = []string{
    "GEOG-AS-SO",
    "GEOG-AS-WE",
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
                    pick == val && val == value && value == lookup &&
                        lookup == geoSetsExpectedOrderBy["InternalCode"]["ASC"],
                    "Wrong Type",
                )
            })
            t.Run("TestTheBehaviorOf`.Pick()`WithWrongProperty", func(t *testing.T) {
                assert.Panics(
                    t,
                    func() { _ = geoCodes.GeoSets().First().Pick("NotExistentPropertyName") },
                    "expected panic on wrong index",
                )
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
            geoCodes.UseDefaultLanguage()
            assert.Equal(
                t,
                geoCodes.GeoSets().WithIndex(geoSetsPrimaryKey).Get().Pick("ORGS-WTO.Name"),
                "World Trade Organization (WTO)",
                "The chosen language does not seem to work",
            )
        })

        t.Run("TestTheIndexSetters", func(t *testing.T) {
            ObjSetters := geoCodes.GeoSets()
            const CheckProperty string = "European Union (EU)"
            t.Run("TestThe`.WithIndex()`Setter", func(t *testing.T) {
                t.Run("TestTheOverrideBehavior", func(t *testing.T) {
                    assert.Equal(
                        t,
                        ObjSetters.WithIndex("Name").WithIndex("InternalCode").Get().Pick("ORGS-EU.Name"),
                        CheckProperty,
                        "The WithIndex override does not seem to work",
                    )
                })
                t.Run("TestTheAllIndexes", func(t *testing.T) {
                    assert.Equal(
                        t,
                        ObjSetters.WithIndex("InternalCode").Get().Pick("ORGS-EU.Name"),
                        CheckProperty,
                        "The index `InternalCode` does not seem to work",
                    )
                    assert.Equal(
                        t,
                        ObjSetters.WithIndex("Name").Get().Pick("European Union (EU).Name"),
                        CheckProperty,
                        "The index `Name` does not seem to work",
                    )
                })

                t.Run("TestTheWrongIndex", func(t *testing.T) {
                    assert.Panics(
                        t,
                        func() { _ = ObjSetters.WithIndex("UnM49") },
                        "expected panic on wrong index",
                    )
                })
            })

            t.Run("TestThe`.OrderBy()`Setter", func(t *testing.T) {
                t.Run("TestTheOverrideBehavior", func(t *testing.T) {
                    assert.Equal(
                        t,
                        ObjSetters.OrderBy("Name", "").OrderBy("InternalCode", "").OrderBy(geoSetsPrimaryKey, "").
                            First().Pick(geoSetsPrimaryKey),
                        geoSetsExpectedOrderBy[geoSetsPrimaryKey]["ASC"],
                        "The OrderBy override does not seem to work",
                    )
                    assert.Equal(
                        t,
                        ObjSetters.OrderBy(geoSetsPrimaryKey, "asc").OrderBy(geoSetsPrimaryKey, "DESC").
                            First().Pick(geoSetsPrimaryKey),
                        geoSetsExpectedOrderBy[geoSetsPrimaryKey]["DESC"],
                        "The OrderBy override does not seem to work",
                    )
                })
                t.Run("TestTheAllIndexesAndDirections", func(t *testing.T) {
                    for prop := range geoSetsExpectedOrderBy {
                        assert.Equal(
                            t,
                            ObjSetters.OrderBy(prop, "asc").First().Pick(prop),
                            geoSetsExpectedOrderBy[prop]["ASC"],
                            "The OrderBy `" + prop + "` (asc) does not seem to work",
                        )
                        assert.Equal(
                            t,
                            ObjSetters.OrderBy(prop, "desc").First().Pick(prop),
                            geoSetsExpectedOrderBy[prop]["DESC"],
                            "The OrderBy `" + prop + "` (desc) does not seem to work",
                        )
                    }
                })
                t.Run("TestTheWrongIndex", func(t *testing.T) {
                    assert.Panics(
                        t,
                        func() { _ = ObjSetters.OrderBy("UnM49", "ASC") },
                        "expected panic on wrong index",
                    )
                })
                t.Run("TestTheWrongDirection", func(t *testing.T) {
                    assert.Panics(
                        t,
                        func() { _ = ObjSetters.OrderBy("InternalCode", "wrong") },
                        "expected panic on wrong direction",
                    )
                })
            })
        })
        t.Run("TestTheSelectSetter", func(t *testing.T) {
            t.Run("TestTheSingleSelect", func(t *testing.T) {
                for _, sel := range geoSetsFields {
                    got := geoCodes.GeoSets().Select(sel).First()

                    selParts := strings.Split(sel, ".")
                    selIsParent := len(selParts) == 1
                    selIsChild  := len(selParts) > 1
                    parentOfSel := selParts[0]

                    for _, f := range geoSetsFields {
                        fParts := strings.Split(f, ".")
                        isSame     := f == sel
                        isChildOfSel := selIsParent  && len(fParts) > 1 && fParts[0] == sel
                        isParentOfSel:= selIsChild   && f == parentOfSel

                        switch {
                        // same property
                        case isSame:
                            assert.NotPanics(
                                t,
                                func() { _ = got.Pick(f) },
                                "Pick(%q) shouldn't panic", f,
                            )

                        // Select father → Pick son
                        case isChildOfSel:
                            assert.NotPanics(
                                t,
                                func() { _ = got.Pick(f) },
                                "Pick(%q) should not panic (father→son)", f,
                            )

                        // Select son → Pick father
                        case isParentOfSel:
                            assert.NotPanics(
                                t,
                                func() { _ = got.Pick(f) },
                                "Pick(%q) should not panic (son→father)", f,
                            )

                        // default
                        default:
                            assert.Panics(
                                t,
                                func() { _ = got.Pick(f) },
                                "Pick(%q) should panic", f,
                            )
                        }
                    }
                }
            })
            t.Run("TestTheAggregateSelect", func(t *testing.T) {
                for i := range geoSetsFields {
                    // I recreate the object from scratch for each i
                    agg := geoCodes.GeoSets()
                    var selected []string

                    // I call .Select() on all fields from 0 to i
                    for j := 0; j <= i; j++ {
                        sel := geoSetsFields[j]
                        agg = agg.Select(sel)
                        selected = append(selected, sel)
                    }

                    got := agg.First()

                    // Test the fields
                    for _, f := range geoSetsFields {
                        if TestLib.ContainsOrRelated(selected, f) {
                            assert.NotPanics(
                                t,
                                func() { _ = got.Pick(f) },
                                "Pick(%q) should not panic because %v is in selected %v",
                                f, f, selected,
                            )
                        } else {
                            assert.Panics(
                                t,
                                func() { _ = got.Pick(f) },
                                "Pick(%q) should panic – selected=%v",
                                f, selected,
                            )
                        }
                    }
                }
            })
        })
        t.Run("TestTheLimitAndOffsetSetter", func(t *testing.T) {
            t.Run("TestOffsetLimitCount", func(t *testing.T) {
                tests := []int{
                    geoSetsTotalCount - 21,
                    27,
                    5,
                    32,
                    0,
                }

                for _, want := range tests {
                    g := geoCodes.GeoSets().
                        Offset(21).
                        Limit(want)

                    got := g.Count()
                    assert.Equalf(
                        t,
                        want,
                        got,
                        "offset=52, limit=%d: expected count == %d, got %d",
                        want, want, got,
                    )
                }
            })
            t.Run("TestLimitAndOffsetPropertiesGet", func(t *testing.T) {
                t.Run("InvalidOffset", func(t *testing.T) {
                    offset := -5
                    limit := 20
                    assert.Panics(
                        t,
                        func() { _ = geoCodes.GeoSets().Offset(offset).Limit(limit) },
                        "expected panic on Offset(%d)",
                        offset,
                    )
                })
                t.Run("InvalidLimit", func(t *testing.T) {
                    offset := 20
                    limit := -5
                    assert.Panics(
                        t,
                        func() { _ = geoCodes.GeoSets().Offset(offset).Limit(limit) },
                        "expected panic on Limit(%d)",
                        limit,
                    )
                })
                t.Run("ValidOffsetLimit", func(t *testing.T) {
                    c := geoCodes.GeoSets().Offset(22).Limit(2)
                    assert.Equal(t, 2, c.Count())
                    got := c.Get()
                    assert.Equal(t, geoSetsExpectedLimit[0], got.Pick("0.InternalCode"))
                    assert.Equal(t, geoSetsExpectedLimit[1], got.Pick("1.InternalCode"))
                })
                t.Run("AliasSkipTake", func(t *testing.T) {
                    c := geoCodes.GeoSets().Skip(22).Take(2)
                    assert.Equal(t, 2, c.Count())
                    got := c.Get()
                    assert.Equal(t, geoSetsExpectedLimit[0], got.Pick("0.InternalCode"))
                    assert.Equal(t, geoSetsExpectedLimit[1], got.Pick("1.InternalCode"))
                })
            })
        })
    })
}


