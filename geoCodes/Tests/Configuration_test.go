package geoCodesTest

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/alibe1971/go-geoCodes/geoCodes"
    "github.com/alibe1971/go-geoCodes/geoCodes/Tests/TestLib"
    "fmt"
)

func TestConfiguration(t *testing.T) {
    t.Run("TestTheConfigurationFunctionality", func(t *testing.T) {

        t.Run("CheckTheAvailableLanguages", func(t *testing.T) {
            languages := geoCodes.GetAvailableLanguages()
            assert.True(t, TestLib.IsInSlice(languages, "en"), fmt.Sprintf("Language `en` not found"))
            assert.True(t, TestLib.IsInSlice(languages, "it"), fmt.Sprintf("Language `it` not found"))
        })

        t.Run("CheckTheDefaultLanguage", func(t *testing.T) {
            languages := geoCodes.GetDefaultLanguage()
            assert.True(t, languages == "en", fmt.Sprintf("Language `en` is not the default language"))
        })

        t.Run("CheckTheCurrentLanguage", func(t *testing.T) {
            languages := geoCodes.GetCurrentLanguage()
            assert.True(t, languages == "en", fmt.Sprintf("Language `en` is not the current language"))
        })

        t.Run("CorrectlyChangeTheDefaultLanguage", func(t *testing.T) {
            geoCodes.SetDefaultLanguage("it")
            languages := geoCodes.GetDefaultLanguage()
            assert.True(t, languages == "it", fmt.Sprintf("Language `it` is not the default language"))
        })

        t.Run("CorrectlyChangeTheCurrentLanguage", func(t *testing.T) {
            geoCodes.UseLanguage("it")
            languages := geoCodes.GetCurrentLanguage()
            assert.True(t, languages == "it", fmt.Sprintf("Language `it` is not the current language"))
        })


        t.Run("TryToSetDefaultLanguageWithANotValidLanguage", func(t *testing.T) {
            defer func() {
                if r := recover(); r != nil {
                   assert.Contains(t, r.(string), "not a valid language")
                   return
                }
                t.Error("Expected panic, but no panic occurred")
            }()
            geoCodes.SetDefaultLanguage("xyz")
        })

        t.Run("TryToSetDefaultLanguageWithANotValidLanguage", func(t *testing.T) {
            defer func() {
                if r := recover(); r != nil {
                  assert.Contains(t, r.(string), "not a valid language")
                  return
                }
                t.Error("Expected panic, but no panic occurred")
            }()
            geoCodes.UseLanguage("xyz")
        })

        t.Run("ResetTheLanguages", func(t *testing.T) {
            geoCodes.ResetLanguages()
            defaultLang := geoCodes.GetDefaultLanguage()
            assert.True(t, defaultLang == "en", fmt.Sprintf("Language `en` is not the default language"))
            currentLang := geoCodes.GetCurrentLanguage()
            assert.True(t, currentLang == "en", fmt.Sprintf("Language `en` is not the current language"))
        })

    })
}


