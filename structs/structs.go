package main

import (
	"fmt"

	"kayodedev.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	adminUser := user.NewAdmin("babalolakayode44@gmail.com", "kaybabs")

	adminUser.OutputUserDetail()
	appUser.OutputUserDetail()
	appUser.ClearUserName()
	appUser.OutputUserDetail()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
