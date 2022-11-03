package main

import "fmt"

func main() {
	fmt.Println("hello from arrays")
	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "orange"
	fruitList[2] = "banana"
	fruitList[3] = "kiwi"

	fmt.Println(fruitList)
	fmt.Print(len(fruitList))

	var vegList = [3]string{"potato", "tomato", "carrot"}

	fmt.Print(vegList)

}
