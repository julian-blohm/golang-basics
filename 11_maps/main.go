package main

import "fmt"

/*
* map is an unordered collection of key/value pairs
* in java called hash tables or in python dicitonaries
* efficient data retrieval/modification with keys
* in go, maps are implemented by hash tables
 */

func main() {
	// declare map; in brackets is the key datatype, then the value datatype; map is currently nil an key/value cannot be added
	var myMapName map[string]int
	fmt.Println("Map declared")
	fmt.Println(myMapName)
	fmt.Println("")

	// declare an initialize map
	myInitializedMap := map[string]string{"en": "English", "fr": "French"}
	fmt.Println("Map initialized & declared")
	fmt.Println(myInitializedMap)
	fmt.Println("length for this initialized map is: ", len(myInitializedMap))
	fmt.Println("")

	// create map with make function
	myMapWithMake := make(map[string]int)
	fmt.Println("Map with make")
	fmt.Println(myMapWithMake)
	fmt.Println("length for this uninitialized map is: ", len(myMapWithMake))
	fmt.Println("")

	//access map using the key
	fmt.Println("access map")
	fmt.Println(myInitializedMap["en"])
	fmt.Println(myInitializedMap["fr"])
	fmt.Println("")

	// when accessing value from a map, two values will be returned, the value itself and a boolean if the key is found or not
	fmt.Println("get key")
	value, found := myInitializedMap["en"]
	fmt.Println(found, value)
	value, found = myInitializedMap["ger"]
	fmt.Println(found, value)
	fmt.Println("")

	// add item to map
	fmt.Println("add value to map")
	myInitializedMap["ger"] = "German"
	fmt.Println(myInitializedMap)
	fmt.Println("")

	// update value
	fmt.Println("update value")
	myInitializedMap["ger"] = "Deutsch"
	fmt.Println(myInitializedMap)
	fmt.Println("")

	// delete key/value
	fmt.Println("delete key/value")
	delete(myInitializedMap, "ger")
	fmt.Println(myInitializedMap)
	fmt.Println("")

	// iterate over map
	// similar as for slices and arrays but instead of index and value we get the key and the value when using range
	fmt.Println("iterate over map")
	for key, value := range myInitializedMap {
		fmt.Println(key, "=>", value)
	}
	fmt.Println("")

	// truncate map

	fmt.Println("truncate map option 1; iterate over it and delete entries")
	for key := range myInitializedMap {
		delete(myInitializedMap, key)
	}
	fmt.Println(myInitializedMap)

	fmt.Println("truncate map option 2; reinitialize it")
	myInitializedMap["ger"] = "German"
	fmt.Println(myInitializedMap)
	myInitializedMap = make(map[string]string)
	fmt.Println(myInitializedMap)

}
