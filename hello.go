package main

import (
	"fmt"
	"sort"
)

func main() {
	//greeting :="hello there friends!"

	//fmt.Println(strings.Contains(greeting,"hello !"))

	//replacing hello with hi
	//fmt.Println(strings.ReplaceAll(greeting,"hello", "hi"))

	//changing to uppercase
	//fmt.Println(strings.ToUpper(greeting))

	//fmt.Println(strings.Index(greeting, "th"))
	//fmt.Println(strings.Split(greeting, " "))

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"tunde", "femi", "bola", "tolu"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names,"bola"))

	//var room = "cave"
	//if room == "cave" {
	//		fmt.Println("You find yourself in a dimly lit cavern.")
	//	} else if room == "entrance" {
	//		fmt.Println("There is a cavern entrance here and a path to the east.")
	//	} else if room == "mountain" {
	//		fmt.Println("There is a cliff here. A path leads west down the mountain.")
	//	} else {
	//		fmt.Println("Everything is white.")

	//}
	// 	age := 35
	// 	name := "shuna"
	// 	// Print
	// 	fmt.Print("hello, ")
	// 	fmt.Print("world! \n")
	// 	fmt.Print("new line \n")

	// 	// Println
	// 	fmt.Println("hello tunde!")
	// 	fmt.Println("goodbye tunde")
	// 	fmt.Println("my age is", age, "and my name is", name)

	// 	//Printf (formatted strings) %_ = format specifier
	// 	fmt.Printf("my age is %v and my name is %v \n", age, name)
	// 	fmt.Printf("my age is %q and my name is %q \n", age, name)
	// 	fmt.Printf("age is of type %T", age)
	// 	fmt.Printf("you scored %f points! \n", 225.55)
	// 	fmt.Printf("you scored %0.1f points! \n", 225.55)

	// 	// Sprinf (save formatted strings)
	// 	var tun = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	// 	fmt.Println("the saved string is:", tun)
	// }

	//package main

	//import (
	//	"fmt"
	//	"math/rand"
	//)

	//func main() {
	//var num = rand.Intn(10) + 1
	//fmt.Println(num)
	//num = rand.Intn(10) + 1
	//fmt.Println(num)

	//strings
	//var nameOne string = "mario"
	//var nameTwo = "tunde"
	//var nameThree string

	//fmt.Println(nameOne, nameTwo, nameThree)

	//nameOne = "peach"
	//nameThree = "bowser"

	//fmt.Println(nameOne, nameTwo, nameThree)

	//nameFour := "yoshi"
	//fmt.Println(nameFour)

	//ints
	//var ageOne int = 20
	//var ageTwo = 30
	//ageThree := 40

	//fmt.Println(ageOne, ageTwo, ageThree)

	//bits & memory
	//var numOne int8 = 25
	//var numTwo int8 = -128
	//var numThree uint16 = 256

	//var scoreOne float32 = 25.98
	//var scoreTwo float64 = 455667999.5

}
