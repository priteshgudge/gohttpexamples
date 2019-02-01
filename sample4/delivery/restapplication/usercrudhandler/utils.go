package usercrudhandler

import "github.com/priteshgudge/gohttpexamples/sample4/domain"

func transformobjListToResponse(resp []*domain.User) UserGetListRespDTO {
	responseObj := UserGetListRespDTO{}
	for _, obj := range resp {
		userObj := UserGetRespDTO{
			ID:        obj.ID,
			FirstName: obj.Firstname,
			LastName:  obj.Lastname,
			Age:       obj.Age,
			CreatedOn: obj.CreatedOn,
		}
		responseObj.Users = append(responseObj.Users, userObj)
	}
	responseObj.Count = len(responseObj.Users)

	return responseObj
}
