package main


import (
  "net/http"
  "log"
   groupie "groupie-tracker/handlers"
   "groupie-tracker/api"
   "groupie-tracker/controllers"
   "github.com/joho/godotenv"
   "time"
   "os"
   "os/signal"
   "syscall"
   "context"

)






func main(){
  if err := godotenv.Load(); err != nil {
    log.Println("No .env file, use environmental variable")
  }

  if err := initialize(); err != nil{
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


  go func(){
    log.Printf("Server listening on port :%s", port)
    if err := server.ListenAndServe(); err != nil && err == http.ErrServerClosed{
      log.Fatal(err)
    } 
  }()
  
  quit := make(chan os.Signal, 1)
  signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
  <-quit
  
  cxt, cancel := context.WithTimeout(context.Background(), 30*time.Second)
  defer cancel()

  if err := server.Shutdown(cxt); err !=nil{
      log.Fatal("Server forced to shutdown:", err)
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