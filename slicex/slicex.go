/*
 * @Author: coller
 * @Date: 2024-03-07 12:35:39
 * @LastEditors: coller
 * @LastEditTime: 2024-04-21 14:49:18
 * @Desc: 参考 github.com/duke-git/lancet/v2/slice
 */
package slicex

import (
	"fmt"
	"reflect"
)

func IntSlice(slice any) []int {
	sv := sliceValue(slice)

	result := make([]int, sv.Len())
	for i := 0; i < sv.Len(); i++ {
		v, ok := sv.Index(i).Interface().(int)
		if !ok {
			panic("invalid element type")
		}
		result[i] = v
	}

	return result
}

func sliceValue(slice any) reflect.Value {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Invalid slice type, value of type %T", slice))
	}
	return v
}

// InterfaceSlice convert param to slice of interface.
// This function is deprecated, use generics feature of go1.18+ for replacement.
// Play: https://go.dev/play/p/FdQXF0Vvqs-
func InterfaceSlice(slice any) []any {
	sv := sliceValue(slice)
	if sv.IsNil() {
		return nil
	}

	result := make([]any, sv.Len())
	for i := 0; i < sv.Len(); i++ {
		result[i] = sv.Index(i).Interface()
	}

	return result
}

// StringSlice convert param to slice of string.
// This function is deprecated, use generics feature of go1.18+ for replacement.
// Play: https://go.dev/play/p/W0TZDWCPFcI
func StringSlice(slice any) []string {
	v := sliceValue(slice)

	result := make([]string, v.Len())
	for i := 0; i < v.Len(); i++ {
		v, ok := v.Index(i).Interface().(string)
		if !ok {
			panic("invalid element type")
		}
		result[i] = v
	}

	return result
}
