package util

import (
	"reflect"
)

func FlatField(structEnt interface{}) (fields []*reflect.StructField, values []*reflect.Value) {
	flatField(structEnt, &fields, &values)
	return
}

func flatField(structEnt interface{}, fields *[]*reflect.StructField, values *[]*reflect.Value) {
	dStructType := reflect.TypeOf(structEnt)
	dStructValue := reflect.ValueOf(structEnt)

	if dStructType.Kind() == reflect.Ptr {
		dStructType = dStructType.Elem()
		dStructValue = dStructValue.Elem()
	}

	for i := 0; i < dStructType.NumField(); i++ {
		field := dStructType.Field(i)
		if field.Tag.Get("recursive") == "yes" {
			var recursiveField reflect.Value
			if dStructValue.Kind() == reflect.Invalid {
				flatField(reflect.New(field.Type).Elem().Interface(), fields, values)
			} else {
				recursiveField = dStructValue.Field(i)
				if recursiveField.CanAddr() {
					flatField(recursiveField.Addr().Interface(), fields, values)
				} else {
					flatField(recursiveField.Interface(), fields, values)
				}
			}
			continue
		}

		*fields = append(*fields, &field)
		if dStructValue.Kind() == reflect.Invalid {
			*values = append(*values, nil)
		} else {
			value := dStructValue.Field(i)
			*values = append(*values, &value)
		}
	}
}
