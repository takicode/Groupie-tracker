package artist

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Relation struct {
	ID            int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type RelationIndex struct {
	Index []Relation `json:"index"`
}

type FullArtistInfo struct {
	Artist
	DatesLocations map[string][]string
	Coordinates map[string]GeoLocation
}

// incoming data
type SearchFilter struct{
	Query string
	Page int
    Decade string
	Location string
	Members string
	SortBy string
}
// outgoing data
type SearchResult struct {
	PaginatedArtists
    Locations []string
	Dates []int
}

type PaginatedArtists struct {
	Artists []FullArtistInfo
	CurrentPage int
    TotalPages int
    TotalArtists int
    NextPage int
    PrevPage int
	Pages []int
	Start int 
    End int
    PageNo int
}


type GeoLocation struct {
    Lat float64 
    Lon float64 
}



type openCageResponse struct{
	Results []struct{
		Geometry struct{
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"geometry"`
	} `json:"results"`
}


type GeoJob struct {
    Location string
}

type GeoResult struct {
    Location string
    Geo      GeoLocation
    Err      error
}