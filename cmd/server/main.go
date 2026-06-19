package main

import (
   "net/http"
   "html/template"
   "log"
   "groupie-tracker/internal/handler"
   "groupie-tracker/internal/artist"
   "groupie-tracker/internal/config"
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

  cfg := config.Load()
  port:= cfg.Port
    
  artistsHandler,singleArtistHandler,homeHandler := initialize(cfg)
      
  mux := http.NewServeMux()
  registerRoutes(mux,artistsHandler,singleArtistHandler,homeHandler )
 
  server := &http.Server{
     Addr :  ":" + port,
     Handler : mux,
     ReadTimeout : 10 * time.Second,
     WriteTimeout : 10 * time.Second,
     IdleTimeout : 60 * time.Second,
  }


  go func(){
    log.Printf("Server listening on port :%s", port)
    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed{
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


func initialize(cfg *config.Config)(*handler.Handler, *handler.Handler, *handler.Handler){

  client := artist.NewClient(cfg.BaseURL)
  store := artist.NewStore(client)
  service := artist.NewService(store)
  ctx := context.Background()
   
  if err := store.Load(ctx); err !=nil{
    log.Fatal(err)
  }

  temp := template.Must(template.ParseGlob("templates/*.html"))
  render := handler.NewRender(temp)
  artistsHandler := handler.NewHandler(temp, service)
	singleArtistHandler := handler.NewArtistHandler(temp, service) 
	homeHandler := handler.NewHomeHandler(temp,service,render)

  return artistsHandler,singleArtistHandler,homeHandler
}



func registerRoutes(mux *http.ServeMux,artists *handler.Handler,artist *handler.Handler,home *handler.Handler){
  mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
  mux.HandleFunc("/artists", artists.AllArtist)
  mux.HandleFunc("/artist", artist.SingleArtist)
  mux.HandleFunc("/", home.Home)
}