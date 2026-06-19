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

type ArtistPageData struct{
	Artist artist.FullArtistInfo
}