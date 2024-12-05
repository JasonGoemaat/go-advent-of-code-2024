package day02

import (
	"fmt"
	"slices"

	"github.com/JasonGoemaat/go-advent-of-code-2024/util"
)

func IsSafeIncreasing(values []int) bool {
	mylog("  IsSafeIncreasing(): %v\n", values)
	for i, value := range values {
		if i > 0 && (value <= values[i-1] || value > values[i-1]+3) {
			mylog("    Fails at %d comparing %d and %d\n", i, value, values[i-1])
			return false
		}
	}
	mylog("  IsSafeIncreasing() returning true")
	return true
}

func IsSafeDecreasing(values []int) bool {
	mylog("  IsSafeDecreasing(): %v\n", values)
	for i, value := range values {
		if i > 0 && (value >= values[i-1] || value < values[i-1]-3) {
			mylog("    Fails at %d comparing %d and %d\n", i, value, values[i-1])
			return false
		}
	}
	mylog("  IsSafeDecreasing() returning true")
	return true
}

// https://stackoverflow.com/a/37335777/369792
// THIS IS TRASH - it affects the original array
func remove_EVIL(slice []int, i int) []int {
	fmt.Printf("REMOVE(%d): %v\n", i, slice)
	a := slice[:i]
	fmt.Printf("   a: %v\n", a)
	b := slice[i+1:]
	fmt.Printf("   a: %v\n", b)
	result := append(a, b...)
	fmt.Printf("   result: %v\n", result)
	return result
}

func remove(slice []int, i int) []int {
	a := slices.Clone(slice)
	return slices.Delete(a, i, i+1)
}

func testIncreasing(a, b int) bool {
	if a >= b {
		return false
	}
	if (b - a) > 3 {
		return false
	}
	return true
}

func IsSafeIncreasingLenient(values []int) bool {
	mylog("IsSafeIncreasingLenient(): %v\n", values)
	for i, value := range values {
		if i > 0 && ((value <= values[i-1]) || (value > values[i-1]+3)) {
			// if i > 0 && !testIncreasing(values[i-1], value) {
			mylog("  failure at %d comparing %d and %d\n", i, value, values[i-1])
			mylog("  (value <= values[%d-1]) or (%d <= %d) is: %v", i, value, values[i-1], value <= values[i-1])
			mylog("  (value > values[%d-1]-3) or (%d > %d) is: %v", i, value, values[i-1], value > values[i-1]-3)
			// try removing index `i` and `i-1` and don't be lenient
			if IsSafeIncreasing(remove(values, i-1)) {
				mylog("  returning true, removing %d worked\n", i-1)
				return true
			}
			if IsSafeIncreasing(remove(values, i)) {
				mylog("  returning true, removing %d worked\n", i)
				return true
			}
			mylog("  returning false, removing %d and %d both failed\n", i, i-1)
			return false
		}
	}
	return true
}

func IsSafeDecreasingLenient(values []int) bool {
	mylog("IsSafeDecreasingLenient(): %v\n", values)
	for i, value := range values {
		if i > 0 && (value >= values[i-1] || value < values[i-1]-3) {
			// try removing index `i` and `i-1` and don't be lenient
			mylog("  failure at %d comparing %d and %d\n", i, value, values[i-1])
			if IsSafeDecreasing(remove(values, i-1)) {
				mylog("  returning true, removing %d worked\n", i-1)
				return true
			}
			if IsSafeDecreasing(remove(values, i)) {
				mylog("  returning true, removing %d worked\n", i)
				return true
			}
			mylog("  returning false, removing %d and %d both failed\n", i, i-1)
			return false
		}
	}
	mylog("  returning true, no change needed\n")
	return true
}

func Solve(filePath string) int {
	numbers := util.LoadNumbers(filePath)
	safeCount := 0
	for _, row := range numbers {
		if IsSafeIncreasing(row) || IsSafeDecreasing(row) {
			safeCount++
		}
	}
	return safeCount
}

func Solve2(filePath string) int {
	numbers := util.LoadNumbers(filePath)
	safeCount := 0
	for _, row := range numbers {
		mylog("\n")
		if IsSafeIncreasingLenient(row) || IsSafeDecreasingLenient(row) {
			safeCount++
		}
	}
	return safeCount
}

func mylog(format string, args ...interface{}) {
	// fmt.Printf(format, args...)
}
