package main

import "fmt"

func main(){

	var a [4]int
	fmt.Println(a)

	a[1] = 2
	a[3] = 4
	//a[4] = 4 // compile time error

	fmt.Println(a)

	b := [5]int{1,2,3,4,5}
	fmt.Println(b)

	c := [2][4] int{{1,2,3,4},{5,6,7,8}}
	fmt.Println(c)

	var d[2][2] int
	for i:=0; i < 2; i++{
		for j:=0; j<2; j++{
			d[i][j] += 1
		}
	}
	fmt.Println(d)

}