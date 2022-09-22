package maps

import "fmt"

func Maps() {
	statePopulations := map[string]int{
		"a": 2,
		"b": 22,
		"c": 1,
		"d": 67,
	}

	// delete(statePopulations, "c")
	fmt.Println(statePopulations["c"]) // 0
	fmt.Println(statePopulations["z"]) // 0

	// we get 0 if it's not in map
	// better way
	pop, ok := statePopulations["c"]

	if !ok {
		panic("No value find for c")
	} else {
		fmt.Println(pop)
	}

	m := map[int]string{}

	fmt.Println(statePopulations, m)
}
