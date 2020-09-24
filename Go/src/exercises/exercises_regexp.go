package exercises

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)
func sumDigits(number int) int {
	remainder := 0
	sumResult := 0
	for number != 0 {
		remainder = number % 10
		sumResult += remainder
		number = number / 10
	}
	return sumResult
}

func readString() string {
	var pat string
	fmt.Print("enter some string:")
	fmt.Scan(&pat)
	return pat
}

func SumNumbers()  {
    pat := readString()
	re := regexp.MustCompile("[0-9]+")
	fmt.Println(re)
    result :=0
	r := re.FindAllString(pat, -1)
	for i :=0; i<len(r); i++ {
		num, err := strconv.Atoi(r[i])
		if err == nil {
			result = result + sumDigits(num)
		}
	}
	fmt.Println(result)

}

func BeginWithCaps() {
	pat := "I like to do Exercise"
	re := regexp.MustCompile("[A-Z]\\w+")
	slice := strings.Split(pat, " ")
	for _, val := range slice {
		r := re.FindAllString(val, -1)
		fmt.Println(r)
	}
}

func BeginWithCapsAndContainvVwel() {
	pat := "Aeiou joke Uae"
	re := regexp.MustCompile("[A-Z]\\w*[aeiou]")
	slice := strings.Split(pat, " ")
	for _, val := range slice {
		r := re.FindAllString(val, -1)
		fmt.Println(r)
	}
}

func WordsHasVowel() {
	pat := "Uae Divya Jjj"
	re := regexp.MustCompile("\\w*[aeiou]\\w*")
	slice := strings.Split(pat, " ")
	for _, val := range slice {
		r := re.FindAllString(val, -1)
		fmt.Println(r)
	}
}

func AllCaps() {
	pat := "Uae Divya hjhj"
	re := regexp.MustCompile("[A-Z][^A-Z]*")
	slice := strings.Split(pat, " ")
	for _, val := range slice {
		r := re.FindAllString(val, -1)
		fmt.Println(r)
	}
}


