package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmpty(t *testing.T) {
	// true
	assert.True(t, Empty(0))
	assert.True(t, Empty(uint(0)))
	assert.True(t, Empty(""))
	assert.True(t, Empty(nil))
	assert.True(t, Empty(int8(0)))
	assert.True(t, Empty(int16(0)))
	assert.True(t, Empty(int32(0)))
	assert.True(t, Empty(int64(0)))
	assert.True(t, Empty(uint64(0)))
	assert.True(t, Empty(uint32(0)))
	assert.True(t, Empty(uint16(0)))
	assert.True(t, Empty(uint8(0)))
	assert.True(t, Empty(byte(0)))
	assert.True(t, Empty((0)))
	assert.True(t, Empty(int(0)))
	assert.True(t, Empty([]int{}))
	assert.True(t, Empty([]string{}))

	// false
	assert.False(t, Empty("dfdf"))
	assert.False(t, Empty(123))
	assert.False(t, Empty([]int{2323, 32}))
	assert.False(t, Empty([]string{"i", "32"}))

}

// testing slice
func TestInSlice(t *testing.T) {
	s1 := []int{1, 2, 3}
	ok1, sv1 := InSlice(s1, 1)
	not_ok1, _ := InSlice(s1, 5)
	assert.Equal(t, sv1, 0)
	assert.Equal(t, ok1, true)
	assert.False(t, not_ok1)

	s2 := []string{"1", "2", "3"}
	ok2, sv2 := InSlice(s2, "1")
	ok2_1, sv2_1 := InSlice("2", s2)
	ok2_2, sv2_2 := InSlice(nil, s2)
	not_ok2, _ := InSlice(s1, 5)
	assert.Equal(t, sv2, 0)
	assert.Equal(t, ok2, true)
	assert.False(t, not_ok2)
	assert.Equal(t, true, ok2_1)
	assert.Equal(t, 1, sv2_1)
	assert.Equal(t, false, ok2_2)
	assert.Equal(t, -1, sv2_2)

}

func TestInSliceInterface(t *testing.T) {
	s3 := []interface{}{1, 2, 3}
	ok3, sv3 := InSlice(s3, 1)
	not_ok3, _ := InSlice(s3, 5)
	not_ok3_rev, _ := InSlice(5, s3)
	ok3_1, sv3_1 := InSlice(2, s3)
	ok3_2, sv3_2 := InSlice(nil, s3)
	assert.Equal(t, sv3, 0)
	assert.Equal(t, ok3, true)
	assert.False(t, not_ok3)
	assert.False(t, not_ok3_rev)
	assert.Equal(t, 1, sv3_1)
	assert.Equal(t, true, ok3_1)
	assert.Equal(t, -1, sv3_2)
	assert.Equal(t, false, ok3_2)

}

