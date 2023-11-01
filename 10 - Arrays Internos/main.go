package main

import "fmt"

func main() {
	slice1 := make([]float32, 10, 50)
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
}
