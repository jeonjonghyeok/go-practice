package main

import (
	"fmt"
)

func inputSubNum() (num int) {
	fmt.Scanln(&num)

	if num > 0 {
		return num
	}

	panic(fmt.Errorf("잘못된 과목 수입니다."))
}

func average(num int) (float64) {
	var score, total int

	for i := 0; i < num; i++ {
		fmt.Scanln(&score)
		if score < 0 || score > 100 {
			panic(fmt.Errorf("잘못된 점수입니다."))
		}
		total+=score

	}

	avg := float64(total)/float64(num)

	return avg
}

func main() {
	defer func (){
		if r:= recover(); r!= nil {
			fmt.Print(r)
		}
	}()


	num := inputSubNum()



	result := average(num)

	fmt.Println(result)
}
