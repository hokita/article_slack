package main

import (
	"encoding/json"
	"fmt"
	"github.com/nlopes/slack"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	webhookURL = ""
	endpoint   = ""
	rankCount  = 10
)

// Ranking structure
type Ranking struct {
	Entries []Article `json:"entries"`
}

// Article structure
type Article struct {
	Title      string `json:"title"`
	URL        string `json:"url"`
	Categories string `json:"categories"`
	Image      string `json:"image"`
}

func main() {
	articles, err := getArticles()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = postArticle(articles, rankCount)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func postArticle(articles []Article, count int) error {
	for i, article := range articles[:count] {
		err := postSlack(article, i+1)
		if err != nil {
			return err
		}
	}
	return nil
}

// postSlack function
func postSlack(article Article, rank int) error {
	payload := &slack.WebhookMessage{
		Attachments: []slack.Attachment{
			{
				Blocks: []slack.Block{
					getArticleSectionBlock(article, rank),
				},
			},
		},
	}

	err := slack.PostWebhook(webhookURL, payload)
	if err != nil {
		return err
	}

	return nil
}

// getArticleSectionBlock function
func getArticleSectionBlock(article Article, rank int) slack.Block {
	rankText := "rank: " + strconv.Itoa(rank)
	titleText := "*" + article.Title + "*"
	categoryText := "category: " + article.Categories
	text := rankText + "\n" + titleText + "\n" + categoryText + "\n" + article.URL
	textBlockObject := slack.NewTextBlockObject(
		"mrkdwn",
		text,
		false,
		false,
	)
	imageElement := slack.NewImageBlockElement(article.Image, "alt image")
	section := slack.NewSectionBlock(
		textBlockObject,
		nil,
		slack.NewAccessory(imageElement),
	)
	return section
}

// get article data from api
func getArticles() ([]Article, error) {
	resp, err := http.Get(endpoint)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ranking Ranking
	if err := json.Unmarshal(body, &ranking); err != nil {
		return nil, err
	}

	return ranking.Entries, nil
}
