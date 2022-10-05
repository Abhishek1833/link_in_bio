package handler

/*func Login(w http.ResponseWriter, r *http.Request) {

}

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("user_name")
	alphaNumeric := true
	for _, char := range username {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			alphaNumeric = false
			break
		}
	}

	if !alphaNumeric || !(5 <= len(username) && len(username) <= 50) {
		json.NewEncoder(w).Encode("please choose a valid user name")
		return
	}
	password := r.FormValue("password")
	if len(password) < 5 || len(password) > 50 {
		json.NewEncoder(w).Encode("please keep password lenght b/w 5 to 50")
	}
	raw := models.Isexist(username)
	//var uID string

	if raw.Err() == sql.ErrNoRows {
		json.NewEncoder(w).Encode("username already exist")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	fmt.Println("hash: ", hash)
	fmt.Println("hash string: ", string(hash))
	db := models.Getdb()
	insertstmt, err := db.Prepare("INSERT INTO user_detail(username,hash) VALUES(?,?);")
	if err != nil {
		json.NewEncoder(w).Encode("error in prepare statement")
		return
	}
	_, err = insertstmt.Exec(username, hash)
	defer insertstmt.Close()
	if err != nil {
		json.NewEncoder(w).Encode("error in inserting new user")
		return
	}
	json.NewEncoder(w).Encode("congrats, user registered succesfully")
}*/
