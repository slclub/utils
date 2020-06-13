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
