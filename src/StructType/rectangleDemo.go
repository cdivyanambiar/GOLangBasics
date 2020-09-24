package StructType

import "fmt"

type rectagle struct {
	breadth,len float32
}
func (r rectagle) String() string {
	return fmt.Sprintf("rectangle{len: %0.1f, breadth: %0.1f}", r.len, r.breadth)
}
func (r rectagle) Area() float32 {
	return r.len * r.breadth
}

func (r rectagle) IsSquare() bool {
	return r.len == r.breadth
}

func (r rectagle) Perimeter() float32 {
	return 2 * (r.len + r.breadth)
}

func (r rectagle) Scale(len, breadth float32) {
	if len + r.len > 0 {
		r.len = r.len + len
	}
	if breadth + r.breadth > 0 {
		r.breadth = r.breadth + breadth
	}
}

func (r *rectagle) ScaleWithPointer(len, breadth float32) {
	if len + r.len > 0 {
		r.len = r.len + len
	}
	if breadth + r.breadth > 0 {
		r.breadth = r.breadth + breadth
	}
}

func (r rectagle) Add(t rectagle) rectagle {
	return rectagle{r.len+t.len, r.breadth+t.breadth}
}

func (r rectagle) Equal(t rectagle) bool  {
	return r.len == t.len && r.breadth == t.breadth
}
func main() {
	r1 := rectagle{}
	fmt.Println(r1)

	r2 := rectagle{2,3}
	fmt.Println(r2)

	r3 := rectagle{len: 3, breadth: 4}
	fmt.Println(r3)

	a2 := r2.Area()
	fmt.Println(a2)

	if r2.IsSquare() {fmt.Println("Square")} else {fmt.Println("Not Square")}
/*
This will not make any change and it will print 3 and 2.
 */
	r2.Scale(1,2)
	fmt.Println(r2)

	r4 := &r2
	r5 := r2
	/*
	This will not make any change and it will print 3 and 2.
	*/
	r2.ScaleWithPointer(1,2)
	fmt.Println(r2)

	fmt.Println(*r4)
	fmt.Println(r5)
	fmt.Println(r2.Add(r3))
	fmt.Println(r2.Equal(r3))
}
