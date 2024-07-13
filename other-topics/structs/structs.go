package main

import (
	"fmt"
	"time"
)

// Defining struct here. The `user` is in lower case becuase it is not being exported(in other words...it will not be imported anywhere else), it will only be used in this file structs.go
type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	userFirstName := inputUserData("Enter first name here: ")
	userLastName := inputUserData("Enter last name here: ")
	userBirthDate := inputUserData("Enter birthdate here (dd/mm/yyyy): ")
	// fmt.Printf("%v %v %v", userFirstName, userLastName, userBirthDate)

	// declaring the appUser of type `user`(whihc is a struct)
	var appUser user

	// creating an instance of struct `user`
	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthDate,
		createdAt: time.Now(),
	}

	// calling function which accepts appUser as an argument. appUser is an instance of the struct `user`
	outputUserData(appUser)
}

func outputUserData(u user) {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func inputUserData(text string) string {
	fmt.Print(text)
	var value string
	fmt.Scan(&value)
	return value
}
