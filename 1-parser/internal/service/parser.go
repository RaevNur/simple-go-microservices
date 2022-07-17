package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/RaevNur/simple-go-microservices-parser/internal/models"
)

type PostJSON struct {
	Meta struct {
		Pagination struct {
			Total int `json:"total"`
			Pages int `json:"pages"`
			Page  int `json:"page"`
			Limit int `json:"limit"`
			Links struct {
				Previous interface{} `json:"previous"`
				Current  string      `json:"current"`
				Next     string      `json:"next"`
			} `json:"links"`
		} `json:"pagination"`
	} `json:"meta"`
	Data []struct {
		ID     int    `json:"id"`
		UserID int    `json:"user_id"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	} `json:"data"`
}

func (s *Service) startParse() {
	log.Println("Parsing started")
	urlParse := "https://gorest.co.in/public/v1/posts?page=%d"
	urls := make([]string, 0, 10)

	for i := 1; i <= 10; i++ {
		urls = append(urls, fmt.Sprintf(urlParse, i))
	}

	chIsFinished := make(chan bool)

	for _, url := range urls {
		go s.fetchUrl(url, chIsFinished)
		time.Sleep(time.Second)
	}

	for i := 0; i < 10; {
		<-chIsFinished
		i++
	}

	s.mu.Lock()
	s.av = true
	s.mu.Unlock()
	log.Println("Parsing is over")
}

func (s *Service) fetchUrl(url string, chIsFinished chan bool) {
	defer func() {
		chIsFinished <- true
	}()

	response, err := http.Get(url)
	if err != nil || response.StatusCode != 200 {
		log.Println(err)
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var result PostJSON
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println(err)
		return
	}

	for _, res := range result.Data {
		post := models.Post{
			UserId: int64(res.UserID),
			Title:  res.Title,
			Body:   res.Body,
		}

		err := s.repo.Add(context.Background(), &post)
		if err != nil {
			log.Println(err)
		}
	}
}
