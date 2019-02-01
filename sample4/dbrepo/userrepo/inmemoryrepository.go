package userrepo

import "github.com/priteshgudge/gohttpexamples/sample4/domain"

type UserInMemRepository struct {
	userObjs []*domain.User
}

func NewUserInMemRepository() *UserInMemRepository {
	return &UserInMemRepository{}
}

// type Reader interface {
// 	GetAll() ([]*domain.User, error)
// 	GetByID() (*domain.User, error)
// }

// type Writer interface {
// 	Create(*domain.User) (string, error)
// 	Update(*domain.User) error
// 	Archive(*domain.User) error
// }

func (in *UserInMemRepository) GetAll() ([]*domain.User, error) {
	var userObjResp []*domain.User
	for _, u := range in.userObjs {
		user := &domain.User{
			ID:        u.ID,
			Firstname: u.Firstname,
			Lastname:  u.Lastname,
			Age:       u.Age,
			CreatedOn: u.CreatedOn,
		}
		userObjResp = append(userObjResp, user)
	}
	return userObjResp, nil
}

func (in *UserInMemRepository) GetByID(ID string) (*domain.User, error) {
	for _, u := range in.userObjs {
		if u.ID == ID {
			user := &domain.User{
				ID:        u.ID,
				Firstname: u.Firstname,
				Lastname:  u.Lastname,
				Age:       u.Age,
				CreatedOn: u.CreatedOn,
			}
			return user, nil
		}

	}

	return nil, domain.DomainErrorNotFound
}

func (in *UserInMemRepository) Create(u *domain.User) (string, error) {
	user := &domain.User{
		ID:        u.ID,
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Age:       u.Age,
		CreatedOn: u.CreatedOn,
	}

	in.userObjs = append(in.userObjs, user)

	return user.ID, nil
}

func (in *UserInMemRepository) Update(inp *domain.User) error {

	for _, u := range in.userObjs {
		if u.ID == inp.ID {
			u.Firstname = inp.Firstname
			u.Lastname = inp.Lastname
			u.Age = inp.Age
			return nil
		}

	}

	return domain.DomainErrorNotFound
}
