package main

import (
	"fmt"
)

func main() {

	fmt.Println(plusOne[int64](10))
	fmt.Println(plusOne[float64](10.30))
}

func plusOne[T int | float64 | int64 | float32](t T) T {
	return t * t
}
