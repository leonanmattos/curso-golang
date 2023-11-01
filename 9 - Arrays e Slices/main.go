package main

import (
	"fmt"
	"reflect"
)

func main() {

	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]string{"txt1", "txt2", "txt3", "txt4", "txt5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

}
