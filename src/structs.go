package src

import "fmt"

type Person struct {
	Id      int
	Name    string
	Address string
}

type Account struct {
	Id      int
	Name    string
	Cleaner func(string) string
	Owner   Person
	Person
}

func main() {
	var acc = Account{
		Id:   1,
		Name: "rvasily",
		Person: Person{
			Name:    "Serj",
			Address: "Sochi",
		},
	}
	fmt.Printf("%#v\n", acc)

	acc.Owner = Person{2, "Roman Vasily", "Moscow"}
	fmt.Printf("%#v\n", acc)

	fmt.Println(acc.Address)
	fmt.Println(acc.Name)
	fmt.Println(acc.Person.Name)
}
