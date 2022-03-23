package configuration

import (
	"attendance-backend/services"
	"fmt"
)

const dataDir = "models/ai_models"

func face_detection() {
	err := services.Rec.Init(dataDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	services.Rec.Tolerance = 1
	services.Rec.UseGray = true
	services.Rec.UseCNN = true
	services.LoadDataset(&services.Rec)
}
