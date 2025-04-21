package routes

import (
	"go-gin-chat/controller"
	"go-gin-chat/services/friends"
	"go-gin-chat/services/session"
	"go-gin-chat/static"
	"go-gin-chat/ws/primary"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRoute() *gin.Engine {
	//router := gin.Default()
	router := gin.New()

	if viper.GetString(`app.debug_mod`) == "false" {
		// live 模式 打包用
		router.StaticFS("/static", http.FS(static.EmbedStatic))
	} else {
		// dev 开发用 避免修改静态资源需要重启服务
		router.StaticFS("/static", http.Dir("static"))
	}

	sr := router.Group("/", session.EnableCookieSession())
	{
		sr.GET("/", controller.Index)

		sr.POST("/login", controller.Login)
		sr.GET("/logout", controller.Logout)
		sr.GET("/ws", primary.Start)
		sr.POST("/friend/add", friends.AddFriend)
		sr.POST("/friend/list", friends.AddFriend)
		sr.POST("/friend/recommend", friends.AddFriend)

		authorized := sr.Group("/", session.AuthSessionMiddle())
		{
			//authorized.GET("/ws", ws.Run)
			authorized.GET("/home", controller.Home)
			authorized.GET("/room/:room_id", controller.Room)
			authorized.GET("/private-chat", controller.PrivateChat)
			authorized.POST("/img-kr-upload", controller.ImgKrUpload)
			authorized.GET("/pagination", controller.Pagination)
		}

	}

	return router
}
