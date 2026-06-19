package handler

import (
	"net/http"
	"log"
	"html/template"
	"groupie-tracker/internal/artist"
)

type Handler struct{
	templates *template.Template
	service ArtistService
	render *Render
}

func NewHomeHandler(templates *template.Template,service ArtistService,  render *Render) *Handler{
	return &Handler{
		templates:templates,
		service:service,
		render:render,
	}
}

type ArtistService interface{
	Artists() []artist.FullArtistInfo 
	ArtistByID(ID int)(artist.FullArtistInfo, error)
	Search(filter artist.SearchFilter)[]artist.FullArtistInfo
}


func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	
	if r.URL.Path != "/" {
		h.render.Render404(w)
		return
	}

	err := h.templates.ExecuteTemplate(w, "base.html", nil)
	if err != nil {

		log.Printf("Error executing base template: %v", err)
		h.render.Render500(w)
		return
	}
}


