package classin

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Students client
type students struct {
	client
}
type Students interface {
	AddStudentToSchool(studentAccount string, studentName string, school int) (StudentResponse, error)
	ChangeStudentName(classinAccountId string, studentName string, school int) (StudentResponse, error)
	ChangeStudentPassword(classinAccountId string, phone string, newPassword string, school int) (StudentResponse, error)
	AddStudentToCourse(classinAccountId string, classinCourseId string, studentName string, school int) (StudentResponse, error)
	RegisterAccountId(phoneNumber string, contactId string, school int) (StudentResponse, error)
	RemoveStudentInCourse(accountUid []string, courseId string, school int) (DeleteStudentsResponse, error)
}

// Students constructor (from Client)
func (c client) Students() Students {
	return &students{
		client: c,
	}
}

type StudentResponse struct {
	Data      int       `json:"data,omitempty"`
	ErrorInfo ErrorInfo `json:"error_info"`
}

type ErrorInfo struct {
	ErrNo int    `json:"errno"`
	Error string `json:"error"`
}

type DeleteStudentsResponse struct {
	Data      []ErrorInfo `json:"data,omitempty"`
	ErrorInfo ErrorInfo   `json:"error_info"`
}

// AddStudentToSchool POST
func (c students) AddStudentToSchool(studentAccount string, studentName string, school int) (StudentResponse, error) {
	var r StudentResponse
	country := Country
	params := map[string]string{
		"studentAccount": country + "-" + studentAccount,
		"studentName":    studentName,
	}
	err := c.client.Request("POST", "/partner/api/course.api.php?action=addSchoolStudent", params, &r, school)
	return r, err
}

// ChangeStudentName POST
func (c students) ChangeStudentName(classinAccountId string, studentName string, school int) (StudentResponse, error) {
	var r StudentResponse
	params := map[string]string{
		"studentUid":  classinAccountId,
		"studentName": studentName,
	}
	err := c.client.Request("POST", "/partner/api/course.api.php?action=editSchoolStudent", params, &r, school)
	return r, err
}

// ChangeStudentPassword POST
func (c students) ChangeStudentPassword(accountId string, phone string, newPassword string, school int) (StudentResponse, error) {
	var r StudentResponse
	country := Country
	params := map[string]string{
		"uid":       accountId,
		"password":  newPassword,
		"telephone": country + "-" + phone,
	}

	err := c.client.Request("POST", "/partner/api/course.api.php?action=modifyPasswordByTelephone", params, &r, school)
	return r, err
}

// AddStudentToCourse POST
func (c students) AddStudentToCourse(accountId string, courseId string, studentName string, school int) (StudentResponse, error) {
	var r StudentResponse
	params := map[string]string{
		"courseId":    courseId,
		"identity":    "1",
		"studentUid":  accountId,
		"studentName": studentName,
	}
	err := c.client.Request("POST", "/partner/api/course.api.php?action=addCourseStudent", params, &r, school)
	return r, err
}

// RegisterAccountId POST
func (c students) RegisterAccountId(phoneNumber string, contactId string, school int) (StudentResponse, error) {
	var r StudentResponse
	classinCountry := "0065"
	params := map[string]string{
		"telephone": classinCountry + "-" + phoneNumber,
		"password":  contactId + DefaultPasswordSuffix,
	}
	err := c.client.Request("POST", "/partner/api/course.api.php?action=register", params, &r, school)
	return r, err
}

func (c students) RemoveStudentInCourse(accountUids []string, courseId string, school int) (DeleteStudentsResponse, error) {
	var r DeleteStudentsResponse
	studentUidJson := "["
	studentUidJson += strings.Join(accountUids, ",")
	studentUidJson += "]"
	params := map[string]string{
		"courseId":       courseId,
		"identity":       "1",
		"studentUidJson": studentUidJson,
	}
	paramsByte, _ := json.Marshal(params)
	fmt.Println(string(paramsByte))
	err := c.client.Request("POST", "/partner/api/course.api.php?action=delCourseStudentMultiple", params, &r, school)
	return r, err
}
