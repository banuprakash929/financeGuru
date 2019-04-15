package usermanagement

import (
	"errors"

	"github.com/banuprakash929/financeGuru/internal/database"
	"github.com/banuprakash929/financeGuru/internal/platform/model"
	"github.com/banuprakash929/financeGuru/internal/platform/properties"
)

//AddUser to load user to db
func AddUser(u model.User) error {

	if u.ID == "" {
		return errors.New("ENTER ID")
	}
	if u.ID != "" {
		isUserExists := database.CheckifUserExist(u.ID)
		if isUserExists {
			return errors.New(properties.ErrorUserAlreadyExist)
		}
	}
	if u.FirstName == "" {
		return errors.New("Please enter FirstName")
	}
	if u.LastName == "" {
		return errors.New("Please enter LastName")
	}
	if u.Email == "" {
		return errors.New("Please enter Email")
	}
	if u.Contact == "" {
		return errors.New("Please enter Contact Number")
	}

	err := database.SaveUser(u)
	if err != nil {
		return err
	}
	return nil
}

//GetUser to fetch user details
func GetUser(u string) (model.User, error) {
	user, err := database.GetUser(u)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

//DeleteUser to remove user form db
func DeleteUser(u model.User) error {
	err := database.DeleteUser(u)
	if err != nil {
		return err
	}
	return nil
}
