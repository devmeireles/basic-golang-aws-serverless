package helpers

import (
	"reflect"
)

func HandleTestKey(structItem interface{}) []string {
	var keys []string

	val := reflect.ValueOf(structItem).Elem()
	for i := 0; i < val.NumField(); i++ {
		keys = append(keys, string(val.Type().Field(i).Tag))
	}

	return keys
}
