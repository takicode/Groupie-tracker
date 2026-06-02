package main


import (
  "net/http"
  "log"
   groupie "groupie-tracker/handlers"
   api "groupie-tracker/api"
)


func main(){
  err := api.LoadData()
   if err != nil{
    log.Fatal(err)
  }
  http.HandleFunc("/artists", groupie.AllartistsHandler)
  http.HandleFunc("/artist", groupie.ArtistHandler)
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
  
  // Routes
  http.HandleFunc("/", groupie.HomeHandler)
  

 log.Println("Server listening on port 8080")
 err= http.ListenAndServe(":8080", nil)
  if err != nil{
    log.Fatal(err)
  }
}



