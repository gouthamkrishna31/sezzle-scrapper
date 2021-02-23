package main

import "time"

type ImdbDetails struct {
	Context       string      `json:"@context"`
	Type          string      `json:"@type"`
	URL           string      `json:"url"`
	Name          string      `json:"name"`
	Image         string      `json:"image"`
	Genre         interface{} `json:"genre"`
	ContentRating string      `json:"contentRating"`
	Actor         []struct {
		Type string `json:"@type"`
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"actor"`
	Director interface{} `json:"director"`
	Creator  []struct {
		Type string `json:"@type"`
		URL  string `json:"url"`
		Name string `json:"name,omitempty"`
	} `json:"creator"`
	Description     string `json:"description"`
	DatePublished   string `json:"datePublished"`
	Keywords        string `json:"keywords"`
	AggregateRating struct {
		Type        string `json:"@type"`
		RatingCount int    `json:"ratingCount"`
		BestRating  string `json:"bestRating"`
		WorstRating string `json:"worstRating"`
		RatingValue string `json:"ratingValue"`
	} `json:"aggregateRating"`
	Review struct {
		Type         string `json:"@type"`
		ItemReviewed struct {
			Type string `json:"@type"`
			URL  string `json:"url"`
		} `json:"itemReviewed"`
		Author struct {
			Type string `json:"@type"`
			Name string `json:"name"`
		} `json:"author"`
		DateCreated  string `json:"dateCreated"`
		InLanguage   string `json:"inLanguage"`
		Name         string `json:"name"`
		ReviewBody   string `json:"reviewBody"`
		ReviewRating struct {
			Type        string `json:"@type"`
			WorstRating string `json:"worstRating"`
			BestRating  string `json:"bestRating"`
			RatingValue string `json:"ratingValue"`
		} `json:"reviewRating"`
	} `json:"review"`
	Duration string `json:"duration"`
	Trailer  struct {
		Type      string `json:"@type"`
		Name      string `json:"name"`
		EmbedURL  string `json:"embedUrl"`
		Thumbnail struct {
			Type       string `json:"@type"`
			ContentURL string `json:"contentUrl"`
		} `json:"thumbnail"`
		ThumbnailURL string    `json:"thumbnailUrl"`
		Description  string    `json:"description"`
		UploadDate   time.Time `json:"uploadDate"`
	} `json:"trailer"`
}

type Data struct {
	Title            string      `json:"title"`
	MovieReleaseYear int         `json:"movie_release_year"`
	ImdbRating       float64     `json:"imdb_rating"`
	Summary          string      `json:"summary"`
	Duration         string      `json:"duration"`
	Genre            interface{} `json:"genre"`
}
