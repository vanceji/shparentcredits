package api

import (
	"net/url"
	"strconv"
)

type Student struct {
	Guid               string `json:"guid"`
	UserGuid           string `json:"userguid"`
	Username           string `json:"username"`
	Points             string `json:"points"`
	SerialNumber       string `json:"SerialNumber"`
	TestedCourseCount  string `json:"cnum"`
	VisitedCourseCount string `json:"snum"`
}

func GetStudents(classId int32) *Response[Student] {
	data := url.Values{
		"method":   {"classdetail"},
		"classid":  {strconv.FormatInt(int64(classId), 10)},
		"keywords": {""},
		"pagesize": {"999"},
		"pagenum":  {"1"},
	}
	return Post(PATH_REPORT, data, &Response[Student]{})
}
