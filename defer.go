package main

import "fmt"

func getSomeVars() string {
	fmt.Println("getSomeVars execution")
	return "getSomeVars result"
}

func deferTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happened FIRST:", err)
		}
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happened SECOND:", err)
			panic(2)
		}
	}()
	fmt.Println("Some useful work")
	panic("something bad happened")
	fmt.Println("WTF?")
}

func main() {
	deferTest()
	defer fmt.Println("After work")
	//defer fmt.Println(getSomeVars())
	defer func() {
		fmt.Println(getSomeVars())
	}()
	fmt.Println("Some useful work")
}
