package artist


import (
	"context"
	"net/http"
	"time"
	"fmt"
	"encoding/json"
)

type Client struct{
	httpClient *http.Client
	baseURL string
}

func NewClient() *Client{
	return &Client{
		httpClient:&http.Client{
			Timeout:10 * time.Second,
		},
		baseURL:"https://groupietrackers.herokuapp.com/api",
	}
}


func(c *Client) GetArtists(ctx context.Context) ([]Artist, error){
  req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/artists", nil)

  if err != nil{
	return nil, fmt.Errorf("create request:%w",err)
  }
  

  resp, err := c.httpClient.Do(req)

  if err != nil{
	return nil, fmt.Errorf("send request:%w", err)
  }

  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK{
    return nil, fmt.Errorf("Unexpected status code:%d", resp.StatusCode)
  }

  var artists []Artist
  

 err = json.NewDecoder(resp.Body).Decode(&artists)

 if err != nil{
	return nil, fmt.Errorf("decode response:%w", err)
 }

 return artists, nil

}

func (c *Client) GetRelations(ctx context.Context)([]Relation, error){
   req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/relation", nil)

   if err != nil{
		return nil, fmt.Errorf("create request:%w", err)
 	}

	resp, err:=c.httpClient.Do(req)

	if err != nil{
		return nil, fmt.Errorf("send request:%w", err)
 	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
	 	return nil, fmt.Errorf("Unexpected status code:%d", resp.StatusCode)
	}

	var relIndex RelationIndex
	 
	if err := json.NewDecoder(resp.Body).Decode(&relIndex); err !=nil{
		return nil, fmt.Errorf("decode response:%w", err)
	}
	 

	return relIndex.Index, nil
}
