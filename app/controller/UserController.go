package controller

import (
	"poetry/app/dao"
	"poetry/app/form"
	"poetry/app/util"

	"github.com/gin-gonic/gin"
)



type UserController struct{

}


func (UserController) UserLogin(c *gin.Context){

	userLoginForm:=form.UserLoginForm{}
	if err := c.ShouldBind(&userLoginForm); err != nil {
		util.HandleValidatorError(c, err)
		return
	}
	//验证码
	boolCheck:=store.Verify(userLoginForm.CaptchaId,userLoginForm.Captcha,true)
	if(!boolCheck){
		util.Error(c,"验证码验证错误")
		return
	}
	util.Success(c,"登录成功")
	return
}


func (UserController) UserList(c *gin.Context){

	UserListForm := form.UserListForm{}
	if err := c.ShouldBind(&UserListForm); err != nil {
		util.HandleValidatorError(c, err)
		return
	}
	// 获取数据
	total, userlist := dao.GetUserListDao(UserListForm.Page, UserListForm.Pagesize)
	// 判断
	if (total + len(userlist)) == 0 {
		util.Error(c, "未获取到到数据")
		return
	}
	util.Success(c, map[string]interface{}{
		"total":    total,
		"userlist": userlist,
	})

}