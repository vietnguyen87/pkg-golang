package classin

// Teachers client
type teachers struct {
	client
}
type Teachers interface {
	RegisterTeacher(phoneNumber string, contactId string, school int, country string) (TeacherResponse, error)
}

// Teachers constructor (from Client)
func (c client) Teachers() Teachers {
	return &teachers{
		client: c,
	}
}

type TeacherResponse struct {
	Data      int       `json:"data,omitempty"`
	ErrorInfo ErrorInfo `json:"error_info"`
}

// RegisterTeacher POST
func (c teachers) RegisterTeacher(phoneNumber string, password string, school int, country string) (TeacherResponse, error) {
	var r TeacherResponse
	if country == "" {
		country = Country
	}
	params := map[string]string{
		"telephone": country + "-" + phoneNumber,
		"password":  password,
	}
	err := c.client.request("POST", "/partner/api/course.api.php?action=register", params, &r, school)
	return r, err
}

// AddTeacher POST
func (c teachers) AddTeacher(phoneNumber string, name string, school int, country string) (TeacherResponse, error) {
	var r TeacherResponse
	if country == "" {
		country = Country
	}
	params := map[string]string{
		"teacherAccount": country + "-" + phoneNumber,
		"teacherName":    name,
	}
	err := c.client.request("POST", "/partner/api/course.api.php?action=addTeacher", params, &r, school)
	return r, err
}
