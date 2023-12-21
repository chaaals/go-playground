package main

import (
	"fmt"
	"go-playground/pkg/population"
)

func main() {
	// letterGrade := getgrade.GetGrade(95,90,93)
	// fmt.Printf("%q", letterGrade)

	test := population.NbYear(1500, 5, 100, 5000)
	fmt.Printf("%v",test)
}