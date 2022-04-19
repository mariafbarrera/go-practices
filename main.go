package main

import (
	"Tareasconmibebe/models"
	"fmt"
)

func main(){
	user := new(models.User)
	user.Name = "user 1"
	user2 := new(models.User)
	user2.Name = "user 2"
	user3 := models.User{}
	user3.Name = "user3"
	var user4 models.User
	user4.Name = "user4"
	user5 := (models.User{})
	user5.Name = "user5"

	fmt.Println(user.GetName())
	user.SetName("Mari es el amor de mi vida, te amo tantooooooooooooooooooooo, " +
		"gracias por estar conmigo <3 lo aprecio mucho")
	fmt.Println(user.GetName())
}

