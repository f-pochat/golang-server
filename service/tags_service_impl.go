package service

import (
	"f-pochat/golang-server/data/request"
	"f-pochat/golang-server/data/response"
	"f-pochat/golang-server/model"
	"f-pochat/golang-server/repository"

	"github.com/f-pochat/golang-base/helper"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type TagsServiceImpl struct {
	TagRepository repository.TagsRepository
	Validate      *validator.Validate
}

func NewTagServiceImpl(tagRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{
		TagRepository: tagRepository,
		Validate:      validate,
	}
}

func (t TagsServiceImpl) Create(tag request.CreateTagsRequest) {
	err := t.Validate.Struct(tag)
	helper.ErrorPanic(err)
	tagModel := model.Tags{
		Name: tag.Name,
	}
	t.TagRepository.Save(tagModel)
}

func (t TagsServiceImpl) Update(tag request.UpdateTagsRequest) {
	tagData, err := t.TagRepository.FindById(tag.Id)
	helper.ErrorPanic(err)
	tagData.Name = tag.Name
	t.TagRepository.Update(tagData)
}

func (t TagsServiceImpl) Delete(tagId uuid.UUID) {
	t.TagRepository.Delete(tagId)
}

func (t TagsServiceImpl) FindById(tagId uuid.UUID) response.TagsResponse {
	tagData, err := t.TagRepository.FindById(tagId)
	helper.ErrorPanic(err)

	tagResponse := response.TagsResponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}
	return tagResponse
}

func (t TagsServiceImpl) FindAll() []response.TagsResponse {
	result := t.TagRepository.FindAll()

	var tags []response.TagsResponse
	for _, value := range result {
		tag := response.TagsResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}
	return tags
}
