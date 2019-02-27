package main

import "fmt"

func main() {
	var p *int
	i := 42
	p = &i
	fmt.Println(*p) // read i through the pointer p
	*p = 21         // set i through the pointer p
	fmt.Println(p, *p)

	type Vertex struct {
		X, Y int
	}

	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		v4 = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, v2, v3, v4)
}
