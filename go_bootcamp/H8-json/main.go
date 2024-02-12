package main

import (
	"fmt"
	"net/url"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	// var jsonString = `{
	// "full_name" : "Airell Jordan",
	// "email": " airell@mail.com",
	// "age": 23
	// }
	// `

	// var result Employee/

	// var err = json.Unmarshal([]byte(jsonString), &result)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println("full_name:", result.FullName)
	// fmt.Println("email:", result.Email)
	// fmt.Println("age:", result.Age)

	// var object = []Employee{{"john wick", "slebew@main.com", 27}, {"ethan hunt", "slebew@main.com", 32}}

	// var jsonData, err = json.Marshal(object)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// var jsonString = string(jsonData)
	// fmt.Println(jsonString)
	var urlString = "http://developer.com:80/hello?name=airell&age=23"
	var u, e = url.Parse(urlString)

	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Printf("url: %s\n", urlString)
	fmt.Printf("protocol: %s\n", u.Scheme)
	fmt.Printf("host: %s\n", u.Host)
	fmt.Printf("path: %s\n", u.Path)

	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]
	fmt.Printf("name: %s, age: %s\n", name, age)
}
