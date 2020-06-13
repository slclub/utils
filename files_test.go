package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileExist(t *testing.T) {
	dir_path := "/tmp"
	ret, err := IsFileExist(dir_path)

	assert.True(t, ret)
	assert.Nil(t, err)

	dir_path = "/tmp/nofile"
	ret, err = IsFileExist(dir_path)
	fmt.Println("[UTIL][FILE][EXIST]", err)
	assert.False(t, ret)
}

func TestReadAll(t *testing.T) {

	s1, ok := ReadAll("/tmp/noexist")
	assert.False(t, ok)
	assert.Empty(t, s1)

	s1, ok = ReadAll("files.go")
	assert.True(t, ok)
	assert.NotEmpty(t, s1)
}
