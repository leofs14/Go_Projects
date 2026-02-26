package maps

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])

	websites["Linkedin"] = "https://linkedin.com" // Add a new item
	fmt.Println(websites)
 
	delete(websites, "Google") // Deletes an item
	fmt.Println(websites)
}