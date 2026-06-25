package handler


import (
	"html/template"
	"groupie-tracker/internal/artist"
	"net/http"
	"encoding/json"
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
  	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	q := r.URL.Query().Get("id")
	id, err := strconv.Atoi(q)

	if err != nil || id <= 0 {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	artistInfo, err  := h.service.ArtistByID(id)

		if err != nil {
			render := NewRender(h.templates)

			if errors.Is(err, artist.ErrArtistNotFound) {
				render.Render404(w)
				return
			}

			log.Printf("Error fetching artist by id %d: %v",id, err)
			render.Render500(w)
			return
		}

    coordJSON, err := json.Marshal(artistInfo.Coordinates)
	
		if err != nil {
			log.Printf("Error encoding coordinates for artist %d: %v", id, err)
			http.Error(w, "failed to encode coordinates", http.StatusInternalServerError)
			return
		}

	data := ArtistPageData{
		artistInfo.Artist,
		template.JS(coordJSON),
	}

	err = h.templates.ExecuteTemplate(w, "artist.html", data)

	if err != nil{
		log.Printf("execute template: %v",err,)
		return
	}
	
}

