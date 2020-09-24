package main

import (
	"exercises"
	"fmt"
)

func main() {
	problems := `
1. write a code to find the smallest integer out of three integers
2. write a code to read N integers from keyborad and find the smallest integer along with its position
3. write a code to read a string and find out the number of even length words in it
4. write a code to swap first element with last element, second element with last but one element, and so on from a given list of elements
5. write a code to remove duplicate words from a given string
6. write a code to find the frequence occurrence of words in a string 
7. write a code to test a given IP address is valid or not
8. write a code to read N IP addresses and classify them into 6-classes
	a) Valid IPs, b) Invalid IPs, c) ClassA, d) ClassB, e) ClassC, f) ClassD
9. write a code to create and display a matrix (use slice of slices)
10. write a code to find the sum of two matrices
11. write a code to find the sum of N matrices
12. write a code to read N strings from keyborad and find the frequence occurrence of characters string wise.
`
	fmt.Println(problems)

	exercises.Print_q("1. write a code to find the smallest integer out of three integers")
	exercises.FindSmallestOfAaray()

	exercises.Print_q("2. write a code to read N integers from keyborad and find the smallest integer along with its position")
	exercises.FindSmallest_userInput()

	exercises.Print_q("3. write a code to read a string and find out the number of even length words in it")
	exercises.ReadStringAndPrintEven()

	exercises.Print_q("4. write a code to swap first element with last element, second element with last but one element, and so on from a given list of elements")
    exercises.SwapElementsOfArray()

	exercises.Print_q("5. write a code to remove duplicate words from a given string")
	//exercises.re


}
