package main

import (
	"fmt"
	"math"
)

type I interface {
	Scale(f float64)
}

type Vertex struct {
	X, Y float64
}
// can only define func for the receiver which exists in the same package 

// Abs value receiver, methods operates on a copy of receiver
func (v Vertex) Abs() float64 {
	fmt.Printf("location 3 of v is %v \n", &v.X)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale point receiver, can modify the value which receiver points to
// reasons to use pointer receiver
// 1. can modify the receiver value
// 2. avoid copying the value on each method call, more effective.
func (v *Vertex) Scale(f float64) {
	if v == nil {
		fmt.Println("nil concrete type")
		return
	}
	fmt.Printf("location 2 of v is %+v \n", &v.X)
	v.X = v.X * f
	v.Y = v.Y * f
}
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main()  {
	v := Vertex{2,3}
	v1 := &Vertex{2,3}
	fmt.Printf("location 1 of v is %v \n", &v.X)

	// value receiver can accept either value or pointer as receiver
	v.Abs()
	v1.Abs()  // -> (*v1).Abs()


	// pointer receiver  can accept pointer or value as the receiver
	v.Scale(1) // -> (&v).Scale(1)
	v1.Scale(1)

	// pointer argument must take a pointer as argument
	Scale(&v, 10)
	Scale(v1, 10)

	// value argument must take a value as argument
	Abs(v)
	Abs(*v1)
	// ========

	var i1 I
	var i2 I
	var v2 *Vertex
	i1 = v2
	// i1 is a nil interface with a non-nil concrete type
	// i2 is a nil interface
	i1.Scale(1) // this will have no error, but need handle nil in concrete type
	i2.Scale(2) // this is a runtime error
}