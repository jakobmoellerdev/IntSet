package main

import (
	"fmt"
)

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(14)
	x.Add(13)
	x.Add(543)
	fmt.Println(x.String())
	fmt.Println(x.Exists(1))
	x.Remove(1)
	fmt.Println(x.String())
	fmt.Println(x.Exists(1))

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	fmt.Println("x", x)
	fmt.Println("y", y)
	fmt.Println("union", *y.Union(&x))
	fmt.Println("intersection", *y.Intersection(&x))
	fmt.Println("symmetric difference", *x.SymmetricDifference(&y))

	fmt.Println("x", x)
	fmt.Println("y", y)

	fmt.Println(x.Exists(9), x.Exists(123))
}
