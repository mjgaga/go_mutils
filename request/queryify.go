package request

import (
	"fmt"
	"mutils/util"
	"net/url"
	"reflect"
	"strings"
	"time"
)

func Quertify(i interface{}) string {
	rt := reflect.TypeOf(i)
	rv := reflect.ValueOf(i)

	if rt.Kind() == reflect.Ptr {
		return Quertify(rv.Elem().Interface())
	}

	if rt.Kind() == reflect.Struct {
		return structQuertify(i)
	} else if rt.Kind() == reflect.Map {
		return mapQuertify(i)
	}

	return ""
}

func structQuertify(i interface{}) string {
	fields, values := util.FlatField(i)

	fmap := make(map[string]interface{})
	for i, field := range fields {
		fName := field.Tag.Get("json")
		if fName == "" {
			fName = field.Name
		}
		if values[i] != nil {
			fmap[fName] = values[i].Interface()
		} else {
			fmap[fName] = nil
		}
	}

	return mapQuertify(fmap)
}

func mapQuertify(i interface{}) string {
	fmap := i.(map[string]interface{})
	strList := []string{}
	for item, v := range fmap {
		structValue := reflect.ValueOf(v)
		if v == nil || (structValue.Kind() == reflect.Ptr && structValue.IsNil()) {
			continue
		}

		strList = append(strList, fmt.Sprintf("%s=%s", item, url.QueryEscape(getInterfaceValue(v))))
	}

	return strings.Join(strList, "&")
}

func getInterfaceValue(i interface{}) string {
	structValue := reflect.ValueOf(i)
	if structValue.Kind() == reflect.Ptr {
		if !structValue.IsNil() {
			return getInterfaceValue(structValue.Elem().Interface())
		} else {
			return "null"
		}
	}

	switch i.(type) {
	case *time.Time, time.Time:
		var t time.Time
		if structValue.Kind() == reflect.Ptr {
			t = *(i.(*time.Time))
		} else {
			t = i.(time.Time)
		}

		text, _ := t.MarshalText()
		return string(text)
	}

	valueKind := structValue.Kind()
	switch valueKind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%d", i)
	// case reflect.String:
	// 	return fmt.Sprintf("%s", i)
	// case reflect.Bool:
	// 	return fmt.Sprintf("%b", i)
	default:
		return fmt.Sprintf("%s", i)
	}
}
