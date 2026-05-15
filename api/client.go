package api

import (
  "fmt"
  "net/http"
  "encoding/json"
  "log"
)

func AllArtist ()[]FullArtistInfo{
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
      Artist:artist,
      DateLocations: rel.DatesLocations,
    }
    artistsInfo = append(artistsInfo, info)
  }
  
  log.Println(artistsInfo)

  return artistsInfo
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