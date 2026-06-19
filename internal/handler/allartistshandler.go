package handler

import (
	"net/http"
	"groupie-tracker/internal/artist"
	"log"
	"strings"
	"html/template"
)


func NewHandler(templates *template.Template,service ArtistService) *Handler{
	return &Handler{
		service:service,
		templates:templates,
	}
}


func (h *Handler) AllArtist( w http.ResponseWriter, r *http.Request){
    // search
	query := r.URL.Query().Get("search")

	// page number
    pageString := r.URL.Query().Get("page")
	pageNo, err:= strconv.Atoi(pageString)
	if err != nil || pageNo < 1 {
	  pageNo = 1
	}


	
	var artists []artist.SearchResult

   
	filter := artist.SearchFilter{
		Query:query,
		Page:pageNo,
	}
	artists = h.service.Search(filter)

   

	data:= HomePageData{
		artists,
    	Search:query,
	}


	err := h.templates.ExecuteTemplate(w, "artists.html", data)

	if err != nil{
		log.Printf("execute template: %v",err,)
		return
	}

}


