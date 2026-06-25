package handler

import(
	"groupie-tracker/internal/artist"
	"html/template"
)

type HomePageData struct{
	artist.SearchResult
    Search string
	Decade string
	Location string
	Members string
	SortBy string
}

type ArtistPageData struct{
	artist.ArtistDetails
	Coordinates template.JS
}