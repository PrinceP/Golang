package main

import "fmt"

func main(){

	fmt.Println("Slices")

	var s []string
    fmt.Println("uninit:", s, " Check null = ", s == nil, " Length: ", len(s) == 0)

	//To create an empty slice with non-zero length, 
	//use the builtin make. 
	//Here we make a slice of strings of length 3 (initially zero-valued). 
	//By default a new slice’s capacity is equal to its length; 
	//if we know the slice is going to grow ahead of time, it’s possible to pass a capacity explicitly as an additional parameter ot make.
	s = make([]string, 3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])
	
	s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("append:", s)

	c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

	l := s[2:5]
    fmt.Println("sl1:", l)

	l = s[:5]
    fmt.Println("sl2:", l)

	l = s[2:]
    fmt.Println("sl3:", l)


	t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

}	