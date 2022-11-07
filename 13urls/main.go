package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=@#Ee34443ffff4232"

func main() {
	res, err := url.Parse(myurl)
	checkNilErrors(err)
	fmt.Println(res)
	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Path)
	fmt.Println(res.Port())
	fmt.Println(res.RawQuery)
	qparams := res.Query()
	fmt.Printf("the type is %T and the value is %v", qparams, qparams) //key-value pairs
	fmt.Println(qparams["coursename"])
	//loop
	for _, val := range qparams {
		fmt.Println("params is:", val)
	}
	anotherUrl := partsofurl.String()
	fmt.Println(anotherUrl)
}

func checkNilErrors(err error) {
	if err != nil {
		panic(err)
	}
}

var partsofurl = &url.URL{
	Scheme:  "https",
	Host:    "lco.dev",
	Path:    "/learn",
	RawPath: "user=prajjawal",
}
