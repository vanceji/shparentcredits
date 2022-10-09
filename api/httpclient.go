package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var Domain string
var Token string

type Response[T string | Mode | Course | ClassRanking | Student | MonthlyPoints] struct {
	Code    string `json:"retCode"`
	Message string `json:"retMsg"`
	Data    []T    `json:"retObj"`
}

func Post[T string | Mode | Course | ClassRanking | Student | MonthlyPoints](path string, data url.Values, response *Response[T]) *Response[T] {
	api := Domain + path
	result, _ := NewPost(api, data, Token)
	err := json.Unmarshal(result, response)
	if err != nil {
		log.Fatalf("Unmarshal response error, url: %s, response: %v, error: %v", api, string(result), err)
	}
	return response
}

func NewPost(url string, data url.Values, cookie string) ([]byte, error) {
	body := strings.NewReader(data.Encode())
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatalf("request %v with data %v error %v", Domain, data, err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	req.Header.Set("Cookie", cookie)
	clt := http.Client{}
	resp, err := clt.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)
	if err != nil {
		log.Fatalf("request %v with data %v error %v", Domain, data, err)
		return nil, err
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("request %v with data %v response %v error %v", Domain, data, content, err)
		return nil, err
	}

	return content, nil
}
