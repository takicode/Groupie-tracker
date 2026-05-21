package handlers

import (
  "strings"
  "fmt"
  "net/http"
  "groupie-tracker/api"
  "strconv"
)


type FilterArtistInfo struct{
   Artists []api.FullArtistInfo
   NextPage int
   PrevPage int
   Pages []int
   Start int
   End int
   TotalArtists int
   PageNo int
   TotalPages int
}



func FilterArtist(w http.ResponseWriter, r *http.Request)FilterArtistInfo{
  artists := api.AllArtist()

  // search, filter and page Queries
  search := r.URL.Query().Get("search")
  // decadeFilter := r.URL.Query().Get("decade")
  // locationFilter := r.URL.Query().Get("location")
  // sortFilter := r.URL.Query().Get("sort_by")
  pageString := r.URL.Query().Get("page")


  // Search logic
    
  for _, artist := range artists{
     if strings.ToLower(artist.Artist.Name) != strings.ToLower(search){
       continue
     }
    
     for _, 
  }


  // Pagination logic
  limit := 9

  totalArtists := len(artists)

  pageNo, err:= strconv.Atoi(pageString)

  if err != nil || pageNo < 1 {
    pageNo = 1
  }

  var next int
  var prev int


  start:= (pageNo -1) * limit

  if start >= totalArtists{
    start = 0
    pageNo = 1
  }

  end := start + limit

  if end > totalArtists{
    end = totalArtists
  }

  displayArtist := artists[start:end]

  totalPages := totalArtists / limit
  remainder := totalArtists % limit

  if remainder != 0{
    totalPages += 1
  }


  pages := make([]int, totalPages)


  for i :=0; i < len(pages); i++{
    pages[i] = i+1
  }

  if pageNo <= 1{
  prev = 1 
  }else{
    prev = pageNo -1
  }

  if pageNo >=totalPages{
    next = totalPages
  }else{
    next = pageNo + 1
  }


  data := FilterArtistInfo{
    Artists : displayArtist,
    NextPage:next,
    PrevPage:prev,
    Pages:pages,
    Start:start,
    End:end,
    TotalArtists:totalArtists,
    PageNo:pageNo,
    TotalPages:totalPages, 
  } 
  return data
}