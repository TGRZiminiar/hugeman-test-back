package todo

import (
	"time"
)

type (
	TodoShowcase struct {
		Id          string    `json:"_id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		Image       string    `json:"image"`
		Status      string    `json:"status"`
	}

	CreateTodoReq struct {
		Title       string `json:"title" form:"title" validate:"required,max=100"`
		Description string `json:"description" form:"description" validate:"required,max=64"`
		Image       string `json:"image"`
		Status      string `json:"status" form:"status" validate:"required,max=64"`
	}

	UpdateTodoReq struct {
		Id          string `json:"_id" form:"_id" validate:"required"`
		Title       string `json:"title" form:"title" validate:"required,max=100"`
		Description string `json:"description" form:"description" validate:"required,max=64"`
		Image       string `json:"image"`
		Status      string `json:"status" form:"status" validate:"required,max=64"`
	}
)

// | Field  | Data Type | Notes |
// | ------ | --------- | ----- |
// | ID | UUID |  |
// | Title | String |  |
// | Description | String |  |
// | Created At Date Time | Date Time with Time Zone |  |
// | Image | String | base64 |
// | Status | String  | Accept: `IN_PROGRESS` \| `COMPLETED` |
