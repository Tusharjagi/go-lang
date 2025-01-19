package main

import "fmt"

func main () {
	fmt.Println("Welcome to array in go langs")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Grapes"
	fruitList[3] = "PineApple"

	fmt.Println("Fruit list is: ",  fruitList)
	fmt.Println("Fruit list length: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}

	fmt.Println("Vegy list is:", vegList)
	fmt.Println("Vegy list is:", len(vegList))


}