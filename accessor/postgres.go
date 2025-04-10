package accessor

import "github.com/avaneesh-ravat/decoupling/model"

type Postg map[int]model.Person

func (pg Postg) Save(n int, p model.Person) {
	pg[n] = p
}

func (pg Postg) Retrieve(n int) model.Person {
	return pg[n]
}
