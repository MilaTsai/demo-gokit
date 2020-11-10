package Services

import "errors"

type IUserService interface {
	GetName(userid int) string
	DelName(userid int) error
}

type UserService struct {
}

func (this UserService) GetName(userid int) string {

	if userid == 101 {
		return "Mila"
	}

	return "guest"
}

func (this UserService) DelName(userid int) error {

	if userid == 101 {
		return errors.New("無權限")
	}

	return nil
}
