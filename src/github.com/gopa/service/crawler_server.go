package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopa/database"
	"io"
	"net/http"
	"net/url"
	"sort"
	"time"
)

// 昵称对应的真实姓名
var UserMap = map[string]string{
	"jile":            "周冠涛",
	"AiLY":            "李留芹",
	"‘Mr.Liu":         "刘德宝",
	"fei39":           "王鹏飞",
	"13759723989edou": "邓晴",
	"overfloatGame":   "骆建威",
	"roger-yuan":      "罗杰元",
	"manman":          "黄嫚嫚",
}

var userUrl = []string{
	"https://api.github.com/repos/isJILE/mine/commits",
	"https://api.github.com/repos/Elyxiya/RLXY/commits",
	"https://api.github.com/repos/Mr-LiuDeBao/zuopin/commits",
	"https://api.github.com/repos/d2bz/project/commits",
	"https://api.github.com/repos/CSDAndroid/dem/commits",
	"https://api.github.com/repos/CSDAndroid/musicPlayer/commits",
	"https://api.github.com/repos/CSDAndroid/TestingGround/commits",
	"https://api.github.com/repos/CSDAndroid/note/commits",
}
var gitUrl = map[string]string{
	"周冠涛": "https://github.com/isJILE/mine",
	"李留芹": "https://github.com/Elyxiya/RLXY",
	"刘德宝": "https://github.com/Mr-LiuDeBao/zuopin",
	"王鹏飞": "https://github.com/d2bz/project",
	"邓晴":  "https://github.com/dem",
	"骆建威": "https://github.com/CSDAndroid/musicPlayer",
	"罗杰元": "https://github.com/CSDAndroid/TestingGround",
	"黄嫚嫚": "https://github.com/CSDAndroid/note",
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
	proxyURL, err := url.Parse("http://127.0.0.1:7890")
	if err != nil {
		panic(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}
	var errs []error
	location, _ := time.LoadLocation("Asia/Shanghai")
	//循环链接爬取所有人信息
	for _, url := range userUrl {
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:120.0) Gecko/20100101 Firefox/120.0")
		req.Header.Add("Host", "api.github.com")
		req.Header.Add("Accept", "application/vnd.github.v3+json")
		req.Header.Add("Authorization", "Bearer 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")
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				panic(err)
			}
		}(res.Body)
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
		meberInformation.Name = UserMap[commits[0].Commit.Author.Name]
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
			return MemberInformations[a].Information[i].Data.Before(MemberInformations[a].Information[j].Data)
		})
	}
	return MemberInformations, nil
}
