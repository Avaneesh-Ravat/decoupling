package personservice

import (
	"fmt"
	"github.com/avaneesh-ravat/decoupling/accessor"
	"github.com/avaneesh-ravat/decoupling/model"
)

type PersonService struct {
	A accessor.Accessor
}

func (ps PersonService) get(n int) (model.Person, error) {
	p := ps.A.Retrieve(n)
	if p.First == "" {
		return model.Person{}, fmt.Errorf("no person found at key: %v in %v", n, ps)
	}
	return p, nil
}

func (ps PersonService) put(n int, p model.Person) error {
	ps.A.Save(n, p)
	return nil
}
