package variables

import "fmt"

const (
	_  = iota // ignore first value
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	a = iota
	b
	c
)

func Const() {
	const i string = "Nitin"
	// i = "anuj"  -> cannot assign as it is consant
	fmt.Printf("%v, %T\n", i, i)

	// iota - counter
	fmt.Printf("%v\n", a) // 0
	fmt.Printf("%v\n", b) // 1
	fmt.Printf("%v\n", c) // 2

	fileSize := 40000000.

	fmt.Printf("%.2fGB\n", fileSize/GB)

}
