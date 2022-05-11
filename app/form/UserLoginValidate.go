package form

type UserLoginForm struct {
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
	//用户名
	Phone string `form:"phone" json:"phone" binding:"required"`

	Captcha   string `form:"captcha" json:"captcha" binding:"required"`       // 验证码
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"` // 验证码id
}
