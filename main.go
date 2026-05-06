package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  "strings"
  "log"
)

// Artist
type Artist struct{
  Id      int     `json:"id"`
  Name    string  `json:"name"`
  Image        string   `json:"image"`
  Members      []string `json:"members"`
  CreationDate int      `json:"creationDate"`
  FirstAlbum   string   `json:"firstAlbum"` 
}

// Location
type LocationIndex struct {
  Index  []Location  `json:"index"`
}

type Location struct{
  Id         int     `json:"id"`
  Locations  []string  `json:"locations"`
}

// Date
type DateIndex struct {
  Index  []Date  `json:"index"`
}

type Date struct{
  Id      int     `json:"id"`
  Dates  []string  `json:"dates"`
}

// Relation
type RelationIndex struct {
  Index  []Relation  `json:"index"`
}

type Relation struct{
  Id      int     `json:"id"`
  DatesLocations  map[string][]string  `json:"datesLocations"`
}

type FullArtistInfo struct{
  Id int
  Name string
  DateLocations map[string][]string
}

func main (){
  artists, err :=getArtists()
  if err != nil {
    log.Fatal("failed to get artists:", err)
  }

  relations, err := getRelations()
  if err != nil {
    log.Fatal("failed to get relations:", err)
  }

  relMap := make(map[int]Relation)

  var artistsInfo []FullArtistInfo

  for _, rel := range relations{
    relMap[rel.Id] = rel
  }
  

  for _, artist := range artists{
    id := artist.Id
    rel, ok := relMap[id]
      if !ok {
        log.Printf("no relation found for artist %d", id)
        continue
      }

    info := FullArtistInfo{
      Id:            id,
      Name:          artist.Name,
      DateLocations: rel.DatesLocations,
    }
    artistsInfo = append(artistsInfo, info)
  }

  for _, info := range artistsInfo{
      fmt.Printf("artist:%s\n", info.Name)

      for place, loc := range info.DateLocations{
        fmt.Println(" ", place, "=>", loc)
      }
      fmt.Println(strings.Repeat("=",20))
  }

}


func getArtists()([]Artist, error){
    data,err:= http.Get("https://groupietrackers.herokuapp.com/api/artists")
    if err != nil{
      return nil, err
    }
   defer data.Body.Close()

    if data.StatusCode != http.StatusOK{
     return nil, fmt.Errorf("bad status: %s", data.Status)
   }
   
   var artists []Artist


   err = json.NewDecoder(data.Body).Decode(&artists)
    if err != nil{
     return nil, err
    }
    

    return artists, nil
}

func getRelations()([]Relation, error){
    data,err:= http.Get("https://groupietrackers.herokuapp.com/api/relation")
    if err != nil{
      return nil, err
    }
   defer data.Body.Close()

   if data.StatusCode != http.StatusOK{
     return nil, fmt.Errorf("bad status: %s", data.Status)
   }
   
   var relation RelationIndex


   err = json.NewDecoder(data.Body).Decode(&relation)
    if err != nil{
     return nil, err
    }
    
    return relation.Index, nil
}