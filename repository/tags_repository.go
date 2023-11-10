package repository

import (
	"f-pochat/golang-server/model"

	"github.com/f-pochat/golang-base/repository"
)

type TagsRepository interface {
	repository.BaseRepository[model.Tags]
}
