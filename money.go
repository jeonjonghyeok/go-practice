package main

import "fmt"

/*타입문 작성은 선택입니다*/

func calCoin(price int) func(coinnum int) int  {
	return func(coinnum int) int { //클로저
		return coinnum*price
	}
}

func main() {
	var coin10, coin50, coin100, coin500 int
	fmt.Scan(&coin10, &coin50, &coin100, &coin500)


	if coin10 < 0 || coin50 <0 || coin100 <0 || coin500 < 0 {
		fmt.Println("잘못된입력입니다.")
		return
	}

	add10 := calCoin(10)
	add50 := calCoin(50)
	add100 := calCoin(100)
	add500 := calCoin(500)

	totalmoney := add10(coin10)+add50(coin50)+add100(coin100)+add500(coin500)
	fmt.Println(totalmoney)
}
