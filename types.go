package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Type conversion: int(var), cast var as a type
// Type assertion: var.(int)
// Type switch: var.(type) (literally 'type')

type Box interface {
	Number() float64
}

type NumberBox struct {
	Value float64
}

func (n NumberBox) Number() float64 {
	return n.Value
}

type FancyNumberBox struct {
	Value string
}

func (f FancyNumberBox) Number() float64 {
	num_str := strings.Split(f.Value, ".")
	num, err := strconv.Atoi(num_str[0])
	if err != nil {
		fmt.Printf("failed to parse num_str[0] = %s, num_str = %s, err = %v\n", num_str[0], num, err)
	}
	dec := 0
	if len(num_str) > 1 {
		dec, err = strconv.Atoi(num_str[1])
		if err != nil {
			fmt.Printf("failed to parse num_str[1] = %s, num_str = %s, err = %v\n", num_str[1], num, err)
		}
	}
	return float64(num) + float64(dec)
}

func DescribeNumberHelper(num float64) string {
	num_str := strconv.Itoa(int(num))
	dec := math.Round(num) - num
	dec *= 10
	dec_str := strconv.Itoa(int(dec))
	return num_str + "." + dec_str
}

func DescribeNumber(num float64) string {
	var sb strings.Builder
	sb.WriteString("This is the number ")
	sb.WriteString(DescribeNumberHelper(num))
	return sb.String()
}

func DescribeNumberBox(n NumberBox) string {
	var sb strings.Builder
	sb.WriteString("This is a box containing the number ")
	sb.WriteString(DescribeNumberHelper(n.Number()))
	return sb.String()
}

func ExtractFancyNumber(b Box) int {
	fnb, ok := b.(FancyNumberBox)
	if !ok {
		//fmt.Printf("%v not parsable as a FancyNumberBox", b)
		return 0
	}
	return int(fnb.Number())
}

func DescribeFancyNumberBox(f FancyNumberBox) string {
	var sb strings.Builder
	sb.WriteString("This is a fancy box containing the number ")
	sb.WriteString(DescribeNumberHelper(f.Number()))
	return sb.String()
}

func DescribeAnything(x interface{}) string {
	// Check supported primitives
	i, ok := x.(int)
	if ok {
		return DescribeNumber(float64(i))
	}
	f, ok := x.(float64)
	if ok {
		return DescribeNumber(f)
	}

	// Check box types
	nb, ok := x.(NumberBox)
	if ok {
		return DescribeNumberBox(nb)
	}
	fnb, ok := x.(FancyNumberBox)
	if ok {
		return DescribeFancyNumberBox(fnb)
	}

	// Unknown type
	//fmt.Printf("%v is an unsupported type", x)
	return "Return to sender"
}

func TypesTest() {
	float := -12.3456
	fmt.Printf("%.4f is described as: '%s'\n", float, DescribeNumber(float))
	nb := NumberBox{12}
	fmt.Printf("number box nb %v is described as: '%s'\n", nb, DescribeNumberBox(nb))
	fnb := FancyNumberBox{"10"}
	fmt.Printf("fancy number box fnb %v is described as: '%s'\n", fnb, DescribeFancyNumberBox(fnb))
	fmt.Printf("extracting fancy number from nb %v: %d (0 because nb is not fancy)\n", nb, ExtractFancyNumber(nb))
	fmt.Printf("extracting fancy number from fnb %v: %d\n", fnb, ExtractFancyNumber(fnb))
	i := 2000
	s := "hello descriptor"
	fmt.Printf("describe %v: %s\n", i, DescribeAnything(i))
	fmt.Printf("describe %v: %s\n", float, DescribeAnything(float))
	fmt.Printf("describe '%v': %s\n", s, DescribeAnything("hello descriptor"))
	fmt.Printf("describe %v: %s\n", nb, DescribeAnything(nb))
	fmt.Printf("describe %v: %s\n", fnb, DescribeAnything(fnb))

}
