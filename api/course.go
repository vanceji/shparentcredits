package api

import (
	"net/url"
	"strconv"
)

type Course struct {
	Id                string `json:"guid"`
	Name              string `json:"title"`
	VideoLink         string `json:"courseurl"`
	ImgLink           string `json:"imgurl"`
	TopLevelModeId    string `json:"modeid"`
	SecondLevelModeId string `json:"titleid"`
	IsVisited         string `json:"isfinish"`
	IsFree            string `json:"isfree"`
	Types             string `json:"types"`
	Status            string `json:"status"`
}

func GetCourse(parentModeId, childModeId, userCode string) *Response[Course] {
	data := url.Values{
		"method":   {"coursesel"},
		"modeid":   {parentModeId},
		"titleid":  {childModeId},
		"usercode": {userCode},
		"guid":     {""},
	}
	return Post(PATH_COURSE, data, &Response[Course]{})
}

func VisitCourse(userCode string, videoId string) bool {
	data := url.Values{
		"method":   {"coursecredit"},
		"usercode": {userCode},
		"videoid":  {videoId},
	}
	response := Post(PATH_POINTS, data, &Response[string]{})
	return response.Message != "5"
}

func HowManyCoursesToVisitToday(userCode string, years int, months int) int {

	response := GetPointsOfMonth(userCode, years, months)
	for _, pm := range response.Data {
		if pm.IsToday() {
			pointsToday, _ := strconv.Atoi(pm.Points)
			return MAX_COURSE_TO_VISIT_PER_DAY - (pointsToday / POINTS_PER_COURSE)
		}
	}
	return MAX_COURSE_TO_VISIT_PER_DAY
}
