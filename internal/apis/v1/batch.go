package v1

import "time"

type CreateBatchRequest struct {
	Name            string  `json:"name" binding:"required"`
	Quota           string  `json:"quota"`
	MinCGPARequired float64 `json:"minCGPARequired"`
	MinCHRequired   uint16  `json:"minCHRequired" binding:"required"`

	TeamRegDeadline time.Time `json:"teamRegDeadline" binding:"required"`
	MaxTeamMember   uint8     `json:"maxTeamMember" binding:"required"`
	MaxTeacherPref  uint8     `json:"maxTeacherPref" binding:"required"`
	PreDefenceAt    time.Time `json:"preDefenceAt" binding:"required"`
	DefenceAt       time.Time `json:"defenceAt" binding:"required"`
}

type UpdateBatchRequest struct {
	Name            string  `json:"name"`
	Quota           string  `json:"quota"`
	MinCGPARequired float64 `json:"minCGPARequired"`
	MinCHRequired   uint16  `json:"minCHRequired"`

	TeamRegDeadline *time.Time `json:"teamRegDeadline" `
	MaxTeamMember   uint8      `json:"maxTeamMember"`
	MaxTeacherPref  uint8      `json:"maxTeacherPref"`
	PreDefenceAt    *time.Time `json:"preDefenceAt"`
	DefenceAt       *time.Time `json:"defenceAt"`
}

type BatchInfo struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Quota string `json:"quota"`

	MinCGPARequired float64   `json:"minCGPARequired"`
	MinCHRequired   uint16    `json:"minCHRequired"`
	TeamRegDeadline time.Time `json:"teamRegDeadline"`
	MaxTeamMember   uint8     `json:"maxTeamMember"`
	MaxTeacherPref  uint8     `json:"maxTeacherPref"`
	PreDefenceAt    time.Time `json:"preDefenceAt"`
	DefenceAt       time.Time `json:"defenceAt"`
	Closed          bool      `json:"closed"`

	CreatedBy *UserInfo `json:"createdBy,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type BatchResponse struct {
	Response
	Data BatchInfo
}
