package studyGolang

import "fmt"
var csy string

const pai = 3.1415
func foo()(string,int) {
	return "ajax",1000
}
func Test() {
	var x int
	//自动推导数据类型
	x = 111
	fmt.Println(x)
	//一次性声明多个变量
	var(
		a int
		b string
		c bool
	)
	fmt.Println(a,b,c)

	aa,bb :=foo();

	fmt.Println(aa,bb)
//_匿名变量,用于接收不需要的值
	cc,_ :=foo();
	fmt.Println(cc)
}

func Ofe()  {
	const e  = 2.718
	fmt.Println(pai,e)
	const(
		aa = iota
		bb
		cc
		dd
	)
	fmt.Println(aa,bb,cc,dd)
	var x [10]int
	x[0] = 1
	x[9] = 10	
	xx := [10]int8 {1,4,6,2,8}
	fmt.Println(x)
	fmt.Println(xx)
	var slice []int
	slice = append(slice,1,2,3)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func Ma()  {
	student := make(map[string]int)
	student["zhansan"] = 11
	fmt.Println(student)
}

func Cir()  {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i	
	}
	fmt.Println(sum)
	x:=[5]int{1,2,3,4,5}
	//range用于迭代
	for i,v:=range x{
		fmt.Println(i,v)
	}

	a := 1
	b := 2
	a,b = swap(a,b)
	fmt.Println(a,b)
}

func swap(a  int,b int)(int,int){
	return b,a
}