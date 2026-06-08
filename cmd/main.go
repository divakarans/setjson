package main

import (
	"fmt"

	"github.com/divakarans/setjson/setter"
)

func main() {

	data := `{
		"name": "Asha Rao",
		"email":"asha@example.com",
		"age": 28,
		"active": true
	}`

	fmt.Println("Before:")
	fmt.Println(data)

	update := setter.Setjson(data, "age", 30)

	fmt.Println("After:")
	fmt.Println(update)
}
