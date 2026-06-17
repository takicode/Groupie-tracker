package artist
 
import (
	"fmt"
	"strings"
	"errors"
)

var ErrArtistNotFound = errors.New("artist not found")

type Service struct{
	store ArtistStore
}

func NewService(store *ArtistStore) *Service{
    return &Service{
		store:store,
	}
}

type ArtistStore interface{
	Artists() []FullArtistInfo
}

func (s *Service) Artists() []FullArtistInfo {
	return s.store.Artists()
}


func(s *Service)ArtistByID(Id int)(FullArtistInfo, error){
   artists := s.store.Artists()

   for _, artist := range artists{
		if artist.Id == Id{
			return artist, nil
		}
   }

   return FullArtistInfo{}, ErrArtistNotFound
}


func (s *Service)search(filter SearchFilter)[]FullArtistInfo{
   artists := s.store.Artists()
   result := make([]FullArtistInfo, 0)

   query := strings.TrimSpace(strings.ToLower(filter.Query))

   if query ==""{
       return s.Artists()
   }
   
   for _, artist := range artists{
	  if matchArtist(query, artist){
          result = append(result, artist)
	  }
   }
   
  return result
}

func matchArtist(query string, artist FullArtistInfo)bool{

	if strings.Contains(strings.ToLower(artist.Name), query){
		return true
	}

	for _,  member:= range artist.Members{
		if strings.Contains(strings.ToLower(member), query){
			return true
		}
	}

	for loc:= range artist.DatesLocations{
		if strings.Contains(strings.ToLower(loc), query){
			return true
		}
	}

	return false
}

func (s *Service) paginate(){
     
} 