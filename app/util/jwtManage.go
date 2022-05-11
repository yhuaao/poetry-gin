package util

import (
	"errors"
	"poetry/global"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)
type CustomClaims struct {
	ID          int   //
	Phone    string //
	AuthorityId uint   //
	jwt.StandardClaims
}

type JWT struct {
	JwtSec []byte
}

func NewJWT()*JWT {
	return &JWT{
		[]byte(global.Settings.JWTConfig.JwtSec),
	}
}

func (j *JWT) CreateJwt(claims CustomClaims)(string, error ){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtSec)
}


func (j *JWT) ParseToken(tokenString string)(*CustomClaims,error){
	token, err := jwt.ParseWithClaims(
		tokenString, 
		&CustomClaims{},
		func(token *jwt.Token) (i interface{}, e error) {
			return j.JwtSec, nil
		})
		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					return nil, TokenMalformed
				} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
					// Token is expired
					return nil, TokenExpired
				} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
					return nil, TokenNotValidYet
				} else {
					return nil, TokenInvalid
				}
			}
		}
		// claims, ok :=token.Claims.(*CustomClaims);
		// fmt.Print(claims, ok)
		if token != nil {
			if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
				return claims, nil
			}
			return nil, TokenInvalid

		} else {
			return nil, TokenInvalid
		}
}



func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		 func(token *jwt.Token) (interface{}, error) {
		return j.JwtSec, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
		return j.CreateJwt(*claims)
	}
	return "", TokenInvalid
}


//将token加入黑名单
