package handler

import (
  "net/http"
  "html/template"
  "log"
)

type HomeHandler struct{
	templates *template.Template
}


func NewHomeHandler(templates *template.Template) *homeHandler{
	return &homeHandler{
		templates:templates,
	}
}

func(h *HomeHandler) Home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		h.render404(w)
		return
	}
    
    err := h.templates.ExecuteTemplate(w, "base.html", nil)
   	if err != nil{
		log.Printf("execute template: %v",err,)
		http.Error(w,"internal server error",http.StatusInternalServerError)
		return
	}
}