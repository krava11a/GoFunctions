package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

type VertexMaps struct {
	Lat, Long float64
}

var m map[string]VertexMaps

func TestMaps() {
	//first Map
	m = make(map[string]VertexMaps)
	m["Bell Labs"] = VertexMaps{
		40.68433, -74.39967,
	}
	m["Everest"] = VertexMaps{
		27.988763274346873, 86.9249749485967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Everest"])

	//Literals
	var m2 = map[string]VertexMaps{
		"Munku Sardyuk": {
			51.72012186236767, 100.59741398773271,
		},
		"Elbrus": {
			43.35105975460793, 42.445158345662236,
		},
	}
	fmt.Println(m2)

	//Mutating maps

	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	//Exercise MAPS
	wc.Test(WordCount)
}

func WordCount(s string) map[string]int {
	count := strings.Fields(s)
	var res = make(map[string]int)
	for i := 0; i < len(count); i++ {

		res[count[i]] += 1
	}
	return res
}
