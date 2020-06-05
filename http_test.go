package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestGetStringFromUrl(t *testing.T) {
	u, err := url.Parse("http://www.baidu.com/search?q=dotnet")

	assert.Nil(t, err)

	mq := u.Query()
	q, ret := GetStringFromUrl(mq, "q")
	assert.True(t, ret)
	fmt.Println(q)
}

func TestGetArrayFromUrl(t *testing.T) {
	u, err := url.Parse("http://www.baidu.com/search?q[]=dotnet&q[]=com")
	assert.Nil(t, err)
	mq := u.Query()
	//fmt.Println(mq)
	q, ret := GetArrayFromUrl(mq, "q")
	assert.True(t, ret)
	assert.Equal(t, "dotnet", q[0])
	assert.Equal(t, "com", q[1])
}

func TestGetMapFromUrl(t *testing.T) {
	u, err := url.Parse("http://www.baidu.com/search?q[a]=dotnet&q[b]=com")
	assert.Nil(t, err)
	mq := u.Query()
	q, ret := GetMapFromUrl(mq, "q")
	assert.True(t, ret)
	assert.Equal(t, "dotnet", q["a"])
	assert.Equal(t, "com", q["b"])
}

func TestGetPartFilterTrimOrSemicolon(t *testing.T) {
	s := "application/html;"
	s_ret := GetPartFilterTrimOrSemicolon(s)
	assert.Equal(t, s[:len(s)-1], s_ret)

	s = "application/html"
	s_ret = GetPartFilterTrimOrSemicolon(s)
	assert.Equal(t, s, s_ret)

}

func TestUrlPartErrorException(t *testing.T) {

	u, _ := url.Parse("http://www.baidu.com/search?q=dotnet")
	// get string
	q, ret := GetStringFromUrl(nil, "q")
	assert.False(t, ret)
	assert.Equal(t, "", q)

	q, ret = GetStringFromUrl(u.Query(), "not exsit key")
	assert.False(t, ret)
	assert.Equal(t, "", q)

	// get array
	qs, ret := GetArrayFromUrl(nil, "no array key")
	assert.False(t, ret)
	assert.Nil(t, qs)

	qs, ret = GetArrayFromUrl(u.Query(), "no array key")
	assert.False(t, ret)
	assert.Nil(t, qs)
}
