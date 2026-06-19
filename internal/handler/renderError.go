package handler


func (h *Handler) render404(w http.ResponseWriter){
	 w.WriteHeader(http.StatusNotFound)

    err := h.templates.ExecuteTemplate(w,"404.html",nil)

    if err != nil {
        log.Printf("render 404: %v", err)
        http.Error(w,"internal server error",http.StatusInternalServerError)
    }
}

func (h *Handler) render500(w http.ResponseWriter){
	 w.WriteHeader(http.StatusBadRequest)
	log.Printf("render 500: %v", err)
}
