package main

import (
	"fmt"
	"log"
	"time"

	"github.com/banuprakash929/financeGuru/internal/usermanagement"

	"github.com/banuprakash929/financeGuru/internal/platform/model"
)

func main() {
	fmt.Println("Stating service...")

	user := model.User{}
	user.ID = "TejaK"
	user.FirstName = "Teja"
	user.LastName = "Kalagara"
	user.Age = 29
	user.Contact = "2345677777"
	user.Email = "tk@gmail.com"
	user.CreatedTimestamp = time.Now()

	for {
		err := usermanagement.AddUser(user)
		if err != nil {
			log.Println(err)
		} else {
			value, err := usermanagement.GetUser(user.ID)
			if err != nil {
				log.Println(err)
			} else {
				log.Println(value)
			}

		}
		time.Sleep(30 * time.Second)
	}

}
