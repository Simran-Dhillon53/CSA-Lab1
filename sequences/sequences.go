package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) {
	slice = append(slice, slice...)
	fmt.Println(slice)
}

func mapSlice(f func(a int) int, slice []int) {
	for i := range slice {
		slice[i] = f(slice[i])
	}
}

func mapArray(f func(a int) int, array *[5]int) {

	for i := range array {
		array[i] = f(array[i])
	}
}

// Do not understand question 3e(had ti put print into the function to double it).
// Question 3f explain question(
// -explain how arrays and slices are passed into functions
// explain how append() works
// give use cases for arrays and slices
func main() {
	s := []int{1, 2, 3, 4, 5}
	newslice := s[1:3]
	//mapSlice(addOne, s)
	//fmt.Println(newslice)
	//fmt.Println(s)
	double(newslice)
	//double(s)
	//mapSlice(double, newslice)
	//fmt.Println(newslice)
	//fmt.Println(s)

	//regular array
	a := [5]int{1, 2, 3, 4, 5}

	//pointer to array
	//var ptr *[3]int = &a

	mapArray(addOne, &a)

	//fmt.Println(s)
	//fmt.Println(a)
}
