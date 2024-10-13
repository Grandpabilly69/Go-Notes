package main

import "fmt"

func main() {
	createMap()
	intMap()
}

// creates map
// note map is a key value pair
func createMap() {
	//all keys must be same type and all values must be same type
	menu := map[string]float64{
		//left side is key
		//right side is value
		"soup":   4.99,
		"pie":    7.99,
		"salad":  6.99,
		"toffee": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["salad"])

	loopMap(menu)

}

// looping through map is same as array
func loopMap(menu map[string]float64) {
	for key, value := range menu {
		fmt.Println(key, value)
	}
}

// int map uses ints instead of strings in this case
// this shows that maps can have
// different types for both keys and values
func intMap() {
	phonebook := map[int]string{
		23433433: "matt",
		45646566: "Amdul",
		29868534: "mGoof",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[23433433])
	updateMap(phonebook)
}

// updating is similar as array
// but use key instead of index
func updateMap(phonebook map[int]string) {
	phonebook[23433433] = "Clown"
	fmt.Println(phonebook)
}
