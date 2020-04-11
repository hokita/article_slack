package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ArticleRepository class
type ArticleRepository struct {
}

// FindRanking func
func (a ArticleRepository) FindRanking() ([]Article, error) {
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
