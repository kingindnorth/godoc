package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello from time")
	//get date
	presentTime := time.Now()
	fmt.Println(presentTime)
	presentTime.Format("01-02-2006")
	fmt.Println(presentTime)
	//create date
	createDate:=time.Date(2020,time.November,4,9,20,3,9,time.UTC)
	// fmt.Println(createDate)
	createDate.Format("01-02-2006")
	fmt.Println(createDate)

}
