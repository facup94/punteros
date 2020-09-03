package nullable

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/google/go-cmp/cmp"
)

type Nullable interface {
	HasValue() bool
}

func HasValueStruct(i interface{}) bool {
	if i == nil {
		return false
	}
	iType := reflect.TypeOf(i)
	zero := reflect.Zero(iType)
	return !cmp.Equal(i, zero.Interface())
}

func MarshalStruct(e Nullable) ([]byte, error) {
	if !e.HasValue() {
		return json.Marshal(nil)
	}

	first := true

	resultJSON := "{"

	structValue := reflect.Indirect(reflect.ValueOf(e))
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Type().Field(i)
		value := structValue.Field(i).Interface()

		if shouldOmit(field, value) {
			continue
		}

		fieldJSON, err := getFieldJSON(field, value)
		if err != nil {
			return nil, err
		}

		if !first {
			resultJSON += ","
		}

		resultJSON += fieldJSON

		first = false
	}

	resultJSON += "}"

	return []byte(resultJSON), nil
}

func shouldOmit(field reflect.StructField, value interface{}) bool {
	tag := field.Tag.Get("json")
	return strings.Contains(tag, "omitempty") && !HasValueStruct(value)
}

func getFieldName(field reflect.StructField) string {
	tag := field.Tag.Get("json")
	if i := strings.Index(tag, ","); i != -1 {
		return tag[:i]
	}
	return tag
}

func getFieldJSON(field reflect.StructField, value interface{}) (fieldJSON string, err error) {
	name := getFieldName(field)
	bytes, err := json.Marshal(value)
	if err != nil {
		return
	}

	fieldJSON = `"` + name + `":` + string(bytes)

	return
}

func marshall(hasValue bool, value interface{}) ([]byte, error) {
	if !hasValue {
		return json.Marshal(nil)
	}

	return json.Marshal(value)
}

func unmarshal(data []byte, hasValue *bool, value interface{}) (err error) {
	var v interface{}

	if err = json.Unmarshal(data, &v); err != nil {
		return
	}

	if v == nil {
		*hasValue = false
	} else {
		err = json.Unmarshal(data, value)
		*hasValue = err == nil
	}

	return
}

func toString(hasValue bool, value interface{}) string {
	if !hasValue {
		return ""
	}
	return fmt.Sprintf("%v", value)
}
