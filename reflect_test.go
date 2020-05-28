package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var _domain = "github.com/slclub/utils"

func Test_FUNC_NAME(t *testing.T) {
	fn := _FUNC_NAME(needGetMyNameFunc)
	assert.Equal(t, _domain+".needGetMyNameFunc", fn)
}

func Test_IS_FUNC(t *testing.T) {
	fn := needGetMyNameFunc
	assert.True(t, _IS_FUNC(fn))
	fn2 := "string"
	assert.False(t, _IS_FUNC(fn2))
}

func Test_IS_STRUCT(t *testing.T) {
	u1 := stu{}
	assert.True(t, _IS_STRUCT(u1))
	assert.False(t, _IS_STRUCT(2))
	assert.False(t, _IS_STRUCT("I am not "))
}

func needGetMyNameFunc() {}

type stu struct{}
