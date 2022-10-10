package api

import (
	"net/url"
	"strconv"
)

type Paper struct {
	PaperId      string `json:"guid"`
	PaperName    string `json:"title"`
	IsShow       string `json:"isshow"`
	ParentModeId string `json:"modeid"`
	Status       string `json:"status"`
	SubmitTime   string `json:"submittime"` // e.g. 2022/3/25 10:07:23
	SecondModeId string `json:"titleid"`
	UserCode     string `json:"usercode"`
}

type Question struct {
	SerialNumber  string `json:"SerialNumber"`
	Answers       string `json:"answers"`
	QuestionId    string `json:"guid"`
	ParentModeId  string `json:"modeid"`
	PaperId       string `json:"paperid"`
	Questions     string `json:"questions"`
	Status        string `json:"status"`
	SubmitTime    string `json:"submittime"`
	TitleId       string `json:"titleid"`
	CorrectAnswer string `json:"trueanswer"`
	Type          string `json:"types"`
	UserAnswer    string `json:"useranswer"`
}

func GetPapers(userCode, secondModeId string) *Response[Paper] {
	data := url.Values{
		"method":   {"papersel"},
		"usercode": {userCode},
		"titleid":  {secondModeId},
	}
	return Post(PATH_TEST, data, &Response[Paper]{})
}

func GetQuestions(userCode, paperId string) *Response[Question] {
	data := url.Values{
		"method":   {"questionsel"},
		"usercode": {userCode},
		"paperid":  {paperId},
		"pagenum":  {"1"},
		"pagesize": {"999"},
	}
	return Post(PATH_TEST, data, &Response[Question]{})
}

func Answer(schoolId int, userCode, paperId, questionId, answers string) *Response[string] {
	data := url.Values{
		"method":   {"answeradd"},
		"schid":    {strconv.Itoa(schoolId)},
		"usercode": {userCode},
		"paperid":  {paperId},
		"guid":     {questionId},
		"answers":  {answers},
		"types":    {"1"},
	}
	return Post(PATH_TEST, data, &Response[string]{})
}
