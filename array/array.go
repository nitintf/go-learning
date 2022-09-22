package array

import "fmt"

func Array() {
	grades := [3]int{12, 34, 56}

	fmt.Printf("Grades: %v\n", grades)

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Nitin"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Number of students: %v\n", len(students))

	var identityMatrix [3][3]int = [3][3]int{{1, 0, 0}, {1, 2, 3}, {4, 5, 6}}
	fmt.Println(identityMatrix)

	a := [...]int{1, 2, 4}
	b := a // make a copy
	b[0] = 2
	fmt.Println(a)
	fmt.Println(b)

	// Slice

	c := []int{1, 3, 5}
	fmt.Println(len(c))
	fmt.Println(cap(c))

	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(d[:])
	fmt.Println(d[:5])
	fmt.Println(d[3:])
	fmt.Println(d[4:8])
	fmt.Println(d[:1])

	e := make([]int, 3)
	fmt.Println(e)

	f := []int{}
	f = append(f, 1)
	f = append(f, []int{1, 2, 3, 4}...)
	fmt.Println(f)
}
