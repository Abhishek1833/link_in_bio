package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/Abhishek1833/link_in_bio/Models"
	"github.com/Abhishek1833/link_in_bio/controller"
)

func GetUrl(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("username")
	newUser := models.User{}
	newUser.UserName = name
	for _, char := range name {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') && (char < '1' || char > '9') {
			//panic("non alphanumeric string")
			json.NewEncoder(w).Encode("username should be a alphanumeric string")
			return
		}
	}
	fmt.Println("nefor is exist")
	if models.Isexist(name) {
		json.NewEncoder(w).Encode("this name has been taken")
		return
	}
	fmt.Println("AFTER IS EXIST")

	controller.CreateURL(&newUser)
	json.NewEncoder(w).Encode(newUser)
}
