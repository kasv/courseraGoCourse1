package main

import (
	"fmt"
	"visibility/person"
)

func main() {
	p := person.NewPerson(1, "rvasily", "secret")
	secret := person.GetSecret(p)
	fmt.Println("GetSecret", secret)
	p.UpdateSecret("supersecret")
	fmt.Println("GetSecret", person.GetSecret(p))
}
