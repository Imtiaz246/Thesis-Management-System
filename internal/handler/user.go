package handler

import (
	"context"
	"fmt"
	apisv1 "github.com/Imtiaz246/Thesis-Management-System/internal/apis/v1"
	userservice "github.com/Imtiaz246/Thesis-Management-System/internal/service/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	*Handler
	userService userservice.Service
}

func NewUserHandler(handler *Handler, userService userservice.Service) *UserHandler {
	return &UserHandler{
		Handler:     handler,
		userService: userService,
	}
}

// ReqRegister godoc
// @Summary User request-register
// @Schemes
// @Description To get the user info from the vendor and cache it
// @Tags User module
// @Accept json
// @Produce json
// @Param request body v1.ReqRegister true "params"
// @Success 200 {object} v1.Response
// @Router /api/v1/students/request-register [post]
func (h *UserHandler) ReqRegister(ctx *gin.Context) {
	req := new(apisv1.ReqRegister)
	if err := ctx.ShouldBindJSON(req); err != nil {
		apisv1.HandleError(ctx, apisv1.ErrBadRequest, err.Error())
		return
	}

	studentInfo, err := h.userService.ReqRegister(context.TODO(), req)
	if err != nil {
		h.logger.WithContext(ctx).Error("userService.ReqRegister", zap.Error(err))
		apisv1.HandleError(ctx, err, nil)
		return
	}

	successMsg := fmt.Sprintf("Please verify your email `%s` to complete the further registration process", studentInfo.Email)
	apisv1.HandleSuccess(ctx, successMsg)
}

// VerifyEmail godoc
// @Summary User email verification
// @Schemes
// @Description Confirms the email and sends the pre-saved cache student data got from IIUC server
// @Tags User module
// @Accept json
// @Produce json
// @Param token query string true "Email confirmation token"
// @Success 200 {object} v1.Response
// @Router /api/v1/students/verify-email [post]
func (h *UserHandler) VerifyEmail(ctx *gin.Context) {
	token := ctx.Query("token")
	studentInfo, err := h.userService.VerifyEmail(ctx, token)
	if err != nil {
		h.logger.WithContext(ctx).Error("userService.VerifyEmail", zap.Error(err))
		apisv1.HandleError(ctx, err, nil)
		return
	}

	apisv1.HandleSuccess(ctx, studentInfo)
}

// Register godoc
// @Summary User registration
// @Schemes
// @Description To register user account to the system
// @Tags User module
// @Accept json
// @Produce json
// @Param request body v1.RegisterRequest true "params"
// @Success 200 {object} v1.Response
// @Router /api/v1/students/register [post]
func (h *UserHandler) Register(ctx *gin.Context) {
	req := new(apisv1.RegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		apisv1.HandleError(ctx, apisv1.ErrBadRequest, err.Error())
		return
	}

	if err := h.userService.Register(ctx, req, ctx.Query("token")); err != nil {
		h.logger.WithContext(ctx).Error("userService.Register", zap.Error(err))
		apisv1.HandleError(ctx, err, nil)
		return
	}

	apisv1.HandleSuccess(ctx, nil)
}

// Login godoc
// @Summary User login
// @Schemes
// @Description
// @Tags User module
// @Accept json
// @Produce json
// @Param request body v1.LoginRequest true "params"
// @Success 200 {object} v1.LoginResponse
// @Router /api/v1/login [post]
func (h *UserHandler) Login(ctx *gin.Context) {
	var req apisv1.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apisv1.HandleError(ctx, apisv1.ErrBadRequest, nil)
		return
	}

	userTokens, err := h.userService.Login(ctx, &req)
	if err != nil {
		apisv1.HandleError(ctx, err, nil)
		return
	}

	apisv1.HandleSuccess(ctx, userTokens)
}

// GetProfile godoc
// @Summary Get user information
// @Schemes
// @Description
// @Tags User module
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} v1.GetProfileResponse
// @Router /user [get]
func (h *UserHandler) GetProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		apisv1.HandleError(ctx, apisv1.ErrUnauthorized, nil)
		return
	}

	user, err := h.userService.GetProfile(ctx, userId)
	if err != nil {
		apisv1.HandleError(ctx, err, nil)
		return
	}

	apisv1.HandleSuccess(ctx, user)
}

// UpdateProfile godoc
// @Summary Update user information
// @Schemes
// @Description
// @Tags User module
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200
// @Router /user [put]
func (h *UserHandler) UpdateProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)

	var req apisv1.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apisv1.HandleError(ctx, apisv1.ErrBadRequest, nil)
		return
	}

	if err := h.userService.UpdateProfile(ctx, userId, &req); err != nil {
		apisv1.HandleError(ctx, err, nil)
		return
	}

	apisv1.HandleSuccess(ctx, nil)
}
