// A map is a data type maps keys to values
// The zero value of a map is nil
package main

import "fmt"

func main() {
	// init map if you have initial values

	m := map[string]bool{
		"x":true,
		"y":false,
	}

	fmt.Println("m", m)
	// init a map, no need for length

	players := make(map[string]int)

	fmt.Println("p", players)

	// add val to the map
	players["Ahmed"] = 100
	players["jack"] = 99
	players["sara"]=10

	fmt.Println("players", players)

	// get data from map
	fmt.Println("players", players["Ahmed"], "next", players["sara"])

	// del val from map
	delete(players, "jack")
	fmt.Println("p", players)

	// delete woring key does not affect the map
		delete(players, "x")
		fmt.Println(players)
	
	// check if key exist and get value
	// value will be the default value if key does not exist

	v, ok := players["y"]
	fmt.Println("missing key, v", v, "ok", ok)
	
	
	v, ok = players["sara"]
	fmt.Println("valid key, v", v, "ok", ok)

	v = players["y"]
	fmt.Println("Valid key, V:", v)


	// map of slices

	mp := map[string][]string{
		"admins":[]string{"ahmed", "jack"},
		"users":[]string{"x","y","z"},
	}
   
	fmt.Println("mp", mp)

}