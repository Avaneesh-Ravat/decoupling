package personService

import (
	// "fmt"
	"testing"

	// "github.com/avaneesh-ravat/decoupling/accessor"
	"github.com/avaneesh-ravat/decoupling/model"
	"github.com/golang/mock/gomock"
)

func TestPut(t *testing.T) {
	ctl := gomock.NewController(t)
	acc := NewMockAccessor(ctl)

	// mdb := accessor.Mongo{}
	p := model.Person{
		First: "James",
	}

	acc.EXPECT().Save(1, p).MinTimes(1).MaxTimes(1)

	ps := PersonService{
		A: acc,
	}

	ps.Put(1, p)

	ctl.Finish()
}

// func ExamplePersonService_Put() {
// 	mdb := accessor.Mongo{}
// 	p := model.Person{
// 		First: "James",
// 	}

// 	ps := PersonService{
// 		A: mdb,
// 	}

// 	ps.Put(1, p)

// 	got := mdb.Retrieve(1)

// 	fmt.Println(got)
// 	// Output: {James}
// }
