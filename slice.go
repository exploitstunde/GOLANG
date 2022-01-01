package main

import "fmt"

func main() {

	// //
	// keywords := []string{}
	// fmt.Println(keywords)

	// var x []string
	// fmt.Println(x)

	// // init slice with value
	// src := []string{"x", "y", "z"}
	// fmt.Println(src)

	// //get/set values by index

	// fmt.Println(src[0])

	// src[1] = "b"

	// fmt.Println(src)

	// // add new

	// src = append(src, "d")
	// fmt.Println(src)

	// // another way to init slie with predefined sizze

	// dest := make([]string, 2)
	// fmt.Println(dest)

	//  // copy

	//  copy(dest, src)
	//  fmt.Println(dest)

	// // change source after copy it

	// src[0] = "a"
	// fmt.Println(src)
	// fmt.Println(dest)

	// // Since Slice is mainly a pointer to underline array
	// // slicing a slice does not create a copy of values
	// // this will create a slice that point to the same
	// //values in the original array

	// s := src[1:3]
	// fmt.Println(s)

	// // updating the source
	// src[2] = "c"
	// fmt.Println(src)
	// fmt.Println(s)

	// // src[:x] will create a slice point to values start at index 0 end at x-1
	// // src[x:y] will create a slice point to values start at index x end at y-1
	// // src[x:] will create a slice point to values start at index x end at len(s)-1

	// fmt.Println(src)
	// fmt.Println("src :3", src[:3])
	// fmt.Println("src 0:2", src[0:2])
	// fmt.Println("src 2:", src[2:])

	gameMap := [][]string{
		make([]string, 3),
		make([]string, 5),
		[]string{"-", "-"},
	}

	fmt.Println(gameMap)

	gameMap[0] = make([]string, 6)
	gameMap[1] = make([]string, 10)

	fmt.Println(gameMap)

	gameMap[1][2] = "-"
	fmt.Println(gameMap)

}
