package main

import "fmt"

func add(a int, b int) int{
	return a + b
}

func add_three(a, b, c int) int{
	return a + b + c
}

func vals() (int, int) {
    return 3, 7
}

func sum(nums ...int) int{
	fmt.Print(nums, " = ")
	total := 0

	for _,num := range nums{
		total += num
	}

	return total
}

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main(){

	fmt.Println(add(1,2))

	fmt.Println(add_three(1, 2, 3))
	
	a,b := vals()
	fmt.Println(a, b)

	_, c := vals()
    fmt.Println(c)

	fmt.Println(sum(1,2))
	fmt.Println(sum(1,2,3))
	
	nextInt := intSeq() // Closure
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	
	fmt.Println(fact(7))

	var fib func(n int) int
	fib = func(n int) int{
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

}