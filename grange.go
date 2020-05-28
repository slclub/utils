package utils

import (
	//"fmt"
	"sync"
)

// go generator.
// generato numbers can be used by range loop.
// Get numbers iteratively
var do_once_range sync.Once

func grange(start int, after int) chan int {
	var ch = make(chan int, 1)
	if after < start {
		start, after = after, start
	}
	f1 := func() {

		go func(start int, after int) {
			for i := start; i < after; i++ {
				ch <- i
			}
			close(ch)
		}(start, after)
	}

	do_once_range.Do(f1)
	return ch
}
