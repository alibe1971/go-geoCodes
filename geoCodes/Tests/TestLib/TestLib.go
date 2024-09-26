package TestLib

import (
    "fmt"
//     "bytes"
    "encoding/json"
    "encoding/xml"
//     "sync"
//     "os"
//     "log"
//     "io"
//     "runtime/debug"
//     "io/ioutil"
)

var dataMap = make(map[string]interface{})

func LoadData(key string, jsonData string, v interface{}) {
    err := json.Unmarshal([]byte(jsonData), v)
    if err != nil {
        fmt.Printf("Failed to unmarshal JSON for %s: %v\n", key, err)
        return
    }
    dataMap[key] = v
}
func GetData(key string) interface{} {
    if data, ok := dataMap[key]; ok {
            return data
        }
    fmt.Printf("Data with key %s not found\n", key)
    return nil
}
func IsInMap(m map[string]struct{}, value string) bool {
    _, exists := m[value]
    return exists
}
func IsValidSVG(svgContent string) bool {
    var svg struct {
        XMLName xml.Name `xml:"svg"`
    }
    return xml.Unmarshal([]byte(svgContent), &svg) == nil
}
func IsInSlice(slice []string, value string) bool {
    for _, item := range slice {
        if item == value {
            return true
        }
    }
    return false
}

