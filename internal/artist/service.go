package artist

import (
	"strings"
  "strconv"
  "slices"
  "cmp"
)

type Service struct {
	store ArtistStore
}

func NewService(store ArtistStore) *Service {
	return &Service{
		store: store,
	}
}

type ArtistStore interface {
	Artists() []FullArtistInfo
	ArtistID(ID int) (FullArtistInfo, error)
	CoordinatesForLocations(locations []string,) (map[string]GeoLocation)
}

func (s *Service) Artists() []FullArtistInfo {
	return s.store.Artists()
}

func (s *Service) ArtistByID(ID int) (ArtistDetails, error) {
	artist, err := s.store.ArtistID(ID)
	if err != nil{
		return ArtistDetails{},err 
	}
   
	var locations  []string

	for loc := range artist.DatesLocations{
		locations = append(locations, loc)
	}

	coords := s.store.CoordinatesForLocations(locations)


	return ArtistDetails{
		artist,
		Coordinates:coords,
	}, nil
	
}


func (s *Service) Search(filter SearchFilter) SearchResult {
  artists := s.store.Artists()
  locations := getLocations(artists)
  decades:=getDecades(artists)
  
	result := make([]FullArtistInfo, 0, len(artists))
	

	for _, artist := range artists {
		if matchArtist(filter, artist) {
			result = append(result, artist)
		}
	}
  
  sortArtists(result, filter.SortBy)
  pagination:= paginate(filter.Page, result)

	return SearchResult{
    PaginatedArtists:pagination,
    Locations:locations,
    Dates:decades,
  }
}

func sortArtists(result []FullArtistInfo, sortFilter string) {
	switch sortFilter{
  case "A-Z":
      slices.SortFunc(result, func(a, b FullArtistInfo)int{
        return cmp.Compare(a.Name, b.Name)
      })
  case "Z-A":
      slices.SortFunc(result, func(a, b FullArtistInfo)int{
        return cmp.Compare(b.Name, a.Name)
      })
  case "New":
      slices.SortFunc(result, func(a, b FullArtistInfo)int{
        return cmp.Compare(b.CreationDate, a.CreationDate)
      })
  case "Old":
      slices.SortFunc(result, func(a, b FullArtistInfo)int{
        return cmp.Compare(a.CreationDate, b.CreationDate)
      })
  }
}


func paginate(pageNum int, artists []FullArtistInfo) PaginatedArtists {
	limit := 9
	totalArtists := len(artists)
	totalPages := (totalArtists + limit - 1) / limit 
	if totalPages == 0 {
		totalPages = 1
	}

	if pageNum < 1 {
		pageNum = 1
	} else if pageNum > totalPages {
		pageNum = totalPages
	}

	start := (pageNum - 1) * limit
	end := start + limit
	if end > totalArtists {
		end = totalArtists
	}

	displayArtist := []FullArtistInfo{}
	if start < totalArtists {
		displayArtist = artists[start:end]
	}

	pages := make([]int, totalPages)
	for i := range pages {
		pages[i] = i + 1
	}

	prev := pageNum - 1
	if prev < 1 {
		prev = 1
	}

	next := pageNum + 1
	if next > totalPages {
		next = totalPages
	}

	return PaginatedArtists{
		Artists:      displayArtist,
		NextPage:     next,
		PrevPage:     prev,
		Pages:        pages,
		Start:        start,
		End:          end,
		TotalArtists: totalArtists,
		PageNo:       pageNum,
		TotalPages:   totalPages,
	}
}

func matchArtist(filter SearchFilter, artist FullArtistInfo) bool {
	
	query := strings.TrimSpace(strings.ToLower(filter.Query))
	if query != "" {
		matchFound := strings.Contains(strings.ToLower(artist.Name), query)
		
		for _, member := range artist.Members {
			if matchFound {
				break
			}
			if strings.Contains(strings.ToLower(member), query) {
				matchFound = true
			}
		}
		for loc := range artist.DatesLocations {
			if matchFound {
				break
			}
			if strings.Contains(strings.ToLower(loc), query) {
				matchFound = true
			}
		}
		if !matchFound {
			return false
		}
	}


	if filter.Location != "" {
		locMatch := false
		for loc := range artist.DatesLocations {
			if strings.Contains(strings.ToLower(loc), strings.ToLower(filter.Location)) {
				locMatch = true
				break
			}
		}
		if !locMatch {
			return false
		}
	}

	
	if filter.Decade != "" {
		targetDecade, err := strconv.Atoi(filter.Decade)
    if err != nil {
    return false
    }
		artistDecade := (artist.CreationDate / 10) * 10
		if artistDecade != targetDecade {
			return false
		}
	}

	if filter.Members != "" {
		membersNo, err:= strconv.Atoi(filter.Members)
      if err != nil {
      return false
      }
		if len(artist.Members) > membersNo {
			return false
		}
	}

	return true
}


func getLocations(artists []FullArtistInfo) []string{
  
  locMap := make(map[string]bool)
  
	for _, artist := range artists{
		for loc := range artist.DatesLocations{
      locMap[loc] = true
      } 
    }
    
  locations := make([]string, 0, len(locMap))
	for loc := range locMap{
        locations = append(locations, loc)
	}
    
	slices.Sort(locations)
	return locations
}




func getDecades(artists []FullArtistInfo)[]int{

  dateMap := make(map[int]bool)
  
	for _, artist := range artists{
    decade := (artist.CreationDate/ 10) * 10
		dateMap[decade] = true
	}
  
  dates := make([]int, 0, len(dateMap))
	for date := range dateMap{
        dates = append(dates, date)
	}


	slices.Sort(dates)

	return dates
}


