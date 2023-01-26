package routes

import (
	"app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api")
	backService := services.ServiceGroupApp.BackServiceGroup.UserService
	{
		v1.POST("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "接口连接成功",
			})
		})
		v1.POST("/login", backService.Login)
		v1.POST("/register", backService.Register)
		v1.POST("/userinfo", backService.QueryUserInfo)
		users := v1.Group("/users")
		{
			users.GET("/list", backService.GetUserList)
		}
	}
	return r
}
