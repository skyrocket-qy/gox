package main

import (
	"fmt"
)

func HeavyWork() {
	x := make([]int, 0)
	for i := 0; i < 100000; i++ {
		x = append(x, i*i)
	}
}

func main() {
	result, err := ProfileFunc(HeavyWork, true, "cpu.prof", 10000)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Time (ns):       %d\n", result.ElapsedNS)
	fmt.Printf("Alloc Delta:     %d bytes\n", result.MemAllocDelta)
	fmt.Printf("GC Cycles:       %d\n", result.NumGC)
	fmt.Printf("CPU Profile File: %s\n", result.CPUProfileFile)
}
