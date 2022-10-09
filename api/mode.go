package api

import "net/url"

type Mode struct {
	Id   string `json:"guid"`
	Name string `json:"title"`
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
