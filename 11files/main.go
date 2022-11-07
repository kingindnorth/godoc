package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	content :="this needs to go in a file lco.in"
	file,err:=os.Create("./gofile.txt")
	checkNilErr(err)	
	fmt.Println(file) //its a pointer variable (stores addess)
	len,err:=io.WriteString(file,content)
	checkNilErr(err)
	fmt.Println(len)
	defer file.Close()

	readFile(file.Name())

}

func checkNilErr(err error){
	if err!=nil{
		panic(err)
	}
}

//read the file
func readFile(filename string){
	dataByte,err:=ioutil.ReadFile(filename) //byte format data
	checkNilErr(err)
	fmt.Println(dataByte)
	fmt.Println(string(dataByte)) 

}