package user

import (
	"bytes"
	"fmt"
	"github.com/Imtiaz246/Thesis-Management-System/internal/apis/v1"
	"github.com/Imtiaz246/Thesis-Management-System/internal/templates"
	"html/template"
	"time"
)

func renderEmailVerifyTemplate(url, universityId, token string, expirationTime time.Time) ([]byte, error) {
	tmpl, err := template.New("email_verification").Parse(templates.EmailVerifyTmpl)
	if err != nil {
		return nil, err
	}
	data := map[string]interface{}{
		"Link":    fmt.Sprintf("%s?universityId=%s,token=%s", url, universityId, token),
		"ExpTime": expirationTime.Format(time.DateTime),
	}
	var out bytes.Buffer
	if err = tmpl.Execute(&out, data); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func mockStudentInfoApi(universityId string) (*v1.StudentInfo, error) {
	return &v1.StudentInfo{
		UniversityId:     universityId,
		Email:            fmt.Sprintf("%s@ugrad.iiuc.ac.bd", universityId),
		Department:       "CSE",
		CGPA:             3.50,
		Batch:            47,
		CompletedCredits: 160,
		Country:          "Bangladesh",
		Mobile:           "01773820154",
		AlternateMobile:  "01773820154",
	}, nil
}
