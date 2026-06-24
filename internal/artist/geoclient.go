package artist

import(
	"net/http"
	"time"
	"context"
	"net/url"
	"encoding/json"
	"fmt"
	"strings"
)

type GeoClient struct{
	client *http.Client
	apiKey string
	baseApiUrl string
}

func NewGeoClient(apiKey string, baseApiUrl string) *GeoClient{
   return &GeoClient{
			client:&http.Client{
				Timeout:10 *time.Second,
			},
			apiKey:strings.TrimSpace(apiKey),
			baseApiUrl:baseApiUrl,
   }
}

func (g *GeoClient)GetCoordinates(ctx context.Context,location string) (GeoLocation, error){
   params := url.Values{}

   params.Set("q", location)
   params.Set("key", g.apiKey)

   endpoint:= g.baseApiUrl + "?" + params.Encode()
  
   req, err:= http.NewRequestWithContext(ctx,http.MethodGet,endpoint, nil )

   if err != nil {
		return GeoLocation{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := g.client.Do(req)

	if err != nil{
		return GeoLocation{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return GeoLocation{}, fmt.Errorf("api request failed with status: %s (code %d)", resp.Status, resp.StatusCode)
	}


	var data openCageResponse 

	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil{
        return GeoLocation{}, err
	}


	if len(data.Results) == 0 {
		return GeoLocation{}, fmt.Errorf("no coordinates found for %s",location)
    }
    
	result := data.Results[0]

	geo:= GeoLocation{
        Lat:result.Geometry.Lat,
   		Lon:result.Geometry.Lng,
	}

   return geo, nil
}