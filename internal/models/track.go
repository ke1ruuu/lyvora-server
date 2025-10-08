package models

type Track struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Artist    string `json:"artist"`
	Artwork   string `json:"artwork"`
	StreamURL string `json:"stream_url"`
}
