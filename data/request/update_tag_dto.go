package request

import "github.com/google/uuid"

type UpdateTagsRequest struct {
	Id   uuid.UUID `validate:"required"`
	Name string    `validate:"required,max=200,min=1" json:"name"`
}
