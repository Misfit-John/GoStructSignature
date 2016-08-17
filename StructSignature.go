package StructSignature

import (
	"crypto/md5"
	"fmt"
	"reflect"
)

//GetSignature accept any input and will generate a md5 signature in string for the stuct
func GetSignature(i interface{}) string {
	refleType := reflect.TypeOf(i)
	searchResult := searchInterface(refleType)
	data := []byte(searchResult)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func searchInterface(i reflect.Type) string {
	switch i.Kind() {
	case reflect.Func:
		return ""
	case reflect.Ptr:
		return fmt.Sprintf("%d:%s", i.Kind(), searchInterface(i.Elem()))
	case reflect.Struct:
		ret := ""
		for index := 0; index < i.NumField(); index++ {
			field := i.Field(index)
			ret += fmt.Sprintf("%s:%s", field.Name, searchInterface(field.Type))
		}
		return fmt.Sprintf("%s:%s", i.String(), ret)
	case reflect.Map:
		keyType := searchInterface(i.Key())
		valueType := searchInterface(i.Elem())
		return fmt.Sprintf("%s_%s", keyType, valueType)
	case reflect.Slice, reflect.Array:
		valueType := searchInterface(i.Elem())
		return fmt.Sprintf("%d:%s", i.Kind(), valueType)
	default:
		return fmt.Sprintf("%d", i.Kind())
	}
}
