package services

import (
	"encoding/json"
	"net/http"
)

type Artwork struct {
    X150  string `json:"150x150"`
    X480  string `json:"480x480"`
    X1000 string `json:"1000x1000"`
}

type User struct {
    Name string `json:"name"`
}

type Track struct {
    ID      string  `json:"id"`
    Title   string  `json:"title"`
    Artwork Artwork `json:"artwork"`
    User    User    `json:"user"`
}


func FetchTracks() ([]map[string]string, error) {
    resp, err := http.Get("https://discoveryprovider.audius.co/v1/tracks/trending?app_name=LyvoraDemo")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var data struct {
        Data []Track `json:"data"`
    }

    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return nil, err
    }

    var formatted []map[string]string
    for _, t := range data.Data {
        img := t.Artwork.X1000
        if img == "" {
            img = t.Artwork.X480
        }
        if img == "" {
            img = t.Artwork.X150
        }

        formatted = append(formatted, map[string]string{
            "id":      t.ID,
            "title":   t.Title,
            "artist":  t.User.Name,
            "artwork": img,
        })
    }

    return formatted, nil
}


func GetStreamURL(id string) string {
	return "https://discoveryprovider.audius.co/v1/tracks/" + id + "/stream?app_name=lyvora"
}
