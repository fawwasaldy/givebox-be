package controller

import (
	"github.com/gin-gonic/gin"
	request_chat "givebox/application/request/chat"
	"givebox/application/service"
	"givebox/platform/pagination"
	"givebox/presentation"
	"givebox/presentation/message"
	"net/http"
)

type (
	ChatController interface {
		ChatToDonor(ctx *gin.Context)
		SendMessage(ctx *gin.Context)
		GetAllConversationsByUserIDWithPagination(ctx *gin.Context)
		GetAllMessagesByConversationIDWithPagination(ctx *gin.Context)
	}

	chatController struct {
		chatService service.ChatService
	}
)

func NewChatController(chatService service.ChatService) ChatController {
	return &chatController{
		chatService: chatService,
	}
}

func (c chatController) ChatToDonor(ctx *gin.Context) {
	var req request_chat.FirstConversation
	if err := ctx.ShouldBind(&req); err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetDataFromBody, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	recipientID := ctx.MustGet("user_id").(string)

	result, err := c.chatService.ChatToDonor(ctx.Request.Context(), recipientID, req)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedChatToDonor, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessChatToDonor, result)
	ctx.JSON(http.StatusCreated, res)
}

func (c chatController) SendMessage(ctx *gin.Context) {
	var req request_chat.MessageSend
	if err := ctx.ShouldBind(&req); err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetDataFromBody, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	senderID := ctx.MustGet("user_id").(string)

	result, err := c.chatService.SendMessage(ctx.Request.Context(), senderID, req)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedSendMessage, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessSendMessage, result)
	ctx.JSON(http.StatusCreated, res)
}

func (c chatController) GetAllConversationsByUserIDWithPagination(ctx *gin.Context) {
	var req pagination.Request
	if err := ctx.ShouldBind(&req); err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetDataFromBody, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userID := ctx.MustGet("user_id").(string)

	result, err := c.chatService.GetAllConversationsByUserIDWithPagination(ctx.Request.Context(), userID, req)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetAllConversationsByUserIDWithPagination, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessGetAllConversationsByUserIDWithPagination, result)
	ctx.JSON(http.StatusOK, res)
}

func (c chatController) GetAllMessagesByConversationIDWithPagination(ctx *gin.Context) {
	var req pagination.Request
	if err := ctx.ShouldBind(&req); err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetDataFromBody, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	myUserID := ctx.MustGet("user_id").(string)

	conversationID := ctx.Param("conversation_id")
	result, err := c.chatService.GetAllMessagesByConversationIDWithPagination(ctx.Request.Context(), conversationID, myUserID, req)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetAllMessagesByConversationIDWithPagination, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessGetAllMessagesByConversationIDWithPagination, result)
	ctx.JSON(http.StatusOK, res)
}
