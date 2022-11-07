package main

import "fmt"

type User struct{
	Name string
	Email string
	Status bool
	Age int
}

func main(){
	user:=User{
		"prajjawal",
		"prajjawal@gmail.com",
		true,
		22,
	}
	fmt.Printf("%v\n",user)
	fmt.Printf("%+v\n",user)
	fmt.Printf("Name is: %v\n",user.Name)

}	