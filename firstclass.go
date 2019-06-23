package main

import "fmt"

func main() {
	prt := func(in string) {
		fmt.Println("anon func out:", in)
	}
	prt("nobody")

	type strFuncType func(string)

	worker := func(callback strFuncType) {
		callback("as callback")
	}
	worker(prt)

	prefixer := func(prefix string) strFuncType {
		return func(in string) {
			fmt.Printf("[%s] %s\n", prefix, in)
		}
	}

	successLogger := prefixer("SUCCESS")
	successLogger("expected behaviour")
}
