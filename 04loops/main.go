package main

import "fmt"

func main() {
	days := []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}

	//loop 1
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	//loop 2
	for i := range days {
		fmt.Println(days[i])
	}

	//loop 3
	for index, value := range days {
		fmt.Printf("value of index:%v is:%v\n", index, value)
	}

	//loop 4
	val := 1
	for val < 10 {
		fmt.Println("value is:", val)
		val++
	}

}
