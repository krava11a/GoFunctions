package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"strings"
)

func TestArrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "And GO"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 4, 5, 6, 7}
	fmt.Println(primes)

	//slices
	var op []int = primes[1:4]
	fmt.Println(op)

	//Slices are like references to arrays
	names := [4]string{
		"Atos",
		"Portos",
		"Aramis",
		"D'artanian",
	}
	fmt.Println(names)

	aa := names[0:2]
	bb := names[1:3]
	fmt.Println(aa, bb)

	bb[0] = "MILEDI"
	fmt.Println(aa, bb)
	fmt.Println(names)

	//Slice literals
	q := []int{2, 3, 6, 7, 8, 9, 11}
	fmt.Println(q)
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{7, true},
		{55, true},
		{11, false},
		{18, true},
	}
	fmt.Println(s)

	//Slice defaults
	fmt.Println(q)
	qs := q[:]
	fmt.Println(qs)
	qs = qs[1:4]
	fmt.Println(qs)
	qs = qs[:2]
	fmt.Println(qs)
	qs = qs[1:]
	fmt.Println(qs)

	//Slice length and capacity
	qc := q
	printSlice(qc)
	qc = qc[:0]
	printSlice(qc)
	qc = qc[:4]
	printSlice(qc)
	qc = qc[2:]
	printSlice(qc)

	//nil slices
	var snil []int
	printSlice(snil)
	if snil == nil {
		fmt.Println("nil!")
	}

	//creating a slice with make
	fmt.Println("creating a slice with make")
	am := make([]int, 5)
	printSlice(am)
	bm := make([]int,0,5)
	printSlice(bm)
	cm := bm[:2]
	printSlice(cm)
	dm := cm[2:5]
	printSlice(dm)

	//slices of slices
	fmt.Println("slices of slices")
	board := [][]string{
		[]string{"_","_","_"},
		[]string{"_","_","_"},
		[]string{"_","_","_"},
	}
	board[0][0]="X"
	board[2][2]="0"
	board[1][2]="X"
	board[1][0]="0"
	board[0][2]="X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s \n", strings.Join(board[i]," "))
	}

	//appending to a slice
	fmt.Println("appending to a slice")
	var ats []int
	printSlice(ats)
	ats = append(ats,0)
	printSlice(ats)
	ats = append(ats,2)
	printSlice(ats)
	ats = append(ats,4,8,16)
	printSlice(ats)

	//range
	testRange(ats)

	//rangeContinue
	rangeContinued()

	//Exercise:Slices
	pic.Show(Pic)

}

func printSlice(s []int) {
	fmt.Printf("len = %d cap=%d %v \n", len(s), cap(s), s)
}

func testRange(at []int) {
	fmt.Println("range test")
	for i,v := range at {
		fmt.Printf("2**%d = %d \n",i,v)
	}
}

func rangeContinued()  {
	fmt.Println("range continued test")
	pow := make([]int,10)
	for i, _ :=range pow{
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d \n", value)
	}

}

func Pic(dx,dy int) [][]uint8 {
	var result = make([]([]uint8),dy)
	for i :=0; i <len(result); i++ {
		result[i]= make([]uint8,dx)
		for j := 0; j < len(result[i]); j++ {
			result[i][j]=uint8((i + j)/2)
		}
	}
	return result
}
