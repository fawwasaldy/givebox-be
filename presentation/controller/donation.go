package controller

import (
	"github.com/gin-gonic/gin"
	request_donation "givebox/application/request/donation"
	"givebox/application/service"
	"givebox/platform/pagination"
	"givebox/presentation"
	"givebox/presentation/message"
	"net/http"
	"strconv"
)

type (
	DonationController interface {
		GetAllDonatedItemsWithPagination(ctx *gin.Context)
		GetAllDonatedItemsByCategoryIDWithPagination(ctx *gin.Context)
		GetAllDonatedItemsByConditionWithPagination(ctx *gin.Context)
		GetAllDonatedItemsByStatusWithPagination(ctx *gin.Context)
		GetAllDonatedItemsBeforeDateWithPagination(ctx *gin.Context)
		GetDonatedItemByID(ctx *gin.Context)
		OpenDonatedItem(ctx *gin.Context)
		RequestDonatedItem(ctx *gin.Context)
		AcceptDonatedItem(ctx *gin.Context)
		RejectDonatedItem(ctx *gin.Context)
		TakenDonatedItem(ctx *gin.Context)
		UpdateDonatedItem(ctx *gin.Context)
		DeleteDonatedItem(ctx *gin.Context)
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

func (c *donationController) GetAllDonatedItemsByConditionWithPagination(ctx *gin.Context) {
	var req pagination.Request
	if err := ctx.ShouldBind(&req); err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetDataFromBody, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	condition, err := strconv.Atoi(ctx.Param("condition"))
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetDataFromBody, "Invalid condition parameter", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	result, err := c.donationService.GetAllDonatedItemsByConditionWithPagination(ctx.Request.Context(), condition, req)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetAllDonatedItemsByConditionWithPagination, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.Response{
		Status:  true,
		Message: message.SuccessGetAllDonatedItemsByConditionWithPagination,
		Data:    result.Data,
		Meta:    result.Response,
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *donationController) GetAllDonatedItemsByStatusWithPagination(ctx *gin.Context) {
	var req pagination.Request
	if err := ctx.ShouldBind(&req); err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetDataFromBody, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	status := ctx.Param("status")
	result, err := c.donationService.GetAllDonatedItemsByStatusWithPagination(ctx.Request.Context(), status, req)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetAllDonatedItemsByStatusWithPagination, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.Response{
		Status:  true,
		Message: message.SuccessGetAllDonatedItemsByStatusWithPagination,
		Data:    result.Data,
		Meta:    result.Response,
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *donationController) GetAllDonatedItemsBeforeDateWithPagination(ctx *gin.Context) {
	var req pagination.Request
	if err := ctx.ShouldBind(&req); err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetDataFromBody, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	date := ctx.Param("date")
	result, err := c.donationService.GetAllDonatedItemsBeforeDateWithPagination(ctx.Request.Context(), date, req)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetAllDonatedItemsBeforeDateWithPagination, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.Response{
		Status:  true,
		Message: message.SuccessGetAllDonatedItemsBeforeDateWithPagination,
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
	var req request_donation.DonationItemOpen
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

func (c *donationController) RequestDonatedItem(ctx *gin.Context) {
	var req request_donation.DonationItemRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetDataFromBody, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	recipientID := ctx.MustGet("user_id").(string)
	result, err := c.donationService.RequestDonatedItem(ctx.Request.Context(), recipientID, req)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedRequestDonatedItem, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessRequestDonatedItem, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *donationController) AcceptDonatedItem(ctx *gin.Context) {
	var req request_donation.DonationItemAccept
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

func (c *donationController) RejectDonatedItem(ctx *gin.Context) {
	var req request_donation.DonationItemReject
	if err := ctx.ShouldBind(&req); err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetDataFromBody, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.donationService.RejectDonatedItem(ctx.Request.Context(), req)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedRejectDonatedItem, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessRejectDonatedItem, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *donationController) TakenDonatedItem(ctx *gin.Context) {
	var req request_donation.DonationItemTaken
	if err := ctx.ShouldBind(&req); err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetDataFromBody, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.donationService.TakenDonatedItem(ctx.Request.Context(), req)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedTakenDonatedItem, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessTakenDonatedItem, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *donationController) UpdateDonatedItem(ctx *gin.Context) {
	var req request_donation.DonationItemUpdate
	if err := ctx.ShouldBind(&req); err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetDataFromBody, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.donationService.UpdateDonatedItem(ctx.Request.Context(), req)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedUpdateDonatedItem, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessUpdateDonatedItem, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *donationController) DeleteDonatedItem(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.donationService.DeleteDonatedItem(ctx.Request.Context(), id); err != nil {
		res := presentation.BuildResponseFailed(message.FailedDeleteDonatedItem, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessDeleteDonatedItem, nil)
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
