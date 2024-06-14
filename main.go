package main

import (
	"fmt"
	"time"
)

func main() {
	C := 1500
	E := 1000

	ts := []int{}
	t := 0
	for {
		time.Sleep(5 * time.Millisecond)
		t += 1
		ts = append(ts, t)
		if t%C == 0 {
			i := 0
			for ; i < len(ts); i++ {
				if ts[i]+E <= t {
					i++
				} else {
					break
				}
			}
			fmt.Println(float32(i) / float32(len(ts)))
			ts = ts[i:]
		}
	}
}
