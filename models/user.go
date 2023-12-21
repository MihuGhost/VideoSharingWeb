package models

import (
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

type User struct {
	Id       int
	Name     string
	Password string
	Status   int
	AddTime  int64
	Mobile   string
	Avatar   string
}

type UserInfo struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	AddTime int64  `json:"addTime"`
	Avatar  string `json:"avatar"`
}

func init() {
	orm.RegisterModel(new(User))
}

func IsUserMobile(mobile string) bool {
	o := orm.NewOrm()
	user := User{Mobile: mobile}
	err := o.Read(&user, "Mobile")
	if err == orm.ErrNoRows {
		log.Printf("IsUserMobile: no row found")
		return false
	} else if err == orm.ErrMissPK {
		log.Printf("IsUserMobile: missed pk value")
		return false
	}
	return true
}

func UserSave(mobile string, password string) error {
	o := orm.NewOrm()
	user := User{
		Name:     "",
		Password: password,
		Status:   1,
		AddTime:  time.Now().Unix(),
		Mobile:   mobile,
	}
	_, err := o.Insert(&user)
	return err
}
