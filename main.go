package main

import "fmt"

type Person struct {
	First string
}

type Mongo map[int]Person
type Postg map[int]Person

func (m Mongo) Save(n int, p Person) {
	m[n] = p
}

func (m Mongo) Retrieve(n int) Person {
	return m[n]
}

func (pg Postg) Save(n int, p Person) {
	pg[n] = p
}

func (pg Postg) Retrieve(n int) Person {
	return pg[n]
}

type Accessor interface {
	Save(n int, p Person)
	Retrieve(n int) Person
}

type PersonService struct {
	a Accessor
}

func (ps PersonService) get(n int) (Person, error) {
	p := ps.a.Retrieve(n)
	if p.First == "" {
		return Person{}, fmt.Errorf("no person found at key: %v in %v", n, ps)
	}
	return p, nil
}

func (ps PersonService) put(n int, p Person) error {
	ps.a.Save(n, p)
	return nil
}

func main() {
	p1 := Person{
		First: "John",
	}

	p2 := Person{
		First: "James",
	}

	dbm := Mongo{}
	dbg := Postg{}

	ps1 := PersonService{
		a: dbm,
	}

	ps2 := PersonService{
		a: dbg,
	}

	fmt.Println("Adding in mongo \n======================")
	ps1.put(1, p1)
	ps1.put(2, p2)

	fmt.Println("Getting data from mongo ==============")
	fmt.Println(ps1.get(1))
	fmt.Println(ps1.get(2))

	fmt.Println("Adding in postgress \n======================")
	ps2.put(1, p1)
	ps2.put(2, p2)

	fmt.Println("Getting data from postgress ==============")
	fmt.Println(ps2.get(1))
	fmt.Println(ps2.get(2))

	fmt.Println(ps2.get(3))

}
