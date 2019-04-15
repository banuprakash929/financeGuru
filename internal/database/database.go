package database

import (
	"errors"
	"log"
	"sync"

	"github.com/banuprakash929/financeGuru/internal/platform/properties"

	"github.com/banuprakash929/financeGuru/internal/platform/model"
)

var db sync.Map

func init() {
	createAndLoadDB()
}

func createAndLoadDB() {
	// TODO: get database to afile

}

//SaveUser to add or update user in db
func SaveUser(u model.User) error {
	_, exists := db.LoadOrStore(u.ID, u)
	// uu = model.User(uu)

	// if uu.ID != u.ID {
	// 	return errors.New(fmt.Sprintf("Failed to save the user with id %v, but saved with id %v", u.ID, uu.ID))
	// }

	if exists {
		log.Printf("User %v got updated", u.ID)
	} else {
		log.Printf("User %v got added", u.ID)
	}
	return nil
}

//DeleteUser to delete user from db
func DeleteUser(u model.User) error {

	eu, exists := db.Load(u.ID)
	if !exists {
		log.Printf("User %v doesnot exist", u.ID)
		return errors.New(properties.ErrorUserNotFound)
	}

	uValue, ok := eu.(model.User)

	if !ok {
		return errors.New("Expecting User type")
	}

	if len(uValue.Loans) != 0 {
		for _, v := range uValue.Loans {
			if v.BalancePrincipleAmount != 0 || v.BalanceInterestAmount != 0 {
				return errors.New("User " + uValue.ID + " cant be deleted due to existing debt")
			}
		}
	}

	db.Delete(uValue.ID)
	log.Printf("User %v deleted", uValue.ID)
	return nil
}

//GetUser to fetch user details from db
func GetUser(u string) (model.User, error) {
	value, exists := db.Load(u)
	if !exists {
		return model.User{}, errors.New(properties.ErrorUserNotFound)
	}

	uValue, ok := value.(model.User)
	if !ok {
		return model.User{}, errors.New("Failed to retrive a value with typer user with" + u)
	}
	return uValue, nil

}

//CheckifUserExist to validate userID already exists
func CheckifUserExist(u string) bool {
	_, exists := db.Load(u)
	return exists

}
