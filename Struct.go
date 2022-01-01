// Structs are a collection of fields.
// They’re useful for grouping data together to form records.
// the closest thing in Go to objects/OOP
package main

import "fmt"

type player struct {
	firstname string
	lastname string
	score int
}

type game struct {
	players map[string]*player
}


func (g *game) getWinner() *player {
	var maxScore int
	var winner *player

	for _,value := range g.players {
		if value.score > maxScore {
			winner = value
			maxScore = value.score
		}
	}

	return winner
}


type chain struct {
	base
	value int
	next *chain
}

func (c *chain) sayOK(){
	fmt.Println("chain OK")
}

type base struct {
}

func (b *base) sayHi(){
	fmt.Println("base hiii")
}

func (c *base) sayOK(){
	fmt.Println("base OK")
}




 func main() {

//   g := &game {
// 	  players : make(map[string]*player),
//   }

//   p1 := player{}
//   fmt.Println("p1", p1)

//   g.players["p1"] = &p1
//   fmt.Println("game", g)
  
//   var p2 player

//   p2 = player{}
//   p2.score = 4
//   fmt.Println("p2", p2)

//   g.players["p2"] = &p2
//   fmt.Println("game", g)

//   p3 := player{
// 	  lastname :"khalf",
// 	  firstname: "ahmed",
// 	  // 
//   }
//   fmt.Println("p3", p3)
//   p3.score = 10
//   fmt.Println("p3", p3)

//   g.players["p3"] = &p3
//   fmt.Println("game", g)



//   fmt.Println("g.p[2]", g.players["p2"])
//   fmt.Println("g.p[2]*", *g.players["p2"])


//   winner := g.getWinner()
//   fmt.Println("winner", winner)
//   fmt.Println("*winner", *winner)


//   b := base{}
//   b.sayHi()
//   b.sayOK()

 c1 := chain{value :100}
fmt.Println("c1", c1)

c1.sayHi()

c2 := &chain{value:200}
c1.next = c2
fmt.Println("c1", c1)
fmt.Println(" c1.next", *c1.next)

c1.next.sayHi()
c1.sayOK()


}
