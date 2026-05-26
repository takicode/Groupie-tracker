package handlers

import (
	"log"
	"net/http"
	"groupie-tracker/api"
	"strconv"
)


func ArtistHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		w.WriteHeader(http.StatusNotFound)
		err:= templ.ExecuteTemplate(w, "notFound.html", nil)
			if err != nil{
				log.Println("Error accessing not found page", err)
			}
		return
	}

	id , err:= strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
	http.Error(w, "Invalid artist ID", http.StatusBadRequest)
	return
}

	artists := api.AllArtist()
	var singleArtist  api.FullArtistInfo
    found:= false

	for _,artist := range artists{
		if artist.Id == id{
		   	singleArtist = artist
			found = true
			break
		}
	}

	if !found {
	w.WriteHeader(http.StatusNotFound)

	err := templ.ExecuteTemplate(w, "notFound.html", nil)
	if err != nil {
		log.Println(err)
	}
	return
}


	err = templ.ExecuteTemplate(w, "artist.html", singleArtist)
	if err != nil{
		log.Println("Error executing artist template:", err)
		return
	}
}