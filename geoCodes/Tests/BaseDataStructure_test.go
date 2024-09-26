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

    t.Run("DataBase initialization", func(t *testing.T) {
        t.Run("Config data initialization", func(t *testing.T) {
            TestLib.LoadData("config", data.Config, &Structs.Config{})
            cfg = TestLib.GetData("config").(*Structs.Config)
            if cfg == nil {
                t.Fatal("Config data is not initialized")
            }
        })
        t.Run("Countries data initialization", func(t *testing.T) {
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
        t.Run("GeoSets data initialization", func(t *testing.T) {
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
        t.Run("Currencies data initialization", func(t *testing.T) {
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
                isoAlphaMap[cu.ISOAlpha] = struct{}{}
            }
        })
    })


    t.Run("Tests on Config data", func(t *testing.T) {

        t.Run("Check the data Config structure", func(t *testing.T) {
            translations = make([]string, 0, len(cfg.Settings.Languages.InPackage))
            for key, locale := range cfg.Settings.Languages.InPackage {
                assert.True(t, regexp.MustCompile(`^[a-z]{2}(_[A-Za-z]+)*(_[A-Z]{2})?$`).MatchString(locale),
                    fmt.Sprintf("Wrong format for locale: %s", locale))
                translations = append(translations, key)
            }
            assert.Equal(t, "en", cfg.Settings.Languages.Default, "Default should be 'en'")
        })
    })

    t.Run("Tests on Countries data", func(t *testing.T) {
        t.Run("Check the data Countries structure", func(t *testing.T) {
            seenAlpha2 := []string{}
            seenAlpha3 := []string{}
            seenUnM49  := []string{}
            uniqueKeysOfficialName := make(map[string]bool)
            for _, cc := range *CC {
                t.Run("Check the data structure for country " + cc.Alpha2, func(t *testing.T) {

                    /** alpha2 **/
                    t.Run("Check the alpha2 property", func(t *testing.T) {
                        assert.True(t, regexp.MustCompile(`^[A-Z]{2}$`).MatchString(cc.Alpha2),
                            "Wrong format for alpha2")
                        assert.False(t, TestLib.IsInSlice(seenAlpha2, cc.Alpha2),
                            fmt.Sprintf("Duplicate alpha2 found: %s", cc.Alpha2))
                        seenAlpha2 = append(seenAlpha2, cc.Alpha2)
                    })

                    /** alpha3 **/
                    t.Run("Check the alpha3 property", func(t *testing.T) {
                        assert.True(t, regexp.MustCompile(`^[A-Z]{3}$`).MatchString(cc.Alpha3),
                            "Wrong format for alpha3")
                        assert.False(t, TestLib.IsInSlice(seenAlpha3, cc.Alpha3),
                            fmt.Sprintf("Duplicate alpha3 found: %s", cc.Alpha3))
                        seenAlpha3 = append(seenAlpha3, cc.Alpha3)
                    })

                    /** unM49 **/
                    t.Run("Check the unM49 property", func(t *testing.T) {
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
                    t.Run("Check the flag property", func(t *testing.T) {
                        t.Run("Check the flag.svg property", func(t *testing.T) {
                            assert.NotEmpty(t, cc.Flags.SVG, "flags.SVG must not be empty")
                            assert.True(t, TestLib.IsValidSVG(cc.Flags.SVG), "Flags.SVG must be a valid SVG")
                        })
                    })

                    /** dependency **/
                    t.Run("Check the dependency property", func(t *testing.T) {
                        if cc.Dependency != nil {
                            dependencyStr := *cc.Dependency
                            assert.True(t, regexp.MustCompile(`^[A-Z]{2}$`).MatchString(dependencyStr),
                                "Wrong format for dependency")
                            assert.True(t, TestLib.IsInMap(alpha2Map, dependencyStr),
                                "The value of dependency must match an existing alpha2")
                        }
                    })

                    /** officialName **/
                    t.Run("Check the officialName property", func(t *testing.T) {
                        assert.NotEmpty(t, cc.OfficialName, "officialName must not be empty")
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
                    t.Run("Check the mottos property", func(t *testing.T) {
                         // Nothing to do
                    })

                    /** currencies **/
                    t.Run("Check the currency property", func(t *testing.T) {
                        t.Run("Check the currency.legalTender property", func(t *testing.T) {
                            if len(cc.Currencies.LegalTenders) > 0 {
                                for _, currency := range cc.Currencies.LegalTenders {
                                    assert.True(t, regexp.MustCompile(`^[A-Z]{3}$`).MatchString(currency),
                                        "Wrong format for LegalTenders")
                                    assert.True(t, TestLib.IsInMap(isoAlphaMap, currency),
                                        fmt.Sprintf("The currency '%s' for LegalTenders isn't in the currency database", currency))
                                }
                            }
                        })
                        t.Run("Check the currency.widelyAccepted property", func(t *testing.T) {
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
                    t.Run("Check the dialCodes property", func(t *testing.T) {
                        t.Run("Check the dialCodes.main values have the right format", func(t *testing.T) {
                            if len(cc.DialCodes.Main) > 0 {
                                for _, dial := range cc.DialCodes.Main {
                                    assert.True(t, regexp.MustCompile(`^\+\d+$`).MatchString(dial),
                                        "Wrong format for dialCodes.main")
                                }
                            }
                        })
                        t.Run("Check the dialCodes.exceptions values have the right format", func(t *testing.T) {
                            if len(cc.DialCodes.Exceptions) > 0 {
                                for _, dial := range cc.DialCodes.Exceptions {
                                    assert.True(t, regexp.MustCompile(`^\+\d+$`).MatchString(dial),
                                        "Wrong format for dialCodes.exceptions")
                                }
                            }
                        })
                    })

                    /** ccTld **/
                    t.Run("Check the ccTld property", func(t *testing.T) {
                        if cc.CcTLD != nil {
                            CcTld := *cc.CcTLD
                            assert.True(t, regexp.MustCompile(`^\.[a-z]{2}$`).MatchString(CcTld),
                                "Wrong format for ccTLD")
                        }
                    })

                    /** timeZones **/
                    t.Run("Check the timeZones property", func(t *testing.T) {
                         assert.NotEmpty(t, cc.TimeZones, "The timeZones cannot be empty")
                    })

                    /** languages **/
                    t.Run("Check the languages property", func(t *testing.T) {
                         // [TODO]
                    })

                    /** locales **/
                    t.Run("Check the locales property", func(t *testing.T) {
                         assert.NotEmpty(t, cc.Locales, "The timeZones locales be empty")
                    })

                })
            }
        })

        t.Run("Tests on Countries Translation data", func(t *testing.T) {
            TranslationsCC := make(map[string]*Structs.TransCountries)
            t.Run("Translation Countries data initialization", func(t *testing.T) {
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
            t.Run("Translation Countries check data", func(t *testing.T) {
                for lang, trans := range TranslationsCC {
                    t.Run(fmt.Sprintf("Test for the language: %s", lang), func(t *testing.T) {
                        var translationKeys []string
                        for cc, tr :=range *trans {
                            t.Run(fmt.Sprintf("Test for the country: %s", cc), func(t *testing.T) {
                                translationKeys = append(translationKeys, cc)
                                t.Run("Check the country exists", func(t *testing.T) {
                                    assert.True(t, TestLib.IsInMap(alpha2Map, cc),
                                        "The country code in the translation must match an existing alpha2")
                                })

                                if lang == cfg.Settings.Languages.Default {
                                    trimmedName := strings.TrimSpace(tr.Name)
                                    t.Run("Check translation name exists for default language", func(t *testing.T) {
                                        assert.True(t, trimmedName != "",
                                            "In default language the property name must exist and not be empty")
                                    })
                                    trimmedFullName := strings.TrimSpace(tr.FullName)
                                    t.Run("Check translation name exists for default language", func(t *testing.T) {
                                        assert.True(t, trimmedFullName != "",
                                            "In default language the property FullName must exist and not be empty")
                                    })
                                }
                            })
                        }
                        if lang == cfg.Settings.Languages.Default {
                            t.Run("Check all the countries are present in the default language", func(t *testing.T) {
                                assert.True(t, len(alpha2Map) == len(translationKeys),
                                    "Not all the countries are present in the translation for default language")
                            })
                        }
                    })
                }
            })
        })
    })

    t.Run("Tests on geoSets data", func(t *testing.T) {
        t.Run("Check the data geoSets structure", func(t *testing.T) {
            seenInternalCode := []string{}
            seenUnM49  := []string{}
            seenGEOG := make(map[string][]string)
            geogGr := []string{}
            geoLv := make([][]string, 2)
            for _, gs := range *GS {
                t.Run("Check the data structure for geoSet " + gs.InternalCode, func(t *testing.T) {
                    /** internalCode **/
                    t.Run("Check the internalCode property", func(t *testing.T) {
                        isValidFormat := regexp.MustCompile(`^[A-Z]+(-[A-Z0-9]+){1,4}$`).MatchString(gs.InternalCode)
                        assert.True(t, isValidFormat, "Wrong format for InternalCode")
                        isDuplicate := TestLib.IsInSlice(seenInternalCode, gs.InternalCode)
                        assert.False(t, isDuplicate, fmt.Sprintf("Duplicate InternalCode found: %s", gs.InternalCode))
                        seenInternalCode = append(seenInternalCode, gs.InternalCode)
                    })

                    /** unM49 **/
                    if gs.UnM49 != nil {
                        UnM49 := *gs.UnM49
                        t.Run("Check the unM49 property", func(t *testing.T) {
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

                    /** timeZones **/
                    t.Run("Check the tags property", func(t *testing.T) {
                         assert.NotEmpty(t, gs.Tags, "The tags cannot be empty")
                    })

                    /** countryCodes **/
                    t.Run("Check the countryCodes property", func(t *testing.T) {
                         assert.NotEmpty(t, gs.CountryCodes, "The countryCodes cannot be empty")

                         for _,cc := range gs.CountryCodes {
                            t.Run("Check the country code exists", func(t *testing.T) {
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
            t.Run("Test that the geographic group have all the countries", func(t *testing.T) {
                str := fmt.Sprintf("%d", len(geogGr))
                assert.True(t, len(alpha2Map) == len(geogGr),
                    "The Geographic Groups haven't inside all the countries " + str)
            })
        })

        t.Run("Tests on GeoSets Translation data", func(t *testing.T) {
            TranslationsGS := make(map[string]*Structs.TransGeneric)
            t.Run("Translation GeoSets data initialization", func(t *testing.T) {
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
            t.Run("Translation GeoSets check data", func(t *testing.T) {
                for lang, trans := range TranslationsGS {
                    t.Run(fmt.Sprintf("Test for the language: %s", lang), func(t *testing.T) {
                        var translationKeys []string
                        for gs, tr :=range *trans {
                            t.Run(fmt.Sprintf("Test for the geoSet: %s", gs), func(t *testing.T) {
                                translationKeys = append(translationKeys, gs)
                                t.Run("Check the geoSet exists", func(t *testing.T) {
                                    assert.True(t, TestLib.IsInMap(internalCodeMap, gs),
                                        "The country code in the translation must match an existing internalCode")
                                })

                                if lang == cfg.Settings.Languages.Default {
                                    trimmedName := strings.TrimSpace(tr.Name)
                                    t.Run("Check translation name exists for default language", func(t *testing.T) {
                                        assert.True(t, trimmedName != "",
                                            "In default language the property name must exist and not be empty")
                                    })
                                }
                            })
                        }
                        if lang == cfg.Settings.Languages.Default {
                            t.Run("Check all the geoSets are present in the default language", func(t *testing.T) {
                                assert.True(t, len(internalCodeMap) == len(translationKeys),
                                    "Not all the geosets are present in the translation for default language")
                            })
                        }
                    })
                }
            })
        })
    })

    t.Run("Tests on Currencies data", func(t *testing.T) {
        t.Run("Check the data Currencies structure", func(t *testing.T) {
            seenIsoAlpha := []string{}
            seenIsoNumber := []string{}
            for _, cu := range *CU {
                t.Run("Check the data structure for currency " + cu.ISOAlpha, func(t *testing.T) {

                    /** isoAlpha **/
                    t.Run("Check the isoAlpha property", func(t *testing.T) {
                        isValidFormat := regexp.MustCompile(`^[A-Z]{3}$`).MatchString(cu.ISOAlpha)
                        assert.True(t, isValidFormat, "Wrong format for isoAlpha")
                        isDuplicate := TestLib.IsInSlice(seenIsoAlpha, cu.ISOAlpha)
                        assert.False(t, isDuplicate, fmt.Sprintf("Duplicate isoAlpha found: %s", cu.ISOAlpha))
                        seenIsoAlpha = append(seenIsoAlpha, cu.ISOAlpha)
                    })

                    /** isoNumber **/
                    t.Run("Check the isoNumber property", func(t *testing.T) {
                        isValidFormat := regexp.MustCompile(`^[0-9]{3}$`).MatchString(cu.ISONumber)
                        assert.True(t, isValidFormat, "Wrong format for isoNumber")
                        isDuplicate := TestLib.IsInSlice(seenIsoNumber, cu.ISONumber)
                        assert.False(t, isDuplicate, fmt.Sprintf("Duplicate isoNumber found: %s", cu.ISONumber))
                        seenIsoAlpha = append(seenIsoNumber, cu.ISONumber)
                    })
                })
            }
        })

        t.Run("Tests on Currencies Translation data", func(t *testing.T) {
            TranslationsCU := make(map[string]*Structs.TransGeneric)
            t.Run("Translation Currencies data initialization", func(t *testing.T) {
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
            t.Run("Translation Currencies check data", func(t *testing.T) {
                for lang, trans := range TranslationsCU {
                    t.Run(fmt.Sprintf("Test for the language: %s", lang), func(t *testing.T) {
                        var translationKeys []string
                        for cu, tr :=range *trans {
                            t.Run(fmt.Sprintf("Test for the currency: %s", cu), func(t *testing.T) {
                                translationKeys = append(translationKeys, cu)
                                t.Run("Check the currency exists", func(t *testing.T) {
                                    assert.True(t, TestLib.IsInMap(isoAlphaMap, cu),
                                        "The currency code in the translation must match an existing internalCode")
                                })

                                if lang == cfg.Settings.Languages.Default {
                                    trimmedName := strings.TrimSpace(tr.Name)
                                    t.Run("Check translation name exists for default language", func(t *testing.T) {
                                        assert.True(t, trimmedName != "",
                                            "In default language the property name must exist and not be empty")
                                    })
                                }
                            })
                        }
                        if lang == cfg.Settings.Languages.Default {
                            t.Run("Check all the currencies are present in the default language", func(t *testing.T) {
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
