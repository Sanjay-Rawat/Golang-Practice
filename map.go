package main

import "fmt"

func mapHandler() {
	user := map[string]string{"name": "sanjay"}
	fmt.Println(user)
	fmt.Println(user["name"])

	user["age"] = "age"
	fmt.Println(user)
	fmt.Println(user)

	value, ok := user["name"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("key does not exits")
	}

	for key, value := range user {
		fmt.Print((value), key)
	}

}
