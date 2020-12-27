package main
import (
	. "fmt"
	"time"
)
type buyer struct {
	mileage int
	article map[string]int
}
type article struct {
	name string
	price int
	quantity int
}
type delivery struct {
	status string
	onedelivery map[string]int
}
func newDelivery() delivery {
	d:= delivery{}
	d.onedelivery = map[string]int{}
	return d
}
func newBuyer() *buyer {
	b:=buyer{}
	b.mileage=1000000
	b.article = map[string]int{}
	return &b
}
func (a *article) remainQuantity() int{
	return a.quantity
}
func (b *buyer) remainMileage() int{
	return b.mileage
}
func buying(itemchoice int, b *buyer, a []article, numby *int, ch chan bool, temp map[string]int) {
				defer func() {
					if r:= recover(); r!=nil {
						Println(r)
					}
				}()
	amount := 0
	buy:=0
	Println("구매할 수량을 선택하세요")
	Scanln(&amount)
	if amount < 0 {
		panic("올바른 수량을 입력하세요.")
	}
	if b.mileage < a[itemchoice-1].price*amount || a[itemchoice-1].quantity<amount {
		panic("주문이 불가능 합니다.")
	} else {
	for {
		Println("1. 바로주문\n2. 장바구니 담기")
		Scanln(&buy)
		if buy == 1 {
			if *numby < 5 {
				b.mileage -= a[itemchoice-1].price*amount
				a[itemchoice-1].quantity-=amount
				temp[a[itemchoice-1].name]=amount
				ch<-true
				*numby++
				Println("상품이 주문 접수 되었습니다.")
				break
			} else {
				Println("배송 한도를 초과했습니다. 배송이 완료되면 주문하세요.")
			}

		} else if buy == 2 {
			duplicate := false

			for ba := range b.article {
				if ba == a[itemchoice-1].name {
					duplicate = true
				}
			}
			if duplicate == true {
				temp:= b.article[a[itemchoice-1].name] + amount
				if temp > a[itemchoice-1].quantity {
					Println("물품의 잔여 수량을 초과했습니다.")
					break
				} else {
				b.article[a[itemchoice-1].name]+=amount
			}

			} else{
				b.article[a[itemchoice-1].name]=amount
			}
			Println("상품이 장바구니에 추가되었습니다.")
			break
		} else {
			Println("다시 입력하세요\n")
		}
	}
}
}
func emptyBucket(b *buyer) {
		if len(b.article) == 0 {
			Println("장바구니가 비었습니다.")
		} else {
			for key, val:=range b.article{
				Println(key,val)
			}
		}

}
func buyBucket(b *buyer, a []article, numby *int, ch chan bool, temp map[string]int) {
	defer func() {
		if r:= recover(); r!=nil {
			Println(r)
		}
	}()
	if len(b.article) == 0 {
		panic("주문 가능한 목록이 없습니다.")
	}
	if *numby < 5 {
	for key,val:=range b.article {
		temp[key]=val
		for i:=0;i<5;i++ {
			if a[i].name == key {
				a[i].quantity -= val
				b.mileage -= val*a[i].price
			}
		}
	}
	b.article = map[string]int{}
	Println("주문 접수 되었습니다.")
	*numby++
	ch<-true
} else {
	Println("배송 한도를 초과했습니다. 배송이 완료되면 주문하세요.")
}

}
func requiredPoint(b *buyer, a []article) bool {
	bucketpoint:=0
	for key,val:=range b.article {
		for i:=0;i<5;i++ {
			if key==a[i].name {
				bucketpoint += val*a[i].price
			}
		}
	}
	Println("필요 마일리지: ",bucketpoint)
	Println("보유 마일리지: ",b.mileage)
	if bucketpoint > b.mileage {
		Printf("마일리지가 %d점 부족합니다.",bucketpoint-b.mileage)
		return false
	}
	return true
}
func excessAmount(b *buyer, a []article) bool {
	for key,val:=range b.article {
		for i:=0;i<5;i++ {
			if key==a[i].name {
				if a[i].quantity < val {
					Println("%s, %d개초과",key,val-a[i].quantity)
					return false
				}
			}
		}

	}
	return true
}
func deliveryStatus(deliverylist []delivery, i int, numby *int, ch chan bool,temp *map[string]int) {
	if <-ch {
		for key,val:=range *temp{
			deliverylist[i].onedelivery[key]=val
		}
		*temp = map[string]int{}
	deliverylist[i].status = "주문접수"
	time.Sleep(time.Second*10)
	deliverylist[i].status = "배송중"
	time.Sleep(time.Second*30)
	deliverylist[i].status = "배송완료"
	time.Sleep(time.Second*10)

	deliverylist[i].status = ""
	deliverylist[i].onedelivery = map[string]int{}


	*numby--
}


}
func main() {
	temp := make(map[string]int)
	deliverylist := make([]delivery,5)
	ch:= make(chan bool)
	numby:=0

	for i:=0;i<5;i++ {
		deliverylist[i] = newDelivery()
	}
	for j:=0;j<5;j++ {
		time.Sleep(time.Millisecond)
		go deliveryStatus(deliverylist,j,&numby,ch,&temp)
	}
	defer func() {
		if r:=recover(); r!=nil {
			Println(r)
		}
	}()
	b := newBuyer()
	articles:= make([]article,5)
	articles[0] = article{"텀블러", 1000, 30}
	articles[1] = article{"롱패딩", 500000, 20}
	articles[2] = article{"루미 백팩", 4000000, 20}
	articles[3] = article{"나이키 운동화", 150000, 50}
	articles[4] = article{"빼빼로", 1200, 500}
	for {
		Println("1. 구매")
		Println("2. 잔여 수량 확인")
		Println("3. 잔여 마일리지 확인")
		Println("4. 배송 상태 확인")
		Println("5. 장바구니 확인")
		Println("6. 프로그램 종료")

	var num int
	Scanln(&num)
	switch (num) {
	case 1:
		Printf("************구매***********\n\n")
		for {
			for i:=0;i<5;i++ {
				Println(i+1,articles[i].name,", 가격:",articles[i].price)
			}
			var itemchoice int
			Println("구매하실 제품의 번호를 입력하세요")
			Scanln(&itemchoice)
			if itemchoice < 1 || itemchoice > 5 {
				Println("없는 제품입니다.")
			} else{
				buying(itemchoice, b, articles,&numby,ch,temp)
				break
			}
		}
	case 2:
		Printf("************잔여 수량 확인************\n\n")
		for i:=0;i<5;i++ {
			Printf("%s, 잔여 수량: %d\n\n",articles[i].name,articles[i].quantity)

		}
		Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
		Scanln()
	case 3:
		Printf("*********잔여 마일리지 확인*********\n\n")
		Println("현재 잔여 마일리지는",b.mileage, "점입니다.")
		Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
		Scanln()
	case 4:
		Printf("***********배송 상태 확인***********\n\n")

		total:=0
		for i:=0;i<5;i++ {
			total+=len(deliverylist[i].onedelivery)
		}
		if total == 0 {
			Println("현재 배송중인 상품이 없습니다.")
		}
		for j:=0;j<5;j++ {
			if len(deliverylist[j].onedelivery) > 0 {
			for key,val:=range deliverylist[j].onedelivery {
				Println(key,val)
			}
			Println("배송상태:",deliverylist[j].status)
		}
		}

		Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
		Scanln()
	case 5:
		Printf("***********장바구니 확인***********\n\n")
		var bucket int
		for {
			emptyBucket(b)
			Println("1. 장바구니 물품 주문\n2. 장바구니 초기화\n3. 매뉴 화면으로 돌아가기")
			Scanln(&bucket)
			if bucket == 1 {
				requiredbuy := requiredPoint(b,articles)
				excessbuy := excessAmount(b,articles)
				if requiredbuy && excessbuy {
					buyBucket(b,articles,&numby,ch,temp)
					break
				} else {
					Println("구매할 수 없습니다.")
					break
				}
			} else if bucket ==2 {
				b.article = map[string]int{}
				Println("장바구니를 초기화했습니다.")
				break
			} else if bucket ==3 {
				break
			} else {
				Println("잘못된 입력입니다.")
			}
		}
		Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
		Scanln()
	case 6:
		panic("프로그램 종료")
		Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
		Scanln()
	default:
		Println("잘못 입력하였습니다.")
		Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
		Scanln()
	}

	}
}
