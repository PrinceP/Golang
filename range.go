package main

import "fmt"

func main(){

	nums:= []int{1,2,3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

	n := map[string]int{"foo": 1, "bar": 2}
	for k,v := range n{
		fmt.Println(k, ": ", v)
	}

	for k := range n {
        fmt.Println("key:", k)
    }

	for i, c := range "go" {
        fmt.Println(i, c)
    }


}