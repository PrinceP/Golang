package main

import "fmt"

func main(){

	fmt.Println("For loop")
	i := 1
	for i <= 3 {
		fmt.Print(i, " ")
		i++ // i = i+1
	}
	

	for j:= 4; j <= 8; j++{
		fmt.Println(j)
	}

	for {
        fmt.Println("loop")
        break
    }
	
	for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }


}