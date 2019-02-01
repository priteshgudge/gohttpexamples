package userrepo

import (
	"github.com/priteshgudge/gohttpexamples/sample4/domain"
	"github.com/priteshgudge/gohttpexamples/sample4/utils"
)

type Factory struct {
}

func (f *Factory) NewUser(firstname string, lastname string, age int) *domain.User {
	return &domain.User{
		Firstname: firstname,
		Lastname:  lastname,
		Age:       age,
		CreatedOn: utils.GetUTCTimeNow(),
	}

}
