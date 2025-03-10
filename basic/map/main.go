package main

import "fmt"

func main () {
	fmt.Println("Maps in go lang")

	languages := make(map[string]string)

	languages["JS"] = "javaScript"
	languages["RB"] = "ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages", languages)
	fmt.Println("JS shorts", languages["JS"])
	
	delete(languages, "RB")
	fmt.Println("List of all languages", languages)

	// loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}