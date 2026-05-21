package handlers

import (
	"log"
	"net/http"
)




func ArtistHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		w.WriteHeader(http.StatusNotFound)
		err:= templ.ExecuteTemplate(w, "notFound.html", nil)
		if err != nil{
			log.Fatal("Error accessing not found page", err)
		}
		return
	}

	displayInfo := Pagination(w, r)

	
	err := templ.ExecuteTemplate(w, "artists.html", displayInfo)
	if err != nil{
			log.Println("Error executing artist template:", err)
			return
		}
}