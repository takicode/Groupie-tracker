package artist

import (
	"slices"
)

func UniqueLocations(relations []Relation) []string{

relMap := make(map[string]bool)
   
 for _, rel := range relations{
	for loc := range rel.DatesLocations{
          relMap[loc] = true
	}
 }

 locations := make([]string, 0, len(relMap)) 

 for loc :=  range relMap{
	locations = append(locations, loc)
 } 


slices.Sort(locations)

return locations

}