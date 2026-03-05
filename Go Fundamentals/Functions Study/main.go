package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(1, 10 ,15, 40, -5)
	otherSum := sumup(1, numbers...)

	fmt.Println(sum)
	fmt.Println(otherSum)
}

func sumup(startingValue int, numbers ...int) int { //Variadic functions, you can have as many parameters as you want, numbers here will be a slice
	sum := 0					//You can put more parameters separated here

	for _, val := range numbers {
		sum += val // sum = sum + val
	}

	return sum
}