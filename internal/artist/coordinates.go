package artist

import (
    "encoding/json"
    "os"
)

type CoordinateMap map[string]GeoLocation

func LoadCoordinates(filename string) (CoordinateMap, error){
	
}

func SaveCoordinates(filename string, coords CoordinateMap) error{

}