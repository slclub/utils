package utils

import (
	"reflect"
	"runtime"
)

var EOL = "\n"

func init() {
	if !IsUnix() {
		EOL = "\r\n"
	}
}

// Judging unix system.
// Unix*
func IsUnix() bool {
	sys_type := runtime.GOOS

	if sys_type == "windows" {
		return false
	}
	return true
}

// Verify that different types are empty. especially interface{}
// int 0 is return true
func Empty(v interface{}) bool {
	if v == nil {
		return true
	}
	if num, ok := v.(string); ok {
		return num == ""
	}

	if num, ok := ForceInt64(v); ok {
		return num == 0
	}

	if s, ok := v.([]interface{}); ok {
		return len(s) == 0
	}

	if s, ok := v.([]int); ok {
		return len(s) == 0
	}
	if s, ok := v.([]string); ok {
		return len(s) == 0
	}

	return false
}

func ForceInt64(value interface{}) (int64, bool) {
	if v, ok := value.(int); ok {
		return int64(v), true
	}

	if v, ok := value.(int64); ok {
		return v, true
	}

	if v, ok := value.(int32); ok {
		return int64(v), true
	}

	if v, ok := value.(int16); ok {
		return int64(v), true
	}

	if v, ok := value.(int8); ok {
		return int64(v), true
	}

	if v, ok := value.(uint64); ok {
		return int64(v), true
	}
	if v, ok := value.(uint32); ok {
		return int64(v), true
	}
	if v, ok := value.(uint16); ok {
		return int64(v), true
	}
	if v, ok := value.(uint8); ok {
		return int64(v), true
	}

	if v, ok := value.(rune); ok {
		return int64(v), true
	}

	if v, ok := value.(uint); ok {
		return int64(v), true
	}

	return 0, false
}

// ==========================================InSlice=======================================
// in array check
// Either nil returns failure
// @return	@1	success or fail
// @return  @2	index of item.
func InSlice(a interface{}, b interface{}) (bool, int) {
	index := -1

	// speed up []int searching.
	index = InSliceIntQuickDynamic(a, b)
	if index >= 0 {
		return true, index
	}
	// speed up []string searching.
	index = InSliceStringDynamic(a, b)
	if index >= 0 {
		return true, index
	}

	// speed up []interface{} searching.
	index = InSliceInterfaceDynamic(a, b)
	if index >= 0 {
		return true, index
	}

	// It includes all of the above situations.
	index = InSliceIntDynamic(a, b)
	if index >= 0 {
		return true, index
	}
	return false, -1
}

// If exist return the key; else return -1
// String slice search.
// @param	@1	 []string	lookup slice
// @param	@2	 string		lookup value.
// @return	int	 index of the value.
func InSliceString(arr []string, value string) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}
func InSliceStringDynamic(a interface{}, b interface{}) int {
	if _, ok := b.([]string); ok {
		a, b = b, a
	}
	v, ok := a.([]string)
	if !ok {
		return -1
	}
	if vb, okb := b.(string); okb {
		index := InSliceString(v, vb)
		return index
	}

	return -1
}

// int
func InSliceInt(arr []int, value int) int {
	for i, v := range arr {
		if (v) == value {
			return i
		}
	}
	return -1
}

// It's not necessary to exist for performance
func InSliceIntQuickDynamic(a interface{}, b interface{}) int {
	if _, ok := b.([]int); ok {
		a, b = b, a
	}
	v, ok := a.([]int)
	if !ok {
		return -1
	}
	if vb, okb := b.(int); okb {
		index := InSliceInt(v, vb)
		return index
	}

	return -1

}

//删除切片
//func DeleteSlice(slice interface{}, index int) interface{} {
//	sliceValue := reflect.ValueOf(slice)
//	length := sliceValue.Len()
//	if slice == nil || length == 0 || (length-1) < index {
//		return nil
//	}
//	if length-1 == index {
//		return sliceValue.Slice(0, index).Interface()
//	} else if (length - 1) >= index {
//		return reflect.AppendSlice(sliceValue.Slice(0, index), sliceValue.Slice(index+1, length)).Interface()
//	}
//	return nil
//}

// all int
func InSliceIntDynamic(a interface{}, b interface{}) int {
	int64_b, ok := ForceInt64(b)
	if !ok {
		a, b = b, a
	}
	int64_b, ok = ForceInt64(b)
	if !ok {
		return -1
	}
	if !IS_SLICE(a) {
		return -1
	}

	slice_value := reflect.ValueOf(a)
	length := slice_value.Len()

	for i := 0; i < length; i++ {
		int_tmp := slice_value.Index(i).Interface()
		if v, ret := ForceInt64(int_tmp); ret && int64_b == v {
			return i
		}
	}

	return -1
}

func InSliceInterfaceDynamic(a interface{}, b interface{}) int {
	if v, ok := b.([]interface{}); ok {
		return InSliceInterface(v, a)
	}

	if v, ok := a.([]interface{}); ok {
		return InSliceInterface(v, b)
	}

	return -1
}
func InSliceInterface(arr []interface{}, value interface{}) int {
	if value == nil {
		return -1
	}

	value_int, value_ret := ForceInt64(value)

	for i, v := range arr {
		// string check
		if vv, ok := value.(string); ok {
			if vv2, ok2 := v.(string); ok2 && vv2 == vv {
				return i
			}
			continue
		}
		// int
		if vv, ok := ForceInt64(value); ok && value_ret && value_int == vv {
			continue
		}

	}
	return -1
}

// ==========================================InSlice=======================================
