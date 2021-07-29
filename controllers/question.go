package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/miriko-channel/NekoBox/models"
	"strconv"
)

type QuestionController struct {
	beego.Controller
}

// Question is the page of a question.
func (this *QuestionController) Question() {
	domain := this.Ctx.Input.Param(":domain")
	id := this.Ctx.Input.Param(":id")

	questionID, err := strconv.Atoi(id)
	if err != nil {
		this.Redirect("/", 302)
		return
	}
	question, err := models.GetQuestionByDomainID(domain, uint(questionID))
	if err != nil {
		this.Redirect("/", 302)
		return
	}

	// public user can't get the no answer question.
	isLogin := this.Ctx.Input.GetData("isLogin").(bool)
	if question.Answer == "" {
		if !isLogin || this.Ctx.Input.GetData("user").(*models.User).PageID != question.PageID {
			this.Redirect("/", 302)
			return
		} else {
			if this.Ctx.Input.Query("err") != "" {
				this.Data["error"] = "回答问题失败！"
			}
		}
	}

	user, _ := models.GetUserByPage(question.PageID)
	page, _ := models.GetPageByDomain(domain)
	questions := models.GetQuestionsByPageID(question.PageID, false)
	this.Data["userContent"] = user
	this.Data["pageContent"] = page
	this.Data["questionsContent"] = questions
	this.Data["questionContent"] = question
	this.Data["title"] = fmt.Sprintf("%s的提问箱 | %s", user.Name, beego.AppConfig.String("title"))
	this.TplName = "question.tpl"
}

// QuestionList show the owner's all questions.
func (this *QuestionController) QuestionList() {
	isLogin := this.Ctx.Input.GetData("isLogin").(bool)
	if !isLogin {
		this.Redirect("/login", 302)
		return
	}
	user := this.Ctx.Input.GetData("user").(*models.User)
	questions := models.GetQuestionsByPageID(user.PageID, true)
	this.Data["questionContent"] = questions
	this.TplName = "questionlist.tpl"
}

// AnswerQuestion is the answer question handler.
func (this *QuestionController) AnswerQuestion() {
	this.TplName = "questionlist.tpl"
	isLogin := this.Ctx.Input.GetData("isLogin").(bool)
	if !isLogin {
		this.Redirect("/login", 302)
		return
	}

	domain := this.Ctx.Input.Param(":domain")
	id := this.Ctx.Input.Param(":id")
	questionID, err := strconv.Atoi(id)
	if err != nil {
		this.Redirect("/", 302)
		return
	}

	question, err := models.GetQuestionByDomainID(domain, uint(questionID))
	if err != nil || question.Answer != "" {
		this.Redirect("/", 302)
		return
	}

	// make sure the question belong to this user
	loginUser := this.Ctx.Input.GetData("user").(*models.User)
	if loginUser.PageID != question.PageID {
		this.Redirect("/", 302)
		return
	}

	questionURL := "/_/" + domain + "/" + id

	// parse form
	a := new(models.AnswerForm)
	if err := this.ParseForm(a); err != nil {
		this.Redirect(questionURL+"?err=1", 302)
		return
	}

	valid := validation.Validation{}
	b, err := valid.Valid(a)
	if err != nil {
		this.Redirect(questionURL+"?err=1", 302)
		return
	}

	if !b {
		this.Redirect(questionURL+"?err=1", 302)
		return
	}

	question = &models.Question{
		Answer: a.Answer,
	}

	err = models.AnswerQuestion(uint(questionID), question)
	if err != nil {
		this.Redirect(questionURL+"?err=1", 302)
		return
	}
	this.Redirect(questionURL, 302)
}

func (this *QuestionController) QuestionDelete() {
	this.TplName = "questionlist.tpl"
	isLogin := this.Ctx.Input.GetData("isLogin").(bool)
	if !isLogin {
		this.Redirect("/login", 302)
		return
	}
	user := this.Ctx.Input.GetData("user").(*models.User)

	domain := this.Ctx.Input.Param(":domain")
	id := this.Ctx.Input.Param(":id")
	questionID, err := strconv.Atoi(id)
	if err != nil {
		this.Redirect("/", 302)
		return
	}

	question, err := models.GetQuestionByDomainID(domain, uint(questionID))
	if err != nil {
		this.Redirect("/", 302)
		return
	}

	if question.PageID != user.PageID {
		this.Redirect("/", 302)
		return
	}

	models.DeleteQuestion(question.ID)
	this.Redirect("/_/"+domain, 302)
}
