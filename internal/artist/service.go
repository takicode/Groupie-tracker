package artist
 
import (
	"strings"
)


type Service struct{
	store ArtistStore
}

func NewService(store ArtistStore) *Service{
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


func(s *Service)ArtistByID(ID int)(FullArtistInfo, error){
   artists := s.store.Artists()

   for _, artist := range artists{
		if artist.ID == ID{
			return artist, nil
		}
   }

   return FullArtistInfo{}, ErrArtistNotFound
}


func (s *Service)Search(filter SearchFilter)SearchResult{
   artists := s.store.Artists()
   result := make([]FullArtistInfo, 0)

   query := strings.TrimSpace(strings.ToLower(filter.Query))

   if query ==""{
       return paginate(filter.Page, artists)
   }
   
   for _, artist := range artists{
	  if matchArtist(query, artist){
          result = append(result, artist)
	  }
   }
   
  return paginate(filter.Page, result)
}

func paginate(pageNum int, artists []FullArtistInfo)SearchResult{
  limit := 9
  totalArtists := len(artists)


  var next int
  var prev int


  start:= (pageNum - 1) * limit

  if start >= totalArtists{
    start = 0
    pageNum = 1
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

  if pageNum <= 1{
  prev = 1 
  }else{
    prev = pageNum -1
  }

  if pageNum >=totalPages{
    next = totalPages
  }else{
    next = pageNum + 1
  }


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
