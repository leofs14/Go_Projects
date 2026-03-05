package recursion

import "fmt"

func main() {

	fact := factorial(5)
	fmt.Println(fact)

}

func factorial(number int) int { 
	if number == 0 {
		return 1					// Recursive solution (using the func inside the func)
	}

	return number * factorial(number - 1) 

	// result := 1

	// for i := 1; i <= number; i++ { // Loop based solution
	// 	result = result * i
	// }

	// return result
}