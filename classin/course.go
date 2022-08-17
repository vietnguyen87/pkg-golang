package classin

// Courses client
type courses struct {
	client
}

type Courses interface {
	CreateCourse(courseName string, school int) (CourseResponse, error)
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

// AddStudentToSchool POST
func (c courses) CreateCourse(courseName string, school int) (CourseResponse, error) {
	var r CourseResponse
	params := map[string]string{
		"courseName": courseName,
	}
	err := c.client.request("POST", "/partner/api/course.api.php?action=addCourse", params, &r, school)
	return r, err
}
