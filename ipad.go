package main

import "fmt"


func inputNums() []int {

	var score int
	nums := make([]int,5)

	for i := 0; i<5; i++{
		fmt.Scanf("%d",&score)
		nums[i] = score
	}

	return nums
}


func calExam(arr []int) (int, int, int) {

	var sum,num90,num70 int
	num90 =0
	num70 = 0
	for i:= 0; i<5; i++ {
		if arr[i] >= 90 {
			num90++
		}
		if arr[i] < 70 {
			num70++
		}

		sum += arr[i]
	}

	return sum, num90, num70
}



func printResult (sum int,num90 int,num70 int ) {
	var result bool = true

	if sum < 400 {
		fmt.Println("총점이 400점 미만입니다.")
		result = false

	}
	if num90 < 2 {

		fmt.Println("90점 이상 과목 수가 2개 미만입니다.")
		result =false
	}
	if num70 > 0 {
		fmt.Println("70점 미만 과목이 있습니다.")
		result = false

	}

	if result == false {
		fmt.Println("아이패드를 살 수 없습니다.")
	}else{
		fmt.Println("아이패드를 살 수 있습니다.")
	}
}

func main() {
	printResult(calExam(inputNums()))

}
