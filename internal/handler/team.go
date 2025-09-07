package handler

import (
	v1 "github.com/Imtiaz246/Thesis-Management-System/internal/apis/v1"
	teamservice "github.com/Imtiaz246/Thesis-Management-System/internal/service/team"
	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	*Handler
	teamService teamservice.Service
}

func NewTeamHandler(handler *Handler, teamService teamservice.Service) *TeamHandler {
	return &TeamHandler{
		Handler:     handler,
		teamService: teamService,
	}
}

func (h *TeamHandler) CreateTeam(ctx *gin.Context) {
	batchId, err := ParseUintParam(ctx, "id")
	if err != nil {
		v1.HandleError(ctx, v1.ErrBadRequest, err.Error())
		return
	}

	req := new(v1.CreateTeamRequest)
	if err = ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, v1.ErrBadRequest, err.Error())
		return
	}

	requesterUniId := GetUserUniIdFromCtx(ctx)
	team, err := h.teamService.CreateTeam(ctx, batchId, req, requesterUniId)
	if err != nil {
		v1.HandleError(ctx, err, nil)
		return
	}

	v1.HandleSuccess(ctx, team)
}

func (h *TeamHandler) GetJoinedTeam(ctx *gin.Context) {
	batchId, err := ParseUintParam(ctx, "id")
	if err != nil {
		v1.HandleError(ctx, v1.ErrBadRequest, err.Error())
		return
	}

	requesterUniId := GetUserUniIdFromCtx(ctx)
	team, err := h.teamService.GetJoinedTeam(ctx, batchId, requesterUniId)
	if err != nil {
		v1.HandleError(ctx, err, nil)
		return
	}

	v1.HandleSuccess(ctx, team)
}

func (h *TeamHandler) LeaveTeam(ctx *gin.Context) {
	teamId, err := ParseUintParam(ctx, "team_id")
	if err != nil {
		v1.HandleError(ctx, v1.ErrBadRequest, err.Error())
		return
	}
	requesterUniId := GetUserUniIdFromCtx(ctx)

	if err = h.teamService.LeaveTeam(ctx, teamId, requesterUniId); err != nil {
		v1.HandleError(ctx, err, nil)
		return
	}
}

func (h *TeamHandler) SendInvitation(ctx *gin.Context) {
	batchId, err := ParseUintParam(ctx, "id")
	if err != nil {
		v1.HandleError(ctx, v1.ErrBadRequest, err.Error())
		return
	}
	teamId, err := ParseUintParam(ctx, "team_id")
	if err != nil {
		v1.HandleError(ctx, v1.ErrBadRequest, err.Error())
		return
	}
	requesterUniId := GetUserUniIdFromCtx(ctx)
	targetStudentUniId := ctx.Param("target_student_uni_id")

	if err = h.teamService.SendInvitation(ctx, batchId, teamId, requesterUniId, targetStudentUniId); err != nil {
		v1.HandleError(ctx, err, nil)
		return
	}
}

func (h *TeamHandler) RejectInvitation(ctx *gin.Context) {
	teamId, err := ParseUintParam(ctx, "team_id")
	if err != nil {
		v1.HandleError(ctx, v1.ErrBadRequest, err.Error())
		return
	}
	requesterUniId := GetUserUniIdFromCtx(ctx)

	if err = h.teamService.RejectInvitation(ctx, teamId, requesterUniId); err != nil {
		v1.HandleError(ctx, err, nil)
		return
	}
}

func (h *TeamHandler) AcceptInvitation(ctx *gin.Context) {
	batchId, err := ParseUintParam(ctx, "id")
	if err != nil {
		v1.HandleError(ctx, v1.ErrBadRequest, err.Error())
		return
	}
	teamId, err := ParseUintParam(ctx, "team_id")
	if err != nil {
		v1.HandleError(ctx, v1.ErrBadRequest, err.Error())
		return
	}
	requesterUniId := GetUserUniIdFromCtx(ctx)

	if err = h.teamService.AcceptInvitation(ctx, batchId, teamId, requesterUniId); err != nil {
		v1.HandleError(ctx, err, nil)
		return
	}
}

func (h *TeamHandler) ListInvitations(ctx *gin.Context) {
	teamId, err := ParseUintParam(ctx, "team_id")
	if err != nil {
		v1.HandleError(ctx, v1.ErrBadRequest, err.Error())
		return
	}
	requesterUniId := GetUserUniIdFromCtx(ctx)

	teamInvitationInfos, err := h.teamService.ListInvitations(ctx, teamId, requesterUniId)
	if err != nil {
		v1.HandleError(ctx, err, nil)
		return
	}

	v1.HandleSuccess(ctx, teamInvitationInfos)
}
