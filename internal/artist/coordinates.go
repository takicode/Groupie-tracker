package artist

import (
	"encoding/json"
	"os"
)

const coordinateFile = "./data/coordinates.json"

type CoordinateMap map[string]GeoLocation

func LoadCoordinates() (CoordinateMap, error) {
	file, err := os.ReadFile(coordinateFile)

	if err != nil {
		if os.IsNotExist(err) {
			return CoordinateMap{}, nil
		}
		return CoordinateMap{}, err
	}

	coord := make(CoordinateMap)

	err = json.Unmarshal(file, &coord)

	if err != nil {
		return CoordinateMap{}, err
	}
	return coord, nil

}

func SaveCoordinates(coords CoordinateMap) error {

	if err := os.MkdirAll("./data", 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(coords, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(coordinateFile, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
