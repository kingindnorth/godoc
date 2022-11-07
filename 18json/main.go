package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// EncodeJson()
	DecodeJson()
}

type Course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string `json:"-"` //password removed while marshaling
	Tags     []string
}

func EncodeJson() {
	lcocourses := []Course{
		{
			"reactjs", 299, "lco", "abc123", []string{
				"web-dev", "js",
			}},
		{
			"nodejs", 299, "lco", "abc123", []string{
				"back-end", "js",
			}},
		{
			"go", 299, "lco", "abc123", []string{
				"back-end", "go",
			}},
	}
	//package this data as json data
	finalJson, err := json.Marshal(lcocourses)
	checkNilError(err)
	fmt.Println(string(finalJson))
	fmt.Printf("%s", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename":"reactjs-BootCamp",
			"price":299,
			"website":"lco.in",
			"tags":["web-dev","js"]
		}
	`)
	var lcocourse Course
	checkvalid := json.Valid(jsonDataFromWeb)
	if checkvalid {
		fmt.Println("json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcocourse)
		fmt.Println(lcocourse)
		fmt.Printf("%#v\n", lcocourse) //gets key value pairs
	} else {
		fmt.Println("json was not valid")
	}

	//if you want to add data to key value(maps)
	var myonlinedata map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myonlinedata)
	fmt.Println(myonlinedata)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
