package accessor

import "github.com/avaneesh-ravat/decoupling/model"

type Mongo map[int]model.Person

func (m Mongo) Save(n int, p model.Person) {
	m[n] = p
}

func (m Mongo) Retrieve(n int) model.Person {
	return m[n]
}
