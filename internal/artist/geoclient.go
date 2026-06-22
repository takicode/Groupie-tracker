package geocache

import(
	"net/http"
	"time"
	"context"
)

type GeoClient struct{
	client *http.Client
	apiKey string
}

func NewGeoClient(apiKey string) *GeoClient{
   return &GeoClient{
			client:&http.Client{
				Timeout:10 *time.Second
			}
   }
}

func (g *GeoClient)GetCoordinates(ctx context.Context,location string) (GeoLocation, error){
   
}