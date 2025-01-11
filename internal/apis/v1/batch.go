package v1

import "time"

type CreateBatchRequest struct {
	Name         string    `json:"name" binding:"required"`
	Quota        string    `json:"quota"`
	PreDefenceAt time.Time `json:"preDefenceAt" binding:"required"`
	DefenceAt    time.Time `json:"defenceAt" binding:"required"`
}

type UpdateBatchRequest struct {
	Name         string    `json:"name,omitempty"`
	Quota        string    `json:"quota,omitempty"`
	PreDefenceAt time.Time `json:"preDefenceAt,omitempty"`
	DefenceAt    time.Time `json:"defenceAt,omitempty"`
}

type BatchInfo struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Quota        string    `json:"quota"`
	PreDefenceAt time.Time `json:"preDefenceAt"`
	DefenceAt    time.Time `json:"defenceAt"`

	CreatedBy *UserInfo `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type BatchResponse struct {
	Response
	Data BatchInfo
}
