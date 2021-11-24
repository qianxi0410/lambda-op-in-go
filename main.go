package main

import (
	"fmt"
	"qianxi/lambda-go/stream"
)

type Person struct {
	Name string
	Age  string
}

func main() {
	var p []Person
	stream.Init([]Person{
		{
			Name: "qianxi",
			Age:  "21",
		},
		{
			Name: "",
			Age:  "",
		},
	}).Map(func(idx int, obj stream.Elem) stream.Elem {
		return &Person{
			Name: "map",
			Age:  "22",
		}
	}).ObjectSlice(&p)

	fmt.Println(p)

}
