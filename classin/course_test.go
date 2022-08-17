package classin

import (
	"errors"
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
}
