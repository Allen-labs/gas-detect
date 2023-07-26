package util

import "reflect"

// StructCopy  复制一个指针类型struct到另外一个指针类型struct
func StructCopy(binding, value interface{}) {
	bVal := reflect.ValueOf(binding).Elem()
	vVal := reflect.ValueOf(value).Elem()
	//获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		if bVal.FieldByName(name).IsValid() && !vVal.Field(i).IsZero() {
			if vVal.Field(i).Type().Kind().String() == "ptr" {
				//if bVal.Field(i).Type().Kind().String() == "ptr" {
				if bVal.FieldByName(name).Type().Kind().String() == "ptr" {
					bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
				} else {
					bVal.FieldByName(name).Set(vVal.Field(i).Elem())
				}
			} else {
				//if bVal.Field(i).Type().Kind().String() == "ptr" {
				if bVal.FieldByName(name).Type().Kind().String() == "ptr" {
					bVal.FieldByName(name).Set(vVal.Field(i).Addr())
				} else {
					bVal.FieldByName(name).Set(vVal.Field(i))
				}
			}
		}
	}
}
