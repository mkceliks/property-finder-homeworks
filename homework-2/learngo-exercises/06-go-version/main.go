package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.Version()) // OUTPUT = go1.18.3
}
