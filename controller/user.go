package controller

import (
	models "github.com/Abhishek1833/link_in_bio/Models"
)

func CreateURL(User *models.User) {
	User.UserUrl = "http://localhost:3102/" + "fetch" + "?username=" + User.UserName
	err := models.AddUser(User)
	if err != nil {
		panic(err)
	}

}
func AddLink(link *models.Links) error {
	err := models.AddLink(link)
	return err
}
