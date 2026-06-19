package main

import (
   "net/http"
   "html/template"
   "log"
   "groupie-tracker/internal/handler"
   "groupie-tracker/internal/artist"
   // "github.com/joho/godotenv"
   "time"
   "os"
   "os/signal"
   "syscall"
   "context"

)



func main(){
   // if err := godotenv.Load(); err != nil {
   //    log.Println("No .env file, use environmental variable")
   // }
   // port := os.Getenv("PORT")
   // if port == ""{
   //   port = "8080"
   // }

   ctx := context.Background()
  client := artist.NewClient()
  store := artist.NewStore(client)
  
  if err := store.Load(ctx); err !=nil{
   //  return err
    log.Fatal(err)
  }

//   if err := initialize(); err != nil{
//   } 

  mux := http.NewServeMux()
  registerRoutes(mux,store)
 
  server := &http.Server{
     Addr :  ":" + "8080",
     Handler : mux,
     ReadTimeout : 10 * time.Second,
     WriteTimeout : 10 * time.Second,
     IdleTimeout : 60 * time.Second,
  }


  go func(){
    log.Printf("Server listening on port :%s", "8080")
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


// func initialize()error{
  
//   log.Println("Loading artists data...")
  
//   return nil  
// }


func registerRoutes(mux *http.ServeMux, store *artist.Store){
  service := artist.NewService(store)
  temp := template.Must(template.ParseGlob("templates/*.html"))

  render := handler.NewRender(temp)
  artistsHandler := handler.NewHandler(temp, service)
	singleArtistHandler := handler.NewArtistHandler(temp, service) 
	homeHandler := handler.NewHomeHandler(temp,service,render)


  mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

  mux.HandleFunc("/artists", artistsHandler.AllArtist)
  mux.HandleFunc("/artist", singleArtistHandler.SingleArtist)
  mux.HandleFunc("/", homeHandler.Home)
}