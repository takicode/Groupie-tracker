package api

import (
  "fmt"
  "net/http"
  "encoding/json"
  "log"
  "sync"
)

var artistsInfo []FullArtistInfo

func LoadData()error{
  var wg sync.WaitGroup

    var artists []Artist
    var relations []Relation
    var artistErr error
    var relationErr error

    wg.Add(2)

    go func() {
        defer wg.Done()
        artists, artistErr = getArtists()
    }()
    go func() {
        defer wg.Done()
        relations, relationErr = getRelations()
    }()

    wg.Wait()

    if artistErr != nil {
        return artistErr
    }

    if relationErr != nil {
        return relationErr
    }
  

  relMap := make(map[int]Relation)

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

  return nil
}
func AllArtist ()[]FullArtistInfo{
  return artistsInfo
}


func getArtists()([]Artist, error){
    data,err:= http.Get("https://groupietrackers.herokuapp.com/api/artists")
    if err != nil{
      return []Artist{}, err
    }
   defer data.Body.Close()

    if data.StatusCode != http.StatusOK{
     return []Artist{}, fmt.Errorf("bad status: %s", data.Status)
   }
   
   var artists []Artist


   err = json.NewDecoder(data.Body).Decode(&artists)
    if err != nil{
     return []Artist{}, err
    }
    

    return artists, nil
}

func getRelations()([]Relation, error){
    data,err:= http.Get("https://groupietrackers.herokuapp.com/api/relation")
    if err != nil{
      return []Relation{}, err
    }
   defer data.Body.Close()

   if data.StatusCode != http.StatusOK{
     return []Relation{}, fmt.Errorf("bad status: %s", data.Status)
   }
   
   var relation RelationIndex


   err = json.NewDecoder(data.Body).Decode(&relation)
    if err != nil{
     return []Relation{}, err
    }
    
    return relation.Index, nil
}