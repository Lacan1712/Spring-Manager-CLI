package json

import (
    "io/ioutil"
    "os"
	"encoding/json"
)

func ReadJson(path string) ([]byte, error) {
    jsonFile, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer jsonFile.Close()

    byteValue, err := ioutil.ReadAll(jsonFile)
    if err != nil {
        return nil, err
    }

    return byteValue, nil
}

func MappingStructToJson(path string, result interface{}) error{
	byteValue, err := ReadJson(path)
    if err != nil {
        return err
    }

	err = json.Unmarshal(byteValue, result)
    if err != nil {
        return err
    }

	return nil
}