package main

import (
	"fmt"
)

func main() {
	example()

	//forangeSliceExample1()
}

func example(){
	h := hello

	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
	//s:=new([5]int)
	//s[0]=10
	//fmt.Println(s)
	//fmt.Printf("a.%T\n",s)
	//a := make([]int,5)
	//a = append(a,1,2,3,4)
	//fmt.Println(a)
	//fmt.Printf("2.%T\n",a)
	//s := make([]int, 5)
	//s = append(s, 1, 2, 3)
	//fmt.Println(s)
	//for i := 0; i < 5; i++ {
	//	defer fmt.Printf("%d ", i)
	//}
	//strs := []string{"one", "two", "three"}
	//
	//for _, s := range strs {
	//	go func(s string) {
	//		time.Sleep(1 * time.Second)
	//		fmt.Printf("%s ", s)
	//	}(s)
	//}
	//time.Sleep(3 * time.Second)
}

//forangeSliceExample1  切片for range
func forangeSliceExample1(){
	var c  =[]string{"a","b","c"}
	for i := range c {
		fmt.Println(i)
	}

	//0
	//1
	//2
}

func hello() []string {
	return nil
}