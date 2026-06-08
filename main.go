package main


import (
  "net/http"
  "log"
   groupie "groupie-tracker/handlers"
   "groupie-tracker/api"
   "groupie-tracker/controllers"
   "github.com/joho/godotenv"
   "time"
)


func main(){
  if err := godotenv.Load(); err != nil {
    log.Println("No .env file, use environmental variable")
  }

  if err := initialize() err != nil{
    log.Fatal(err)
  } 

  port := os.Getenv("PORT")
  if port == ""{
    port = "8080"
  }

  mux := http.NewServeMux()
  registerRoutes(mux)
 
  server := &http.Server{
     Addr :  ":" + port,
     Handler : mux,
     ReadTimeout : 10 * time.Second,
     WriteTimeout : 10 * time.Second,
     IdleTimeout : 60 * time.Second,
  }

  log.Printf("Server listening on port :%s", port)

  err := server.ListenAndServe(":8080", mux); err != nil{
    log.Fatal(err)
  }
}


func initialize()error{
  
  log.Println("Loading artists data...")
  
  if err:= api.LoadData(); err !=nil{
    return err
  }
  
  log.Println("Loading coordinates cache...")

  locations := controllers.GetLoc()
   
  if err := controllers.LoadOrBuildCache(locations); err != nil {
      return err
  }

  return nil  
}


func registerRoutes(mux *http.ServeMux){
  mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
  mux.HandleFunc("/artists", groupie.AllartistsHandler)
  mux.HandleFunc("/artist", groupie.ArtistHandler)
  mux.HandleFunc("/", groupie.HomeHandler)
}