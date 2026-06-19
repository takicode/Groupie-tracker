package controllers

import (
  "strings"
  "net/http"
  "groupie-tracker/api"
  "strconv"
)



func FilterArtist(w http.ResponseWriter, r *http.Request)[]api.FullArtistInfo{
	artists := api.AllArtist()

	// search, filter and page Queries
	search := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("search")))
	decadeFilter:= r.URL.Query().Get("decade")
	locationFilter := r.URL.Query().Get("location")
	sortFilter := r.URL.Query().Get("sort_by")
	members := r.URL.Query().Get("members")
	 

	var result []api.FullArtistInfo

	for _, artist := range artists{

		//   Search logic
		if search != ""{
		matched := false

		if strings.Contains(strings.ToLower(artist.Artist.Name), search){
			matched = true;
		}


		for _, member := range artist.Artist.Members{
			if strings.Contains(strings.ToLower(member),search){
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

		if decadeFilter != ""{
			DateCreated,_ := strconv.Atoi(decadeFilter)
		
			if DateCreated != (artist.Artist.CreationDate/ 10) * 10 {
				continue
			}
		}

		if locationFilter != ""{
			found:= false
			for loc := range artist.DateLocations{
				if strings.Contains(loc, locationFilter){
                   found = true
				   break
				}
			}
			if !found{
				continue
			}
		}

		if members != ""{
			membersNo, _ :=strconv.Atoi(members)

			if membersNo < len(artist.Artist.Members){
				continue
			}
		}


		result = append(result, artist)
	}
    
	// Sorting
	for i := 0; i < len(result)-1; i++{
		for j:= 0; j < len(result)-1-i; j++{
	        if sortFilter == "A-Z" && result[j].Artist.Name > result[j+1].Artist.Name{
                result[j], result[j+1] = result[j+1], result[j]
			}

			if sortFilter == "Z-A" && result[j].Artist.Name < result[j+1].Artist.Name{
                result[j], result[j+1] = result[j+1], result[j]
			}

			if sortFilter == "New" && result[j].Artist.CreationDate < result[j+1].Artist.CreationDate{
                result[j], result[j+1] = result[j+1], result[j]
			}

			if sortFilter == "Old" && result[j].Artist.CreationDate > result[j+1].Artist.CreationDate{
                result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	
	
			
		
	
	
	 
		
	return result

}