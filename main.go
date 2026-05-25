package main


import (
  "net/http"
  "log"
   groupie "groupie-tracker/handlers"
   api "groupie-tracker/api"
)


func main(){
  http.HandleFunc("/artists", groupie.ArtistsHandler)
  http.HandleFunc("/artist", groupie.ArtistHandler)
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
  // Routes
  http.HandleFunc("/", groupie.HomeHandler)
  
  err := api.LoadData()
   if err != nil{
    log.Fatal(err)
  }

 log.Println("Server listening on port 8080")
 err= http.ListenAndServe(":8080", nil)
  if err != nil{
    log.Fatal(err)
  }
}



