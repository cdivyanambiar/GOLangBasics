package exercises

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func WriteFile(filein, fileout  string) {
	fin,err := os.Open(filein)
	if err != nil { log.Panic(err) }
	// defer will be executed before exiting the main
	defer fin.Close()

	fout, err := os.OpenFile(fileout, os.O_APPEND, 0644)
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

func CopyFile(){
	if len(os.Args) < 3 {
		fmt.Println("Should pass atleast 2 files")
		os.Exit(1)
	}

	i := 0
	nthfile := os.Args[len(os.Args)-1]

	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
		if i <= len(os.Args)-1 {
			WriteFile(arg, nthfile)
		} else {
			break
		}
		i = i + 1
	}

}
func UniqueStrcount(filein string) {
	fin,err := ioutil.ReadFile(filein)

	if err != nil { log.Panic(err) }
	// defer will be executed before exiting the main
	sliceData := strings.Split(string(fin), " ")
	fmt.Println(sliceData)
	keys := make(map[string]int)
	for _, entry := range sliceData {
		keys[entry]++
	}
	fmt.Println(keys)
}

func UniqueStrfromFile(filein string) {
	fin,err := ioutil.ReadFile(filein)

	if err != nil { log.Panic(err) }
	// defer will be executed before exiting the main
	sliceData := strings.Split(string(fin), " ")
	fmt.Println(sliceData)
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range sliceData {
		value := keys[entry];
		if  !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	fmt.Println(list)
}


