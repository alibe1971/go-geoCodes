package geoCodes

import(
    "reflect"
    "encoding/json"
    "crypto/rand"
    "time"
    "github.com/google/uuid"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "log"
    "os"
    "runtime"
    "bytes"
    "strings"
    "strconv"
    Structs "github.com/alibe1971/go-geoCodes/geoCodes/Structs"
)

func strToLower(str string) string {
    return strings.ToLower(str)
}

func strToUpper(str string) string {
    return strings.ToUpper(str)
}

func toString(value interface{}) string {
    switch v := value.(type) {
        case string:
            return v
        case int:
            return strconv.Itoa(v)
        case float64:
            return strconv.FormatFloat(v, 'f', -1, 64)
        default:
            logPanicWithStackTrace(fmt.Sprintf("Unsupported type: %T", v))
            return ""
    }
}

func setElementsToStrings(values []interface{}) []string {
    var stringValues []string
    for _, value := range values {
        stringValues = append(stringValues, toString(value))
    }
    return stringValues
}

func setElementsToIntegers(values []interface{}) []int {
    var intValues []int
    for _, value := range values {
        switch v := value.(type) {
            case int:
                intValues = append(intValues, v)
            case string:
                if intVal, err := strconv.Atoi(v); err == nil {
                    intValues = append(intValues, intVal)
                } else {
                    logPanicWithStackTrace(fmt.Sprintf("Unable to convert `%s` to int", v))
                }
            default:
                logPanicWithStackTrace(fmt.Sprintf("Unsupported type for integer conversion: %T", v))
        }
    }
    return intValues
}

func loadData(key string, jsonData string, v interface{}) {
    err := json.Unmarshal([]byte(jsonData), v)
    if err != nil {
        fmt.Printf("Failed to unmarshal JSON for %s: %v\n", key, err)
        return
    }
    dataMap[key] = v
}


func getData(key string) interface{} {
    if data, ok := dataMap[key]; ok {
            return data
        }
    fmt.Printf("Data with key %s not found\n", key)
    return nil
}

func isInSlice(slice []string, value string) bool {
    for _, item := range slice {
        if item == value {
            return true
        }
    }
    return false
}


func isInMap(m map[string]interface{}, value string) bool {
    _, exists := m[value]
    return exists
}

func getStructPropertiesNamesSingleLevel(s interface{}) []string {
    var fieldNames []string
    t := reflect.TypeOf(s)
    if t.Kind() != reflect.Struct {
        return nil
    }
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        fieldNames = append(fieldNames, field.Name)
    }
    return fieldNames
}

func getStructPropertiesNamesMultilevel(s interface{}, prefix string) []string {
	var fieldNames []string
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Struct {
		return nil
	}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldName := field.Name
		fullName := fieldName
		if prefix != "" {
			fullName = prefix + "." + fieldName
		}
		fieldNames = append(fieldNames, fullName)

		fieldType := field.Type
		if fieldType.Kind() == reflect.Struct {
			nestedStruct := reflect.New(fieldType).Elem().Interface()
			fieldNames = append(fieldNames, getStructPropertiesNamesMultilevel(nestedStruct, fullName)...)
		}

		if fieldType.Kind() == reflect.Ptr && fieldType.Elem().Kind() == reflect.Struct {
			nestedStruct := reflect.New(fieldType.Elem()).Elem().Interface()
			fieldNames = append(fieldNames, getStructPropertiesNamesMultilevel(nestedStruct, fullName)...)
		}
	}
	return fieldNames
}


func structToMap(v reflect.Value) map[string]interface{} {
	result := make(map[string]interface{})
	vType := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := vType.Field(i).Name
		result[fieldName] = field.Interface()
	}
	return result
}

func generateUniqueString() Structs.GeoCodeReference {
    uuidStr := uuid.New().String()
    timestamp := time.Now().UnixNano()
    randomBytes := make([]byte, 8)
    _, err := rand.Read(randomBytes)
    if err != nil {
        panic(err)
    }
    data := fmt.Sprintf("%s-%d-%x", uuidStr, timestamp, randomBytes)
    hash := sha256.Sum256([]byte(data))
    reference := Structs.GeoCodeReference(hex.EncodeToString(hash[:]))
    return reference
}


func logFatalWithTrace(message string) {
    _, file, line, ok := runtime.Caller(3)
    if ok {
        log.SetOutput(os.Stderr)
        log.Fatalf("Critical: %s\nFile: %s\nLine: %d\n", message, file, line)
    } else {
        log.SetOutput(os.Stderr)
        log.Fatalln("Critical: %s", message)
    }
}

func logPanicWithStackTrace(message string) {
    const depth = 15
    buf := make([]byte, 1024)
    n := runtime.Stack(buf, false)
    stackTrace := string(buf[:n])
    reversedStackTrace := reverseStackTrace(stackTrace)
    log.SetOutput(os.Stderr)
    panic(fmt.Sprintf("Critical Error: %s\nStack Trace:\n%s", message, reversedStackTrace))
}


func reverseStackTrace(trace string) string {
    var reversed bytes.Buffer
    lines := strings.Split(trace, "\n")
    for i := len(lines) - 1; i >= 0; i-- {
        reversed.WriteString(lines[i] + "\n")
    }
    return reversed.String()
}


// ALIBE
func WriteDataToFile(data interface{}) error {
    // Crea il file
    var filename = "/Users/aliberati/ALIBE/test.log"
    file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return fmt.Errorf("failed to create file %s: %w", filename, err)
    }
    defer file.Close()

    // Converti i dati in una stringa JSON formattata
    jsonData, err := json.MarshalIndent(data, "", "    ")
    if err != nil {
        return fmt.Errorf("failed to marshal data to JSON: %w", err)
    }

    // Scrivi la stringa JSON nel file
    if _, err := file.WriteString(string(jsonData) + "\n"); err != nil {
        return fmt.Errorf("failed to write data to file: %w", err)
    }

    return nil
}