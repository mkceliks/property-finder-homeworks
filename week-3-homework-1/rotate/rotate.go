package main

import (
	"fmt"
)

func left_rotate(sentence string, num int) { // left rotation for sentence.

	begin := sentence[0:num] // begin of new sentence.
	end := sentence[num:]    // end of new sentence.

	fmt.Printf("Left Rotation result is : %s\n", end+begin)
}

func right_rotate(sentence string, num int) { // right rotation for sentence.

	begin := sentence[0 : len(sentence)-num] // begin of new sentence.
	end := sentence[len(sentence)-num:]      // end of new sentence.

	fmt.Printf("Right Rotation result is : %s \n", end+begin)

}
func main() {

	var sentence string
	var num int
	fmt.Printf("Enter a sentence and number to rotate : ")
	fmt.Scanf("%s %d", &sentence, &num)
	left_rotate(sentence, num)
	right_rotate(sentence, num)

}
