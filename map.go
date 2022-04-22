package main

import "fmt"

func mapHandler() {
	user := map[string]string{"name": "sanjay"}
	fmt.Println(user)
	fmt.Println(user["name"])

	user["age"] = "age"
	fmt.Println(user)

}
