package main

import (
	"fmt"
	"sort"
)

func main(){
	var fruitList= []string{"apple","mango","kiwi"}
	fmt.Println(fruitList)
	fruitList = append(fruitList, "banana","orange")
	fmt.Println(fruitList)


	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	//-------------------------------//

	highScore:=make([]int,4)

	highScore[0]=234
	highScore[1]=945
	highScore[2]=476
	highScore[3]=867

	fmt.Println(highScore)
 	highScore = append(highScore, 555,666,321)
	fmt.Println(highScore)
	isSorted := sort.IntsAreSorted(highScore)
	fmt.Println(isSorted)
	sort.Ints(highScore)
	isSorted = sort.IntsAreSorted(highScore)
	fmt.Println(isSorted)
	fmt.Println(highScore)


	//how to remove a value from slice based on index //

	courses:= []string{"reactjs","nodejs","golang","springboot","docker","git"}
	fmt.Println(courses)
	courses = append(courses[:3],courses[4:]...)
	fmt.Println(courses)


}