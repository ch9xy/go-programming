package main

import (
	"fmt"
)

// First we need a summation function

func sum(xi ...int) int {
	var total int
	for _, v := range xi {
		total = total + v
	}
	return total
}

// Second we need even number control and and we are going to use the summing function above us.


// To be able to use the sum() function we are building the infrastructure by using an anonymous func as an parameter.
func even(f func(xi ...int) int, y ...int) int {
	var total_Slice []int
	for _, v := range y {
		if v%2 == 0 {
			total_Slice = append(total_Slice, v)
		}
	}
	total_Val := f(total_Slice...)
	return total_Val
}

func main() {


	//		TEST STUFF
	//a := []int{1, 2, 3}
	//fmt.Println(sum(a...))
	//		TEST STUFF


	
	t := even(sum, []int{1,2,3,4,5,6,7,8,9}...)
	fmt.Println(t)
}
