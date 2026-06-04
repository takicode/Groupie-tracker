package controllers

import (
	"groupie-tracker/api"
	// "encoding/json"
	"strconv"
	// "fmt"
)

func GeoJson(geoResp []api.GeoResponse)(api.GeoLocation,error){
	var geo api.GeoLocation

	result := geoResp[0]

	lat,err := strconv.ParseFloat(result.Lat,64)
	if err != nil {
	 return api.GeoLocation{}, err
	}
	long,err := strconv.ParseFloat(result.Lon,64)
	if err != nil {
	 return api.GeoLocation{}, err
	}

	geo = api.GeoLocation{
		Lat :lat,
		Long:long,
		Name : result.DisplayName,
	}
// 	fmt.Printf("%+v\n", geo)
//    fmt.Println(geo)
   return geo, nil
}

