package classin

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCourses(t *testing.T) {
	classinClient := NewClient(NewClientConfig(Host, SID1, Secret1, SID2, Secret2))
	courseName := "MRT IELTS"

	result, _ := classinClient.Courses().CreateCourse(courseName, 1)
	if result.ErrorInfo.ErrNo != 1 && result.ErrorInfo.ErrNo != 135 {
		assert.Error(t, errors.New(result.ErrorInfo.Error))
	}

	newCourseName := "MRS IELTS"
	courseId := result.Data
	result, _ = classinClient.Courses().UpdateCourse(newCourseName, strconv.Itoa(courseId), 1)
	if result.ErrorInfo.ErrNo != 1 && result.ErrorInfo.ErrNo != 135 {
		assert.Error(t, errors.New(result.ErrorInfo.Error))
	}

	result, _ = classinClient.Courses().EndCourse(strconv.Itoa(courseId), 1)
	if result.ErrorInfo.ErrNo != 1 && result.ErrorInfo.ErrNo != 135 {
		assert.Error(t, errors.New(result.ErrorInfo.Error))
	}
}
