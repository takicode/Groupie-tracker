package handler

import (
	"net/http"
	"groupie-tracker/internal/artist"
	"log"
	"strings"
	"html/template"
)

// type Handler struct{
// 	service ArtistService
// 	templates *template.Template
// }

func NewHandler(templates *template.Template,service ArtistService) *Handler{
	return &Handler{
		service:service,
		templates:templates,
	}
}

// type ArtistServices interface{
// 	Artists() []artist.FullArtistInfo 
// 	ArtistByID(ID int)(artist.FullArtistInfo, error)
// 	Search(filter artist.SearchFilter)[]artist.FullArtistInfo
// }


func (h *Handler) AllArtist( w http.ResponseWriter, r *http.Request){

	q := r.URL.Query().Get("search")
	query := strings.TrimSpace(strings.ToLower(q))
	
	var artists []artist.FullArtistInfo

   if query ==""{
       artists = h.service.Artists()
   }else{
		filter := artist.SearchFilter{
		Query:query,
		}
		artists = h.service.Search(filter)

   }

	data:= HomePageData{
		Artists:artists,
    	Search:query,
	}


	err := h.templates.ExecuteTemplate(w, "artists.html", data)

	if err != nil{
		log.Printf("execute template: %v",err,)
		// http.Error(w,"internal server error",http.StatusInternalServerError)
		return
	}

}


