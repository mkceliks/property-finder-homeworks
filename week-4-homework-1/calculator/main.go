package main

import "fmt"

type Degrees struct { // Defined a struct that includes the conversion types
	Celcius    float64
	Kelvin     float64
	Fahrenheit float64
}

type Conversions interface { // Defined an interface that includes the functions about conversions.
	Celsius2Fahrenheit()
	Celsius2Kelvin()
	Kelvin2Celsius()
	Kelvin2Fahrenheit()
	Fahrenheit2Celsius()
	Fahrenheit2Kelvin()
}

func (a Degrees) Celsius2Fahrenheit() float64 { // Celsius to Fahrenheit conversion
	return (a.Celcius*9/5 + 32)
}

func (a Degrees) Celsius2Kelvin() float64 { // Celsius to Kelvin conversion
	return (a.Celcius + 273.15)
}

func (a Degrees) Kelvin2Celsius() float64 { // Kelvin to Celsius conversion
	return (a.Kelvin - 273.15)
}

func (a Degrees) Kelvin2Fahrenheit() float64 { // Kelvin to Fahrenheit conversion
	return (a.Kelvin*9/5 - 459.67)
}

func (a Degrees) Fahrenheit2Celsius() float64 { // Fahrenheit to Celsius conversion
	return ((a.Fahrenheit - 32) * 5 / 9)
}

func (a Degrees) Fahrenheit2Kelvin() float64 { // Fahrenheit to Kelvin conversion
	return ((a.Fahrenheit + 459.67) * 5 / 9)
}

func main() {
	// Initializing the variables that the user gonna enter.
	var cel float64
	var fah float64
	var kel float64

	fmt.Printf("Enter a Celcius Degree :")
	fmt.Scanln(&cel)
	fmt.Printf("Enter a Kelvin Degree :")
	fmt.Scanln(&kel)
	fmt.Printf("Enter a Fahrenheit Degree :")
	fmt.Scanln(&fah)
	// interface values declared
	a := Degrees{
		Celcius:    cel,
		Kelvin:     kel,
		Fahrenheit: fah,
	}

	// All functions here...
	fmt.Printf(" This is your Celcius to Fahrenheit Degree %v \n", a.Celsius2Fahrenheit())
	fmt.Printf(" This is your Celcius to Kelvin Degree %v \n", a.Celsius2Kelvin())
	fmt.Printf(" This is your Kelvin to Celcius Degree %v \n", a.Kelvin2Celsius())
	fmt.Printf(" This is your Kelvin to Fahrenheit Degree %v \n", a.Kelvin2Fahrenheit())
	fmt.Printf(" This is your Fahrenheit to Celcius Degree %v \n", a.Fahrenheit2Celsius())
	fmt.Printf(" This is your Fahrenheit to Kelvin Degree %v \n", a.Fahrenheit2Kelvin())

}
