package artist

import (
	"strings"
  "strconv"
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
}

func (s *Service) Artists() []FullArtistInfo {
	return s.store.Artists()
}

func (s *Service) ArtistByID(ID int) (FullArtistInfo, error) {
	artists := s.store.Artists()

	for _, artist := range artists {
		if artist.ID == ID {
			return artist, nil
		}
	}

	return FullArtistInfo{}, ErrArtistNotFound
}


func (s *Service) Search(filter SearchFilter) SearchResult {
	artists := s.store.Artists()
	result := make([]FullArtistInfo, 0, len(artists))
  locations := getLocations(artists)
  decades:=GetDecades(artists)
  
	query := strings.TrimSpace(strings.ToLower(filter.Query))
  
	if query == "" {
    
		return SearchResult{
    PaginationInfo:paginate(filter.Page, artists, filter.SortBy),
    Locations:locations,
    Dates:decades,
  }
	}

	for _, artist := range artists {
		if matchArtist(filter, artist) {
			result = append(result, artist)
		}
	}
  pagination:= paginate(filter.Page, result, filter.SortBy)

	return SearchResult{
    PaginationInfo:pagination,
    Locations:locations,
    Dates:decades,
  }
}

func SortBy(result []FullArtistInfo, sortFilter string)[]FullArtistInfo{

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


func paginate(pageNum int, artists []FullArtistInfo, sortBy string) PaginationInfo {
	limit := 9
	totalArtists := len(artists)

  if pageNum < 1 {
    pageNum = 1
}

	var next int
	var prev int

  displayArtist := []FullArtistInfo{}

	start := (pageNum - 1) * limit
	end := start + limit
  
	if end > totalArtists {
		end = totalArtists
	}

	displayArtist = artists[start:end]

	totalPages := totalArtists / limit
	remainder := totalArtists % limit

	if remainder != 0 {
		totalPages += 1
	}

  if totalPages == 0 {
    totalPages = 1
}

if pageNum > totalPages {
  pageNum = totalPages
}
	pages := make([]int, totalPages)

	for i := 0; i < len(pages); i++ {
		pages[i] = i + 1
	}

	if pageNum <= 1 {
		prev = 1
	} else {
		prev = pageNum - 1
	}


	if pageNum >= totalPages {
		next = totalPages
	} else {
		next = pageNum + 1
	}

  displayInfo := SortBy(displayArtist, sortBy)

  return PaginationInfo{
   Artists: displayInfo,
   NextPage:next,
   PrevPage:prev,
   Pages:pages,
   Start:start,
   End:end,
   TotalArtists:totalArtists,
   PageNo:pageNum,
   TotalPages:totalPages,
}

}

func matchArtist(filter SearchFilter, artist FullArtistInfo) bool {

	if strings.Contains(strings.ToLower(artist.Name), filter.Query) {
		return true
	}

	for _, member := range artist.Members {
		if strings.Contains(strings.ToLower(member), filter.Query) {
			return true
		}
	}

	for loc := range artist.DatesLocations {
		if strings.Contains(strings.ToLower(loc), filter.Query) {
			return true
		}
	}

  if filter.Decade != ""{
			DateCreated,_ := strconv.Atoi(filter.Decade)
			if DateCreated == (artist.Artist.CreationDate/ 10) * 10 {
				return true
			}
		}

  for loc := range artist.DatesLocations{
		if strings.Contains(loc,filter.Location ){
			return true
		}
	}


  if filter.Members != ""{
			membersNo, _ :=strconv.Atoi(filter.Members)
			if membersNo >= len(artist.Artist.Members){
			 return true
			}
		}

		

	return false
}


func getLocations(artists []FullArtistInfo) []string{
  
  locMap := make(map[string]bool)
	var locations []string

	for _, artist := range artists{
		for loc := range artist.DatesLocations{
			locMap[loc] = true
		} 
	}

	for loc := range locMap{
        locations = append(locations, loc)
	}
    
	for i := 0; i < len(locations)-1; i++{
		for j:= 0; j < len(locations)-1-i; j++{
	        if  locations[j] > locations[j+1]{
                locations[j], locations[j+1] = locations[j+1], locations[j]
			}
		}
	}


	return locations
}




func GetDecades(artists []FullArtistInfo)[]int{

    dateMap := make(map[int]bool)
	var dates []int

	for _, artist := range artists{
		decade := (artist.Artist.CreationDate/ 10) * 10
		dateMap[decade] = true
	}

	for date := range dateMap{
        dates = append(dates, date)
	}


	for i := 0; i < len(dates)-1; i++{
		for j:= 0; j < len(dates)-1-i; j++{
	        if  dates[j] > dates[j+1]{
                dates[j], dates[j+1] = dates[j+1], dates[j]
			}
		}
	}	

	return dates
}
