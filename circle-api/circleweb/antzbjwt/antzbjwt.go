package antzbjwt

import (
	"circleweb/services"
	"errors"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"time"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
type User struct {
	Username string
	Uid string
}
const identityKey string="userInfo"
func Do(s *services.Service) (*jwt.GinJWTMiddleware, error) {

   return jwt.New(&jwt.GinJWTMiddleware{
    	Realm:"circle-antzb",
    	Key:[] byte("Antzb168"),
    	Timeout:time.Hour*24,
    	MaxRefresh:time.Hour,
    	IdentityKey:identityKey,
    	PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v,ok :=data.(*User);ok{
				return jwt.MapClaims{
					"uid":v.Uid,
					"userName":v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims :=jwt.ExtractClaims(c)
			return &User{
				Uid:claims["uid"].(string),
				Username:claims["userName"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (i interface{}, e error) {
			var loginVals login
			if err :=c.ShouldBind(&loginVals);err!=nil{
				return "",jwt.ErrMissingLoginValues
			}
			println(loginVals.Username)
			uname:=loginVals.Username
			pwd:=loginVals.Password
			//判断登录
			apuser,ismatch,err:=s.Login(uname,pwd)
			if err!=nil{
				println(err.Error())
				return nil,jwt.ErrFailedAuthentication
			}
			if !ismatch{
				return nil,errors.New("用户名或者密码不对！")
			}
			return &User{
				Uid:apuser.Uid,
				Username:apuser.Username,
			},nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v,ok:=data.(*User);ok&&len(v.Uid)>0{
				c.Set("uid",v.Uid)
				c.Set("uname",v.Username)

			  	return true
			  }
			  return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code,gin.H{
				"code":code,
				"message":message,
			})
		},
		TokenLookup:"header: Authorization, query: token, cookie: antzbjwt",
		TokenHeadName:"Bearer",
		TimeFunc:time.Now,
	})
}