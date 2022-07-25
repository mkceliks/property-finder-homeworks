package main

import (
	"flag"
	"fmt"
	"io/fs"
	"path/filepath"
)

func visit(path string, di fs.DirEntry, err error) error {

	fmt.Printf("Visited: %s\n", path) // visited Path
	return nil                        // it is just writing the given path to the console.

}

func main() {

	flag.Parse()
	root := flag.Arg(0)
	err := filepath.WalkDir(root, visit)                // usage of filepath.WalkDir() function
	fmt.Printf("filepath.WalkDir() returned %v\n", err) // it returnes all of the files to the console that takes data from filepath.WalkDir() function

}
