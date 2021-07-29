package models

import (
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	validation.SetDefaultMessage(map[string]string{
		"Required":  "不能为空",
		"MinSize":   "长度最小值是 %d",
		"MaxSize":   "长度最大值是 %d",
		"Length":    "长度需要为 %d",
		"Email":     "格式不正确",
		"AlphaDash": "只能包含字符或数字或横杠 -_",
	})

	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4,utf8&parseTime=True&loc=Local",
			beego.AppConfig.String("db_user"),
			beego.AppConfig.String("db_password"),
			beego.AppConfig.String("db_addr"),
			beego.AppConfig.String("db_name"),
		))

	if err != nil {
		log.Fatalln(err)
	}
	DB = db

	DB.AutoMigrate(&User{}, &Page{}, &Question{}, &EmailValidation{})
}

type UserRegisterForm struct {
	Recaptcha      string `form:"g-recaptcha-response" valid:"Required" label:"Recaptcha"`
	Name           string `form:"name" valid:"Required; MaxSize(20)" label:"昵称"`
	Password       string `form:"password" valid:"Required; MinSize(8); MaxSize(30)" label:"密码"`
	RepeatPassword string `form:"repeat_password"`
	Email          string `form:"email" valid:"Required; Email; MaxSize(100)" label:"电子邮箱"`
	Domain         string `form:"domain" valid:"Required; AlphaDash; MinSize(3); MaxSize(10)" label:"个性域名"`
}

func (f *UserRegisterForm) Valid(v *validation.Validation) {
	if f.Password != f.RepeatPassword {
		_ = v.SetError("Password", "两次输入的密码不相同")
	}
}

type UserLoginForm struct {
	Recaptcha string `form:"g-recaptcha-response" valid:"Required" label:"Recaptcha"`
	Email     string `form:"email" valid:"Required; Email; MaxSize(100)" label:"电子邮箱"`
	Password  string `form:"password" valid:"Required; MinSize(8); MaxSize(30)" label:"密码"`
}

type EmailValidationForm struct {
	Email string `form:"email" valid:"Required; Email; MaxSize(100)" label:"电子邮箱"`
}

type PasswordRecoveryForm struct {
	Password       string `form:"password" valid:"Required; MinSize(8); MaxSize(30)" label:"密码"`
	RepeatPassword string `form:"repeat_password"`
}

func (f *PasswordRecoveryForm) Valid(v *validation.Validation) {
	if f.Password != f.RepeatPassword {
		_ = v.SetError("Password", "两次输入的密码不相同")
	}
}

type QuestionForm struct {
	Recaptcha string `form:"g-recaptcha-response" valid:"Required" label:"Recaptcha"`
	PageID    uint
	Content   string `form:"content" valid:"Required; MaxSize(300)" label:"问题内容"`
}

type UpdateForm struct {
	Name     string `form:"name" valid:"Required; MaxSize(20)" label:"昵称"`
	Password string `form:"password" label:"密码"`
	Intro    string `form:"intro" valid:"MaxSize(40)" label:"留言板介绍"`
}

type AnswerForm struct {
	Answer string `form:"answer" valid:"Required; MaxSize(300)" label:"回答内容"`
}

type UploadCallBack struct {
	Success 		bool 		`json:"success"`
	Code 				string 	`json:"code"`
	Message  		string 	`json:"message"`
	Data struct {
		FileID		int 		`json:"file_id"`
		Width			int 		`json:"width"`
		Height		int 		`json:"height"`
		FileName	string 	`json:"filename"`
		StoreName	string 	`json:"storename"`
		Size			int			`json:"size"`
		Path			string 	`json:"path"`
		Hash     	string 	`json:"hash"`
		Url				string 	`json:"url"`
		DeleteUrl	string 	`json:"delete"`
		PageUrl		string	`json:"page"`
	} `json:"data"`
	RequestId 	string  `json:"RequestId"`
}

type RecaptchaResponse struct {
	Success bool `json:"success"`
}

type User struct {
	gorm.Model
	Name     string
	Password string
	Email    string
	Avatar   string
	PageID   uint
}

type Page struct {
	gorm.Model
	Domain     string
	Background string
	Intro      string
}

type Question struct {
	gorm.Model
	PageID  uint
	Content string
	Answer  string
}

// EmailValidation used to save the email validation data.
type EmailValidation struct {
	gorm.Model
	UserID uint
	Email  string
	Code   string
	Type   string
}
