package main

import "fmt"

func bubbleSort([]int nums) {
	var temp int
	temp = nums[0]
	for i:=0;i<len(nums);i++ {
		for j:=1;j<len(nums)-1;i++ {
			if nums[j] < temp {

			}
		}
	}

}

func inputNums() []int {
	nums := make([]int, 100, 100)


	var n,num int

	fmt.Scanln(&n)
	for i:=0;i<n;i++ {
		fmt.Scanln(&num)
		nums= append(nums,num)
	}
	return nums
}

func outputNums() {

}

func main() {
	nums := inputNums()
	bubbleSort()
	outputNums()
}
