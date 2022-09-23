package main

import (
	"github.com/nitintf/go-learning/array"
	"github.com/nitintf/go-learning/conditions"
	"github.com/nitintf/go-learning/maps"
	"github.com/nitintf/go-learning/variables"
)

func main() {
	print("\n\n\n\n")
	print("Variables\n\n")
	variables.Variables()

	print("\n\n\n\n")
	print("Primitives\n\n")
	variables.Primitives()

	print("\n\n\n\n")
	print("Const\n\n")
	variables.Const()

	print("\n\n\n\n")
	print("Array\n\n")
	array.Array()

	print("\n\n\n\n")
	print("Maps\n\n")
	maps.Maps()

	print("\n\n\n\n")
	print("Struct\n\n")
	maps.Struct()

	print("\n\n\n\n")
	print("Conditional Statements\n\n")
	conditions.IF()
}
