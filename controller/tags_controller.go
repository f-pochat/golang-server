package controller

import (
	"f-pochat/golang-server/data/request"
	"f-pochat/golang-server/data/response"
	"f-pochat/golang-server/service"
	"net/http"

	"github.com/f-pochat/golang-base/helper"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TagController struct {
	tagService service.TagsService
}

func NewTagController(service service.TagsService) *TagController {
	return &TagController{tagService: service}
}

func (controller *TagController) Create(ctx *gin.Context) {
	createTagRequest := request.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&createTagRequest)
	helper.ErrorPanic(err)

	controller.tagService.Create(createTagRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TagController) Update(ctx *gin.Context) {
	updateTagRequest := request.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")

	updateTagRequest.Id, err = uuid.Parse(tagId)
	helper.ErrorPanic(err)

	controller.tagService.Update(updateTagRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TagController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := uuid.Parse(tagId)
	helper.ErrorPanic(err)
	controller.tagService.Delete(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *TagController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := uuid.Parse(tagId)
	helper.ErrorPanic(err)

	tagResponse := controller.tagService.FindById(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TagController) FindAll(ctx *gin.Context) {
	tagResponse := controller.tagService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
