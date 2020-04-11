package main

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
