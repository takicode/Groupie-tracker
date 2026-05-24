package handlers

import (
	"groupie-tracker/api"
)

func GetLoc()[]string{
	artists := api.AllArtist()

    locMap := make(map[string]bool)
	var locations []string

	for _, artist := range artists{
		for loc := range artist.DateLocations{
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


func GetDates()[]int{
	artists := api.AllArtist()

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

