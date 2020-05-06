package utils

import (
	"reflect"
	"strings"
)

// All fields marked with the tag `hidden:"true"` will be set to their zero value.
// For example, this allows to send user models to the client without their password field.
func OmitHiddenFields(model interface{}) {
	if model == nil {
		return
	}
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		if t.Kind() == reflect.Struct {
			value := reflect.ValueOf(model).Elem()
			for i := 0; i < t.NumField(); i++ {
				field := value.Field(i)
				fieldType := t.Field(i)

				if !field.CanSet() {
					continue
				}

				if field.Kind() == reflect.Struct && fieldType.Anonymous {
					OmitHiddenFields(field.Addr().Interface())
					continue
				}

				tagName := strings.Split(fieldType.Tag.Get("hidden"), ";")
				if ContainsString(tagName, "true") {
					field.Set(reflect.Zero(fieldType.Type))
				}
			}
		}
	}
}
