package main

import "fmt"

func main() {
	var names []string
	var name string

	for {
		fmt.Scanln(&name)
		if name == "0" {
			break
		} else {
			names = append(names,name)
		}
	}

	for i:=0;i<len(names);i++ {
		defer fmt.Println(names[i])
	}
}
