package variables

import (
	"fmt"
	"strconv"
)

// declare all var in one block
var (
	name       string = "Nitin"
	age        int    = 11
	occupation string = "SWE"
)

// can't declare using :=, has to manually declare with type
var global int = 5

func Variables() {
	// decalre var method 1 (when you don't want to assign)
	var a int
	a = 20
	fmt.Printf("%v, %T\n", a, a)

	// init method 2 (when you want to assign specific type by yourself)
	var b int = 10
	b = 12
	fmt.Printf("%v, %T\n", b, b)

	var d float32 = 34
	fmt.Printf("%v, %T\n", d, d)

	// init method 3 (autotype)
	c := 3
	fmt.Printf("%v, %T\n", c, c)

	fmt.Printf("%v, %T\n", global, global)
	fmt.Printf("%v, %T\n", name, name)
	fmt.Printf("%v, %T\n", age, age)
	fmt.Printf("%v, %T\n", occupation, occupation)

	var e int = 12
	fmt.Printf("%v, %T \n", e, e)

	var j float32
	j = float32(e)

	fmt.Printf("%v, %T\n", j, j)

	var h int = 12
	var s string
	s = string(h) // print * as string is an alias for string of bytes
	fmt.Printf("%v, %T\n", s, s)
	// use strconv fro string conversions and method
	s1 := strconv.Itoa(h)
	fmt.Printf("%v, %T\n", s1, s1)
}

// var GlobalInt int = 34 - Declaring in capital makes it is expose to other packages
