package main

import (
	"fmt"
	// "math/rand"
	// "strings"
	// "sync"
	"time"
)

type animal struct {
	age  int
	name string
}

type pegion struct {
	age  int
	name string
}

func (e animal) print_name() {
	fmt.Print(e.name)
}

type bird interface {
	print_name()
}

// var wg = sync.WaitGroup{}
// var dbarr = []string{"db1", "db2", "db3", "db4", "db5"}

func main() {
	// var num int
	// num=1
	// fmt.Print(1,num,"\n")
	// var str string

	// str="adjsk"
	// fmt.Print(str,"\n")

	// var str1 string
	// str1="ddddd"

	// str=str1

	// fmt.Print(str,"\n")

	// var flt float
	// flt=31.0

	// fmt.Println(flt)

	// fmt.Print("hello world\n")

	// var num,div int
	// num=2
	// div=1
	// fmt.Println(divi(num,div))

	// var number=123
	// var stockcode="2025-4-2"
	// var url="Code=%d&Enddate=%s"
	// var target=fmt.Sprintf(url,number,stockcode)
	// fmt.Println(target)

	// var intarr []int32 = []int32{1, 2, 3, 4}
	// fmt.Printf("The length is%d,the Capacity is%v\n", len(intarr), cap(intarr))
	// intarr = append(intarr, 6)
	// fmt.Printf("\nThe length is%d,the Capacity is%v\n", len(intarr), cap(intarr))

	// var intslice []int32 = make([]int32, 3, 8)
	// intslice = append(intslice, 1)
	// fmt.Printf("first is %d,second is %d", intslice[0], intslice[3])

	// var stringbuffer strings.Builder
	// stringbuffer.WriteString("hello")
	// fmt.Print(stringbuffer.String())

	// var myaninmal bird = animal{18, "kitty"}
	// // fmt.Print(myaninmal.name)

	// myaninmal.print_name()

	// t0 := time.Now()
	// for i := 0; i < len(dbarr); i++ {
	// 	wg.Add(1)
	// 	go delay_print(i)
	// }
	// wg.Wait()
	// fmt.Println("total time is ", time.Since(t0))

	var c = make(chan int)
	go process(c)
	for i := range c {
		fmt.Println(i)
	}
	time.Sleep(time.Second * 10)
	// fmt.Println(<-c)
	fmt.Println("Now stop")

	go process2()
}

// func divi(num int,div int) (int,int){
// 	var err error=nil
// 	if div==0{
// 		err=errors.New("call div 0")
// 	}
// 	var res int =num/div
// 	var rem int =num%div
// 	return res,rem
// }

// func delay_print(i int) {
// 	var delay float32 = rand.Float32() * 2000
// 	time.Sleep(time.Duration(delay) * time.Millisecond)
// 	fmt.Println("The value in database is ", dbarr[i])
// 	wg.Done()
// }

func process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		var input int
		fmt.Scanln(&input)
		c <- input
	}

	fmt.Println("Exiting process")
}

func process2() {
	fmt.Println("This is a process2")
}
