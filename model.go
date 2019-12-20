package main

// Data stores JSON data returned by the HN API
type Data struct {
	IsDeleted          bool      `json:"deleted"`
	ID                 float64   `json:"id"`
	CreatedAtUnix      float64   `json:"time"`
	Score              float64   `json:"score"`
	CommentCount       float64   `json:"descendants"`
	CommentIDs         []float64 `json:"kids"`
	RelatedPollOptions []float64 `json:"parts"`
	Type               string    `json:"type"`
	By                 string    `json:"by"`
	Text               string    `json:"text"`
	IsDead             string    `json:"dead"`
	Parent             string    `json:"parent"`
	Poll               string    `json:"poll"`
	URL                string    `json:"url"`
	Title              string    `json:"title"`
}

type User struct {
	Id        string    `json:"id"`
	Delay     string    `json:"delay"`
	Created   string    `json:"created"`
	Karma     float64   `json:"karma"`
	About     string    `json:"about"`
	Submitted []float64 `json:"submitted"`
}
