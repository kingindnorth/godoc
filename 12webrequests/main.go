package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main(){
	res,err := http.Get(url)
	checkNilError(err)
	fmt.Printf("response is %v and its type is %T",res,res)
	defer res.Body.Close() //callers responsibility to close the connection
	dataBytes,err:=ioutil.ReadAll(res.Body)
	checkNilError(err)
	fmt.Println(string(dataBytes)) //html content

}

func checkNilError(err error){
	if err!=nil{
		panic(err)
	}
}