package main

import "fmt"

const g float32 = 9.8

type object struct {
	m float32
	v float32
	h float32
	ke float32
	pe float32
	me float32
}

func (o object)keMethod() (ke float32){
	ke = o.m*o.v*o.v*0.5
	return ke
}

func (o object) peMethod() (pe float32){
	pe = o.m*o.h*g
	return pe
}

func main() {
	var cnt int
	fmt.Scanln(&cnt)

	o:=make([]object,cnt)

	for i:=0;i<cnt;i++ {
		fmt.Scanln(&o[i].m,&o[i].v,&o[i].h)
		o[i].ke=o[i].keMethod()
		o[i].pe=o[i].peMethod()
		o[i].me=o[i].ke+o[i].pe
	}


	for i:=0;i<cnt;i++ {
		fmt.Println(o[i].ke,o[i].pe,o[i].me)
	}

}


