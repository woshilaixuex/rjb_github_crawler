package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopa/database"
	"io"
	"net/http"
	"time"
)

// 昵称对应的真实姓名
var UserMap = map[string]string{
	"woshilaixuex": "林煜钢",
}

var userUrl = []string{
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

// 成员信息s
var MemberInformations []database.MemberInformation
var NullInformation = database.Information{
	Message: "啥也没有？这么摸!",
	Data:    time.Time{},
	Url:     "啥也没有？这么摸!",
}

func GetCommit() ([]database.MemberInformation, error) {
	method := "GET"
	client := &http.Client{}
	var errs []error
	location, _ := time.LoadLocation("Asia/Shanghai")
	//循环链接爬取所有人信息
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
		meberInformation := new(database.MemberInformation)
		meberInformation.Name = UserMap[commits[0].Author.Login]
		information := new(database.Information)
		for _, commit := range commits {
			information.Data, _ = time.ParseInLocation(time.RFC3339, commit.Commit.Author.Date, location)
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
	//按时间排序（并发用）
	//for a, _ := range MemberInformations {
	//	sort.Slice(MemberInformations[a].Information, func(i, j int) bool {
	//		return MemberInformations[a].Information[j].Data.Before(MemberInformations[a].Information[i].Data)
	//	})
	//}
	return MemberInformations, nil
}
