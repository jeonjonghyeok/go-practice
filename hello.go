package main

import (
	"fmt"
	"os"
)
func hello() {
	file,err := os.Open("hello.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	buf:= make([]byte, 1024)
	if _, err := file.Read(buf); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}

func main() {
	hello()
}
