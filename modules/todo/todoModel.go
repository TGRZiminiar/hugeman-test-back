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
		Title       string `json:"title"`
		Description string `json:"description"`
		Image       string `json:"image"`
		Status      string `json:"status"`
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
