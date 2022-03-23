package services

import (
	"attendance-backend/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/leandroveronezi/go-recognizer"
	"gorm.io/datatypes"
)

var Rec recognizer.Recognizer

func AddFile(rec *recognizer.Recognizer, Path, Id string) error {
	err := rec.AddImageToDataset(Path, Id)
	if err != nil {
		return err
	}
	return nil
}
func RegisterFace(imagePath string, name string) error {
	if err := AddFile(&Rec, imagePath, name); err != nil {
		return err
	}
	Rec.SetSamples()
	if err := SaveDataset(&Rec); err != nil {
		return err
	}
	return nil
}
func DetectFace(imagePath string) (string, error) {
	Rec.SetSamples()
	face, err := Rec.Classify(imagePath)
	if err != nil {
		return "", fmt.Errorf("no face detected")
	}
	fmt.Println(face[0].Data.Id)
	// Rec.DrawFaces("faces.jpg", face)
	img, err := Rec.DrawFaces(imagePath, face)
	if err != nil {
		return "", err
	}
	Rec.SaveImage("faces.jpg", img)
	return face[0].Data.Id, nil
}
func SaveDataset(rec *recognizer.Recognizer) error {
	data_byte, err := json.Marshal(rec.Dataset)
	data := datatypes.JSON(data_byte)
	fmt.Println(string(data))
	if err != nil {
		return err
	}
	if data_check := models.GetDataset(); len(data_check) == 0 {
		if _, err := models.AddDataset(data); err != nil {
			return err
		}
		return nil
	}
	if err := models.UpdateDataset(data); err != nil {
		return err
	}
	return nil
}
func LoadDataset(rec *recognizer.Recognizer) error {
	data := models.GetDataset()
	ioutil.WriteFile("/tmp/dataset.json", data, 0777)
	file, err := os.OpenFile("/tmp/dataset.json", os.O_RDONLY, 0777)
	if err != nil {
		return err
	}

	Dataset := make([]recognizer.Data, 0)
	err = json.NewDecoder(file).Decode(&Dataset)
	if err != nil {
		return err
	}

	rec.Dataset = append(rec.Dataset, Dataset...)
	return nil
}
