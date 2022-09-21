package variables

import "fmt"

func Primitives() {

	print("\n\n\n\n")
	print("Primitives Package\n\n")

	var isCompleted bool = false

	fmt.Printf("%v, %T\n", isCompleted, isCompleted)

	isEqual := 1 == 3
	fmt.Printf("%v, %T\n", isEqual, isEqual)

	// complex type
	var a complex128 = 1 + 2i // complex(1, 2)
	fmt.Printf("%v, %T\n", real(a), real(a))
	fmt.Printf("%v, %T\n", imag(a), imag(a))

	// rune - using single quotes
	b := 'a' // var b rune = 'a
	fmt.Printf("%v, %T\n", b, b)

}
