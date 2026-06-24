package handler

import (
	"net/http"
	"groupie-tracker/internal/artist"
	"log"
	"strconv"
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
    decade := r.URL.Query().Get("decade")
    location := r.URL.Query().Get("location")
	members:= r.URL.Query().Get("members")
	sortFilter := r.URL.Query().Get("sort_by")

	// page number
    pageString := r.URL.Query().Get("page")
	pageNo, err:= strconv.Atoi(pageString)
	if err != nil || pageNo < 1 {
	  pageNo = 1
	}


	
	var artists artist.SearchResult

   
	filter := artist.SearchFilter{
		Query:query,
		Page:pageNo,
		Decade:decade,
		Location:location,
		Members:members,
		SortBy:sortFilter,
	}
	artists = h.service.Search(filter)

   

	data:= HomePageData{
		SearchResult:artists,
    	Search:filter.Query,
		Decade:filter.Decade,
		Location:filter.Location,
		Members:filter.Members,
		SortBy:filter.SortBy,
	}


	err = h.templates.ExecuteTemplate(w, "artists.html", data)

	if err != nil{
		log.Printf("execute template: %v",err,)
		return
	}

}


