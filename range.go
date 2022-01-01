// Range is used within the the for loop,
// iterates over a array, slice or map to get key/value
package main

import "fmt"

func main() {
    // Range on array
	// arr := [4]string{"Eslam", "Mahmoud", "X", "Y"}

	// for k, v := range arr {
	// 	fmt.Println("range on arr key:", k, "val", v)
	// }

	// Range on slice
	// slc := []string{"ahmed", "jack", "X"}

	
	// for key, value := range slc {
	// 	fmt.Println("range on arr key:", key, "val", value)
	// }


	// mp := map[string]string{
	// 	"admin": "Ahmed",
	// 	"user":  "sara",
	// }

	// for key, value := range mp {
	// 	fmt.Println("Range on map:, key:", key, "Value:", value)
	// }

	// Range on map of slices
	mapOfSlices := map[string][]string{
			"Admins": []string{"ahmed", "jack"},
			"Users":  []string{"X", "Y", "Z"},
	}

	for key, value := range mapOfSlices {
		fmt.Println("rage on map of slices :key", key, "val", value)

		for k, v := range value {
			fmt.Println("range on the slicees in map key", k,"val", v)
		}

	}
		

}