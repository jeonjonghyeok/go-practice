package main

import (
	"fmt"
	"time"
)

func main() {
	ch:=make(chan string)

	go sendMessage(ch)

	for i:=10 ;; i--{
		select {
		case a:=<-ch:
	fmt.Println(a,"메세지가 발송되었습니다.")
			return
		default:
			if i==0 {
				fmt.Println("메시지 발송에 실패했습니다.")
				return
			}
			fmt.Printf("%d초안에 메시지를 입력하세요.\n",i)
			time.Sleep(time.Duration(1)*time.Second)

		}
}
}
func sendMessage(ch chan string) {
	var message string
		fmt.Scanln(&message)
	ch<-message

}
