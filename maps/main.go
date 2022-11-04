package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["js"] = "javaSript"
	languages["py"] = "python"
	languages["go"] = "golang"

	fmt.Println(languages)
	fmt.Println(languages["js"])

	//delete

	delete(languages, "py")
	fmt.Println(languages)

	//loops in maps

	for key,value := range languages{
		fmt.Printf("for the key: %v, the value is: %v \n",key,value)
	}

}
