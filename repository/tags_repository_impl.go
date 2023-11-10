package repository

import (
	"f-pochat/golang-server/model"

	"github.com/f-pochat/golang-base/repository"
	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	TagsRepository
}

func NewTagsRepositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{TagsRepository: repository.NewBaseRepositoryImpl[model.Tags](Db)}
}
