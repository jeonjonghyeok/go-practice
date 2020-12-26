package main

import (
	"fmt"
"time"
)

func main() {

	ch:= make(chan bool)

	go func() {
		for i:=0;i<20;i++ {
			ch<-true
		}
		fmt.Println("송신 루틴 완료")
	}()

	go func() {
		for i:=0;i<20;i++ {
			<-ch
			fmt.Println("수신한 데이터:",i)
		}
	}()





	time.Sleep(time.Second*3)

}
