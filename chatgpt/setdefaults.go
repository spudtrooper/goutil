package chatgpt

import (
	"reflect"
)

func SetDefaults(obj interface{}) {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	fillStruct(t, v)
}

func fillStruct(t reflect.Type, v reflect.Value) {
	if v.Kind() != reflect.Ptr && v.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !field.CanSet() {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			if field.String() == "" {
				field.SetString("the " + t.Field(i).Name)
			}
		case reflect.Slice:
			if field.IsNil() {
				sliceType := reflect.SliceOf(t.Field(i).Type.Elem())
				slice := reflect.MakeSlice(sliceType, 1, 1)
				fillStruct(t.Field(i).Type.Elem(), slice.Index(0))
				field.Set(slice)
			}
		case reflect.Map:
			if field.IsNil() {
				mapType := reflect.MapOf(t.Field(i).Type.Key(), t.Field(i).Type.Elem())
				newMap := reflect.MakeMap(mapType)
				newValue := reflect.New(t.Field(i).Type.Elem()).Elem()
				fillStruct(t.Field(i).Type.Elem(), newValue)
				newMap.SetMapIndex(reflect.ValueOf(t.Field(i).Name+" default key"), newValue)
				field.Set(newMap)
			}
		case reflect.Struct:
			fillStruct(field.Type(), field)
		case reflect.Ptr:
			if field.IsNil() {
				newStruct := reflect.New(field.Type().Elem())
				fillStruct(field.Type().Elem(), newStruct.Elem())
				field.Set(newStruct)
			} else {
				fillStruct(field.Type().Elem(), field.Elem())
			}
		}
	}
}
