package main

import "fmt"

func main() {
	for printNos(n){
        if(n > 0) {
            printNos(n - 1);
            continue (n + " ")
        }
        return;
    }
     
    fmt.println(100);