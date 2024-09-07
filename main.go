package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Precompute arrays outside the loop
	fizzArr := [2]string{"", "Fizz"}
	buzzArr := [2]string{"", "Buzz"}

	numArr := make([]string, 101)
	for i := 1; i <= 100; i++ {
		numArr[i] = strconv.Itoa(i)
	}

	fizzMod := []int{1, 0, 0}
	buzzMod := []int{1, 0, 0, 0, 0}

	for i := 1; i <= 100; i++ {
		fizz := fizzArr[fizzMod[i%3]]
		buzz := buzzArr[buzzMod[i%5]]

		combinedFlag := fizzMod[i%3] | buzzMod[i%5]
		fmt.Println(fizz + buzz + numArr[i*(1-combinedFlag)])
	}
}
