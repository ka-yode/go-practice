package main

import (
	"fmt"
)

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	hobbies := [3]string{"Coding", "Anime", "Manga"}
	fmt.Println("Hobbies: ", hobbies)

	fmt.Println("First hobby: ", hobbies[1])

	lastHobbies := hobbies[1:]
	fmt.Println("Last two hobbies: ", lastHobbies)

	firstTwo := hobbies[0:2]
	fmt.Println("First two hobbies: ", firstTwo)

	firstTwo = hobbies[1:]
	fmt.Println("Second and last hobbies: ", firstTwo)

	var courseGoals []string = []string{"be better at go", "land a new job"}
	courseGoals[1] = "Learn to cook"
	courseGoals = append(courseGoals, "be better in all aspects")
	fmt.Println("Course Goals: ", courseGoals)

	var products []Product = []Product{{title: "Samsung", id: "7", price: 500.00},
		{title: "Iphone", id: "4", price: 800.07}}

	products = append(products, Product{title: "Itel", id: "02", price: 350.90})
	fmt.Println("Products: ", products)

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
