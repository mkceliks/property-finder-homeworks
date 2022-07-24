package main

import (
	"errors"
	"fmt"
	mathFunction "fundamentals/ch07/math"
	"os"
	"strings"
)

var flag bool = true

type Calculator struct {
	functions []mathFunction.MathFunction
}

func (c *Calculator) addMathFunction(m mathFunction.MathFunction) {
	c.functions = append(c.functions, m)
}

func (c *Calculator) doCalculation(name string, arg float64) (float64, error) { // math function name entered by the user is lower case of name registered by the calculator.
	var result float64
	for _, f := range c.functions {
		if strings.ToLower(name) == strings.ToLower(f.GetName()) { // this To.Lower() allows the program that works correctly even though user enters lower case string.
			result = f.Calculate(arg)
			return result, nil
		}
	}
	return 0, errors.New("no such function exists:" + name)
}

func main() {

	startCalculator()

}

func startCalculator() {
	myCalculator := Calculator{}
	myCalculator.addMathFunction(mathFunction.Sin{"Sine"})
	myCalculator.addMathFunction(mathFunction.Cos{"Cosine"})
	myCalculator.addMathFunction(mathFunction.Log{"Log"})

	fmt.Println("\nCalculator started.")
	fmt.Println("You can calculate using following functions")
	for _, f := range myCalculator.functions {
		fmt.Println(f.GetName())
	}

	for flag {
		var fName string
		var arg float64
		fmt.Println("> Enter name of the calculation or enter x to exit:")
		_, err := fmt.Scanf("%s", &fName) // error handlings for flags.
		if err != nil {                   // error handlings for flags.
			fmt.Println(err)
			os.Exit(0)
		}
		//
		if fName == "x" {
			flag = false
		} else {
			fmt.Println("> Enter a value for the calculation:")
			_, err := fmt.Scanf("%f", &arg) // error handlings for flags.
			if err != nil {                 // error handlings for flags.
				fmt.Println(err)
				os.Exit(0)
			}
			value, err := myCalculator.doCalculation(fName, arg) // error handlings for flags.
			if err != nil {                                      // error handlings for flags.
				fmt.Println(err)
			} else {
				fmt.Printf("Result of %s of %f : %f\n", fName, arg, value)
			}
		}
	}
	println("Bye!")
}
