package main

import "fmt"

func main(){
	one:=2
	var ptr *int
	fmt.Println(ptr) //nil
	ptr = &one
	fmt.Println(ptr) //memory address
	fmt.Println(*ptr) //value at memory address


}