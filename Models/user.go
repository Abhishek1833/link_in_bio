package models

import "github.com/astaxie/beego/orm"

type User struct {
	//UserId   int    `json:"userid"; orm:"column(user_id)"`
	UserName string `orm:"pk" json:"username"; orm:"column(user_name)"`
	UserUrl  string `json:"userurl"; orm:"column(user_url)"`
}

func (t *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}

func Isexist(name string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	var user1 User
	err := qs.Filter("user_name", name).RelatedSel().One(&user1)
	return err == nil
}
func AddUser(newuser *User) error {
	o := orm.NewOrm()
	_, err := o.Insert(newuser)
	return err
}

func GetNameFromLink(str string) User {
	o := orm.NewOrm()
	var v User
	err := o.QueryTable(new(User)).Filter("brand_url", str).RelatedSel().One(&v)
	if err != nil {
		panic(err.Error())
	}
	return v
}
