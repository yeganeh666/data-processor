package handlers

import (
	quotaApi "IofIPOS/quota/api/src"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandlePreUpload
// @Summary HandlePreUpload
// @Description check quota and object details before upload
// @Tags Objects
// @Accept json
// @Produce json
// @Param request body api.PreUploadObjectReq true "query params"
// @Success 200 {object} api.PreUploadObjectRes
// @Router /pre-upload [post]
func (h *ObjectHandlerImpl) HandlePreUpload(ctx *gin.Context) {
	model := new(quotaApi.PreUploadObjectReq)
	err := ctx.ShouldBind(model)
	if err != nil {
		h.log.WithError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	objectDetails, err := h.service.ObjectService.PreUpload(ctx, model)
	if err != nil {
		h.log.WithError(err).Error("error while checking pre upload")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, objectDetails)
}
