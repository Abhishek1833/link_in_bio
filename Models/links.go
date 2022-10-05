package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Links struct {
	Id       int    `orm:"pk" json:"id"; orm:"column(id)"`
	UserName string `json:"username"; orm:"column(user_name)"`
	LinkName string `json:"linkname"; orm:"column(link_name)"`
	LinkUrl  string `json:"linkurl"; orm:"column(link_url)"`
}

func (t *Links) TableName() string {
	return "links"
}

func init() {
	orm.RegisterModel(new(Links))
}

func AddLink(link *Links) error {
	o := orm.NewOrm()
	_, err := o.Insert(link)
	return err
}
func DeleteLink(id int) (err error) {
	o := orm.NewOrm()
	v := Links{Id: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Links{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetLinkFromName(name string) []Links {
	var v []Links
	o := orm.NewOrm()
	_, err := o.QueryTable(new(Links)).Filter("user_name", name).RelatedSel().All(&v)
	if err != nil {
		panic(err)
	}
	return v
}
