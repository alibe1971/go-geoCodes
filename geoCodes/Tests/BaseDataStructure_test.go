package geoCodesTest

import (
    "fmt"
    "strings"
    "testing"
    "github.com/stretchr/testify/assert"
    Structs "github.com/alibe1971/go-geoCodes/geoCodes/Structs"
    data "github.com/alibe1971/go-geoCodes/geoCodes/Data"
    transData "github.com/alibe1971/go-geoCodes/geoCodes/Data/Translations"
    "github.com/alibe1971/go-geoCodes/geoCodes/Tests/TestLib"
    "regexp"
)


func TestBaseDataStructure(t *testing.T) {
    var translations []string
    var cfg *Structs.Config
    var CC *Structs.Countries
        var alpha2Map map[string]struct{}
        var ccUnM49Map map[string]struct{}
    var GS *Structs.GeoSets
        var internalCodeMap map[string]struct{}
        var gsUnM49Map map[string]struct{}
    var CU *Structs.Currencies
        var isoAlphaMap map[string]struct{}

    t.Run("DataBaseInitialization", func(t *testing.T) {
        t.Run("ConfigDataInitialization", func(t *testing.T) {
            TestLib.LoadData("config", data.Config, &Structs.Config{})
            cfg = TestLib.GetData("config").(*Structs.Config)
            if cfg == nil {
                t.Fatal("Config data is not initialized")
            }
        })
        t.Run("CountriesDataInitialization", func(t *testing.T) {
            CC = new(Structs.Countries)
            TestLib.LoadData("countries", data.Countries, CC)
            loadedData := TestLib.GetData("countries")
            var ok bool
            CC, ok = loadedData.(*Structs.Countries)
            assert.True(t, ok, "Failed to convert data to *geoCodes.Countries")
            if CC == nil {
                t.Fatal("Countries data is not initialized")
            }
            alpha2Map = make(map[string]struct{})
            ccUnM49Map = make(map[string]struct{})
			for _, cc := range *CC {
				alpha2Map[cc.Alpha2] = struct{}{}
				ccUnM49Map[cc.UnM49] = struct{}{}
			}
        })
        t.Run("GeoSetsDataInitialization", func(t *testing.T) {
            GS = new(Structs.GeoSets)
            TestLib.LoadData("geoSets", data.GeoSets, GS)
            loadedData := TestLib.GetData("geoSets")
            var ok bool
            GS, ok = loadedData.(*Structs.GeoSets)
            assert.True(t, ok, "Failed to convert data to *geoCodes.GeoSets")
            if GS == nil {
                t.Fatal("GeoSets data is not initialized")
            }
            internalCodeMap = make(map[string]struct{})
            gsUnM49Map = make(map[string]struct{})
			for _, gs := range *GS {
			    internalCodeMap[gs.InternalCode] = struct{}{}
			    if gs.UnM49 != nil {
				    gsUnM49Map[*gs.UnM49] = struct{}{}
			    }
			}
        })
        t.Run("CurrenciesDataInitialization", func(t *testing.T) {
            CU = new(Structs.Currencies)
            TestLib.LoadData("currencies", data.Currencies, CU)
            loadedData := TestLib.GetData("currencies")
            var ok bool
            CU, ok = loadedData.(*Structs.Currencies)
            assert.True(t, ok, "Failed to convert data to *geoCodes.Currencies")
            if CU == nil {
                t.Fatal("Currencies data is not initialized")
            }
            isoAlphaMap = make(map[string]struct{})
            for _, cu := range *CU {
                isoAlphaMap[cu.IsoAlpha] = struct{}{}
            }
        })
    })


    t.Run("TestsOnConfigData", func(t *testing.T) {

        t.Run("CheckTheDataConfigStructure", func(t *testing.T) {
            translations = make([]string, 0, len(cfg.Settings.Languages.InPackage))
            for key, locale := range cfg.Settings.Languages.InPackage {
                assert.True(t, regexp.MustCompile(`^[a-z]{2}(_[A-Za-z]+)*(_[A-Z]{2})?$`).MatchString(locale),
                    fmt.Sprintf("Wrong format for locale: %s", locale))
                translations = append(translations, key)
            }
            assert.Equal(t, "en", cfg.Settings.Languages.Default, "Default should be 'en'")
        })
    })

    t.Run("TestsOnCountriesData", func(t *testing.T) {
        t.Run("CheckTheDataCountriesStructure", func(t *testing.T) {
            seenAlpha2 := []string{}
            seenAlpha3 := []string{}
            seenUnM49  := []string{}
            uniqueKeysOfficialName := make(map[string]bool)
            for _, cc := range *CC {
                t.Run("CheckTheDataStructureForCountry:" + cc.Alpha2, func(t *testing.T) {

                    /** alpha2 **/
                    t.Run("CheckTheAlpha2Property", func(t *testing.T) {
                        assert.True(t, regexp.MustCompile(`^[A-Z]{2}$`).MatchString(cc.Alpha2),
                            "Wrong format for alpha2")
                        assert.False(t, TestLib.IsInSlice(seenAlpha2, cc.Alpha2),
                            fmt.Sprintf("Duplicate alpha2 found: %s", cc.Alpha2))
                        seenAlpha2 = append(seenAlpha2, cc.Alpha2)
                    })

                    /** alpha3 **/
                    t.Run("CheckTheAlpha3Property", func(t *testing.T) {
                        assert.True(t, regexp.MustCompile(`^[A-Z]{3}$`).MatchString(cc.Alpha3),
                            "Wrong format for alpha3")
                        assert.False(t, TestLib.IsInSlice(seenAlpha3, cc.Alpha3),
                            fmt.Sprintf("Duplicate alpha3 found: %s", cc.Alpha3))
                        seenAlpha3 = append(seenAlpha3, cc.Alpha3)
                    })

                    /** unM49 **/
                    t.Run("CheckTheUnM49Property", func(t *testing.T) {
                        assert.True(t, regexp.MustCompile(`^[0-9]{3}$`).MatchString(cc.UnM49),
                            "Wrong format for unM49")
                        assert.False(t, TestLib.IsInSlice(seenUnM49, cc.UnM49),
                            fmt.Sprintf("Duplicate unM49 found in the Countries data: %s", cc.UnM49))
                        seenUnM49 = append(seenUnM49, cc.UnM49)
                        // Exception for Antartica (AQ - 010) that is also a continent
                        if cc.Alpha2 != "AQ" {
                            assert.False(t, TestLib.IsInMap(gsUnM49Map, cc.UnM49),
                                fmt.Sprintf("Duplicate unM49 found in the GeoSets data: %s", cc.UnM49))
                        }
                    })

                    /** flags **/
                    t.Run("CheckTheFlagProperty", func(t *testing.T) {
                        t.Run("CheckTheFlagEmojiProperty", func(t *testing.T) {
                            assert.NotEmpty(t, cc.Flags.Emoji, "flags.Emoji must not be empty")
                            assert.Regexp(
                                t,
                                "^[\U0001F1E6-\U0001F1FF]{2}$",
                                cc.Flags.Emoji,
                                "Flags.Svg must be a Regional Indicator Symbols string",
                            )
                        })
                        t.Run("CheckTheFlagSvgProperty", func(t *testing.T) {
                            assert.NotEmpty(t, cc.Flags.Svg, "flags.Svg must not be empty")
                            assert.True(t, TestLib.IsValidSVG(cc.Flags.Svg), "Flags.Svg must be a valid Svg")
                        })
                    })

                    /** dependency **/
                    t.Run("CheckTheDependencyProperty", func(t *testing.T) {
                        if cc.Dependency != nil {
                            dependencyStr := *cc.Dependency
                            assert.True(t, regexp.MustCompile(`^[A-Z]{2}$`).MatchString(dependencyStr),
                                "Wrong format for dependency")
                            assert.True(t, TestLib.IsInMap(alpha2Map, dependencyStr),
                                "The value of dependency must match an existing alpha2")
                        }
                    })

                    /** officialName **/
                    t.Run("CheckTheOfficialNameProperty", func(t *testing.T) {
                        for lang, name := range cc.OfficialName {
                            if name == "" {
                                assert.Fail(t, "Official name is an empty string", "Language: '%s'", lang)
                                continue
                            }
                            keyON := "officialName_" + lang + "_" + name
                            if _, exists := uniqueKeysOfficialName[keyON]; exists {
                                assert.Fail(t, "Duplicate officialName found", "Language: '%s', Name: '%s'", lang, name)
                            } else {
                                uniqueKeysOfficialName[keyON] = true
                            }
                        }
                    })

                    /** mottos **/
                    t.Run("CheckTheMottosProperty", func(t *testing.T) {
                         // Nothing to do
                    })

                    /** currencies **/
                    t.Run("CheckTheCurrencyProperty", func(t *testing.T) {
                        t.Run("CheckTheCurrencyLegalTenderProperty", func(t *testing.T) {
                            if len(cc.Currencies.LegalTenders) > 0 {
                                for _, currency := range cc.Currencies.LegalTenders {
                                    assert.True(t, regexp.MustCompile(`^[A-Z]{3}$`).MatchString(currency),
                                        "Wrong format for LegalTenders")
                                    assert.True(t, TestLib.IsInMap(isoAlphaMap, currency),
                                        fmt.Sprintf("The currency '%s' for LegalTenders isn't in the currency database", currency))
                                }
                            }
                        })
                        t.Run("CheckTheCurrencyWidelyAcceptedProperty", func(t *testing.T) {
                            if len(cc.Currencies.WidelyAccepted) > 0 {
                                for _, currency := range cc.Currencies.WidelyAccepted {
                                    assert.True(t, regexp.MustCompile(`^[A-Z]{3}$`).MatchString(currency),
                                        "Wrong format for WidelyAccepted")
                                    assert.True(t, TestLib.IsInMap(isoAlphaMap, currency),
                                        fmt.Sprintf("The currency '%s' for WidelyAccepted isn't in the currency database", currency))
                                    found := false
                                    for _, ltCurrency := range cc.Currencies.LegalTenders {
                                        if ltCurrency == currency {
                                            found = true
                                            break
                                        }
                                    }
                                    assert.False(t, found,
                                        fmt.Sprintf("The currency '%s' for WidelyAccepted is present also in LegalTenders", currency))
                                }
                            }
                        })
                    })

                    /** dialCodes **/
                    t.Run("CheckTheDialCodesProperty", func(t *testing.T) {
                        t.Run("CheckTheDialCodesMainValuesHaveTheRightFormat", func(t *testing.T) {
                            if len(cc.DialCodes.Main) > 0 {
                                for _, dial := range cc.DialCodes.Main {
                                    assert.True(t, regexp.MustCompile(`^\+\d+$`).MatchString(dial),
                                        "Wrong format for dialCodes.main")
                                }
                            }
                        })
                        t.Run("CheckTheDialCodesExceptionsValuesHaveTheRightFormat", func(t *testing.T) {
                            if len(cc.DialCodes.Exceptions) > 0 {
                                for _, dial := range cc.DialCodes.Exceptions {
                                    assert.True(t, regexp.MustCompile(`^\+\d+$`).MatchString(dial),
                                        "Wrong format for dialCodes.exceptions")
                                }
                            }
                        })
                    })

                    /** ccTld **/
                    t.Run("CheckTheCcTldProperty", func(t *testing.T) {
                        if cc.CcTld != nil {
                            CcTld := *cc.CcTld
                            assert.True(t, regexp.MustCompile(`^\.[a-z]{2}$`).MatchString(CcTld),
                                "Wrong format for CcTld")
                        }
                    })

                    /** timeZones **/
                    t.Run("CheckTheTimeZonesProperty", func(t *testing.T) {
                         assert.NotEmpty(t, cc.TimeZones, "The timeZones cannot be empty")
                    })

                    /** languages **/
                    t.Run("CheckTheLanguagesProperty", func(t *testing.T) {
                         // [TODO]
                    })

                    /** locales **/
                    t.Run("CheckTheLocalesProperty", func(t *testing.T) {
                        // Nothing to do
                    })

                })
            }
        })

        t.Run("TestsOnCountriesTranslationData", func(t *testing.T) {
            TranslationsCC := make(map[string]*Structs.TransCountries)
            t.Run("TranslationCountriesDataInitialization", func(t *testing.T) {
                for _, lang := range translations {
                    var tr = new(Structs.TransCountries)
                    TestLib.LoadData("trans_" + lang + "_countries", transData.Countries[lang], tr)
                    loadedData := TestLib.GetData("trans_" + lang + "_countries")
                    var ok bool
                    tr, ok = loadedData.(*Structs.TransCountries)
                    assert.True(t, ok, fmt.Sprintf("Failed to convert data to Translations %s Countries", lang))
                    if tr == nil {
                        t.Fatal(fmt.Sprintf("Translation %s data is not initialized", lang))
                    }
                    TranslationsCC[lang] = tr
                }
            })
            t.Run("TranslationCountriesCheckData", func(t *testing.T) {
                for lang, trans := range TranslationsCC {
                    t.Run(fmt.Sprintf("TestForTheLanguage:%s", lang), func(t *testing.T) {
                        var translationKeys []string
                        for cc, tr :=range *trans {
                            t.Run(fmt.Sprintf("TestForTheCountry:%s", cc), func(t *testing.T) {
                                translationKeys = append(translationKeys, cc)
                                t.Run("CheckTheCountryExists", func(t *testing.T) {
                                    assert.True(t, TestLib.IsInMap(alpha2Map, cc),
                                        "The country code in the translation must match an existing alpha2")
                                })

                                if lang == cfg.Settings.Languages.Default {
                                    trimmedName := strings.TrimSpace(tr.Name)
                                    t.Run("CheckTranslationNameExistsForDefaultLanguage", func(t *testing.T) {
                                        assert.True(t, trimmedName != "",
                                            "In default language the property name must exist and not be empty")
                                    })
                                    trimmedFullName := strings.TrimSpace(tr.FullName)
                                    t.Run("CheckTranslationNameExistsForDefaultLanguage", func(t *testing.T) {
                                        assert.True(t, trimmedFullName != "",
                                            "In default language the property FullName must exist and not be empty")
                                    })
                                }
                            })
                        }
                        if lang == cfg.Settings.Languages.Default {
                            t.Run("CheckAllTheCountriesArePresentInTheDefaultLanguage", func(t *testing.T) {
                                assert.True(t, len(alpha2Map) == len(translationKeys),
                                    "Not all the countries are present in the translation for default language")
                            })
                        }
                    })
                }
            })
        })
    })

    t.Run("TestsOnGeoSetsData", func(t *testing.T) {
        t.Run("CheckTheDataGeoSetsStructure", func(t *testing.T) {
            seenInternalCode := []string{}
            seenUnM49  := []string{}
            seenGEOG := make(map[string][]string)
            geogGr := []string{}
            geoLv := make([][]string, 2)
            for _, gs := range *GS {
                t.Run("CheckTheDataStructureForGeoSet:" + gs.InternalCode, func(t *testing.T) {
                    /** internalCode **/
                    t.Run("CheckTheInternalCodeProperty", func(t *testing.T) {
                        isValidFormat := regexp.MustCompile(`^[A-Z]+(-[A-Z0-9]+){1,4}$`).MatchString(gs.InternalCode)
                        assert.True(t, isValidFormat, "Wrong format for InternalCode")
                        isDuplicate := TestLib.IsInSlice(seenInternalCode, gs.InternalCode)
                        assert.False(t, isDuplicate, fmt.Sprintf("Duplicate InternalCode found: %s", gs.InternalCode))
                        seenInternalCode = append(seenInternalCode, gs.InternalCode)
                    })

                    /** unM49 **/
                    if gs.UnM49 != nil {
                        UnM49 := *gs.UnM49
                        t.Run("CheckTheUnM49Property", func(t *testing.T) {
                            assert.True(t, regexp.MustCompile(`^[0-9]{3}$`).MatchString(UnM49),
                                "Wrong format for unM49")
                            assert.False(t, TestLib.IsInSlice(seenUnM49, UnM49),
                                fmt.Sprintf("Duplicate unM49 found in the GeoSets data: %s", UnM49))
                            seenUnM49 = append(seenUnM49, UnM49)
                            // Exception for Antartica (AQ - 010) that is also a continent
                            if gs.InternalCode != "GEOG-AQ" {
                                assert.False(t, TestLib.IsInMap(ccUnM49Map, UnM49),
                                    fmt.Sprintf("Duplicate unM49 found in the Countries data: %s", UnM49))
                            }
                        })
                    }

                    /** tags **/
                    t.Run("CheckTheTagsProperty", func(t *testing.T) {
                         assert.NotEmpty(t, gs.Tags, "The tags cannot be empty")
                    })

                    /** countryCodes **/
                    t.Run("CheckTheCountryCodesProperty", func(t *testing.T) {
                         assert.NotEmpty(t, gs.CountryCodes, "The countryCodes cannot be empty")

                         for _,cc := range gs.CountryCodes {
                            t.Run("CheckTheCountryCodeExists", func(t *testing.T) {
                                assert.True(t, TestLib.IsInMap(alpha2Map, cc),
                                    "The country code must match an existing alpha2")
                            })
                            if strings.HasPrefix(gs.InternalCode, "GEOG-") {
                                gArr := strings.Split(gs.InternalCode, "-")
                                Lv := len(gArr) - 2
                                gArr = gArr[:len(gArr)-1]
                                parent := strings.Join(gArr, "-")

                                if _, exists := seenGEOG[parent]; !exists {
                                    seenGEOG[parent] = []string{}
                                }
                                if Lv != 0 {
                                    assert.Contains(
                                        t,
                                        seenGEOG[parent],
                                        cc,
                                        fmt.Sprintf("Inside the country set in the geographic geoSets data, the " +
                                            "value '%s' in '%s' has no correspondence in the parent group '%s'",
                                            cc, gs.InternalCode, parent),
                                    )
                                } else {
                                    geogGr = append(geogGr, cc)
                                }
                                if Lv < 2 {
                                    assert.NotContains(
                                        t,
                                        geoLv[Lv],
                                        cc,
                                        fmt.Sprintf("Inside the country set in the geographic geoSets data, the " +
                                            "value '%s' in '%s' is a duplicated key, because already present in this" +
                                            "or another '%d' region", cc, gs.InternalCode, Lv),
                                    )
                                    seenGEOG[gs.InternalCode] = append(seenGEOG[gs.InternalCode], cc)
                                    geoLv[Lv] = append(geoLv[Lv], cc)
                                }
                            }
                         }
                    })
                })
            }
            t.Run("TestThatTheGeographicGroupsHaveAllTheCountries", func(t *testing.T) {
                str := fmt.Sprintf("%d", len(geogGr))
                assert.True(t, len(alpha2Map) == len(geogGr),
                    "The Geographic Groups haven't inside all the countries " + str)
            })
        })

        t.Run("TestsOnGeoSetsTranslationData", func(t *testing.T) {
            TranslationsGS := make(map[string]*Structs.TransGeneric)
            t.Run("TranslationGeoSetsDataInitialization", func(t *testing.T) {
                for _, lang := range translations {
                    var tr = new(Structs.TransGeneric)
                    TestLib.LoadData("trans_" + lang + "_geosets", transData.GeoSets[lang], tr)
                    loadedData := TestLib.GetData("trans_" + lang + "_geosets")
                    var ok bool
                    tr, ok = loadedData.(*Structs.TransGeneric)
                    assert.True(t, ok, fmt.Sprintf("Failed to convert data to Translations %s GeoSets", lang))
                    if tr == nil {
                        t.Fatal(fmt.Sprintf("Translation %s data is not initialized", lang))
                    }
                    TranslationsGS[lang] = tr
                }
            })
            t.Run("TranslationGeoSetsCheckData", func(t *testing.T) {
                for lang, trans := range TranslationsGS {
                    t.Run(fmt.Sprintf("TestForTheLanguage:%s", lang), func(t *testing.T) {
                        var translationKeys []string
                        for gs, tr :=range *trans {
                            t.Run(fmt.Sprintf("TestForTheGeoSet:%s", gs), func(t *testing.T) {
                                translationKeys = append(translationKeys, gs)
                                t.Run("CheckTheGeoSetExists", func(t *testing.T) {
                                    assert.True(t, TestLib.IsInMap(internalCodeMap, gs),
                                        "The country code in the translation must match an existing internalCode")
                                })

                                if lang == cfg.Settings.Languages.Default {
                                    trimmedName := strings.TrimSpace(tr.Name)
                                    t.Run("CheckTranslationNameExistsForDefaultLanguage", func(t *testing.T) {
                                        assert.True(t, trimmedName != "",
                                            "In default language the property name must exist and not be empty")
                                    })
                                }
                            })
                        }
                        if lang == cfg.Settings.Languages.Default {
                            t.Run("CheckAllTheGeoSetsArePresentInTheDefaultLanguage", func(t *testing.T) {
                                assert.True(t, len(internalCodeMap) == len(translationKeys),
                                    "Not all the geosets are present in the translation for default language")
                            })
                        }
                    })
                }
            })
        })
    })

    t.Run("TestsOnCurrenciesData", func(t *testing.T) {
        t.Run("CheckTheDataCurrenciesStructure", func(t *testing.T) {
            seenIsoAlpha := []string{}
            seenIsoNumber := []string{}
            for _, cu := range *CU {
                t.Run("CheckTheDataStructureForCurrency:" + cu.IsoAlpha, func(t *testing.T) {

                    /** isoAlpha **/
                    t.Run("CheckTheIsoAlphaProperty", func(t *testing.T) {
                        isValidFormat := regexp.MustCompile(`^[A-Z]{3}$`).MatchString(cu.IsoAlpha)
                        assert.True(t, isValidFormat, "Wrong format for isoAlpha")
                        isDuplicate := TestLib.IsInSlice(seenIsoAlpha, cu.IsoAlpha)
                        assert.False(t, isDuplicate, fmt.Sprintf("Duplicate isoAlpha found: %s", cu.IsoAlpha))
                        seenIsoAlpha = append(seenIsoAlpha, cu.IsoAlpha)
                    })

                    /** isoNumber **/
                    t.Run("CheckTheIsoNumberProperty", func(t *testing.T) {
                        isValidFormat := regexp.MustCompile(`^[0-9]{3}$`).MatchString(cu.IsoNumber)
                        assert.True(t, isValidFormat, "Wrong format for isoNumber")
                        isDuplicate := TestLib.IsInSlice(seenIsoNumber, cu.IsoNumber)
                        assert.False(t, isDuplicate, fmt.Sprintf("Duplicate isoNumber found: %s", cu.IsoNumber))
                        seenIsoAlpha = append(seenIsoNumber, cu.IsoNumber)
                    })
                })
            }
        })

        t.Run("TestsOnCurrenciesTranslationData", func(t *testing.T) {
            TranslationsCU := make(map[string]*Structs.TransGeneric)
            t.Run("TranslationCurrenciesDataInitialization", func(t *testing.T) {
                for _, lang := range translations {
                    var tr = new(Structs.TransGeneric)
                    TestLib.LoadData("trans_" + lang + "currencies", transData.Currencies[lang], tr)
                    loadedData := TestLib.GetData("trans_" + lang + "currencies")
                    var ok bool
                    tr, ok = loadedData.(*Structs.TransGeneric)
                    assert.True(t, ok, fmt.Sprintf("Failed to convert data to Translations %s Currencies", lang))
                    if tr == nil {
                        t.Fatal(fmt.Sprintf("Translation %s data is not initialized", lang))
                    }
                    TranslationsCU[lang] = tr
                }
            })
            t.Run("TranslationCurrenciesCheckData", func(t *testing.T) {
                for lang, trans := range TranslationsCU {
                    t.Run(fmt.Sprintf("TestForTheLanguage:%s", lang), func(t *testing.T) {
                        var translationKeys []string
                        for cu, tr :=range *trans {
                            t.Run(fmt.Sprintf("TestForTheCurrency:%s", cu), func(t *testing.T) {
                                translationKeys = append(translationKeys, cu)
                                t.Run("CheckTheCurrencyExists", func(t *testing.T) {
                                    assert.True(t, TestLib.IsInMap(isoAlphaMap, cu),
                                        "The currency code in the translation must match an existing internalCode")
                                })

                                if lang == cfg.Settings.Languages.Default {
                                    trimmedName := strings.TrimSpace(tr.Name)
                                    t.Run("CheckTranslationNameExistsForDefaultLanguage", func(t *testing.T) {
                                        assert.True(t, trimmedName != "",
                                            "In default language the property name must exist and not be empty")
                                    })
                                }
                            })
                        }
                        if lang == cfg.Settings.Languages.Default {
                            t.Run("CheckAllTheCurrenciesArePresentInTheDefaultLanguage", func(t *testing.T) {
                                assert.True(t, len(isoAlphaMap) == len(translationKeys),
                                    "Not all the currencies are present in the translation for default language")
                            })
                        }
                    })
                }
            })
        })
    })

}
