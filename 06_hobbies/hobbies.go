package main

import "fmt"

type products struct {
	title string
	id    int
	price int
}

func main() {
	hobbies := [3]string{"Movies", "Games", "Chess"}

	println("1&2:")

	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	slice1_0 := hobbies[0:2]
	slice1_1 := hobbies[:2]

	println("3:")
	fmt.Println(slice1_0)
	fmt.Println(slice1_1)

	println("4:")
	slice1_1 = slice1_1[1:3]
	fmt.Println(slice1_1)

	println("5:")
	goals := []string{"Finish course", "Make calc project"}
	fmt.Println(goals)

	println("6:")
	goals[1] = "Make REST project"
	goals = append(goals, "Make my own project")
	fmt.Println(goals)

	println("6:")
	productsArr := []products{{"Shirt", 0, 7}, {"Shoe", 1, 43}}
	fmt.Println(productsArr)
	productsArr = append(productsArr, products{"tie", 2, 3})
	fmt.Println(productsArr)
}
