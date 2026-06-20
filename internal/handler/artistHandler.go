package handler


import (
	"html/template"
	"groupie-tracker/internal/artist"
	"net/http"
	"strconv"
	"log"
	"errors"
)


func NewArtistHandler(templates *template.Template,service ArtistService) *Handler{
	return &Handler{
		service:service,
		templates:templates,
	}
}


func(h *Handler)SingleArtist(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query().Get("id")
	id, err := strconv.Atoi(q)

	if err != nil || id <= 0 {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	singleArtist, err  := h.service.ArtistByID(id)

	if err != nil {
    render := NewRender(h.templates)

    if errors.Is(err, artist.ErrArtistNotFound) {
        render.Render404(w)
        return
    }

    log.Printf("artist by id: %v", err)
    render.Render500(w)
    return
}

	data := ArtistPageData{
		singleArtist,
	}

	err = h.templates.ExecuteTemplate(w, "artist.html", data)

	if err != nil{
		log.Printf("execute template: %v",err,)
		// http.Error(w,"internal server error",http.StatusInternalServerError)
		return
	}
	
}