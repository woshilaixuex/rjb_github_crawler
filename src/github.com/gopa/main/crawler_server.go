package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 昵称对应的真实姓名
var UserMap = map[string]string{
	"woshilaixuex": "小明",
}

var userUrl = []string{
	"https://api.github.com/repos/woshilaixuex/woshilaixuex/commits",
	"https://api.github.com/repos/woshilaixuex/tanzu-golang-/commits",
}

// 返回体（有用的）
type Commit struct {
	SHA    string `json:"sha"`
	Commit struct {
		Author struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Date  string `json:"date"`
		} `json:"author"`
		Message string `json:"message"`
	} `json:"commit"`
	Author struct {
		Login string `json:"login"`
	} `json:"author"`
}

// 成员信息
type MemberInformation struct {
	Name        string
	Information []Information
}

// 提交信息
type Information struct {
	Message string
	Data    string
	Url     string
}

// 成员信息s
var MemberInformations []MemberInformation
var NullInformation = Information{
	Message: "啥也没有？这么摸!",
	Data:    "啥也没有？这么摸!",
	Url:     "啥也没有？这么摸!",
}

func GetCommit() ([]MemberInformation, error) {
	method := "GET"
	client := &http.Client{}
	var errs []error
	for _, url := range userUrl {
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		req.Header.Add("User-Agent", "Delyric")
		req.Header.Add("Accept", "*/*")
		req.Header.Add("Host", "api.github.com")
		req.Header.Add("Connection", "keep-alive")
		res, err := client.Do(req)
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				panic(err)
			}
		}(res.Body)
		body, err := io.ReadAll(res.Body)
		var commits []Commit
		err = json.Unmarshal(body, &commits)
		if err != nil {
			errs = append(errs, fmt.Errorf("failed to create request: %w", err))
			continue
		}
		res.Body = io.NopCloser(bytes.NewBuffer(body))
		err = json.Unmarshal(body, &commits)
		if err != nil {
			errs = append(errs, fmt.Errorf("failed to create request: %w", err))
			continue
		}
		meberInformation := new(MemberInformation)
		meberInformation.Name = UserMap[commits[0].Author.Login]
		//拷打用
		information := new(Information)
		for _, commit := range commits {
			information.Data = commit.Commit.Author.Date
			information.Message = commit.Commit.Message
			information.Url = url
			meberInformation.Information = append(meberInformation.Information, *information)
		}
		MemberInformations = append(MemberInformations, *meberInformation)
	}
	//统一处理异常
	if len(errs) > 0 {
		return nil, fmt.Errorf("encountered errors: %w", errs)
	}
	return MemberInformations, nil
}
