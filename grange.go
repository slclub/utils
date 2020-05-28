package utils

import (
//"fmt"
//"sync"
//"fmt"
)

// go generator.
// generato numbers can be used by range loop.
// Get numbers iteratively

func grange(start int, after int) chan int {
	var ch = make(chan int, 1)
	plug := 1
	if after < start {
		//start, after = after, start
		plug = -1
	}
	//fmt.Println("I am runing!!!!!!!!")

	go func() {
		i := start
		for {
			if i == after {
				break
			}
			ch <- i
			i += plug
		}
		close(ch)
	}()

	return ch
}
