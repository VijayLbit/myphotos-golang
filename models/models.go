package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int    `form:"-"`
	Name     string `form:"name,text,name:" valid:"MinSize(5);MaxSize(100)"`
	Email    string `form:"email,text,email:" valid:"Email;EmailUnique"`
	Password string `form:"password,text,password:" valid:"MinSize(5);MaxSize(50)"`
}

type UserAuth struct {
	Email    string `form:"email,text,email:" valid:"Email"`
	Password string `form:"password,text,password:" valid:"Required"`
}

func (u *User) TableName() string {
	return "users"
}

func init() {
	validation.AddCustomFunc("EmailUnique", EmailUnique)
}

// EmailUnique is a custom validation implementing the validation.CustomFunc type
// It checks that a value is not in the list
func EmailUnique(v *validation.Validation, obj interface{}, key string) {
	o := orm.NewOrm()
	inputEmail := obj.(string)
	if exist := o.QueryTable("users").Filter("email", inputEmail).Exist(); exist {
		v.AddError(key, fmt.Sprintf("'%s' already registered", inputEmail))
	}
}

func (u *User) Validate() (isValid bool, valid validation.Validation) {
	isValid, _ = valid.Valid(u)
	if isValid {
		hsh, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
		if err != nil {
			valid.AddError("Password", "Technical error hashing password; please try again")
			isValid = false
		} else {
			u.Password = string(hsh)
		}
	}
	return isValid, valid
}

func (u *User) InsertNew() (msg string, err error) {
	o := orm.NewOrm()
	if id, err := o.Insert(u); err == nil {
		msg = fmt.Sprintf("User inserted with id: %d", id)
	} else {
		msg = fmt.Sprintf("Couldn't insert new user. Reason: %s", err)
	}
	return msg, err
}

func (ua *UserAuth) VerifyLoginRequest() (u User, valid validation.Validation) {
	isValid, err := valid.Valid(ua)
	if isValid {
		if u, err = getUserByEmail(ua.Email); err == nil {
			fmt.Println("verifyLoginRequest", u.Password, ua.Password)
			if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(ua.Password)); err != nil {
				u.Email = ""
				u.Id = 0
				u.Name = ""
				u.Password = ""
			}
		}
	}
	return u, valid
}

func getUserByEmail(email string) (u User, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(&u).Filter("email", email).All(&u)
	if err != nil {
		fmt.Println("getUserByEmail: ", err)
	} else {
		fmt.Println("getUserByEmail", u)
	}
	return u, err
}
