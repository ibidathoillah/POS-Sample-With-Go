package map_json

import "encoding/json"

/*
This function will help you to convert your object from struct to map[string]interface{} based on your JSON tag in your structs.
Example how to use posted in sample_test.go file.
*/
func StructToMap(data interface{}) (map[string]interface{}, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	mapData := make(map[string]interface{})
	err = json.Unmarshal(dataBytes, &mapData)
	if err != nil {
		return nil, err
	}
	return mapData, nil
}
