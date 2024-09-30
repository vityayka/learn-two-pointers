package main

import (
	"fmt"
	"runtime"
	"time"

	"math/rand"
)

//Given n non-negative integers representing the height of vertical lines,
//find two lines that together with the x-axis form a container that holds the most water.
//Example:
//Input: height = [1, 8, 6, 2, 5, 4, 8, 3, 7]
//Output: 49 (between indices 1 and 8)

func main() {

	const cnt = 100_00
	var input [cnt]int

	start := time.Now()

	for i := 0; i < cnt; i++ {
		input[i] = 2*rand.Intn(cnt) - cnt
	}

	fmt.Println("filling the array: ", time.Since(start))
	start = time.Now()

	slice := input[:]

	// slice := []int{-1, 0, 1, 2, -1, -4}

	// slices.Sort(slice)
	// fmt.Println("sorting the slice: ", time.Since(start))
	// start = time.Now()

	// i, j := targetSum(input, 18888)

	// rmDup(slice)
	sumThree(slice, 0)

	// fmt.Printf("%v\n", sum)
	fmt.Println("calculating sum of three: ", time.Since(start))
	// fmt.Println("rm the dups: ", time.Since(start))
	PrintMemUsage()

	// fmt.Println(output)

	// square, leftStick, rightStick, lI, rI := findBiggest(input)
	// fmt.Printf("square=%d, leftStick=%d, rightStick=%d, bottomLen=%d\n", square, leftStick, rightStick, rI-lI)

}

func findBiggest(input []int) (int, int, int, int, int) {
	biggestSquare, biggestI, biggestJ, biggestL, biggestR := 0, 0, 0, 0, 0
	for i, j := 0, len(input)-1; i < j; {
		square := calcSquare(input[i : j+1])
		if square > biggestSquare {
			biggestSquare = square
			biggestL, biggestR, biggestI, biggestJ = input[i], input[j], i, j
		}

		if input[i] < input[j] {
			i++
		} else {
			j--
		}
	}

	return biggestSquare, biggestL, biggestR, biggestI, biggestJ
}

func calcSquare(arr []int) int {
	lastIndex := len(arr) - 1
	littleEdge := arr[0]
	if arr[0] > arr[lastIndex] {
		littleEdge = arr[lastIndex]
	}

	return littleEdge * lastIndex
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
