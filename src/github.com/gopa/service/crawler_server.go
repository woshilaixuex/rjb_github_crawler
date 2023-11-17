package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopa/database"
	"io"
	"net/http"
	"sort"
	"time"
)

// 昵称对应的真实姓名
var UserMap = map[string]string{
	"isJILE":      "周冠涛",
	"Elyxiya":     "李留芹",
	"Mr-LiuDeBao": "刘德宝",
}

var userUrl = []string{
	"https://api.github.com/repos/isJILE/mine/commits",
	"https://api.github.com/repos/Elyxiya/RLXY/commits",
	"https://api.github.com/repos/Mr-LiuDeBao/zuopin/commits",
}
var gitUrl = map[string]string{
	"周冠涛": "https://github.com/isJILE/mine",
	"李留芹": "https://github.com/Elyxiya/RLXY",
	"刘德宝": "https://github.com/Mr-LiuDeBao/zuopin",
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
	for i, url := range userUrl {
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
		if i != 0 {
			meberInformation.Name = UserMap[commits[i].Author.Login]
		} else {
			meberInformation.Name = "周冠涛"
		}
		information := new(database.Information)
		for _, commit := range commits {
			information.Data, _ = time.ParseInLocation(time.RFC3339, commit.Commit.Author.Date, location)
			information.Message = commit.Commit.Message
			information.Url = gitUrl[meberInformation.Name]
			meberInformation.Information = append(meberInformation.Information, *information)
		}
		MemberInformations = append(MemberInformations, *meberInformation)
	}
	//统一处理异常
	if len(errs) > 0 {
		return nil, fmt.Errorf("encountered errors: %w", errs)
	}
	//按时间排序（并发用）
	for a, _ := range MemberInformations {
		sort.Slice(MemberInformations[a].Information, func(i, j int) bool {
			return MemberInformations[a].Information[j].Data.Before(MemberInformations[a].Information[i].Data)
		})
	}
	return MemberInformations, nil
}
