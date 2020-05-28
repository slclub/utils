package utils

import (
	"fmt"
	"testing"
)

func TestGrange(t *testing.T) {
	range_grange()
}

// here we just print all number of the iterator.
func range_grange() {
	fmt.Println("[TEST][GRANGE] 0 -5")
	for i := range grange(0, 5) {
		fmt.Println("[TEST][GRANGE][STEP]", i)
	}
}
