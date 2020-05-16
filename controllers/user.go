package controllers

import (
	"SampleAPI_Bigset/models"
	"encoding/json"
	"log"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

var userM models.User

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {string} models.User.UserID
// @Failure 403 {string} error
// @router / [post]
func (u *UserController) Post() {
	defer u.ServeJSON()
	json.Unmarshal(u.Ctx.Input.RequestBody, &userM)
	// fmt.Printf("%v", userM)
	// log.Println(userM)
	err := userM.Create()
	if err != nil {
		u.Data["json"] = err
		return
	}

	u.Data["json"] = map[string]string{"userid": userM.UserID}
	return
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	defer u.ServeJSON()
	users, _, err := userM.GetAll()
	if err != nil {
		u.Data["json"] = err
	}
	u.Data["json"] = users
	return
}

// @Title GetPaginate
// @Description getPaginate user by numbersf
// @Param   numbersf     query   int32 false       "numbersf id"
// @Param   numbersl    query   int32 false       "numbersl id"
// @Success 200 {object} models.User
// @router /paginate [get]
func (u *UserController) GetPaginate() {
	defer u.ServeJSON()
	var numbersf, numbersl int32
	u.Ctx.Input.Bind(&numbersf, "numbersf")
	u.Ctx.Input.Bind(&numbersl, "numbersl")
	log.Println(numbersf, numbersl, "GetPaginate")
	users, _, err := userM.GetPaginate(numbersf, numbersl)
	if err != nil {
		u.Data["json"] = err
	}
	u.Data["json"] = users
	return

}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {

	defer u.ServeJSON()
	uid := u.GetString(":uid")
	userM.UserID = uid
	if uid != "" {
		log.Println(userM.UserID, ":get")
		user, err := userM.Get()
		if err != nil {
			u.Data["json"] = err.Error()
			return
		} else {
			u.Data["json"] = user
			return
		}
	}
	u.Data["json"] = "nono"
	return

}

// @Title Update
// @Description update the user
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	defer u.ServeJSON()
	json.Unmarshal(u.Ctx.Input.RequestBody, &userM)

	err := userM.PutItem()
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = "success"
	}
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	userM.UserID = uid
	log.Println(userM.UserID, "Delete COntroll")
	userM.UserID = uid
	err := userM.Delete()
	defer u.ServeJSON()
	if err != nil {
		u.Data["json"] = err
		return
	}
	u.Data["json"] = "delete success!"
	return
}
