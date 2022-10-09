package api

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type ClassRanking struct {
	ClassId     string `json:"id"`
	ClassName   string `json:"classname"`
	Ranking     string `json:"SerialNumber"`
	SchoolId    string `json:"schid"`
	StudentsNum string `json:"allnum"`
	AvgPoints   string `json:"avgpoints"`
	SumPoints   string `json:"allpoints"`
}

func (r *ClassRanking) String() string {
	return r.Ranking + "," + r.ClassName + "," + r.AvgPoints
}

type MonthlyPoints struct {
	SerialNumber string `json:"SerialNumber"`
	Points       string `json:"num"`
	UpdatedAt    string `json:"submittime"`
}

func (mp *MonthlyPoints) IsToday() bool {
	now := time.Now().Format("2006-01-02")
	return now == mp.UpdatedAt
}

func GetClassRanking(schoolId int) *Response[ClassRanking] {
	data := url.Values{
		"method":   {"classlist"},
		"schid":    {strconv.Itoa(schoolId)},
		"types":    {"2"},
		"pagesize": {"100"},
		"pagenum":  {"1"},
		"gradeid":  {""},
		"classid":  {""},
	}
	return Post(PATH_REPORT, data, &Response[ClassRanking]{})
}

func PrintClassRanking(schoolId int) {
	result := GetClassRanking(schoolId)
	for _, ranking := range result.Data {
		fmt.Println(ranking.String())
	}
}

func GetPointsOfMonth(userCode string, years int, months int) *Response[MonthlyPoints] {
	data := url.Values{
		"method":   {"pointmonth"},
		"usercode": {userCode},
		"years":    {strconv.Itoa(years)},
		"months":   {strconv.Itoa(months)},
		"pagenum":  {"1"},
		"pagesize": {"31"},
	}
	return Post(PATH_POINTS, data, &Response[MonthlyPoints]{})
}
