package main

import "fmt"

func main() {
	/*basic print*/
	fmt.Println("Hello, world")

	/*static type declaration*/
	var x float64 = 20.1
	fmt.Println(x)
	fmt.Printf("x is of type %T\n", x)
	/*dynamic type declaration*/
	y := 10
	fmt.Printf("y is of type %T\n", y)
	/*mixed variable declaration*/
	var a, b, c = 1, 20.1, "foo"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)

	/*const keyword*/
	const LENGTH int = 10
	const WIDTH int = 5
	var area int = LENGTH * WIDTH
	fmt.Printf("value of area : %d\n", area)

	/*arithmetic operators*/
	fmt.Printf("y + a = %d\n", y+a)
	fmt.Printf("y -a = %d\n", y-a)
	/*relational operators*/
	fmt.Printf("x == b %t\n", x == b)
	/*pointer*/
	var ptr1 *int = &a
	fmt.Printf("*ptr1 is %d\n", *ptr1)
	var ptr2 *float64 = &b
	fmt.Printf("*ptr2 is %f\n", *ptr2)

	/*decision making*/
	if true {
		fmt.Println("if")
	} else {
		fmt.Println("else")
	}
}
