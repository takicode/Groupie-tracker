package handler




import (
	"html/template"
	"net/http"
	"groupie-tracker/internal/artist"
	"strconv"
	"log"
)


type ArtistHandler struct{
    service ArtistServices
	templates template.Template
}


func NewArtistHandler(templates *template.Template,service ArtistService) *Handler{
	return &Handler{
		service:service,
		templates:templates,
	}
}

type ArtistServices interface{
	Artists() []artist.FullArtistInfo 
	ArtistByID(ID int)(artist.FullArtistInfo, error)
	Search(filter artist.SearchFilter)[]artist.FullArtistInfo
}


func(h *ArtistHandler) SingleArtist(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/artist"{
		http.Error(w, "Bad request", http.StatusBadRequest)
		return 
	}

	q := r.URL.Query().Get("id")
	id, err := strconv.Atoi(q)

	if err != nil || id <= 0 {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	artist, err  := h.service.ArtistByID(id)

	data := ArtistData{
		Artist:artist,
	}

	err = h.templates.ExecuteTemplate(w, "artist.html", data)

	if err != nil{
		log.Printf("execute template: %v",err,)
		http.Error(w,"internal server error",http.StatusInternalServerError)
		return
	}
	
}