package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrange(t *testing.T) {
	range_grange()

	// check last item value.
	var i int
	for j := range grange(0, -5) {
		i = j
	}
	fmt.Println("testing grange", i)
	assert.Equal(t, -4, i)
}

// here we just print all number of the iterator.
func range_grange() {
	fmt.Println("[TEST][GRANGE] 0 -5")
	for i := range grange(0, 5) {
		fmt.Println("[TEST][GRANGE][STEP]", i)
	}
}
