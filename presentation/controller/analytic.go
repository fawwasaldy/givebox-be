package controller

import (
	"github.com/gin-gonic/gin"
	response_donation "givebox/application/response/donation"
	"givebox/application/service"
	"givebox/presentation"
	"givebox/presentation/message"
	"net/http"
)

type (
	AnalyticController interface {
		GetTotalDonatedItems(ctx *gin.Context)
		GetTotalUsers(ctx *gin.Context)
		GetTotalOpenedDonatedItems(ctx *gin.Context)
		GetAcceptedDonatedItemPercentage(ctx *gin.Context)
		GetSixCategoriesByMostDonatedItems(ctx *gin.Context)
		GetThreeLatestDonatedItems(ctx *gin.Context)
	}

	analyticController struct {
		analyticService service.AnalyticService
	}
)

func NewAnalyticController(analyticService service.AnalyticService) AnalyticController {
	return &analyticController{
		analyticService: analyticService,
	}
}

func (a analyticController) GetTotalDonatedItems(ctx *gin.Context) {
	var result int64
	var err error

	result, err = a.analyticService.GetTotalDonatedItems(ctx.Request.Context())
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetTotalDonatedItems, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessGetTotalDonatedItems, result)
	ctx.JSON(http.StatusOK, res)
}

func (a analyticController) GetTotalUsers(ctx *gin.Context) {
	var result int64
	var err error

	result, err = a.analyticService.GetTotalUsers(ctx.Request.Context())
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetTotalUsers, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessGetTotalUsers, result)
	ctx.JSON(http.StatusOK, res)
}

func (a analyticController) GetTotalOpenedDonatedItems(ctx *gin.Context) {
	var result int64
	var err error

	result, err = a.analyticService.GetTotalOpenedDonatedItems(ctx.Request.Context())
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetTotalOpenedDonatedItems, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessGetTotalOpenedDonatedItems, result)
	ctx.JSON(http.StatusOK, res)
}

func (a analyticController) GetAcceptedDonatedItemPercentage(ctx *gin.Context) {
	var result float64
	var err error

	result, err = a.analyticService.GetAcceptedDonatedItemPercentage(ctx.Request.Context())
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetAcceptedDonatedItemPercentage, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessGetAcceptedDonatedItemPercentage, result)
	ctx.JSON(http.StatusOK, res)
}

func (a analyticController) GetSixCategoriesByMostDonatedItems(ctx *gin.Context) {
	var result []response_donation.CategoryAnalytic
	var err error

	result, err = a.analyticService.GetSixCategoriesByMostDonatedItems(ctx.Request.Context())
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetSixCategoriesByMostDonatedItems, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessGetSixCategoriesByMostDonatedItems, result)
	ctx.JSON(http.StatusOK, res)
}

func (a analyticController) GetThreeLatestDonatedItems(ctx *gin.Context) {
	var result []response_donation.DonatedItem
	var err error

	result, err = a.analyticService.GetThreeLatestDonatedItems(ctx.Request.Context())
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetThreeLatestDonatedItems, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessGetThreeLatestDonatedItems, result)
	ctx.JSON(http.StatusOK, res)
}
