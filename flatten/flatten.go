package flatten

import (
	"reflect"
)

func Flatten(nested interface{}) []interface{} {

	var result []interface{} = []interface{}{}

	sliceData := reflect.ValueOf(nested)

	if sliceData.Kind() != reflect.Slice {

		return result
	}

	if sliceData.IsNil() {

		return result
	}

	for i := 0; i < sliceData.Len(); i++ {
		val := sliceData.Index(i).Interface()

		isValid := false

		if _, ok := val.(int); ok {
			isValid = true
			result = append(result, val)
		}
		if _, ok := val.(string); ok {
			isValid = true
			result = append(result, val)
		}

		if !isValid {

			res := Flatten(val)

			if len(res) > 0 {

				result = append(result, Flatten(val)...)
			}

		}

	}

	return result
}
