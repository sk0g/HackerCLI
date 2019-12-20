package main

// Data stores JSON comment/ask/show/jobs/poll/parts data
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

// User stores JSON user data
type User struct {
	Username             string    `json:"id"`
	CreationDelayMinutes float64   `json:"delay"`
	CreatedAtUnix        float64   `json:"created"`
	Karma                float64   `json:"karma"`
	About                string    `json:"about"`
	Submitted            []float64 `json:"submitted"`
}
