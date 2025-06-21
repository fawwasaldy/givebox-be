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
		donationGroup.GET("/donated-item/city/:city", donationController.GetAllDonatedItemsByCityWithPagination)
		donationGroup.GET("/donated-item/:id", donationController.GetDonatedItemByID)
		donationGroup.POST("/donated-item/open", middleware.Authenticate(jwtService), donationController.OpenDonatedItem)
		donationGroup.POST("/donated-item/accept", middleware.Authenticate(jwtService), donationController.AcceptDonatedItem)
		donationGroup.GET("/donated-item/image/:donated_item_id", donationController.GetAllImagesByDonatedItemID)
		donationGroup.GET("/category", donationController.GetAllCategories)
	}
}
