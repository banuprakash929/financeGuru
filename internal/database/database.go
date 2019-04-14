package database

import (
	"sync"
)

//"github.com/banuprakash929/financeGuru/internal/model"
var database sync.Map

func init() {
	createAndLoadDB()
}

func createAndLoadDB() {
	// TODO: get database to afile

}

// func SaveUser(u User) {

// }
