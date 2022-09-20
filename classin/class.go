package classin

// Classes client
type classes struct {
	client
}

type Classes interface {
	CreateSingleClass(courseId string, className string, beginTime string, endTime string, teacherUid string, school int) (ClassResponse, error)
	UpdateClass(classId string, courseId string, className string, beginTime string, endTime string, teacherUid string, school int) (ClassResponse, error)
}

// Classes constructor (from Client)
func (c client) Classes() Classes {
	return &classes{
		client: c,
	}
}

type ClassResponse struct {
	Data      int       `json:"data,omitempty"`
	ErrorInfo ErrorInfo `json:"error_info"`
}

// CreateSingleClass POST
func (c classes) CreateSingleClass(courseId string, className string, beginTime string, endTime string, teacherUid string, school int) (ClassResponse, error) {
	var r ClassResponse
	params := map[string]string{
		"courseId":   courseId,
		"className":  className,
		"beginTime":  beginTime,
		"endTime":    endTime,
		"teacherUid": teacherUid,
	}
	err := c.client.request("POST", "/partner/api/course.api.php?action=addCourseClass", params, &r, school)
	return r, err
}

// UpdateClass POST
func (c classes) UpdateClass(classId string, courseId string, className string, beginTime string, endTime string, teacherUid string, school int) (ClassResponse, error) {
	var r ClassResponse
	params := map[string]string{
		"classId":    classId,
		"courseId":   courseId,
		"className":  className,
		"beginTime":  beginTime,
		"endTime":    endTime,
		"teacherUid": teacherUid,
	}
	err := c.client.request("POST", "/partner/api/course.api.php?action=editCourseClass", params, &r, school)
	return r, err
}
