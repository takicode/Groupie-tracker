package main


import (
  "net/http"
  "log"
   groupie "groupie-tracker/handlers"
)


func main(){

  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))


  // Routes
 http.HandleFunc("/", groupie.HomeHandler)
 http.HandleFunc("/artists", groupie.ArtistHandler)

 log.Println("Server listening on port 8080")
 err:= http.ListenAndServe(":8080", nil)
  if err != nil{
    log.Fatal(err)
  }
}



