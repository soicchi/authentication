package main

import (
	"fmt"

	"github.com/soicchi/authentication/pkg/db"
)

func init() {
	if err := db.Connect(); err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
}

func main() {
	fmt.Println("Hello, World!")
}
