package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Print_q(q string) {
fmt.Println("\n------------------------------------------------------------------------\n")
fmt.Println(q)
}

func DeleteNEle(i int) {
	b := []int {1,2,3,4,5,6}
	b = append(b,7)
	fmt.Println(b)
}

/*
In Go method can return multiple values
 */

func Swap(x, y int) (int, int) {
	return y, x
}
/*
Named return values
 */
func Sqare(x int) (result int) {
	result = x * x
	return
}
func FindSmallest_userInput() {
	count := 0
	fmt.Print("Number of elements N=")
	fmt.Scanln(&count)
	fmt.Println("Enter of array elements")
	elements := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scanln(&elements[i])
	}
	fmt.Println("Entered Array of elements:", elements)
	smallestNumber := elements[0]
	pos := 0
	for p, element := range elements {
		if element < smallestNumber {
			smallestNumber = element
			pos = p

		}
	}
	fmt.Println("Smallest number of Array is  ", smallestNumber)
	fmt.Println("Position  ", pos)
}

//  write a code to find the smallest integer out of three integers
func FindSmallestOfAaray() {
	array := []int{10, 7, 8}
	fmt.Println(array)
	smallest := array[0]
	for _, element := range array {
		if element < smallest {
			smallest = element
		}
	}
	fmt.Println("Smallest number of Array is  ", smallest)
}

func ReadStringAndPrintEven()  {
	var str string
	fmt.Print("Enter the string=")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str = scanner.Text()
	fmt.Println("Entered String is:", str)
	wordList := strings.Fields(str)
	count := 0
	for _, word := range wordList {
		l := len(word)
		if l%2 == 0 {
			count++
		}
	}
	fmt.Println("Count of even number word = ", count)
}

func SwapElementsOfArray() {
	array := []int{1, 2, 3, 4, 11, 10}
	fmt.Println(array)

	len := len(array)
	for i:=0; i<len/2; i++ {
		temp1 := array[i]
		temp2 := array[len-1-i]
		array[i] =  temp2
		array[len-1-i] = temp1
	}
	fmt.Println(array)
}
func UniqueStr() {
	var str string
	fmt.Print("Enter the string=")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str = scanner.Text()
	fmt.Println("Entered String is:", str)
	slice := strings.Split(str, " ")
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		value := keys[entry];
		if  !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	fmt.Println(list)
}
func Sum(data ...int) int {
	s := 0
	for _,x := range data {
		s = s +x
	}
	return s
}

func IsValidIpaddress(ip string) (isvalid bool) {

	isvalid = true
	parts := strings.Split(ip, ".")
	if len(parts) !=4 {
		isvalid = false
	}
	for _, part := range parts {
		a, err := strconv.Atoi(part)
		if err!= nil {
			isvalid = false
		}
		if a < 1  || a > 255 {
			isvalid = false
		}
	}
	return
}

func FrequencyOfOccurance() {
	var str string
	fmt.Println("Enter a String")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str = scanner.Text()
	fmt.Println("Entered String is", str)
	slice := strings.Split(str, " ")
	keys := make(map[string]int)
	for _, entry := range slice {
		keys[entry]++
	}
	fmt.Println(keys)
}
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func ClasifyIP(ips ...string) (map[string] []string) {
	m := make(map[string] [] string)
	classa := []string{}
	classb := []string{}
	classc := []string{}
	classd := []string{}
	classe := []string{}
	classinvalid := []string{ "Invalid"}
	for _,ip := range ips {
		if IsValidIpaddress(ip) {
			parts := strings.Split(ip, ".")
			for _, part := range parts {
				a, _ := strconv.Atoi(part)
				if (a >= 1 && a <= 126) {
					_, found := Find(classa, ip)
					if !found {
						classa = append(classa, ip)
					}
				} else if (a >= 128 && a <= 191) {
					_, found := Find(classb, ip)
					if !found {
						classb = append(classb, ip)
					}
				} else if (a >= 192 && a < 223) {
					_, found := Find(classc, ip)
					if !found {
						classc = append(classc, ip)
					}
				} else if (a >= 224 && a <= 239) {
					_, found := Find(classd, ip)
					if !found {
						classd = append(classd, ip)
					}
				} else {
					_, found := Find(classe, ip)
					if !found {
						classe = append(classe, ip)
					}
				}
			}
		} else {
			classinvalid = append(classinvalid,ip)
		}
		m["Invalid"] = classinvalid
		m["A"] = classa
		m["B"] = classb
		m["C"] = classc
		m["D"] = classd
		m["E"] = classe
	}
	return m
}
func ReadMatrix() {
	var matrix1[100][100] int
	var row,col int
	fmt.Print("Enter number of rows: ")
	fmt.Scanln(&row)
	fmt.Print("Enter number of cols: ")
	fmt.Scanln(&col)
	fmt.Println()
	fmt.Println("========== Matrix =============")
	fmt.Println()
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("Enter the element for Matrix1 %d %d :",i+1,j+1)
			fmt.Scanln(&matrix1[i][j])
		}
	}

}
func MetrixAddition(){
	var matrix1[100][100] int
	var matrix2[100][100] int
	var sum[100][100] int
	var row,col int
	fmt.Print("Enter number of rows: ")
	fmt.Scanln(&row)
	fmt.Print("Enter number of cols: ")
	fmt.Scanln(&col)

	fmt.Println()
	fmt.Println("========== Matrix1 =============")
	fmt.Println()
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("Enter the element for Matrix1 %d %d :",i+1,j+1)
			fmt.Scanln(&matrix1[i][j])
		}
	}

	fmt.Println()
	fmt.Println("========== Matrix2 =============")
	fmt.Println()

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("Enter the element for Matrix2 %d %d :",i+1,j+1)
			fmt.Scanln(&matrix2[i][j])
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			sum[i][j] = matrix1[i][j]+matrix2[i][j]
		}
	}

	fmt.Println()
	fmt.Println("========== Sum of Matrix =============")
	fmt.Println()

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf(" %d ",sum[i][j])
			if(j==col-1){
				fmt.Println("")
			}
		}
	}
}
func MetrixAddition_N(){
	var matrix1[100][100] int
	var matrix2[100][100] int
	var sum[100][100] int
	var numMetrix, row,col int
	fmt.Print("Enter number of Matrix: ")
	fmt.Scanln(&numMetrix)
	for num :=0; num <numMetrix; num++ {
		fmt.Print("Enter number of rows: ")
		fmt.Scanln(&row)
		fmt.Print("Enter number of cols: ")
		fmt.Scanln(&col)

		fmt.Println()
		fmt.Println("========== Matrix1 =============")
		fmt.Println()
		for i := 0; i < row; i++ {
			for j := 0; j < col; j++ {
				fmt.Printf("Enter the element for Matrix1 %d %d :", i+1, j+1)
				fmt.Scanln(&matrix1[i][j])
			}
		}
	}

	fmt.Println()
	fmt.Println("========== Matrix2 =============")
	fmt.Println()

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("Enter the element for Matrix2 %d %d :",i+1,j+1)
			fmt.Scanln(&matrix2[i][j])
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			sum[i][j] = matrix1[i][j]+matrix2[i][j]
		}
	}

	fmt.Println()
	fmt.Println("========== Sum of Matrix =============")
	fmt.Println()

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf(" %d ",sum[i][j])
			if(j==col-1){
				fmt.Println("")
			}
		}
	}
}
