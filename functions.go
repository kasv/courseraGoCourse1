package main

import "fmt"

func multipleNamedReturn(ok bool) (rez int, err error) {
	rez = 1
	if ok {
		err = fmt.Errorf("some error happend")
		return
	}
	rez = 2
	return
}

func sum(in ...int) (result int) {
	fmt.Printf("in := %#v \n", in)
	for _, val := range in {
		result += val
	}
	return
}

func main() {
	//fmt.Println(multipleNamedReturn(false))
	nums := []int{1, 2, 3, 4}
	fmt.Println(nums, sum(nums...))
}
