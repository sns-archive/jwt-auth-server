package main

import "fmt"

type Hello struct {
	word string
}

func main() {
	// NOTE: 省略形
	// helloStruct := Hello{"hello"}
	helloStruct := Hello{
		word: "うんち💩",
	}
	// NOTE: +vでvalueだけでなく、keyも表示できる
	fmt.Printf("%+v\n", helloStruct)
}
