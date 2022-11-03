package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello from golang")
	reader := bufio.NewReader(os.Stdin) //gives bytecode
	fmt.Println(reader)
	fmt.Println("enter the value:")
	input, _ := reader.ReadString('\n')
	fmt.Println(input)

	//conversion
	res, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(res+1)

}
