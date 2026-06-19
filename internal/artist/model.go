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
}

type SearchFilter struct{
	Query string
}

type Pagination struct {
	Page  int
	Limit int
	SortBy   string
	SortDesc bool
}

type ArtistQuery struct {
	Search SearchFilter
	Page   Pagination
}
