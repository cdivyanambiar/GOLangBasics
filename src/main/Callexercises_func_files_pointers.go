package main

import (
	"exercises"
	"fmt"
	"os"
)

func main() {
	problems := `
1. write a code copy the content of first N-1 files into Nth file with CLI interface
2. write a code to find the frequency occurrence of words in a file with CLI interface
3. write a code to remove empty lines from a file with CLI interface
4. write a code to find the sum of 2 matrices using files with CLI interface 
5. write a code to compare 2 files with CLI interface
	`
	fmt.Println(problems)

	exercises.CopyFile()
	exercises.UniqueStrfromFile(os.Args[1])
	exercises.UniqueStrcount(os.Args[1])

}
