package api

import "net/url"

type Mode struct {
	Id   string `json:"guid"`
	Name string `json:"title"`
}

type RequiredCourseMode struct {
	ModeId       string `json:"guid"`
	ModeName     string `json:"title"`
	ParentModeId string `json:"modeid"`
}

type ExamMode struct {
	ModeId       string `json:"guid"`
	ModeName     string `json:"title"`
	ParentModeId string `json:"modeid"`
}

func GetTopLevelMode() *Response[Mode] {
	data := url.Values{
		"method": {"modesel"},
		"modeid": {""},
		"uid":    {"1"},
	}
	return Post(PATH_MODE, data, &Response[Mode]{})
}

func GetSecondLevelMode(parentModeId string) *Response[Mode] {
	data := url.Values{
		"method": {"titlesel"},
		"modeid": {parentModeId},
	}
	return Post(PATH_MODE, data, &Response[Mode]{})
}

func GetSecondLevelModeForRequiredCourses(userCode string) *Response[RequiredCourseMode] {
	data := url.Values{
		"method":   {"titlesel_bx"},
		"idenguid": {userCode},
	}
	return Post(PATH_MODE, data, &Response[RequiredCourseMode]{})
}

func GetSecondLevelModeForExam(userCode string) *Response[ExamMode] {
	data := url.Values{
		"method":   {"titlesel"},
		"idenguid": {userCode},
	}
	return Post(PATH_TEST, data, &Response[ExamMode]{})
}
