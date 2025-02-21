package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func formatTitle(name string) string {
	var dashes = strings.Repeat("-", len(name))
	return "\n" + name + "\n" + dashes + "\n"
}

var function_map = map[string]func(){
	"Basics":                          BasicsTest,
	"Numbers & Arithmetic Operations": ArithmeticTest,
	"Booleans":                        BooleansTest,
	"Strings":                         StringsTest,
	"Conditional Ifs & Comparisons":   IfTest,
	"Packages & Formatting Strings":   PackagesTest,
	"Slices & Variadic Functions":     SlicesTest,
	"Conditional Switch":              SwitchTest,
	"Randomness":                      RandomnessTest,
	"For":                             ForTest,
	"Float":                           FloatTest,
	"Time":                            TimeTest,
	"Maps":                            MapsTest,
	"Range Interation (Chess)":        RangeTest,
	"Pointers":                        PointersTest,
	"Structs & Struct Methods":        StructsTest,
	"Runes":                           RunesTest,
	"Regular Expressions":             RegexesTest,
	"Interfaces":                      InterfacesTest,
	"Zero Values":                     ZeroValuesTest,
	"Stringers":                       StringersTest,
	"Type Conversion & Assertion":     TypesTest,
	"Errors":                          ErrorsTest,
	"Functions":                       FunctionsTest,
	"First Class Functions":           FirstClassTest,
}

func PrintKeys(key_array []string) string {
	var sb strings.Builder
	last := 0
	for i, x := range key_array {
		i_str := strconv.Itoa(i)
		sb.WriteString(i_str + ". " + x + "\n")
	}
	last = len(key_array)
	last_str := strconv.Itoa(last)
	sb.WriteString(last_str + ". Run All\n")
	return sb.String()
}

func PromptAndChoice(num_keys int) int {
	choice := 0
	fmt.Printf("\nPlease enter a choice from %d to %d, %d for all, or %d+ to exit.\n", 0, num_keys-1, num_keys, num_keys+1)
	fmt.Scanf("%v", &choice)
	return choice
}

func main() {
	// Sort function descriptions alphabetically
	keys := make([]string, 0, len(function_map))
	for k := range function_map {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Printf("Welcome to the Exercism Go exercises! You can run any of these:\n%v\n", PrintKeys(keys))
	choice := PromptAndChoice(len(keys))
	for choice <= len(function_map) {
		if choice == len(keys) {
			// Run functions by sort
			for _, k := range keys {
				fmt.Println(formatTitle(k))
				function_map[k]()
			}
		} else {
			fmt.Printf(formatTitle(keys[choice]))
			function_map[keys[choice]]()
		}
		choice = PromptAndChoice(len(keys))
	}

	return
}
