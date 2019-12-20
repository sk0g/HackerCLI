package main

// Post stores the JSON data describing a post
type Post struct {
	ID          float64   `json:"id"`
	Deleted     bool      `json:"deleted"`
	Type        string    `json:"type"`
	By          string    `json:"by"`
	Time        float64   `json:"time"`
	Text        string    `json:"text"`
	Dead        string    `json:"dead"`
	Parent      string    `json:"parent"`
	Poll        string    `json:"poll"`
	Kids        []float64 `json:"kids"`
	Url         string    `json:"url"`
	Score       float64   `json:"score"`
	Title       string    `json:"title"`
	Parts       []float64 `json:"parts"`
	Descendants float64   `json:"descendants"`
}
