/*
 * @Author: Coller
 * @Date: 2022-04-29 17:16:35
 * @LastEditTime: 2024-04-21 14:49:04
 * @Desc: 数据转换
 */
package scanx

import (
	"reflect"
)

/**
 * @desc: value 原始值 dst 赋值
 * @return {*}
 */
func StructCopy(value interface{}, dst interface{}) {
	vVal := reflect.ValueOf(value).Elem() //获取reflect.Type类型
	bVal := reflect.ValueOf(dst).Elem()   //获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		fi := vTypeOfT.Field(i)
		name := fi.Name
		if ok := bVal.FieldByName(name).IsValid(); ok {
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
}

// 结构体转map
func StructToMap(val interface{}, tagName string) map[string]interface{} {
	if tagName == "" {
		tagName = "json"
	}
	out := make(map[string]interface{})
	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct { // 判断是不是结构体
		return nil
	}
	t := v.Type()
	// 遍历结构体字段
	// 指定tag值为map中的k，字段值为map中的val
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if fi.Anonymous {
			for ii := 0; ii < v.Field(i).NumField(); ii++ {
				if tagValue := t.Field(i).Type.Field(ii).Tag.Get(tagName); tagValue != "" {
					out[tagValue] = v.Field(i).Field(ii).Interface()
				}
			}
		} else {
			if tagValue := fi.Tag.Get(tagName); tagValue != "" {
				out[tagValue] = v.Field(i).Interface()
			}
		}
	}
	return out
}

// func StructToMap(obj interface{}) map[string]interface{} {
// 	result := make(map[string]interface{})
// 	value := reflect.ValueOf(obj).Elem() // 获取指针的值
// 	typ := value.Type()                  // 获取类型信息

// 	for i := 0; i < typ.NumField(); i++ {
// 		field := typ.Field(i)              // 获取字段信息
// 		tag := field.Tag.Get("json")       // 获取标签（如果有）
// 		if tag != "" && !field.Anonymous { // 只处理非匿名字段且有标签的情况
// 			key := field.Name // 默认使用字段名作为Key
// 			if tag != "-" {   // 若标签不等于-则使用标签作为Key
// 				key = tag
// 			}
// 			result[key] = value.Field(i).Interface() // 存入Map
// 		}
// 	}
// 	return result
// }
