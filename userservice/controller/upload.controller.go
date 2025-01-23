package controller

import (
    "net/http"

    "auth/config" // Import the config package for error handling
    "user/usecase"

    "github.com/gin-gonic/gin"
)

type UploadController struct {
    uploadUC usecase.UploadProfileUsecase
}

func NewUploadController(upload_uc usecase.UploadProfileUsecase) *UploadController {
    return &UploadController{
        uploadUC: upload_uc,
    }
}

// UploadImg handles image upload
// @Summary Upload an image
// @Description Upload an image for a user
// @Tags Upload
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "User ID"
// @Param image formData file true "Image file"
// @Success 200 {string} string "success"
// @Failure 400 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/upload/{id} [post]
func (ctrl *UploadController) UploadImg() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        id := ctx.Param("id")
        file, err := ctx.FormFile("image")
        if err != nil {
            ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        err = ctx.SaveUploadedFile(file, "assets/uploads/"+file.Filename)
        if err != nil {
            ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        err = ctrl.uploadUC.UploadPicture("assets/uploads/"+file.Filename, id)
        if err != nil {
            ctx.IndentedJSON(config.GetStatusCode(err), gin.H{"error": err.Error()})
            return
        }

        ctx.IndentedJSON(http.StatusOK, gin.H{"message": "success"})
    }
}