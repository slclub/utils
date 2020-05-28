package utils

import (
	//"fmt"
	"reflect"
	"runtime"
	//"strconv"
)

// this file use reflect package.
// so these function have a bad performance.
// so we captial the leters of func name and start with "_".

// Get function name.
func _FUNC_NAME(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// Check the param is or not an function.
func _IS_FUNC(fn interface{}) bool {

	if reflect.ValueOf(fn).Kind() == reflect.Func {
		return true
	}

	return false
}

// check type struct
// 判断是结构体
func _IS_STRUCT(stu interface{}) bool {
	if reflect.ValueOf(stu).Kind() == reflect.Struct {
		return true
	}
	return false
}

// check is slice
func _IS_SLICE(s interface{}) bool {
	if reflect.ValueOf(s).Kind() == reflect.Slice {
		return true
	}
	return false
}
