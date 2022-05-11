package form

type UserListForm struct {
	Page int `form:"page" json:"page" binding:"required"`
	//用户名
	Pagesize int `form:"pagesize" json:"pagesize" binding:"required"`
}