package basics

import (
	"fmt"
	"maps"
)

func main() {
	// initializing a map
	myMap := make(map[string]int)
	fmt.Println(myMap)

	// adding key-value pairs
	myMap["apple"] = 5
	myMap["banana"] = 10
	fmt.Println(myMap)

	// Accessing a value
	value := myMap["apple"]
	fmt.Println("Value for 'apple':", value)

	// Deleting a key-value pair
	delete(myMap, "apple")
	fmt.Println("After deleting 'apple':", myMap)

	// clearing a map
	// clear(myMap)
	// fmt.Println("After clearing the map:", myMap)

	value, ok := myMap["banana"]
	// fmt.Println("Value for 'banana':", value, "Exists:", ok)
	if ok {
		fmt.Println("Value for 'banana':", value)
	} else {
		fmt.Println("'banana' does not exist in the map")
	}

	var myMap2 map[string]int = map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("myMap2:", myMap2)

	if maps.Equal(myMap, myMap2) {
		fmt.Println("myMap and myMap2 are equal")
	} else {
		fmt.Println("myMap and myMap2 are not equal")
	}

	// Iterating over a map
	for key, value := range myMap2 {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	var myMap3 map[string]int
	if myMap3 == nil {
		fmt.Println("myMap3 is a nil map")
	} else {
		fmt.Println("myMap3 is not a nil map")
	}

	myMap3 = make(map[string]int)
	myMap3["x"] = 100
	fmt.Println(myMap3)

	fmt.Println("The length of map3 is", len(myMap3))

	// nested maps
	var nestedMap map[string]map[string]int = make(map[string]map[string]int)
	nestedMap["map3"] = myMap3
	fmt.Println("Nested Map:", nestedMap)

}
