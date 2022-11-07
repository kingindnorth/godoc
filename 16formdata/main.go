package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main(){
	performPostFormRequest()
}

func performPostFormRequest(){
	const myurl =""
	//form data
	data:=url.Values{}
	data.Add("firstname","prajjawal")
	data.Add("lastname","tiwari")
	data.Add("email","prajjawal@go.dev")
	res,err:=http.PostForm(myurl,data)
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