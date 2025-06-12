package route

import (
	"github.com/gin-gonic/gin"
	"givebox/application/service"
	"givebox/presentation/controller"
	"givebox/presentation/middleware"
)

func DonationRoute(route *gin.Engine, donationController controller.DonationController, jwtService service.JWTService) {
	donationGroup := route.Group("/api/donation")
	{
		donationGroup.GET("/donated-item", donationController.GetAllDonatedItemsWithPagination)
		donationGroup.GET("/donated-item/category/:category_id", donationController.GetAllDonatedItemsByCategoryIDWithPagination)
		donationGroup.GET("/donated-item/condition/:condition", donationController.GetAllDonatedItemsByConditionWithPagination)
		donationGroup.GET("/donated-item/status/:status", donationController.GetAllDonatedItemsByStatusWithPagination)
		donationGroup.GET("/donated-item/before-date/:date", donationController.GetAllDonatedItemsBeforeDateWithPagination)
		donationGroup.GET("/donated-item/:id", donationController.GetDonatedItemByID)
		donationGroup.POST("/donated-item/open", middleware.Authenticate(jwtService), donationController.OpenDonatedItem)
		donationGroup.POST("/donated-item/request", middleware.Authenticate(jwtService), donationController.RequestDonatedItem)
		donationGroup.POST("/donated-item/accept", middleware.Authenticate(jwtService), donationController.AcceptDonatedItem)
		donationGroup.POST("/donated-item/reject", middleware.Authenticate(jwtService), donationController.RejectDonatedItem)
		donationGroup.POST("/donated-item/taken", middleware.Authenticate(jwtService), donationController.TakenDonatedItem)
		donationGroup.PUT("/donated-item", middleware.Authenticate(jwtService), donationController.UpdateDonatedItem)
		donationGroup.DELETE("/donated-item/:id", middleware.Authenticate(jwtService), donationController.DeleteDonatedItem)
	}
}
