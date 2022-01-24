package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func lenAndUpper2(name string) (length int, uppercase string) {
	defer fmt.Println("im done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	// for index, number := range numbers {
	// 	fmt.Println(index, number)
	// }

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	// fmt.Println(multiply(2, 2))
	// totallength, _ := lenAndUpper("nico")
	// fmt.Println(totallength)

	// a, b := lenAndUpper2("coco")
	// fmt.Println(a, b)

	// repeatMe("nico", "lynn", "ann")

	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)

}
