package controller

import (
	"github.com/gin-gonic/gin"
	request_donation "givebox/application/request/donation"
	"givebox/application/service"
	"givebox/platform/pagination"
	"givebox/presentation"
	"givebox/presentation/message"
	"net/http"
)

type (
	DonationController interface {
		GetAllDonatedItemsWithPagination(ctx *gin.Context)
		GetAllDonatedItemsByCategoryIDWithPagination(ctx *gin.Context)
		GetAllDonatedItemsByCityWithPagination(ctx *gin.Context)
		GetDonatedItemByID(ctx *gin.Context)
		OpenDonatedItem(ctx *gin.Context)
		AcceptDonatedItem(ctx *gin.Context)
		GetAllImagesByDonatedItemID(ctx *gin.Context)
		GetAllCategories(ctx *gin.Context)
	}

	donationController struct {
		donationService service.DonationService
	}
)

func NewDonationController(donationService service.DonationService) DonationController {
	return &donationController{
		donationService: donationService,
	}
}

func (c *donationController) GetAllDonatedItemsWithPagination(ctx *gin.Context) {
	var req pagination.Request
	if err := ctx.ShouldBind(&req); err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetDataFromBody, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.donationService.GetAllDonatedItemsWithPagination(ctx.Request.Context(), req)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetAllDonatedItemsWithPagination, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.Response{
		Status:  true,
		Message: message.SuccessGetAllDonatedItemsWithPagination,
		Data:    result.Data,
		Meta:    result.Response,
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *donationController) GetAllDonatedItemsByCategoryIDWithPagination(ctx *gin.Context) {
	var req pagination.Request
	if err := ctx.ShouldBind(&req); err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetDataFromBody, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	categoryID := ctx.Param("category_id")
	result, err := c.donationService.GetAllDonatedItemsByCategoryIDWithPagination(ctx.Request.Context(), categoryID, req)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetAllDonatedItemsByCategoryIDWithPagination, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.Response{
		Status:  true,
		Message: message.SuccessGetAllDonatedItemsByCategoryIDWithPagination,
		Data:    result.Data,
		Meta:    result.Response,
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *donationController) GetAllDonatedItemsByCityWithPagination(ctx *gin.Context) {
	var req pagination.Request
	if err := ctx.ShouldBind(&req); err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetDataFromBody, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	city := ctx.Param("city")
	result, err := c.donationService.GetAllDonatedItemsByCityWithPagination(ctx.Request.Context(), city, req)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetAllDonatedItemsByCityWithPagination, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.Response{
		Status:  true,
		Message: message.SuccessGetAllDonatedItemsByCityWithPagination,
		Data:    result.Data,
		Meta:    result.Response,
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *donationController) GetDonatedItemByID(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := c.donationService.GetDonatedItemByID(ctx.Request.Context(), id)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetDonatedItemByID, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessGetDonatedItemByID, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *donationController) OpenDonatedItem(ctx *gin.Context) {
	var req request_donation.DonatedItemOpen
	if err := ctx.ShouldBind(&req); err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetDataFromBody, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	donorID := ctx.MustGet("user_id").(string)

	result, err := c.donationService.OpenDonatedItem(ctx.Request.Context(), donorID, req)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedOpenDonatedItem, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessOpenDonatedItem, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *donationController) AcceptDonatedItem(ctx *gin.Context) {
	var req request_donation.DonatedItemAccept
	if err := ctx.ShouldBind(&req); err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetDataFromBody, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.donationService.AcceptDonatedItem(ctx.Request.Context(), req)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedAcceptDonatedItem, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessAcceptDonatedItem, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *donationController) GetAllImagesByDonatedItemID(ctx *gin.Context) {
	id := ctx.Param("donated_item_id")
	images, err := c.donationService.GetAllImagesByDonatedItemID(ctx.Request.Context(), id)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetAllImagesByDonatedItemID, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessGetAllImagesByDonatedItemID, images)
	ctx.JSON(http.StatusOK, res)
}

func (c *donationController) GetAllCategories(ctx *gin.Context) {
	categories, err := c.donationService.GetAllCategories(ctx.Request.Context())
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetAllCategories, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessGetAllCategories, categories)
	ctx.JSON(http.StatusOK, res)
}
