package controller

import (
	"poetry/app/util"
	"poetry/global"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JwtController struct{}

func (JwtController) PaserToken(c *gin.Context){
	//解析出数据
	claim,_:=c.Get("claims")
	// claims, _ := ctx.Get("claims")
		// 获取现在用户信息
	currentUser := claim.(*util.CustomClaims)
	// &CustomClaims{}
	id,_:=c.Get("userId")
	ID,_:=id.(int);
	// fmt.Printf("%T",ID)
	util.Success(c,map[string]string{
		"id":strconv.Itoa(ID),
		"phone":currentUser.Phone,
	})
	return
}

func (JwtController) GetJwt(c *gin.Context) {
	jwtObj := util.NewJWT()
	claims := util.CustomClaims{
		ID:       18,
		Phone:    "16605501720",
		AuthorityId: uint(1),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			// TODO 设置token过期时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, //token -->30天过期
			Issuer:    "test",
		},
	}
	//生成token
	token, err := jwtObj.CreateJwt(claims)
	if err != nil {
		util.Error(c,  "token生成失败,重新再试")
		return 
	}
	util.Success(c,token)
	return 
}



func (JwtController) Blacklist(c *gin.Context)  {

	token:=c.Request.Header.Get("x-token")
	//将token加入黑名单


	global.Redis.Set(c,token,util.RandString(10),3000)


	util.Success(c,"加入黑名单成功");
	return
}

func (JwtController) RefreshToken(c *gin.Context){
	token:=c.Request.Header.Get("x-token")
	global.Redis.Set(c,token,util.RandString(10),3000)
	newToken,_:=util.NewJWT().RefreshToken(token)

	util.Success(c,map[string]string{
		"token":newToken,
	})
	return 
}


