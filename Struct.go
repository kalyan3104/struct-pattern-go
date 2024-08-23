package main

import (
	"fmt"
)

func main() {

	kalyan := User{Name: "kalyan", Email: "kalyan@go.dev", Status: true, Age: 18}
	fmt.Println(kalyan)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
