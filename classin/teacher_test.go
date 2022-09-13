package classin

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeachers(t *testing.T) {
	classinClient := NewClient(NewClientConfig(Host, SID1, Secret1, SID2, Secret2))

	result, _ := classinClient.Students().RegisterAccountId("20004507000", "5417001", 1, "0065")
	if result.ErrorInfo.ErrNo != 1 && result.ErrorInfo.ErrNo != 135 {
		assert.Error(t, errors.New(result.ErrorInfo.Error))
	}
}
