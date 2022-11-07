package main

import "fmt"

func main(){

	//function calls

	result1:=adder(3,5)
	defer fmt.Println(result1)
	result2:=proAdder(1,2,3,4,5,6,7,8,9)
	defer fmt.Println(result2)
	result3:=subtractor(5,10)
	defer fmt.Println(result3)
	result4,result5:=doubleReturn(5)
	defer fmt.Println(result4,result5)


	//method calls

	prajjawal:=User{
		"prajjawal",
		"prajjawal@go.dev",
		true,
		22,
	}
	prajjawal.getStatus()
	prajjawal.getEmail() //test@go.dev

	fmt.Println(prajjawal.Email) //original value did not changed (prajjawal@go.dev)


}

func adder(valOne int , valTwo int) int{
	return valOne + valTwo
}

func subtractor(valOne float32, valTwo float32) float32{
	if valOne > valTwo{
		return valOne - valTwo
	}
	return valTwo - valOne
}

func proAdder(values ...int) int{
	total:=0
	for val := range values {
		total+=val
	}
	return total
}

func doubleReturn(valOne int)(string,int){
	return "iam returning you the same value",valOne
}

//methods
 
type User struct{
	Name string
	Email string
	Status bool
	Age int
}

func (user User) getStatus(){
	fmt.Println("is user active:",user.Status)
}
func (user User) getEmail(){
	user.Email = "test@go.dev"
	fmt.Println("Email of this user is:",user.Email)
}