package seq

import "fmt"

func ExampleSum() {
	s1 := []int{1, 2, 3, 4}
	var s2 []int
	fmt.Println(Sum(FromSlice(s1)))
	fmt.Println(Sum(FromSlice(s2)))
	// Output:
	// 10
	// 0
}

func ExampleProduct() {
	s1 := []int{1, 2, 3, 4}
	var s2 []int
	fmt.Println(Product(FromSlice(s1)))
	fmt.Println(Product(FromSlice(s2)))
	// Output:
	// 24
	// 1
}

func ExampleAverage() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(Average(FromSlice(s1)))
	// Output:
	// 3
}

func ExampleMax() {
	s1 := []int{1, 2, 3}
	s2 := []string{"Huey", "Dewey", "Louie"}
	var s3 []int
	fmt.Println(Max(FromSlice(s1)))
	fmt.Println(Max(FromSlice(s2)))
	fmt.Println(Max(FromSlice(s3)))
	// Output:
	// 3 true
	// Louie true
	// 0 false
}

func ExampleMin() {
	s1 := []int{1, 2, 3}
	s2 := []string{"Huey", "Dewey", "Louie"}
	var s3 []int
	fmt.Println(Min(FromSlice(s1)))
	fmt.Println(Min(FromSlice(s2)))
	fmt.Println(Min(FromSlice(s3)))
	// Output:
	// 1 true
	// Dewey true
	// 0 false
}
