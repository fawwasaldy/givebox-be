package route

import (
	"github.com/gin-gonic/gin"
	"givebox/application/service"
	"givebox/presentation/controller"
	"givebox/presentation/middleware"
)

func ChatRoute(route *gin.Engine, chatController controller.ChatController, jwtService service.JWTService) {
	chatGroup := route.Group("/api/chat")
	{
		chatGroup.POST("/", middleware.Authenticate(jwtService), chatController.ChatToDonor)
		chatGroup.POST("/send", middleware.Authenticate(jwtService), chatController.SendMessage)
		chatGroup.GET("/conversation", middleware.Authenticate(jwtService), chatController.GetAllConversationsByUserIDWithPagination)
		chatGroup.GET("/conversation/:conversation_id/message", middleware.Authenticate(jwtService), chatController.GetAllMessagesByConversationIDWithPagination)
	}
}
