package main

import (
	"app/models"
	"controllers"
	"fmt"
)

func main() {
	// fmt.Println(models.Db)
	// u := &models.User{}
	// u.Name = "gbc"
	// u.Email = "hoge@gbc.co.jp"
	// u.Password = "gbcpassword"
	// fmt.Println(u)
	// u.CreateUser()

	user, _ := models.GetUser(1)
	fmt.Println(user)

	user.Name = "gbc3"
	user.Email = "gbc3@gbc.co.jp"
	user.UpdateUser()
	user, _ = models.GetUser(user.ID)
	fmt.Println(user)
	user.DeleteUser()
	user, _ = models.GetUser(user.ID)
	fmt.Println(user)

	controllers.StartMainServer()
}
