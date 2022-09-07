package classin

// Courses client
type courses struct {
	client
}

type Courses interface {
	CreateCourse(courseName string, school int) (CourseResponse, error)
	UpdateCourse(courseName string, courseId string, school int) (CourseResponse, error)
	EndCourse(courseId string, school int) (CourseResponse, error)
}

// Courses constructor (from Client)
func (c client) Courses() Courses {
	return &courses{
		client: c,
	}
}

type CourseResponse struct {
	Data      int       `json:"data,omitempty"`
	ErrorInfo ErrorInfo `json:"error_info"`
}

// CreateCourse POST
func (c courses) CreateCourse(courseName string, school int) (CourseResponse, error) {
	var r CourseResponse
	params := map[string]string{
		"courseName": courseName,
	}
	err := c.client.request("POST", "/partner/api/course.api.php?action=addCourse", params, &r, school)
	return r, err
}

func (c courses) UpdateCourse(courseName string, courseId string, school int) (CourseResponse, error) {
	var r CourseResponse
	params := map[string]string{
		"courseName": courseName,
		"courseId":   courseId,
	}
	err := c.client.request("POST", "/partner/api/course.api.php?action=editCourse", params, &r, school)
	return r, err
}

func (c courses) EndCourse(courseId string, school int) (CourseResponse, error) {
	var r CourseResponse
	params := map[string]string{
		"courseId": courseId,
	}
	err := c.client.request("POST", "/partner/api/course.api.php?action=endCourse", params, &r, school)
	return r, err
}
