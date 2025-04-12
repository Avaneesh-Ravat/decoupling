package accessor

import "github.com/avaneesh-ravat/decoupling/model"

type Accessor interface {
	Save(n int, p model.Person)
	Retrieve(n int) model.Person
}
