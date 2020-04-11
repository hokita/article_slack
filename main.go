package main

import (
	"fmt"
	"strconv"

	"github.com/slack-go/slack"
)

const (
	webhookURL = ""
	endpoint   = ""
	rankCount  = 10
)

func main() {
	repo := ArticleRepository{}
	articles, err := repo.FindRanking()
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
		Blocks: slack.Blocks{
			BlockSet: []slack.Block{
				getArticleSectionBlock(article, rank),
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
	title := fmt.Sprintf(
		"<%s|%s>", article.URL, article.Title,
	)
	categoryText := "category: " + article.Categories
	text := rankText + "\n" + title + "\n" + categoryText
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
