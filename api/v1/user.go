package v1

type ReqRegister struct {
	UniversityId string `json:"universityId" binding:"required,alphanum,len=7" example:"c183074"`
}

type RegisterRequest struct {
	Name            string `json:"name" binding:"required"`
	Mobile          string `json:"mobile" binding:"required"`
	AlternateMobile string `json:"alternateMobile"`
	Section         string `json:"section" binding:"required"`
	Password        string `json:"password" binding:"required,min=6" example:"123456"`
}

type StudentInfo struct {
	UniversityId     string
	Email            string
	FullName         string
	Country          string
	Gender           string
	Department       string
	CGPA             float64
	Batch            uint8
	CompletedCredits uint8
	Nationality      string
	Mobile           string
	AlternateMobile  string
}

type LoginRequest struct {
	UniversityId string `json:"universityId" binding:"required" example:"c183074"`
	Password     string `json:"password" binding:"required" example:"123456"`
}

type LoginResponseData struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginResponse struct {
	Response
	Data LoginResponseData
}

type UpdateProfileRequest struct {
	Nickname string `json:"nickname" example:"alan"`
	Email    string `json:"email" binding:"required,email" example:"1234@gmail.com"`
}

type GetProfileResponseData struct {
	UserId   string `json:"userId"`
	Nickname string `json:"nickname" example:"alan"`
}

type GetProfileResponse struct {
	Response
	Data GetProfileResponseData
}
