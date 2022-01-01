// Functions are the way to group and re use same logic/code
// functions take zero or more arguments, and return zero or more
package main

import (
	"fmt"
  "errors"
)

// function that does not return anything
func p(s string){
  fmt.Println("hello", s)
}

// function that does not return anything
// and take as many params as we send and print parameters length
func l(params  ...interface{}){

  fmt.Println("len", len(params))

}

// function that take two int arguments and return the sum
func add(x int, y int) int {
  return x+y
}
// lets group types if it is the same
func sum(x,y int) int {
  return x+y
}

// return multiple results

func concat(a,b string) (string, int) {
  r := a+b
  return r, len(r)
}

// named return types

func concat2(a,b string) (r string, l int) {
  r = a+b
  l = len(r)
  // will return the named vars
  return
}

// function that return string & nil or default string value
// & error
// this pattern is widely used in Go
func returnError(b bool)(string, error){
  if b{
    return "", errors.New("this is an go error")
  }
 return "okay", nil
}



func main() {
   p("Ahmed")

   l("hello", 2, "hi", true, 4.5)

  v := add(2, 3)
  fmt.Println("add func", v)

  fmt.Println("sum", sum(2,4))

  res,len := concat("ahmed", "happy")
  fmt.Println("concat", res, len)

  res, len = concat2("ahmed is", "unhappy")
  fmt.Println("con2", res, len)

 
  r,err := returnError(false) // false 
  if err != nil {
    fmt.Println("we have error", err.Error())
  } else {
    fmt.Println("we got", r)
  }

// func type

var x func(int, int) int
x = sum

fmt.Println("x", x(3,5))

}

