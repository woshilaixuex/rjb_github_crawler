package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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
type Member struct {
	Name    string
	Message string
	Data    string
}

func GetCommit() {
	url := "https://api.github.com/repos/woshilaixuex/Tomatoos/commits"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "Delyric")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "api.github.com")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var commits []Commit
	err = json.Unmarshal(body, &commits)
	if err != nil {
		fmt.Println(err)
		return
	}
	res.Body = io.NopCloser(bytes.NewBuffer(body))

	err = json.Unmarshal(body, &commits)
	if err != nil {
		panic(err)
	}
	for _, commit := range commits {
		fmt.Println(commit.Commit.Author.Name)
		fmt.Println(commit.Commit.Author.Date)
		fmt.Println(commit.Commit.Message)
		fmt.Println(commit.Author.Login)
	}
}
