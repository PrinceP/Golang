package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main(){
	fmt.Println(s)

	const n = 500000

	const d = 3e20  / n
	fmt.Println(d) // aribitary precision

	fmt.Print(int64(d))

	fmt.Println(math.Sin(n))


}
