package main

import (
	"fmt"

	"github.com/divakarans/setjson/internal/setter"
)

func main() {
	data := `{
  		"name": "Asha Rao",
  		"email":"asha@example.com",
  		"age": 28,
  		"active": true
		}`

	fmt.Println(data)
	update := setter.Setjson(data)
	fmt.Println(update)

}
