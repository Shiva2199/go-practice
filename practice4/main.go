package main

import "fmt"

func main() {
	defer cleanup()

	counter := count()
	fmt.Println(counter())

	fmt.Println(counter())
	fmt.Println("close the main thread")
	add := sum(4, 5)
	fmt.Println(add)
	fmt.Println(add)

}
func cleanup() {
	fmt.Println("cleanup done")
}

func count() func() int {
	count := 0
	//function clouser
	return func() int {
		count++
		return count
	}
}

var c int

func sum(a, b int) int {
	c += a + b
	return c
}
