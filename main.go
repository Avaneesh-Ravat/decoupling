package main

import (
	"fmt"
	"github.com/avaneesh-ravat/decoupling/accessor"
	"github.com/avaneesh-ravat/decoupling/model"
	"github.com/avaneesh-ravat/decoupling/personservice"
)

func main() {
	p1 := model.Person{
		First: "John",
	}

	p2 := model.Person{
		First: "James",
	}

	dbm := accessor.Mongo{}
	dbg := accessor.Postg{}

	ps1 := personservice.PersonService{
		A: dbm,
	}

	ps2 := personservice.PersonService{
		A: dbg,
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
