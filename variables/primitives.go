package variables

import "fmt"

func Primitives() {
	var isCompleted = false

	fmt.Printf("%v, %T\n", isCompleted, isCompleted)

	isEqual := 1 == 3
	fmt.Printf("%v, %T\n", isEqual, isEqual)

	// complex type
	var a = 1 + 2i // complex(1, 2)

	var b complex64 = 1 + 2i

	fmt.Println("Complex128")
	fmt.Printf("%v, %T\n", real(a), real(a))
	fmt.Printf("%v, %T\n", imag(a), imag(a))

	fmt.Println("Complex64")
	fmt.Printf("%v, %T\n", real(b), real(b))
	fmt.Printf("%v, %T\n", imag(b), imag(b))

	// rune - using single quotes
	ba := 'A' // var b rune = 'a
	fmt.Printf("%v, %T\n", ba, ba)

}
