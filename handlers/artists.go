package handlers

import (
	"log"
	"net/http"
	"groupie-tracker/api"
	"strconv"
)


type artist struct{
  Id           int    
  Name         string   
  Image        string 
  Members      []string 
  CreationDate int     
  FirstAlbum   string   
}


func ArtistsHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		w.WriteHeader(http.StatusNotFound)
		err:= templ.ExecuteTemplate(w, "notFound.html", nil)
			if err != nil{
				log.Fatal("Error accessing not found page", err)
			}
		return
	}

	id , _:= strconv.Atoi(r.URL.Query().Get("id"))


	artists := api.AllArtist()
    
    

	for _,artist := range artists{

		if artist.Id == id{
		   	
		}
	}



	

	err := templ.ExecuteTemplate(w, "artists.html", artists)
	if err != nil{
			log.Println("Error executing artist template:", err)
			return
		}
}