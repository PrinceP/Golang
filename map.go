package main

import "fmt"

func main(){


	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)

	v1 := m["one"]
    fmt.Println("V1:", v1)

	v2 := m["two"]
    // fmt.Println("V2:", v2) // Error if not used at all
	fmt.Println("V2:", v2)


	n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println(n)
}
