package main

import "fmt"

func main() {
	m:=make(map[int]interface{})
	m[1] = nil
	fmt.Println(m)
}
