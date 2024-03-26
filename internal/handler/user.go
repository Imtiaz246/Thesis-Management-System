package handler

import (
	"context"
	"github.com/Imtiaz246/Thesis-Management-System/api/v1"
	"github.com/Imtiaz246/Thesis-Management-System/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type UserHandler struct {
	*Handler
	userService service.UserService
}

func NewUserHandler(handler *Handler, userService service.UserService) *UserHandler {
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
// @Param request body v1.PreRegistrationRequest true "params"
// @Success 200 {object} v1.Response
// @Router /request-register [post]
func (h *UserHandler) ReqRegister(ctx *gin.Context) {
	req := new(v1.ReqRegister)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, err.Error())
		return
	}

	studentInfo, err := h.userService.ReqRegister(context.TODO(), req)
	if err != nil {
		h.logger.WithContext(ctx).Error("userService.ReqRegister", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, studentInfo)
}

// VerifyEmail godoc
// @Summary User email verification
// @Schemes
// @Description Confirms the email using which user opened account in the system
// @Tags User module
// @Accept json
// @Produce json
// @Query token path string true "Email confirmation token"
// @Success 200
// @Router /api/v1/verify_email [post]
func (h *UserHandler) VerifyEmail(ctx *gin.Context) {
	token := ctx.Query("token")
	studentInfo, err := h.userService.VerifyEmail(ctx, token)
	if err != nil {
		h.logger.WithContext(ctx).Error("userService.VerifyEmail", zap.Error(err))
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
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
// @Router /register [post]
func (h *UserHandler) Register(ctx *gin.Context) {
	req := new(v1.RegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, err.Error())
		return
	}

	if err := h.userService.Register(ctx, req, ctx.Query("token")); err != nil {
		h.logger.WithContext(ctx).Error("userService.Register", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
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
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	userTokens, err := h.userService.Login(ctx, &req)
	if err != nil {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	v1.HandleSuccess(ctx, userTokens)
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
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}

	user, err := h.userService.GetProfile(ctx, userId)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	v1.HandleSuccess(ctx, user)
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

	var req v1.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if err := h.userService.UpdateProfile(ctx, userId, &req); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}
