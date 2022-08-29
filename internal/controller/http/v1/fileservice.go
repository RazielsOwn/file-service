package v1

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"

	"file-service/internal/entity"
	"file-service/internal/usecase"
	"file-service/pkg/logger"
)

type fileServiceRoutes struct {
	t usecase.IFileStore
	l logger.Interface
}

func addFileServiceRoutes(handler *gin.RouterGroup, t usecase.IFileStore, l logger.Interface) {
	r := &fileServiceRoutes{t, l}

	h := handler.Group("/fileService")
	{
		h.GET("/getFile", r.getFile)
		h.POST("/uploadFile", r.uploadFile)
	}
}

type fileResponse struct {
	File entity.FileEntity `json:"file"`
}

// @Summary     get file
// @Description get file by id
// @ID          getFile
// @Tags  	    fileService
// @Accept      json
// @Produce     json
// @Param id query int true "File ID"
// @Success     200 {object} fileResponse
// @Failure     500 {object} response
// @Router      /v1/fileService/getFile [get]
func (r *fileServiceRoutes) getFile(c *gin.Context) {
	var params = c.Request.URL.Query()
	idVar, _ := strconv.Atoi(params["id"][0])
	file, err := r.t.GetFileById(c.Request.Context(), idVar)
	if err != nil {
		r.l.Error(err, "http - v1 - getFile")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	file.Path = fmt.Sprintf("/files/%v", file.Id)
	c.JSON(http.StatusOK, fileResponse{file})
}

// type fileJson struct {
// 	Name        string                `form:"name" binding:"required"`
// 	Description string                `form"description"`
// 	fileData    *multipart.FileHeader `form:"file"`
// }

// @Summary     Upload File
// @Tags  	    Upload File
// @Accept      multipart/form-data
// @Produce     json
// @Param file formData file true "Body with file"
// @Success     200 {object} entity.FileEntity
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /v1/fileService/uploadFile [post]
func (r *fileServiceRoutes) uploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	file_name := header.Filename
	file_path := "./public/" + file_name
	if err := os.MkdirAll(filepath.Dir(file_path), 0770); err != nil {
		log.Fatal(err)
		return
	}
	out, err := os.Create(file_path)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
		return
	}

	savedFile, err := r.t.SaveFile(
		c.Request.Context(),
		entity.FileEntity{
			Name: file_name,
			Path: file_path,
		},
	)
	if err != nil {
		r.l.Error(err, "http - v1 - uploadFile")
		errorResponse(c, http.StatusInternalServerError, "upload service problems")

		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("file id:%v uploaded succesfully", savedFile.Id))
}
