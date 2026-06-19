package handler

import(
	"groupie-tracker/internal/artist"
)

type HomePageData struct{
	Artists    []artist.FullArtistInfo
    Search      string
    CurrentPage int
    TotalPages int
}

type ArtistData struct{
	Artist artist.FullArtistInfo
}