package main

import (
	"fmt"
	"qianxi/lambda-go/stream"
)

func main() {
	stream.Init([]int{1, 2, 3, 4}).Foreach(func(e stream.Elem) {
		fmt.Println(fmt.Sprintf("this is %d", e.(int)))
	})
}
