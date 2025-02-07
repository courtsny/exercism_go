package main

import (
	"fmt"
	"strings"
)

func formatTitle(name string) string {
	var dashes = strings.Repeat("-", len(name))
	return "\n" + name + "\n" + dashes + "\n"
}

func main() {
	fmt.Println(formatTitle("Basics"))
	BasicsTest()

	fmt.Println(formatTitle("Numbers & Arithmetic Operations"))
	ArithmeticTest()

	fmt.Println(formatTitle("Booleans"))
	BooleansTest()

	fmt.Println(formatTitle("Strings"))
	StringsTest()

	fmt.Println(formatTitle("Conditional Ifs & Comparisons"))
	IfTest()

	fmt.Println(formatTitle("Packages & Formatting Strings"))
	PackagesTest()

	fmt.Println(formatTitle("Slices & Variadic Functions"))
	SlicesTest()

	fmt.Println(formatTitle("Conditional Switch"))
	SwitchTest()

	fmt.Println(formatTitle("Structs"))
	StructsTest()

	fmt.Println(formatTitle("Randomness"))
	RandomnessTest()

	fmt.Println(formatTitle("For"))
	ForTest()

	fmt.Println(formatTitle("Functions"))
	FunctionsTest()

	fmt.Println(formatTitle("Float"))
	FloatTest()

	fmt.Println(formatTitle("Time"))
	TimeTest()

	fmt.Println(formatTitle("Maps"))
	MapsTest()
	return
}
