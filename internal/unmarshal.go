package internal

import (
	"asher/internal/models"
	"encoding/json"
	"io/ioutil"
)

func ToAsherObject(filePath string) (*models.Asher, error) {
	asherObject := new(models.Asher)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &asherObject)
	if err != nil{
		return nil, err
	}
	return asherObject, nil
}
