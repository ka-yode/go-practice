package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])

	//mutating maps in go, adding new key:value pair to the map
	websites["Linkedin"] = "https://linkedin.com"
	fmt.Println(websites)

	//mutating maps, deleting an item using the inbuild delete method from go
	delete(websites, "Google")
	fmt.Println(websites)

}
