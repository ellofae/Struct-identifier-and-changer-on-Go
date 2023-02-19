package identifier

import (
	"fmt"
	"reflect"
	"strconv"
)

func StructIdentifier(v interface{}) {
	vobj := reflect.ValueOf(v)

	switch vobj.Kind() {
	case reflect.Struct:
	case reflect.Ptr:
		if vobj.Elem().Kind() != reflect.Struct {
			fmt.Printf("Type '%v' has pointer '%v' on value: %v\n", vobj.Type(), vobj, vobj.Elem())
			return
		}
		vobj = vobj.Elem()

	default:
		fmt.Printf("Type '%v' has value: %v\n", vobj.Type(), vobj)
		return
	}

	fmt.Printf("Struct of type %v and number of fields %d:\n", vobj.Type(), vobj.NumField())
	for fieldIndex := 0; fieldIndex < vobj.NumField(); fieldIndex++ {
		v := vobj.Field(fieldIndex)
		fmt.Printf("\tField %v: %v - val: %v\n", vobj.Type().Field(fieldIndex).Name, v.Type(), v)
	}
}

func StructFieldChanger(v interface{}, fieldName string, newFieldValue interface{}) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return
	}

	field := val.FieldByName(fieldName)
	newFieldValueKind := reflect.ValueOf(newFieldValue).Kind()

	if field.IsValid() {
		if field.CanSet() {
			switch field.Kind() {
			case reflect.Int:
				if newFieldValueKind == reflect.Int {
					field.SetInt(int64(newFieldValue.(int)))
					return
				}
			case reflect.String:
				if newFieldValueKind == reflect.String {
					field.SetString(newFieldValue.(string))
					return
				} else if newFieldValueKind == reflect.Int {
					field.SetString(strconv.Itoa(newFieldValue.(int)))
					return
				}
			case reflect.Ptr:
				fieldPtrElementKind := field.Elem().Kind()
				newFieldPtrKind := reflect.ValueOf(newFieldValue).Elem().Kind()

				if fieldPtrElementKind == newFieldPtrKind {
					if fieldPtrElementKind == reflect.Int && newFieldPtrKind == reflect.Int {
						field.Elem().SetInt(reflect.ValueOf(newFieldValue).Elem().Int())
						return
					} else if fieldPtrElementKind == reflect.String && newFieldPtrKind == reflect.String {
						field.Elem().SetString(reflect.ValueOf(newFieldValue).Elem().String())
						return
					}
				}

				fmt.Printf("\nERROR: Cannot assign the value because of mismathed types: %v and %v\n", newFieldValueKind, field.Kind())
				return
			}
		}
	}
}
