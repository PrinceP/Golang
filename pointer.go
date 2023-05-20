package main

import "fmt"

func makezero(input int){
	input = 0
}

func makezero_ptr(input *int){
	*input = 0
}

func main(){
	i := 1
	fmt.Println(i)

	makezero(i)
	fmt.Println(i)

	makezero_ptr(&i)
	fmt.Println(i)

}