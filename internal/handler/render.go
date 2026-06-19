package handler



import(
	"net/http"
	"log"
	"html/template"
)

type Render struct{
	templates *template.Template
}

func NewRender(templates *template.Template) *Render{
    return &Render{
		templates:templates,
	}
}

func (h *Render) Render404(w http.ResponseWriter) {
    w.WriteHeader(http.StatusNotFound)
    err := h.templates.ExecuteTemplate(w,"404.html",nil)

    if err != nil {
        log.Printf("render 404: %v", err)
        http.Error(w,"internal server error",http.StatusInternalServerError,
        )
    }
}

func (h *Render) Render500(w http.ResponseWriter) {
    w.WriteHeader(http.StatusInternalServerError)
    err := h.templates.ExecuteTemplate(w,"500.html",nil)
if err != nil {
		log.Printf("render 500 template failed: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
   
}


