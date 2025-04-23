package TestLib

import (
    "fmt"
    "encoding/json"
    "encoding/xml"
    "gopkg.in/yaml.v3"
    "strings"
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

func ValidateYAML(data []byte) error {
	var y interface{}
	return yaml.Unmarshal(data, &y)
}

func ValidateXML(data []byte) error {
	var x interface{}
	return xml.Unmarshal(data, &x)
}

func ValidateJSON(data []byte) error {
	var js interface{}
	return json.Unmarshal(data, &js)
}

var ContainsOrRelated = func(selected []string, f string) bool {
    for _, sel := range selected {
        if sel == f {
            return true
        }
        // if sel is prefix of f (parent field)
        if strings.HasPrefix(f, sel+".") {
            return true
        }
        // if f is prefix of sel (child field)
        if strings.HasPrefix(sel, f+".") {
            return true
        }
    }
    return false
}
