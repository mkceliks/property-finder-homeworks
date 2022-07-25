package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	tempconv "ex_2.1"
)

type Feet float64
type Meter float64

func (f Feet) String() string  { return fmt.Sprintf("%g(ft)", f) }
func (m Meter) String() string { return fmt.Sprintf("%g(m)", m) }

func FToM(f Feet) Meter { return Meter(f * 0.3048) }
func MToF(m Meter) Feet { return Feet(m * 3.28083) }

func uc(t float64) {

	fmt.Println("Temperature Conversion:")
	f := tempconv.Fahrenheit(t) // taking the fahrenheit value to a variable
	c := tempconv.Celsius(t)    // taking the celcius value to a variable
	fmt.Printf("%s = %s, %s = %s\n\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c)) // convs.

	fmt.Println("Length Conversion:")
	lf := Feet(t)  // lenght of feet
	lm := Meter(t) // length of meter
	fmt.Printf("%s = %s, %s = %s\n\n",
		lf, FToM(lf), lm, MToF(lm)) // convs
}

func main() {
	if len(os.Args[1:]) > 0 { // if there is a os.Args this block running
		for _, arg := range os.Args[1:] { // forr loop for every index of Args
			t, err := strconv.ParseFloat(arg, 64) // parsing the args to string
			if err != nil {
				fmt.Fprintf(os.Stderr, "uc: %v\n", err) // if there is an error
				os.Exit(1)
			}
			uc(t) // if there is no error call the func.
		}
	} else { // if there is no os.Args from command-line
		for {
			input := bufio.NewReader(os.Stdin) // take the input from user
			fmt.Fprintf(os.Stdout, "=> ")
			s, err := input.ReadString('\n') // readstring.
			if err != nil {
				fmt.Fprintf(os.Stderr, "uc: %v\n", err)
				os.Exit(1)
			}
			t, err := strconv.ParseFloat(s[:len(s)-1], 64) // parsing the string of input
			if err != nil {                                // if there is an error block.
				fmt.Fprintf(os.Stderr, "uc: %v\n", err)
				os.Exit(1)
			}
			uc(t) // calling the func.
		}
	}
}
