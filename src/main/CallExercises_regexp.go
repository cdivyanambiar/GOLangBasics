package main

import (
	"exercises"
	"fmt"
)

func main() {
	problems := `
1.	write a code to find the sum of all the numbers available in a string
2.	write a code to extract all the words which begins with Capital letter from a string 
3.	write a code to extract all the words which begins with Capital letter and ends with lower case vowel from a string 
4.	write a code to extract all the words which contains at least a vowel from a string 
5.	write a code to extract all Capital lettered words from a string 
6.	write a code to extract all the numbers between 1 and 255 (inclusive) from a string 
7.	write a code to extract all valid IPs from a string 
8.	write a code to extract both positive and negative integers from a string 
9.	write a code to clean a file with CLI interface (clean means to remove all empty lines, and also replace one or more spaces by a single space within lines if any)
	`
	fmt.Println(problems)
	exercises.Print_q("1.write a code to find the sum of all the numbers available in a string")
	exercises.SumNumbers()

	exercises.Print_q("2.write a code to extract all the words which begins with Capital letter from a string")
	exercises.BeginWithCaps()

	exercises.Print_q("3.write a code to extract all the words which begins with Capital letter and ends with lower case vowel from a string ")
	exercises.BeginWithCapsAndContainvVwel()

	exercises.Print_q("4.write a code to extract all the words which contains at least a vowel from a string ")
	exercises.WordsHasVowel()

	exercises.Print_q("5.write a code to extract all Capital lettered words from a string ")
	exercises.AllCaps()

}
