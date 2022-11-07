package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	performGetRequest()
}

func performGetRequest() {
	const myurl = "https://lco.dev"
	res, err := http.Get(myurl)
	checkNilError(err)
	defer res.Body.Close()
	fmt.Println("status code :", res.StatusCode)
	fmt.Println("content length :", res.ContentLength)
	content, _ := ioutil.ReadAll(res.Body) //bytecode
	fmt.Println(string(content))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
