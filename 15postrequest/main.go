package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main(){
	performPostJsonRequest()
}

func performPostJsonRequest(){
	const myurl = ""
	//json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"golang",
			"price":"50$",
			"platform":"lco.in"
		}
		`)
	res,err:=http.Post(myurl,"application/json",requestBody)
	checkNilError(err)
	defer res.Body.Close()
	content,err:=ioutil.ReadAll(res.Body)
	fmt.Println(string(content))	
}

func checkNilError(err error){
	if err!=nil{
		panic(err)
	}
}