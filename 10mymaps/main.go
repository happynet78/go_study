package main

import "fmt"

func main() {
	fmt.Println("Maps in golangs")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	// loops are interestring in golang

	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}
}
