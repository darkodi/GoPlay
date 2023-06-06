package main

import (
	"fmt"

	"github.com/darkodi/puppy"
)

func main() {
	fmt.Println("Hello Gophers!")
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	s3 := puppy.BigBarks()

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
