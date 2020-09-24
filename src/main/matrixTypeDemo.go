package main

import "fmt"

type matrix struct {
	Rows int
	Cols int 
	Data [][]int
}

func (m matrix) GetDiagonals() [] int  {

	d := make([]int, 0)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			if i == j {
				d = append(d, m.Data[i][j])
			}
		}
	}
	return d
}


func (m matrix) Add(m1 matrix) matrix {

	sum := matrix{m.Rows,m.Cols, m.Data}
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			sum.Data[i][j] = m.Data[i][j] + m1.Data[i][j]
		}
	}
	return sum
}

func (m matrix) GetOrder() (int, int)  {
	return m.Rows, m.Cols
}

func (m matrix) Transpose() matrix {
	var data[][] int
	transpose := matrix{m.Cols,m.Rows, data}
	newArr := make([][]int, len(m.Data))
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			newArr[j] = append(newArr[j], m.Data[i][j])
		}
	}
	transpose.Data = newArr
	return  transpose
}

func (m matrix) GetRows(input [2]int) (rows [][]int) {
	for i:=0; i < m.Rows; i++ {
		for _,val := range input {
			if i == val {
				rows = append(rows, m.Data[i])
			}
		}
	}
	return
}

func (m matrix) GetColumns(input [2] int) (cols [][]int) {
	mt := m.Transpose()
	for i:=0; i < mt.Rows; i++ {
		for _,val := range input {
			if i == val {
				cols = append(cols, mt.Data[i])
			}
		}
	}
	return
}

func (m matrix) IsSquareMatrix() bool  {
	return m.Rows == m.Cols
}

func main() {
	r, c := 3, 3
	a := [][]int {{10, 20, 30}, {40, 50, 60}, {70, 80, 90}}
	m1 := matrix{r, c, a}
	fmt.Println(m1)
	b := [][]int {{11, 22, 33}, {44, 55, 66}, {77, 88, 99}}
	m2 := matrix{r, c, b}
	fmt.Println(m2)

	m3 := m1.Add(m2)
	fmt.Println(m3)

	x, y := m1.GetOrder()
	fmt.Println("Matrix Order is:", x, y)

	m4 := m3.Transpose()
	fmt.Println(m4)

	var row [2]int
	row[0] = 1
	row[1] = 2
	fmt.Println(row)

	rows := m2.GetRows(row)
	fmt.Println("rows :", rows)

	cols := m3.GetColumns(row)
	fmt.Println("cols :", cols)

	diags := m4.GetDiagonals()
	fmt.Println("Diagonals :", diags)

	if m2.IsSquareMatrix() {fmt.Println("Square Matrix")} else {fmt.Println("Not Square Matrix")}

}