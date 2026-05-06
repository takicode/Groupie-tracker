package handlers

import (
	"net/http"
	"html/template"
	"log"
)

var templ  = template.Must(template.New("artist.html").ParseGlob("./templates/*.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request){
    if r.URL.Path != "/"{
		w.WriteHeader(http.StatusNotFound)
		err:= templ.ExecuteTemplate(w, "notFound.html", nil)
		if err != nil{
			log.Fatal("Error accessing not found page", err)
		}
		return
	}
   
    err := templ.ExecuteTemplate(w, "base.html", nil)
	if err != nil{
		log.Fatal("Error accessing not home page", err)
	}
}