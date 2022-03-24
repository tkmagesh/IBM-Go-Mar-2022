package main

import "fmt"

type MyStr string //type aliasing

func (s MyStr) Length() int {
	return len(s)
}

func main() {
	var str MyStr = MyStr("This is a sample string")
	fmt.Println(str.Length())
}
