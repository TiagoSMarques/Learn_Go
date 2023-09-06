package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	var x int = 0
	var y int = 5
	z := 8
	var t int = x + y + z

	// long form declaration of fixed size arrays
	var a [5]int
	// Drop the var and initialize
	b := [5]int{1, 2, 3, 4, 5}
	a[2] = 6
	// Declaration of a "slice" which is an array without a fixed size
	c := []int{1, 2}
	c = append(c, 13) //add elem to the end of the slice
	// Maps [typeOfKeys]typeOfValues
	vert := make(map[string]int)
	vert["triangle"] = 2
	vert["square"] = 4
	vert["circle"] = 0
	d := vert["square"]
	delete(vert, "square")

	//print
	fmt.Printf("Hello %d \n", t)
	fmt.Println(a, b, vert, d)

	//loops - only kind in go is a for loop, which can double as while loop if we remove the init and incre cond in the declaration
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//iterate over the contents of a slice/array
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}
	// iterate over the contents of a map
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	}

	result := sum(2, 3)
	fmt.Println(result)

	res, err := sqrt(-16)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(res)
	}

	main2()
}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, errors.New("Undifined for negative numbers")
	}
	return math.Sqrt(x), nil

}

// A struct is a collection of fields

type person struct {
	name string
	age  int
}

func main2() {
	p := person{name: "Jake", age: 25}
	fmt.Println(p.age)

	//go has pointers

	i := 7
	inc(&i)
	fmt.Println(i)

}

func inc(x *int) {
	*x++
}
