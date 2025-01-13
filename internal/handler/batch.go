package handler

import (
	v1 "github.com/Imtiaz246/Thesis-Management-System/internal/apis/v1"
	batchservice "github.com/Imtiaz246/Thesis-Management-System/internal/service/batch"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type BatchHandler struct {
	*Handler
	batchService batchservice.Service
}

func NewBatchHandler(handler *Handler, batchService batchservice.Service) *BatchHandler {
	return &BatchHandler{
		Handler:      handler,
		batchService: batchService,
	}
}

// ListBatch godoc
// @Summary Get list of batches
// @Schemes
// @Description Retrieves a list of batches with optional pagination
// @Tags Batch module
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param pageSize query int false "Number of records per page"
// @Success 200 {object} v1.Response
// @Router /batch [get]
func (h *BatchHandler) ListBatch(ctx *gin.Context) {
	batches, err := h.batchService.ListBatches(ctx)
	if err != nil {
		h.logger.WithContext(ctx).Error("batchService.ListBatches", zap.Error(err))
		v1.HandleError(ctx, err, nil)
		return
	}

	data := gin.H{
		"batches": batches,
	}
	v1.HandleSuccess(ctx, data)
}

// CreateBatch godoc
// @Summary Create a new batch
// @Schemes
// @Description Creates a new batch
// @Tags Batch module
// @Accept json
// @Produce json
// @Param request body v1.CreateBatchRequest true "Batch creation details"
// @Success 200 {object} v1.Response
// @Router /batch [post]
func (h *BatchHandler) CreateBatch(ctx *gin.Context) {
	req := new(v1.CreateBatchRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, v1.ErrBadRequest, err.Error())
		return
	}

	if req.TeamRegDeadline.After(req.PreDefenceAt) {
		v1.HandleError(ctx, v1.ErrBadRequest, "Team registration deadline should be before pre-defence date")
		return
	}
	if req.PreDefenceAt.After(req.DefenceAt) {
		v1.HandleError(ctx, v1.ErrBadRequest, "Pre-defence date should be before defence date")
		return
	}

	requesterUniId := GetUserUniIdFromCtx(ctx)
	err := h.batchService.CreateBatch(ctx, requesterUniId, req)
	if err != nil {
		h.logger.WithContext(ctx).Error("batchService.CreateBatch", zap.Error(err))
		v1.HandleError(ctx, err, nil)
		return
	}

	v1.HandleSuccess(ctx, "Batch created successfully")
}

// GetBatch godoc
// @Summary Get batch details
// @Schemes
// @Description Retrieves details of a batch
// @Tags Batch module
// @Accept json
// @Produce json
// @Param id path int true "Batch ID"
// @Success 200 {object} v1.BatchResponse
// @Router /batch/{id} [get]
func (h *BatchHandler) GetBatch(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		v1.HandleError(ctx, v1.ErrBadRequest, err.Error())
		return
	}

	batch, err := h.batchService.GetBatchById(ctx, uint(id))
	if err != nil {
		h.logger.WithContext(ctx).Error("batchService.GetBatch", zap.Error(err))
		v1.HandleError(ctx, err, nil)
		return
	}

	v1.HandleSuccess(ctx, batch)
}

// UpdateBatch godoc
// @Summary Update batch details
// @Schemes
// @Description Updates an existing batch
// @Tags Batch module
// @Accept json
// @Produce json
// @Param id path int true "Batch ID"
// @Param request body v1.UpdateBatchRequest true "Batch update details"
// @Success 200 {object} v1.Response
// @Router /batch/{id} [put]
func (h *BatchHandler) UpdateBatch(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		v1.HandleError(ctx, v1.ErrBadRequest, err.Error())
		return
	}

	req := new(v1.UpdateBatchRequest)
	if err = ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, v1.ErrBadRequest, err.Error())
		return
	}

	requesterUniId := GetUserUniIdFromCtx(ctx)
	err = h.batchService.UpdateBatch(ctx, requesterUniId, uint(id), req)
	if err != nil {
		h.logger.WithContext(ctx).Error("batchService.UpdateBatch", zap.Error(err))
		v1.HandleError(ctx, err, nil)
		return
	}

	v1.HandleSuccess(ctx, "Batch updated successfully")
}

// DeleteBatch godoc
// @Summary Delete batch
// @Schemes
// @Description Deletes an existing batch
// @Tags Batch module
// @Accept json
// @Produce json
// @Param id path int true "Batch ID"
// @Success 200 {object} v1.Response
// @Router /batch/{id} [delete]
func (h *BatchHandler) DeleteBatch(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		v1.HandleError(ctx, v1.ErrBadRequest, err.Error())
		return
	}

	requesterUniId := GetUserUniIdFromCtx(ctx)
	err = h.batchService.DeleteBatch(ctx, requesterUniId, uint(id))
	if err != nil {
		h.logger.WithContext(ctx).Error("batchService.DeleteBatch", zap.Error(err))
		v1.HandleError(ctx, err, nil)
		return
	}

	v1.HandleSuccess(ctx, "Batch deleted successfully")
}
