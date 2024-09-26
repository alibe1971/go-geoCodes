package geoCodesTest

import (
    "testing"
//     "github.com/stretchr/testify/assert"
    "github.com/alibe1971/go-geoCodes/geoCodes"
    "github.com/alibe1971/go-geoCodes/geoCodes/Tests/TestLib"
    "fmt"
)

func TestCountries(t *testing.T) {
    t.Run("Test the Countries functionality", func(t *testing.T) {

        t.Run("Check the available languages", func(t *testing.T) {

        })



    })
}


func TestElibeCountries(t *testing.T) {
//      geoCodes.UseLanguage("it")
//     fmt.Printf("Lingua: %v\n", geoCodes.GetAvailableLanguages())
//     fmt.Println("   \n")
//      geoCodes.UseLanguage("itss")
//     stica := "stica"

// TEST Get() IN CASO DI SLICE DA RISTRUTTURARE
//     country0 := geoCodes.Countries().Get().Data.([]map[string]interface{})[0]
//     fmt.Printf("%v", country0)

// TEST Get() IN CASO DI MAP FUNZIONE AsSlice
//     country0 := geoCodes.Countries().Get().AsSlice()[0]
//     fmt.Printf("%v", country0)

// TEST Get() IN CASO DI MAP  DA RISTRUTTURARE
//     country0 := geoCodes.Countries().Get().Data.(map[string]map[string]interface{})["IT"]
//     fmt.Printf("%v", country0)

// TEST Get() IN CASO DI MAP FUNZIONE AsMap
//     country0 := geoCodes.Countries().Get().AsMap()["IT"]
//     fmt.Printf("%v", country0)

// TEST First SE USI AsMap o AsSlice ritorna empty
//     country0 := geoCodes.Countries().First().Data
//     fmt.Printf("%v", country0)

// TEST Count
//     country0 := geoCodes.Countries().Count()
//     country0 := geoCodes.Countries().WithIndex("FullName").Count()
//     fmt.Printf("Conteggio: %v", country0)

// TEST Get().ToJson()
    country0 := geoCodes.Countries().First().ToJson()
//     country0 := geoCodes.Countries().Get().ToJson()
//     country0 := geoCodes.Countries().WithIndex("Alpha2").Get().ToJson()
    fmt.Printf("%v", country0)

// TEST Get().ToYaml()
//     country0 := geoCodes.Countries().First().ToYaml()
//     country0 := geoCodes.Countries().Get().ToYaml()
//     country0 := geoCodes.Countries().WithIndex("Alpha2").Get().ToYaml()
//     fmt.Printf("%v", country0)

// TEST Get().ToXml()
//     country0 := geoCodes.Countries().Get().ToXml()
//     fmt.Printf("%v", country0)

// TEST Get().ToXmlAndValidate()
//     country0 := geoCodes.GeoSets().First().ToXml()
//     country0 := geoCodes.GeoSets().Get().ToXml()
//     country0 := geoCodes.GeoSets().WithIndex("Name").Get().ToXml()
//     fmt.Printf("%v\n", country0)

// TEST Get().ToYaml()
//     country0 := geoCodes.Countries().Get().ToYaml()
//     fmt.Printf("%v", country0)


//     country0 := geoCodes.Countries().WithIndex("FullName").Select("Alpha2", "Alpha3", "Name", "OfficialName").Get().ToXml()
//     fmt.Printf("%v", country0)

//     country0 := geoCodes.Countries().First().Data
//     fmt.Printf("%v", country0)
//     TestLib.WriteDataToFile(country0)


//     country0 := geoCodes.Countries().GetXsd()
//     fmt.Printf("%v\n", country0)
//     country0 = geoCodes.Countries().GetXsdSingle()
//     fmt.Printf("%v\n", country0)
    TestLib.WriteDataToFile("\n")
//     geoset0 := geoCodes.GeoSets()
//     TestLib.WriteDataToFile(geoset0)
//
//     currency0 := geoCodes.Currencies()
//     TestLib.WriteDataToFile(currency0)

}
