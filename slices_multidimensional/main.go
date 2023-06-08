package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Jenny", "Moneypenny", "Guiness", "Wolverine Tracks"}
	fmt.Println(jb)
	fmt.Println(jm)
	// slice [] of slices []string
	xxs := [][]string{jb, jm}
	fmt.Println(xxs)

	for _, v := range xxs {
		fmt.Printf("%v\n", v)
	}
}
