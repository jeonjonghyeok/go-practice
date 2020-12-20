package main

import "fmt"

type student struct {
	name string
	sex string
	score map[string]int
}

func newStudent() student {
	stu := student{}
	stu.score = map[string]int {}
	return stu
}

func main() {
	var stuctn,subctn,score int
	var name,sex,subject string

	fmt.Scanln(&stuctn,&subctn)

	s := make([]student,stuctn)

	for i:=0;i<stuctn;i++ {
		s[i] = newStudent()
		fmt.Scanln(&name,&sex)
		s[i].name=name
		s[i].sex=sex


		for j:=0;j<subctn;j++ {
			fmt.Scanln(&subject,&score)

			s[i].score[subject]=score
		}

	}

	for i:=0;i<stuctn;i++ {
		fmt.Println("----------")
		fmt.Println(s[i].name,s[i].sex)

		for index, val := range s[i].score {
			fmt.Println(index, val)
		}

	}
	fmt.Println("----------")
}
