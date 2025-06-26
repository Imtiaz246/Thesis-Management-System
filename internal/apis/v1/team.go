package v1

import "time"

type CreateTeamRequest struct {
	Name    string `json:"name" binding:"required"`
	Subject string `json:"subject" binding:"required"`
}

type TeamInfo struct {
	ID       uint           `json:"id"`
	BatchID  uint           `json:"batchID"`
	Name     string         `json:"name"`
	Subject  string         `json:"subject"`
	Students []*StudentInfo `json:"students,omitempty"`
	Teachers []*TeacherInfo `json:"teachers,omitempty"`
}

type TeamInvitationInfo struct {
	*TeamInfo
	InvitedAt time.Time `json:"invitedAt"`
}