func TestInSliceInt64(t *testing.T) {
	s3 := []int64{1, 2, 3}
	ok3, sv3 := InSlice(s3, 1)
	not_ok3, _ := InSlice(s3, 5)
	not_ok3_rev, _ := InSlice(5, s3)
	ok3_1, sv3_1 := InSlice(2, s3)
	assert.Equal(t, sv3, 0)
	assert.Equal(t, ok3, true)
	assert.False(t, not_ok3)
	assert.False(t, not_ok3_rev)
	assert.Equal(t, 1, sv3_1)
	assert.Equal(t, true, ok3_1)

}
func TestInSliceUint64(t *testing.T) {
	s3 := []uint64{1, 2, 3}
	ok3, sv3 := InSlice(s3, 1)
	not_ok3, _ := InSlice(s3, 5)
	not_ok3_rev, _ := InSlice(5, s3)
	ok3_1, sv3_1 := InSlice(2, s3)
	assert.Equal(t, sv3, 0)
	assert.Equal(t, ok3, true)
	assert.False(t, not_ok3)
	assert.False(t, not_ok3_rev)
	assert.Equal(t, 1, sv3_1)
	assert.Equal(t, true, ok3_1)

}
func TestInSliceInt32(t *testing.T) {
	s3 := []int32{1, 2, 3}
	ok3, sv3 := InSlice(s3, 1)
	not_ok3, _ := InSlice(s3, 5)
	not_ok3_rev, _ := InSlice(5, s3)
	ok3_1, sv3_1 := InSlice(2, s3)
	assert.Equal(t, sv3, 0)
	assert.Equal(t, ok3, true)
	assert.False(t, not_ok3)
	assert.False(t, not_ok3_rev)
	assert.Equal(t, 1, sv3_1)
	assert.Equal(t, true, ok3_1)

}
func TestInSliceUInt32(t *testing.T) {
	s3 := []uint32{1, 2, 3}
	ok3, sv3 := InSlice(s3, 1)
	not_ok3, _ := InSlice(s3, 5)
	not_ok3_rev, _ := InSlice(5, s3)
	ok3_1, sv3_1 := InSlice(2, s3)
	assert.Equal(t, sv3, 0)
	assert.Equal(t, ok3, true)
	assert.False(t, not_ok3)
	assert.False(t, not_ok3_rev)
	assert.Equal(t, 1, sv3_1)
	assert.Equal(t, true, ok3_1)

}
func TestInSliceUInt16(t *testing.T) {
	s3 := []uint16{1, 2, 3}
	ok3, sv3 := InSlice(s3, 1)
	not_ok3, _ := InSlice(s3, 5)
	not_ok3_rev, _ := InSlice(5, s3)
	ok3_1, sv3_1 := InSlice(2, s3)
	assert.Equal(t, sv3, 0)
	assert.Equal(t, ok3, true)
	assert.False(t, not_ok3)
	assert.False(t, not_ok3_rev)
	assert.Equal(t, 1, sv3_1)
	assert.Equal(t, true, ok3_1)

}
func TestInSliceInt16(t *testing.T) {
	s3 := []int16{1, 2, 3}
	ok3, sv3 := InSlice(s3, 1)
	not_ok3, _ := InSlice(s3, 5)
	not_ok3_rev, _ := InSlice(5, s3)
	ok3_1, sv3_1 := InSlice(2, s3)
	assert.Equal(t, sv3, 0)
	assert.Equal(t, ok3, true)
	assert.False(t, not_ok3)
	assert.False(t, not_ok3_rev)
	assert.Equal(t, 1, sv3_1)
	assert.Equal(t, true, ok3_1)

}
func TestInSliceInt8(t *testing.T) {
	s3 := []int8{1, 2, 3}
	ok3, sv3 := InSlice(s3, 1)
	not_ok3, _ := InSlice(s3, 5)
	not_ok3_rev, _ := InSlice(5, s3)
	ok3_1, sv3_1 := InSlice(2, s3)
	assert.Equal(t, sv3, 0)
	assert.Equal(t, ok3, true)
	assert.False(t, not_ok3)
	assert.False(t, not_ok3_rev)
	assert.Equal(t, 1, sv3_1)
	assert.Equal(t, true, ok3_1)

}
func TestInSliceUInt8(t *testing.T) {
	s3 := []uint8{1, 2, 3}
	ok3, sv3 := InSlice(s3, 1)
	not_ok3, _ := InSlice(s3, 5)
	not_ok3_rev, _ := InSlice(5, s3)
	ok3_1, sv3_1 := InSlice(2, s3)
	assert.Equal(t, sv3, 0)
	assert.Equal(t, ok3, true)
	assert.False(t, not_ok3)
	assert.False(t, not_ok3_rev)
	assert.Equal(t, 1, sv3_1)
	assert.Equal(t, true, ok3_1)

}

func TestInt(t *testing.T) {
	var int_32 uint = 32
	assert_int := func(intval interface{}) bool {

		if v, ok := intval.(int); ok {
			fmt.Println("success", v, ok)
			return true
		}
		return false
	}

	assert.False(t, assert_int(int_32))
}

func TestSwapInterface(t *testing.T) {
	var a interface{} = 2
	var b interface{} = "An String"
	a, b = b, a
	fmt.Println("interface type swap 2, string", a, b)
}
