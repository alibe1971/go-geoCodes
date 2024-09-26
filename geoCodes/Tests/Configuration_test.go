package geoCodesTest

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/alibe1971/go-geoCodes/geoCodes"
    "github.com/alibe1971/go-geoCodes/geoCodes/Tests/TestLib"
    "fmt"
)

func TestConfiguration(t *testing.T) {
    t.Run("Test the Configuration functionality", func(t *testing.T) {

        t.Run("Check the available languages", func(t *testing.T) {
            languages := geoCodes.GetAvailableLanguages()
            assert.True(t, TestLib.IsInSlice(languages, "en"), fmt.Sprintf("Language `en` not found"))
            assert.True(t, TestLib.IsInSlice(languages, "it"), fmt.Sprintf("Language `it` not found"))
        })

        t.Run("Check the default language", func(t *testing.T) {
            languages := geoCodes.GetDefaultLanguage()
            assert.True(t, languages == "en", fmt.Sprintf("Language `en` is not the default language"))
        })

        t.Run("Check the current language", func(t *testing.T) {
            languages := geoCodes.GetCurrentLanguage()
            assert.True(t, languages == "en", fmt.Sprintf("Language `en` is not the current language"))
        })

        t.Run("Correctly change the default language", func(t *testing.T) {
            geoCodes.SetDefaultLanguage("it")
            languages := geoCodes.GetDefaultLanguage()
            assert.True(t, languages == "it", fmt.Sprintf("Language `it` is not the default language"))
        })

        t.Run("Correctly change the current language", func(t *testing.T) {
            geoCodes.UseLanguage("it")
            languages := geoCodes.GetCurrentLanguage()
            assert.True(t, languages == "it", fmt.Sprintf("Language `it` is not the current language"))
        })


        t.Run("Try to set default language with a not valid language", func(t *testing.T) {
            defer func() {
                if r := recover(); r != nil {
                   assert.Contains(t, r.(string), "not a valid language")
                   return
                }
                t.Error("Expected panic, but no panic occurred")
            }()
            geoCodes.SetDefaultLanguage("xyz")
        })

        t.Run("Try to set default language with a not valid language", func(t *testing.T) {
            defer func() {
                if r := recover(); r != nil {
                  assert.Contains(t, r.(string), "not a valid language")
                  return
                }
                t.Error("Expected panic, but no panic occurred")
            }()
            geoCodes.UseLanguage("xyz")
        })

        t.Run("Reset the languages", func(t *testing.T) {
            geoCodes.ResetLanguages()
            defaultLang := geoCodes.GetDefaultLanguage()
            assert.True(t, defaultLang == "en", fmt.Sprintf("Language `en` is not the default language"))
            currentLang := geoCodes.GetCurrentLanguage()
            assert.True(t, currentLang == "en", fmt.Sprintf("Language `en` is not the current language"))
        })

    })
}


