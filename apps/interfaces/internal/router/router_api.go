package router

import (
	"GIM/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	//engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
	publicGroup := engine.Group("open")
	registerOpenRoutes(publicGroup)

	privateGroup := engine.Group("api")
	registerPrivateRouter(privateGroup)
}

// 无需验证
func registerOpenRoutes(group *gin.RouterGroup) {
	registerOpenAuthRouter(group)
}

// 需要验证
func registerPrivateRouter(group *gin.RouterGroup) {
	group.Use(middleware.JwtAuth())
	registerPublicRouter(group)
	registerPrivateAuthRouter(group)
	registerUserRouter(group)
	registerChatMessageRouter(group)
	registerChatMemberRouter(group)
	registerChatInviteRouter(group)
	registerChatRouter(group)
	registerConvoRouter(group)
	registerLbsRouter(group)
	registerRedEnvRouter(group)
	registerRedEnvReceiveRouter(group)
	registerOrderRouter(group)
}

// 开放式api无需验证会话id
func registerPublicRouter(group *gin.RouterGroup) {
	router := group.Group("public")
	registerPublicTestRouter(router)
}

func registerPublicTestRouter(group *gin.RouterGroup) {

}
