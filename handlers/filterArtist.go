package handlers

import (
  "strings"
  // "fmt"
  "net/http"
  "groupie-tracker/api"
//   "strconv"
)



func FilterArtist(w http.ResponseWriter, r *http.Request)[]api.FullArtistInfo{
  artists := api.AllArtist()

//   search, filter and page Queries
  search := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("search")))
//   decadeFilter:= r.URL.Query().Get("decade")
//   locationFilter := r.URL.Query().Get("location")
//   sortFilter := r.URL.Query().Get("sort_by")


var result []api.FullArtistInfo

  
  for _, artist := range artists{

	  //   Search logic
	if search != ""{
		matched := false

		if strings.Contains(strings.ToLower(artist.Artist.Name), search){
          matched = true;
		}


		for _, member := range artist.Artist.Members{
			if strings.Contains(member,search ){
				matched = true
				break
			}
		}


		for loc := range artist.DateLocations{
             if strings.Contains(loc,search ){
				matched = true
				break
			}
		}


		if !matched{
			continue
		}
	} 

    

     result = append(result, artist)

  }
  return result

}