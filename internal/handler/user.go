package handler

import (
	"context"
	"fmt"
	"github.com/Imtiaz246/Thesis-Management-System/internal/apis/v1"
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
// @Router /students/request-register [post]
func (h *UserHandler) ReqRegister(ctx *gin.Context) {
	req := new(v1.ReqRegister)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, v1.ErrBadRequest, err.Error())
		return
	}

	studentInfo, err := h.userService.ReqRegister(context.TODO(), req)
	if err != nil {
		h.logger.WithContext(ctx).Error("userService.ReqRegister", zap.Error(err))
		v1.HandleError(ctx, err, nil)
		return
	}

	msg := fmt.Sprintf("Please verify your email `%s` to complete the further registration process", studentInfo.Email)
	v1.HandleSuccess(ctx, msg)
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
// @Router /students/verify-email [post]
func (h *UserHandler) VerifyEmail(ctx *gin.Context) {
	token := ctx.Query("token")
	studentInfo, err := h.userService.VerifyEmail(ctx, token)
	if err != nil {
		h.logger.WithContext(ctx).Error("userService.VerifyEmail", zap.Error(err))
		v1.HandleError(ctx, err, nil)
		return
	}

	v1.HandleSuccess(ctx, studentInfo)
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
// @Router /students/register [post]
func (h *UserHandler) Register(ctx *gin.Context) {
	req := new(v1.RegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, v1.ErrBadRequest, err.Error())
		return
	}

	if err := h.userService.Register(ctx, req, ctx.Query("token")); err != nil {
		h.logger.WithContext(ctx).Error("userService.Register", zap.Error(err))
		v1.HandleError(ctx, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
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
// @Router /login [post]
func (h *UserHandler) Login(ctx *gin.Context) {
	var req v1.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, v1.ErrBadRequest, nil)
		return
	}

	userTokens, err := h.userService.Login(ctx, &req)
	if err != nil {
		v1.HandleError(ctx, err, nil)
		return
	}

	v1.HandleSuccess(ctx, userTokens)
}

// GetProfile godoc
// @Summary Get user information
// @Schemes
// @Description Retrieves the profile information of a user
// @Tags User module
// @Accept json
// @Produce json
// @Security Bearer
// @Param uniId path string true "University ID"
// @Success 200 {object} v1.UserResponse
// @Failure 400 {object} v1.Response "Bad Request"
// @Failure 404 {object} v1.Response "Not Found"
// @Failure 500 {object} v1.Response "Internal Server Error"
// @Router /users/{uniId}/profile [get]
func (h *UserHandler) GetProfile(ctx *gin.Context) {
	requesterId := GetUserUniIdFromCtx(ctx)
	targetUserId := ctx.Param("uniId")

	profile, err := h.userService.GetProfile(ctx, targetUserId, requesterId)
	if err != nil {
		v1.HandleError(ctx, err, nil)
		return
	}

	v1.HandleSuccess(ctx, profile)
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
	userId := GetUserUniIdFromCtx(ctx)

	var req v1.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, v1.ErrBadRequest, nil)
		return
	}

	if err := h.userService.UpdateProfile(ctx, userId, &req); err != nil {
		v1.HandleError(ctx, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}
