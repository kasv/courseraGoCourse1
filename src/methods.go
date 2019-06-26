package src

import "fmt"

type Person struct {
	Id   int
	Name string
}

func (p Person) UpdateName(name string) {
	p.Name = name
}

func (p *Person) SetName(name string) {
	p.Name = name
}

type Account struct {
	Id   int
	Name string
	Person
}

func (p *Account) SetName(name string) {
	p.Name = name
}

type MySlice []int

func (sl *MySlice) Add(val int) {
	*sl = append(*sl, val)
}

func (sl *MySlice) Count() int {
	return len(*sl)
}

func main() {
	//pers := &Person{1, "Vasily"}
	pers := new(Person)
	pers.SetName("Serj K")
	//(&pers).SetName("Vasiliy Romanov")
	fmt.Printf("updated person: %#v\n", pers)

	var acc = Account{
		Id:   1,
		Name: "rvasily",
		Person: Person{
			Id:   2,
			Name: "Vasily Romanov",
		},
	}
	acc.SetName("romanov.vasily")
	acc.Person.SetName("Test")
	fmt.Printf("updated person: %#v\n", acc)

	sl := MySlice([]int{1, 2})
	sl.Add(3)
	fmt.Println(sl.Count(), sl)
}
