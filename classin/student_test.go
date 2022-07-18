package classin

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStudents(t *testing.T) {
	classinClient := NewClient(NewClientConfig(Host, SID1, Secret1, SID2, Secret2))
	//now := time.Now()
	//phoneNumber := "20004507001"
	//// School 1: 41332568
	//// School 2: 52067164
	//contactId := "34151"
	//_ = contactId
	//_ = phoneNumber
	//_ = classinClient
	result, _ := classinClient.Students().RegisterAccountId("20004507000", "5417001", 1)
	if result.ErrorInfo.ErrNo != 1 && result.ErrorInfo.ErrNo != 135 {
		assert.Error(t, errors.New(result.ErrorInfo.Error))
	}

	resultAddStudentToCourse, _ := classinClient.Students().AddStudentToCourse("41332568", "205340865", "Test student", 1)
	if resultAddStudentToCourse.ErrorInfo.ErrNo != 1 &&
		resultAddStudentToCourse.ErrorInfo.ErrNo != 163 &&
		resultAddStudentToCourse.ErrorInfo.ErrNo != 164 &&
		resultAddStudentToCourse.ErrorInfo.ErrNo != 133 {
		assert.Error(t, errors.New(resultAddStudentToCourse.ErrorInfo.Error))
	}
	//resultRemoveStudentOut, _ := classinClient.Students().RemoveStudentInCourse([]string{"41332568"}, "205340865", 1)
	//if resultRemoveStudentOut.ErrorInfo.ErrNo != 1 {
	//	assert.Error(t, errors.New(resultRemoveStudentOut.ErrorInfo.Error))
	//}
}

func TestClassinAccount(t *testing.T) {
	//+65 20000511708
	// pass: classin123
	classinClient := NewClient(NewClientConfig(Host, SID1, Secret1, SID2, Secret2))
	//result, _ := classinClient.Students().RegisterAccountId("20000511708", "20902101", 1)
	//fmt.Println(result)
	result, _ := classinClient.Students().ChangeStudentName("51618090", "Test 0123", 2)
	fmt.Println(result)
}
