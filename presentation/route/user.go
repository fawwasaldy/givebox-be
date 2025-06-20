package route

import (
	"github.com/gin-gonic/gin"
	"givebox/application/service"
	"givebox/presentation/controller"
	"givebox/presentation/middleware"
)

func UserRoute(route *gin.Engine, userController controller.UserController, jwtService service.JWTService) {
	userGroup := route.Group("/api/user")
	{
		userGroup.POST("/register", userController.Register)
		userGroup.POST("/login", userController.Login)
		userGroup.GET("/me", middleware.Authenticate(jwtService), userController.Me)
		userGroup.POST("/change-password", middleware.Authenticate(jwtService), userController.ChangePassword)
		userGroup.POST("/refresh-token", userController.RefreshToken)
		userGroup.POST("/logout", middleware.Authenticate(jwtService), userController.Logout)
		userGroup.PATCH("/", middleware.Authenticate(jwtService), userController.Update)
		userGroup.DELETE("/", middleware.Authenticate(jwtService), userController.Delete)
	}
}
