package route

import (
	"github.com/gin-gonic/gin"
	"givebox/presentation/controller"
)

func AnalyticRoute(route *gin.Engine, analyticController controller.AnalyticController) {
	analyticGroup := route.Group("/api/analytic")
	{
		analyticGroup.GET("/total-donated-items", analyticController.GetTotalDonatedItems)
		analyticGroup.GET("/total-users", analyticController.GetTotalUsers)
		analyticGroup.GET("/total-opened-donated-items", analyticController.GetTotalOpenedDonatedItems)
		analyticGroup.GET("/accepted-donated-item-percentage", analyticController.GetAcceptedDonatedItemPercentage)
		analyticGroup.GET("/six-categories-by-most-donated-items", analyticController.GetSixCategoriesByMostDonatedItems)
		analyticGroup.GET("/three-latest-donated-items", analyticController.GetThreeLatestDonatedItems)
	}
}
