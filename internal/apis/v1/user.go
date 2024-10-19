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

type UserResponseData struct {
	UniversityId string `json:"universityId"`
	Email        string `json:"email,omitempty"`
	IsAdmin      bool   `json:"isAdmin,omitempty"`
	Role         string `json:"role"`
	Gender       string `json:"gender"`
	IsVerified   bool   `json:"isVerified,omitempty"`
	ChangePass   bool   `json:"changePass,omitempty"`

	Student *StudentInfo `json:"student,omitempty"`
	Teacher *TeacherInfo `json:"teacher,omitempty"`
	Stuff   *StuffInfo   `json:"stuff,omitempty"`
}

type UserResponse struct {
	Response
	Data UserResponseData
}

type StudentInfo struct {
	UniversityId     string  `json:"universityId,omitempty"`
	Email            string  `json:"email,omitempty"`
	FullName         string  `json:"fullName"`
	Country          string  `json:"country"`
	Gender           string  `json:"gender,omitempty"`
	Department       string  `json:"department,omitempty"`
	CGPA             float64 `json:"cgpa,omitempty"`
	Batch            uint16  `json:"batch,omitempty"`
	Section          string  `json:"section,omitempty"`
	CompletedCredits uint16  `json:"completedCredits,omitempty"`
	Mobile           string  `json:"mobile,omitempty"`
	AlternateMobile  string  `json:"alternateMobile,omitempty"`
}

type TeacherInfo struct {
	FullName        string `json:"fullName"`
	Mobile          string `json:"mobile,omitempty"`
	AlternateMobile string `json:"alternateMobile,omitempty"`
	Department      string `json:"department"`
	Designation     string `json:"designation"`
}

type StuffInfo struct {
	FullName        string `json:"fullName"`
	Mobile          string `json:"mobile,omitempty"`
	AlternateMobile string `json:"alternateMobile,omitempty"`
	Department      string `json:"department"`
}

type UpdateProfileRequest struct {
	Nickname string `json:"nickname" example:"alan"`
	Email    string `json:"email" binding:"required,email" example:"1234@gmail.com"`
}
