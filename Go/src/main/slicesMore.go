package main 

import "fmt"

func main() {
	var s1 []int 
	//s1[0] = 10
	s1 = append(s1, 10)
	fmt.Println(s1)

	s2 := []int{}
	//s2[0] = 10
	s2 = append(s2, 20)
	fmt.Println(s2)


	s3 := []int{11, 22, 33}
	s3[0] = 100 //works
	//s3[3] = 200 //fails
	fmt.Println(s3)

	s4 := make([]  int, 5)
	fmt.Println(s4)
	s4[0] = 100
	fmt.Println(s4)
	s4 = append(s4, 200)
	fmt.Println(s4)

	fmt.Println("-----------------------------------")
	s5 := make([] []int, 3)
	fmt.Println(s5) 
	for i:=0; i<3; i=i+1  { s5[i] = make([]int, 4)} 
	fmt.Println(s5) 
	s5[0][0] = 100
	fmt.Println(s5)
}