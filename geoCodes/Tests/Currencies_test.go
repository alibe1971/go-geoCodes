package geoCodesTest

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/alibe1971/go-geoCodes/geoCodes"
    "github.com/alibe1971/go-geoCodes/geoCodes/Tests/TestLib"
    "fmt"
    "math/rand"
)

const currenciesTotalCount int = 180
var currenciesPrimaryKey string = geoCodes.Currencies().GetPrimaryKey()
var currenciesIndexes []string = geoCodes.Currencies().GetIndexes()
var currenciesFields []string = geoCodes.Currencies().GetFields()
var currenciesExpectedOrderBy = map[string]map[string]string{
    "IsoAlpha": {
        "ASC":  "AED",
        "DESC": "ZWL",
    },
    "IsoNumber": {
        "ASC":  "008",
        "DESC": "999",
    },
    "Name": {
        "ASC":  "ADB Unit of Account",
        "DESC": "Zloty",
    },
}


func TestCurrencies(t *testing.T) {
    t.Run("TestTheCurrenciesFunctionalities", func(t *testing.T) {
        t.Run("CheckTheCurrenciesObjectIsCorrectlyInstantiated", func(t *testing.T) {
            assert.NotNil(t, geoCodes.Currencies(), "The object `Currencies` cannot be `nil`")
            t.Run("CheckTheCurrenciesObjectHasTheCorrectNumberOfElements", func(t *testing.T) {
                // Check with the alias commands
                lengthObj := geoCodes.Currencies().Length()
                countObj  := geoCodes.Currencies().Count()
                assert.True(
                    t,
                    countObj == lengthObj && lengthObj == currenciesTotalCount,
                    "The number of the elements in the object must be " + fmt.Sprint(currenciesTotalCount),
                )
            })
        })
        t.Run("CheckTheCurrenciesAsListOfElements:`.Get()`", func(t *testing.T) {
            t.Run("CheckTheCurrenciesAsSliceListOfElements", func(t *testing.T) {
                currencies := geoCodes.Currencies().Get()
                //** Let's work on the first element **//
                currencyTypeAssertion := currencies.Data.([]map[string]interface{})[0]
                assert.Equal(
                    t,
                    currencyTypeAssertion["Name"],
                    currencies.Pick("0.Name"),
                    "Wrong match for `Name`",
                )
                assert.Equal(
                    t,
                    currencyTypeAssertion["IsoAlpha"],
                    currencies.Pick("0.IsoAlpha"),
                    "Wrong match for `IsoAlpha`",
                )
                assert.Equal(
                    t,
                    currencyTypeAssertion["IsoNumber"],
                    currencies.Pick("0.IsoNumber"),
                    "Wrong match for `IsoNumber`",
                )
                assert.Equal(
                    t,
                    currencyTypeAssertion["Symbol"],
                    currencies.Pick("0.Symbol"),
                    "Wrong match for `Symbol`",
                )
                assert.Equal(
                    t,
                    currencyTypeAssertion["Decimal"],
                    currencies.Pick("0.Decimal"),
                    "Wrong match for `Decimal`",
                )
            })

            t.Run("CheckTheCurrenciesAsMapListOfElementsUsing:`.WithIndex()`", func(t *testing.T) {
                currencies := geoCodes.Currencies().WithIndex(currenciesPrimaryKey).Get()
                //** Let's work on the `EUR` element **//
                currencyTypeAssertion := currencies.Data.(map[string]map[string]interface{})["EUR"]
                assert.Equal(
                    t,
                    currencyTypeAssertion["Name"],
                    currencies.Pick("EUR.Name"),
                    "Wrong match for `Name`",
                )
                assert.Equal(
                    t,
                    currencyTypeAssertion["IsoAlpha"],
                    currencies.Pick("EUR.IsoAlpha"),
                    "Wrong match for `IsoAlpha`",
                )
                assert.Equal(
                    t,
                    currencyTypeAssertion["IsoNumber"],
                    currencies.Pick("EUR.IsoNumber"),
                    "Wrong match for `IsoNumber`",
                )
                assert.Equal(
                    t,
                    currencyTypeAssertion["Symbol"],
                    currencies.Pick("EUR.Symbol"),
                    "Wrong match for `Symbol`",
                )
                assert.Equal(
                    t,
                    currencyTypeAssertion["Decimal"],
                    currencies.Pick("EUR.Decimal"),
                    "Wrong match for `Decimal`",
                )
            })
        })

        t.Run("CheckTheCurrencyAsSingleElement:`.First()`", func(t *testing.T) {
            currency := geoCodes.Currencies().First()
            currencyTypeAssertion := currency.Data.(map[string]interface{})
            assert.Equal(
                t,
                currencyTypeAssertion["Name"],
                currency.Pick("Name"),
                "Wrong match for `Name`",
            )
            assert.Equal(
                t,
                currencyTypeAssertion["IsoAlpha"],
                currency.Pick("IsoAlpha"),
                "Wrong match for `IsoAlpha`",
            )
            assert.Equal(
                t,
                currencyTypeAssertion["IsoNumber"],
                currency.Pick("IsoNumber"),
                "Wrong match for `IsoNumber`",
            )
            assert.Equal(
                t,
                currencyTypeAssertion["Symbol"],
                currency.Pick("Symbol"),
                "Wrong match for `Symbol`",
            )
            assert.Equal(
                t,
                currencyTypeAssertion["Decimal"],
                currency.Pick("Decimal"),
                "Wrong match for `Decimal`",
            )

            t.Run("TestThatTheUseOf`.WithIndex()`HasNoInfluenceOn`.First()`", func(t *testing.T) {
                _, okWithIndex := geoCodes.Currencies().
                    WithIndex(currenciesPrimaryKey).
                    First().
                    Data.(map[string]interface{})
                assert.True(
                    t,
                    okWithIndex,
                    "Wrong Type",
                )
                _, okWithoutIndex := geoCodes.Currencies().First().Data.(map[string]interface{})
                assert.True(
                    t,
                    okWithoutIndex,
                    "Wrong Type",
                )
                assert.Equal(
                    t,
                    geoCodes.Currencies().First().Pick(currenciesPrimaryKey),
                    geoCodes.Currencies().WithIndex(currenciesPrimaryKey).First().Pick(currenciesPrimaryKey),
                    "WithIndex().First() is different from First()",
                )
            })
        })

        t.Run("TestThe`.Pick()`Features", func(t *testing.T) {
            t.Run("TestThe`.Pick()`Aliases", func(t *testing.T) {
                currency := geoCodes.Currencies().First()
                pick := currency.Pick(currenciesPrimaryKey)
                val := currency.Val(currenciesPrimaryKey)
                value := currency.Value(currenciesPrimaryKey)
                lookup := currency.Lookup(currenciesPrimaryKey)
                assert.True(
                    t,
                    pick == val && val == value && value == lookup &&
                        lookup == currenciesExpectedOrderBy["IsoAlpha"]["ASC"],
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
                geoCodes.Currencies().First().Pick("NotExistentPropertyName")
            })
        })

        t.Run("TestTheStringEndpoints", func(t *testing.T) {
            t.Run("TestThe`.ToJson()`Endpoint", func(t *testing.T) {
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
            t.Run("TestThe`.ToYaml()`Endpoint", func(t *testing.T) {
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
            t.Run("TestThe`.ToXml()`Endpoint", func(t *testing.T) {
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
            t.Run("TestTheExistenceForTheXsdRelatedToTheList(`.GetXsd()`)Endpoint", func(t *testing.T) {
                xsd := geoCodes.Currencies().GetXsd()
                assert.NotEmpty(t, xsd, "The content of the XSD is empty")
                assert.Contains(
                    t,
                    xsd,
                    "<xs:schema", "The XSD does not contain the <xs:schema> tag, so it may not be valid",
                )
            })
            t.Run("TestTheExistenceForTheXsdRelatedToTheSingleObject(`.GetXsdSingle()`)Endpoint", func(t *testing.T) {
                xsd := geoCodes.Currencies().GetXsdSingle()
                assert.NotEmpty(t, xsd, "The content of the XSD is empty")
                assert.Contains(
                    t,
                    xsd,
                    "<xs:schema", "The XSD does not contain the <xs:schema> tag, so it may not be valid",
                )
            })

            t.Run("TestThe`.ToFlatten()`Endpoint", func(t *testing.T) {
                list := geoCodes.Currencies().Get()
                listFlatten := list.ToFlatten(".")
                for i := 0; i < 5; i++ {
                    key := rand.Intn(currenciesTotalCount)
                    assert.Equal(
                        t,
                        list.Pick(fmt.Sprintf("%d.IsoAlpha", key)),
                        listFlatten[fmt.Sprintf("%d.IsoAlpha", key)],
                        fmt.Sprintf(
                            "The flatten structure does not work (issue on `IsoAlpha` for %v)",
                            list.Pick(fmt.Sprintf("%d.IsoAlpha", key)),
                        ),
                    )
                    assert.Equal(
                        t,
                        list.Pick(fmt.Sprintf("%d.IsoNumber", key)),
                        listFlatten[fmt.Sprintf("%d.IsoNumber", key)],
                        fmt.Sprintf(
                            "The flatten structure does not work (issue on `IsoNumber` for %v)",
                            list.Pick(fmt.Sprintf("%d.IsoNumber", key)),
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
                        list.Pick(fmt.Sprintf("%d.Symbol", key)),
                        listFlatten[fmt.Sprintf("%d.Symbol", key)],
                        fmt.Sprintf(
                            "The flatten structure does not work (issue on `Symbol` for %v)",
                            list.Pick(fmt.Sprintf("%d.Symbol", key)),
                        ),
                    )
                    assert.Equal(
                        t,
                        list.Pick(fmt.Sprintf("%d.Decimal", key)),
                        listFlatten[fmt.Sprintf("%d.Decimal", key)],
                        fmt.Sprintf(
                            "The flatten structure does not work (issue on `Decimal` for %v)",
                            list.Pick(fmt.Sprintf("%d.Decimal", key)),
                        ),
                    )
                }

                t.Run("TestTheBehaviorOf`.ToFlatten()`WithWrongProperty", func(t *testing.T) {
                    assert.Nil(
                        t,
                        geoCodes.Currencies().First().ToFlatten(".")["NotExistentPropertyName"],
                        "The value must be `nil`",
                    )
                })
            })
        })

        t.Run("TestThePackageLanguages", func(t *testing.T) {
            geoCodes.UseLanguage("en")
            assert.Equal(
                t,
                geoCodes.Currencies().WithIndex(currenciesPrimaryKey).Get().Pick("AED.Name"),
                "UAE Dirham",
                "The chosen language does not seem to work",
            )
            geoCodes.UseLanguage("it")
            assert.Equal(
                t,
                geoCodes.Currencies().WithIndex(currenciesPrimaryKey).Get().Pick("AED.Name"),
                "Dirham degli Emirati Arabi Uniti",
                "The chosen language does not seem to work",
            )
            geoCodes.UseDefaultLanguage()
            assert.Equal(
                t,
                geoCodes.Currencies().WithIndex(currenciesPrimaryKey).Get().Pick("AED.Name"),
                "UAE Dirham",
                "The chosen language does not seem to work",
            )
        })

        t.Run("TestTheIndexSetters", func(t *testing.T) {
            ObjSetters := geoCodes.Currencies()
            const CheckProperty string = "Netherlands Antillean Guilder"
            t.Run("TestThe`.WithIndex()`Setter", func(t *testing.T) {
                t.Run("TestTheOverrideBehavior", func(t *testing.T) {
                    assert.Equal(
                        t,
                        ObjSetters.WithIndex("Name").WithIndex("IsoNumber").
                            WithIndex("IsoAlpha").Get().Pick("ANG.Name"),
                        CheckProperty,
                        "The WithIndex override does not seem to work",
                    )
                })
                t.Run("TestTheAllIndexes", func(t *testing.T) {
                    assert.Equal(
                        t,
                        ObjSetters.WithIndex("IsoAlpha").Get().Pick("ANG.Name"),
                        CheckProperty,
                        "The index `IsoAlpha` does not seem to work",
                    )
                    assert.Equal(
                        t,
                        ObjSetters.WithIndex("IsoNumber").Get().Pick("532.Name"),
                        CheckProperty,
                        "The index `IsoNumber` does not seem to work",
                    )
                    assert.Equal(
                        t,
                        ObjSetters.WithIndex("Name").Get().Pick("Netherlands Antillean Guilder.Name"),
                        CheckProperty,
                        "The index `Name` does not seem to work",
                    )
                })

                t.Run("TestTheWrongIndex", func(t *testing.T) {
                    defer func() {
                        if r := recover(); r != nil {
                          assert.Contains(t, r.(string), "not existent or not usable as index")
                          return
                        }
                        t.Error("Expected panic, but no panic occurred")
                    }()
                    ObjSetters.WithIndex("Symbol")
                })
            })

            t.Run("TestThe`.OrderBy()`Setter", func(t *testing.T) {
                t.Run("TestTheOverrideBehavior", func(t *testing.T) {
                    assert.Equal(
                        t,
                        ObjSetters.OrderBy("Name", "").OrderBy("IsoNumber", "").OrderBy(currenciesPrimaryKey, "").
                            First().Pick(currenciesPrimaryKey),
                        currenciesExpectedOrderBy[currenciesPrimaryKey]["ASC"],
                        "The OrderBy override does not seem to work",
                    )
                    assert.Equal(
                        t,
                        ObjSetters.OrderBy(currenciesPrimaryKey, "asc").OrderBy(currenciesPrimaryKey, "DESC").
                            First().Pick(currenciesPrimaryKey),
                        currenciesExpectedOrderBy[currenciesPrimaryKey]["DESC"],
                        "The OrderBy override does not seem to work",
                    )
                })
                t.Run("TestTheAllIndexesAndDirections", func(t *testing.T) {
                    for prop := range currenciesExpectedOrderBy {
                        assert.Equal(
                            t,
                            ObjSetters.OrderBy(prop, "asc").First().Pick(prop),
                            currenciesExpectedOrderBy[prop]["ASC"],
                            "The OrderBy `" + prop + "` (asc) does not seem to work",
                        )
                        assert.Equal(
                            t,
                            ObjSetters.OrderBy(prop, "desc").First().Pick(prop),
                            currenciesExpectedOrderBy[prop]["DESC"],
                            "The OrderBy `" + prop + "` (desc) does not seem to work",
                        )
                    }
                })
                t.Run("TestTheWrongIndex", func(t *testing.T) {
                    defer func() {
                        if r := recover(); r != nil {
                            assert.Contains(t, r.(string), "Attribute `orderBy`.`property` must be usable as index.")
                            return
                        }
                        t.Error("Expected panic, but no panic occurred")
                    }()
                    ObjSetters.OrderBy("Symbol", "")
                })
                t.Run("TestTheWrongDirection", func(t *testing.T) {
                    const errMsgDirection = "Attribute `orderBy`.`direction` must be `ASC` " +
                        "(default if empty string - ``) or `DESC` (case insensitive)"
                    defer func() {
                        if r := recover(); r != nil {
                            assert.Contains(t, r.(string), errMsgDirection)
                            return
                        }
                        t.Error("Expected panic, but no panic occurred")
                    }()
                    ObjSetters.OrderBy("IsoAlpha", "wrong")
                })
            })
        })

    })
}
