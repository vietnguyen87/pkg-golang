package classin

func AllMessages() map[int]string {
	return map[int]string{
		1:   "Successful execution",
		100: "Incomplete or incorrect parameters",
		102: "No permission(Security verification failed.)",
		104: "Operation failure(Unknown error)",
		114: "Server exception",
		134: "Illegal mobile phone number",
		144: "There is no such course in your institution",
		147: "There is no information about the course",
		153: "The course is expired and cannot be edited",
		155: "Array data cannot be null",
		162: "There is no such students in the course",
		369: "Do not support to operate this type of courses and lessons(public course)",
		400: "The requested data is illegal",
	}
}

func (c *client) GetMessage(code int) string {
	return AllMessages()[code]
}
