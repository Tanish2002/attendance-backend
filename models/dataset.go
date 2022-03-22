package models

import "gorm.io/datatypes"

type Dataset struct {
	ID   uint `gorm:"primaryKey"`
	Json datatypes.JSON
}

func AddDataset(json []byte) (*Dataset, error) {
	dataset := &Dataset{
		Json: datatypes.JSON(json),
	}

	if err := DB.Create(dataset).Error; err != nil {
		return nil, err
	}
	return dataset, nil
}
func GetDataset() datatypes.JSON {
	var dataset Dataset
	DB.Take(&dataset, 1)
	return dataset.Json
}
func UpdateDataset(json datatypes.JSON) error {
	if err := DB.Model(&Dataset{
		ID: 1,
	}).UpdateColumns(&Dataset{
		Json: json,
	}).Error; err != nil {
		return err
	}
	return nil
}
