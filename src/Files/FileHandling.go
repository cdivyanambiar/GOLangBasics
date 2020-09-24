package Files

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)
func ArgparamSum() {
	//fmt.Println(os.Args)
	// fmt.Print(os.Args[1])
	if len(os.Args) !=3 {
		fmt.Println("I need 2 integers through CLI")
		os.Exit(1)
	}
	a,_ := strconv.Atoi(os.Args[1])
	b,_ := strconv.Atoi(os.Args[2])
	c := a+b
	fmt.Println("Sum =",c)
}
func PrintcurrentPath() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
}
func WriteIntoFileUsingScan(file string) {
	fin,err := os.Open("temp.dat")
	if err != nil { log.Panic(err) }
	// defer will be executed before exiting the main
	defer fin.Close()
	fout, err := os.Create("temp2.dat")
	if err != nil {
		panic(err)
	}
	defer fout.Close()

	scanner := bufio.NewScanner(fin)
	for scanner.Scan() {
		s := scanner.Text()
		fout.Write([]byte(s+"\n"))
	}
}
func WriteIntoFile(file string) {
	fin,err := os.Open("temp.dat")
	if err != nil { log.Panic(err) }
	// defer will be executed before exiting the main
	defer fin.Close()
	fout, err := os.Create("temp3.dat")
	if err != nil {
		panic(err)
	}
	defer fout.Close()

	b := make([]byte, 10)
	for {
		n, _ := fin.Read(b)
		if n == 0 { break }
		fout.Write(b[:n])
	}
}

