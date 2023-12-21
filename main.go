package main

import (
	"fmt"
	"go-playground/pkg/getgrade"
)

func main() {
	letterGrade := getgrade.GetGrade(95,90,93)
	fmt.Printf("%q", letterGrade)
}