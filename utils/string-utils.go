/*
プロジェクト全体で使用するユーティリティ関数群
*/
package utils

import "reflect"

// 構造体をマップに変換する
//
// Parameters:
//   - obj: 構造体
//
// Returns:
//   - map[string]interface{}: マップ
func StructToMap(obj interface{}) map[string]interface{} {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		panic("変換対象が構造体ではありません")
	}
	n := v.NumField()
	m := make(map[string]interface{})
	for i := 0; i < n; i++ {
		field := v.Type().Field(i)
		name := field.Name
		value := v.Field(i).Interface()
		m[name] = value
	}
	return m
}
