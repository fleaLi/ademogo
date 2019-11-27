package routes

import (
	"circleweb/antzbjwt"
	"circleweb/services"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

type Route struct {
	service *services.Service
}

func Init() (r *gin.Engine) {
	log.SetFlags(log.Ldate|log.Ltime |log.Lshortfile)

	aoss := services.New()
	route := &Route{
		service: aoss,
	}
	r = gin.Default()
     corsmiddlware :=cors.DefaultConfig()
     //corsmiddlware.AllowOrigins=[]string{conf.GetString("fronthost")}
     corsmiddlware.AllowAllOrigins=true
     corsmiddlware.AllowCredentials=true
     corsmiddlware.AddAllowHeaders("x-requested-with")
     corsmiddlware.AddAllowHeaders("content-type")
	corsmiddlware.AddAllowHeaders("authorization")
    corsmiddlware.AddAllowHeaders("origin")
	r.Use(cors.New(corsmiddlware))

	route.register(r)
	return r

}
/**
注册路由
*/
func (rr *Route) register(r *gin.Engine) {
	authMiddleware,err:= antzbjwt.Do(rr.service)
	if err!=nil{
		panic(err)
	}
	r.POST("/login",authMiddleware.LoginHandler)
	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	imputes :=r.Use(authMiddleware.MiddlewareFunc())

	imputes.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 2323, "message": "ok"})
	})
	mediaGroup := r.Group("/media")
	mediaGroup.POST("/img/upload", rr.Upload)
	mediaGroup.POST("/csv/upload",rr.FileUpload)
	taskGroup :=r.Group("/task")
	taskGroup.GET("/constinfo",rr.GetConstInfo)
	taskGroup.POST("/task",rr.CreateTask)
	taskGroup.GET("/items",rr.TaskItems)
	taskGroup.PUT("/change/:taskid",rr.ChangeTaskStatus)
	taskGroup.GET("/one/:tid",rr.GetTask)
	taskGroup.PUT("/one/:tid",rr.UpdateTask)
}

func (rr *Route) check(err error,c *gin.Context)  {
	if err!=nil {
		c.JSON(500,gin.H{"code":500,"msg":err.Error()})
		return
	}
}
