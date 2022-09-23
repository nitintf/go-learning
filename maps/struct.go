package maps

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	name       string
	age        int
	companions []string
}

type Animal struct {
	name   string `required,max:"10"`
	origin string `required,min:"2"`
}

type Bird struct {
	Animal
	speedKPH float32
	canFly   bool
}

func Struct() {
	doctor1 := Doctor{
		name: "Nitin Panwar",
		age:  21,
		companions: []string{
			"Chandler",
			"joey",
		},
	}

	fmt.Println(doctor1)

	// anymonous struct
	doctor2 := struct{ name string }{name: "Ross"}
	fmt.Println(doctor2)

	// b := Bird{
	// 	Animal: Animal{name: "emu", origin: "Australia"},
	// 	canFly: true,
	// 	speedKPH: 12.23,
	// }

	t := reflect.TypeOf(Animal{})

	field, _ := t.FieldByName("name")
	fmt.Println(field.Tag)

	b := Bird{}
	b.canFly = true
	b.speedKPH = 12.23
	b.name = "Emu"
	b.origin = "Australia"

	fmt.Println(b.name)
}
