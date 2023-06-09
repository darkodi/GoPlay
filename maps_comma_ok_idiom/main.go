package main

import "fmt"

func main() {
	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
	}

	fmt.Println("The age of Henry was", am["Henry"])

	fmt.Println(am)
	fmt.Printf("%#v\n", am)

	an := make(map[string]int)
	an["Lucas"] = 28
	an["Steph"] = 37
	an["George"] = 78
	fmt.Println(an)
	fmt.Printf("%#v\n", an)
	fmt.Println(len(an))

	delete(an, "George")

	fmt.Println("--- accessing keys that don't exist")
	delete(an, "Georage")      // won't panic
	fmt.Println(an["Georgey"]) // won't panice
	fmt.Println("------------------------")

	// comma, ok idiom
	v, ok := an["Lucas"]
	if ok {
		fmt.Println("the value prints", v)
	} else {
		fmt.Println("Key didn't exist")
	}

	// statement statement idiom
	if v, ok := an["Georgey"]; !ok {
		fmt.Println("Key didn't exist")
	} else {
		fmt.Println("the value prints", v)
	}
}
