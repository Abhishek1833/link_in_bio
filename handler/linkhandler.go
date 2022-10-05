package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/Abhishek1833/link_in_bio/Models"
	"github.com/Abhishek1833/link_in_bio/controller"
)

func AddLink(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("linkname")
	link := r.URL.Query().Get("url")
	username := r.URL.Query().Get("username")
	newlink := models.Links{}
	newlink.LinkName = name
	newlink.LinkUrl = link
	newlink.UserName = username
	if !models.Isexist(username) {
		json.NewEncoder(w).Encode("user name does not exist,please register or provide a valid user name")
		return
	}
	err := controller.AddLink(&newlink)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
	json.NewEncoder(w).Encode(newlink)
}

func DeleteLink(w http.ResponseWriter, r *http.Request) {
	id := 0
	_, err := fmt.Sscan(r.URL.Query().Get("id"), &id)
	if err != nil {
		json.NewEncoder(w).Encode("please privide a valid id")
		return
	}
	err = models.DeleteLink(id)
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode("link deleted succesfully")
}

func GetLink(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	fmt.Println("kjb", username)
	//username := (models.GetNameFromLink(idstr)).UserName
	links1 := models.GetLinkFromName(username)
	fmt.Println(links1)
	json.NewEncoder(w).Encode(links1)
}
