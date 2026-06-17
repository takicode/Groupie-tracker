package artist

import (
	"sync"
	"context"
)

type Store struct{
	client ArtistClient
    artists []FullArtistInfo
}

type ArtistClient interface{
	GetArtists(ctx context.Context) ([]Artist, error)
	GetRelations(ctx context.Context)([]Relation, error)
}

func NewStore(client ArtistClient) *Store{
	return &Store{
		client:client,
		artists:make([]FullArtistInfo, 0),
	}
}

func (s *Store) Load(ctx context.Context)error {

	var( 
	artists   []Artist
	relations []Relation
	artistErr   error
	relationErr error
	wg sync.WaitGroup
)

result := make([]FullArtistInfo,0,len(artists))
    
	wg.Add(2)

	go func(){
		defer wg.Done()
		artists, artistErr = s.client.GetArtists(ctx)
	}()

	go func(){
		defer wg.Done()
		relations, relationErr = s.client.GetRelations(ctx)  
	}()
    
	wg.Wait()
	
	if artistErr != nil {
		return artistErr
	}

	if relationErr != nil {
		return relationErr
	}

   relMap := make(map[int]Relation, len(relations))

   for _, rel := range relations{
	    relMap[rel.ID] = rel
   }

   for _, artist := range artists{
	rel, ok:=relMap[artist.ID]
	if !ok{
		continue
	}

	info := FullArtistInfo{
		Artist:artist,
		DatesLocations:rel.DatesLocations,
	}

     result = append(result, info)

	}
	s.artists = result
    return nil
}

func (s *Store) Artists() []FullArtistInfo{
      result := make([]FullArtistInfo, len(s.artists))
	  copy(result, s.artists)
	  return result
   }