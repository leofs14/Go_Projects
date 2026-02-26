package main

import "fmt"
import "example.com/array/Product"



func main(){
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	updatedPrices := append(prices, 5.99) // Adds to the dynamic slice and the array in memory but it doesnt change the original slice
	fmt.Println(updatedPrices, prices) // It only updates the original dynamic slice if its reassigned
	fmt.Println("")

	hobbies := [3]string{"read", "game", "soccer"}
	fmt.Println("1 and 2")
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])
	fmt.Println("===================")
	fmt.Println("")

	fmt.Println("3 and 4")
	featuredHobbies := hobbies[:2]
	fmt.Println(featuredHobbies)
	ofeaturedHobbies := hobbies[0:2]
	fmt.Println(ofeaturedHobbies)
	newfeaturedHobbies := hobbies[1:]
	fmt.Println(newfeaturedHobbies)
	fmt.Println("===================")
	fmt.Println("")

	fmt.Println("5 and 6")
	goals := []string{"learn", "work"}
	fmt.Println(goals)
	goals[1] = "golang"
	updatedGoals := append(goals, "enrich")
	fmt.Println(updatedGoals)
	fmt.Println("===================")
	fmt.Println("")

	fmt.Println("7")
	prod1 := product.New("Rice")
	prod2 := product.New("Beans")
	prod3 := product.New("Meat")
	prodList := []product.Product{prod1, prod2}
	fmt.Println(prodList)
	prodList = append(prodList, prod3)
	fmt.Println(prodList)
	fmt.Println("===================")
}

// func main() {
// 	productNames := [4]string{"A book"}
// 	prices := [4]float64{10.99, 9.99, 5.99, 20.0}
// 	fmt.Println(prices)
// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])
// 	// featuredPrices := prices[:3] this makes it shows only until index 3, can do the opposite too (1:)
// 	// fmt.Println(featuredPrices)
// 	featuredPrices := prices[1:3]
// 	fmt.Println(featuredPrices)
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// }