package main

import "fmt"

func bubbleSort(nums []int) {
	var temp int
	for i:=0;i<len(nums);i++ {
		for j:=0;j<len(nums)-i-1;j++ {
			if nums[j] > nums[j+1] {
				temp = nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}

}

func inputNums() []int{
	var n,num int
	var nums []int
	fmt.Scanln(&n)
	for i:=0;i<n;i++ {
		fmt.Scanln(&num)
		nums=append(nums,num)
	}
	return nums

}

func outputNums(nums []int) {
	for i:=0;i<len(nums);i++ {
		fmt.Printf("%d ",nums[i])
	}
}

func main() {
	nums := inputNums()
	bubbleSort(nums)
	outputNums(nums)
}
