package main

import "fmt"

func main() {

	ch:=make(chan bool, 50)

	go func() {
		for i:=0;i<20;i++ {
			ch<-true
		}
		fmt.Println("송신 루틴 완료")
	}()

	for i:=0;i<20;i++ {
		<-ch
		fmt.Println("수신한 데이터:",i)
	}



}
