package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"

	"gopl.io/ch7/eval"
)

func main() {
	e, _ := eval.Parse("sqrt(A / pi)")
	Display("e", e)
}

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

// ! Notice that even unexported fields are visible to reflection.
func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		// iter := v.MapRange()
		// for iter.Next() {
		// 	k := iter.Key()
		// 	display(fmt.Sprintf("%s%s", path, formatAtom(k)), k)

		// 	v := iter.Value()
		// 	display(fmt.Sprintf("%s[%s]", path, formatAtom(k)), v)
		// }
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatKey(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		//? This operation would be safe even if the pointer value is nil,
		//? in which case the result would have kind Invalid, but we use IsNil to detect
		//? nil pointers explicitly so we can print a more appropriate message.
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

// formatAtom formats a value without inspecting its internal structure.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct,		reflect.Interface
		return v.Type().String() + " value"
	}
}

func formatKey(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Struct:
		b := &bytes.Buffer{}
		b.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			if i != 0 {
				b.WriteString(", ")
			}
			fmt.Fprintf(b, "%s: %s", v.Type().Field(i).Name, formatAtom(v.Field(i)))
		}
		b.WriteByte('}')
		return b.String()
	case reflect.Array:
		b := &bytes.Buffer{}
		for i := 0; i < v.Len(); i++ {
			if i != 0 {
				b.WriteString(", ")
			}
			b.WriteString(formatAtom(v.Index(i)))
		}
		return b.String()
	default:
		return formatAtom(v)
	}
}
