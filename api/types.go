package api

// Artist Struct
type Artist struct{
  Id           int      `json:"id"`
  Name         string   `json:"name"`
  Image        string   `json:"image"`
  Members      []string `json:"members"`
  CreationDate int      `json:"creationDate"`
  FirstAlbum   string   `json:"firstAlbum"` 
}


// RelationIndex struct
type RelationIndex struct {
  Index  []Relation  `json:"index"`
}

// Relation struct
type Relation struct{
  Id      int     `json:"id"`
  DatesLocations  map[string][]string  `json:"datesLocations"`
}

// FullArtistInfo struct
type FullArtistInfo struct{
  Artist
  DateLocations map[string][]string
}

