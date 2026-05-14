package handlers

import (
	"log"
	"net/http"
	 "groupie-tracker/api"
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

	artists := api.AllArtist()
    // log.Println(artists)
	err := templ.ExecuteTemplate(w, "artists.html", artists)
	if err != nil{
			log.Println("Error executing artist template:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
}