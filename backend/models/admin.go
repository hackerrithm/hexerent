package models

import (
	"fmt"
)

type Admin struct {
	AdminName    string `json:"adminname"`
	Password     string `json:"password"`
	EmailAddress string `json:"emailaddress"`
}

func NewAdmin(adminname, password, emailaddress string) *Admin {
	admin := new(Admin)
	admin.AdminName = adminname
	admin.EmailAddress = emailaddress
	admin.Password = password

	return admin
}

func (a Admin) SayHi() {
	fmt.Printf("Hi, I am %s you can call me Mr. %s\n", a.AdminName, a.AdminName)
}

// Interface Men implemented by Human, Student and Employee
type Shots interface {
	SayHi()
}
