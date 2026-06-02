package handlers

import (
	"log"
	"net/http"
	"html/template"
	"groupie-tracker/api"
	"groupie-tracker/controllers"
	"encoding/json"
	"strconv"
	"strings"
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
	var singleArtist api.FullArtistInfo
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

	var locations []api.GeoLocation

    for loc, date := range singleArtist.DateLocations{
		location:= strings.ReplaceAll(loc,"_"," ")
		location = strings.ReplaceAll(location,"-"," ")


		locationResp,err := controllers.GeoLocation(location)
			if err != nil {
			log.Println("unable to collect location:", err)
			continue
		}
		if len(locationResp) == 0{
			 log.Println("location not found")
			continue
		}
         
		geo, err := controllers.GeoJson(locationResp, date)
			if err != nil {
				continue
			}

    	locations = append(locations, geo)

	}
    
	locJson, err := json.Marshal(locations)

	if err != nil {
    log.Println("failed to marshal locations:", err)
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
}

	data := api.ArtistPage{
		Artist:singleArtist,
        LocationJSON :template.JS(string(locJson)),
	}


	err = templ.ExecuteTemplate(w, "artist.html", data)

	if err != nil{
		log.Println("Error executing artist template:", err)
		return
	}
}