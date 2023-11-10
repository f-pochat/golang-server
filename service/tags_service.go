package service

import (
	"f-pochat/golang-server/data/request"
	"f-pochat/golang-server/data/response"

	"github.com/google/uuid"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId uuid.UUID)
	FindById(tagsId uuid.UUID) response.TagsResponse
	FindAll() []response.TagsResponse
}
