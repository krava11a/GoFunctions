package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1,2}
	v2 = Vertex{X:31}
	v3 = Vertex{}
	p2 = &Vertex{1,2}
)

func TestStructs() {
	//Using dot
	v := Vertex{2, 3}
	fmt.Println(v)
	fmt.Println(Vertex{1, 2})
	v.X = 4
	fmt.Println(v)

	//struct with pointers
	p := &v
	p.Y=13e2
	fmt.Println(v)

	//Struct Literals
	fmt.Println(v1,p2,v2,v3)

}
