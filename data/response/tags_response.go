package response

import "github.com/google/uuid"

type TagsResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
