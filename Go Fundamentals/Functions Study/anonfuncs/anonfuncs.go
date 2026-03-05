package anonfuncs

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3}

	double := createtransformer(2) // Factory Functions
	triple := createtransformer(3)

	transformed := transformNumbers(&numbers, func (number int) int {
		return number * 2  // Anonymous functions, usually used when you'll only need a function once. You can't call this anywhere else
	})

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int { //func as parameter
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createtransformer(factor int) func(int) int{
	return func (number int) int {
		return number * factor // Closure
	}
}
