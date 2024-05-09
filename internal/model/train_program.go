package model

import "github.com/kirillmc/gin_test_server/internal"

type TrainProgram struct {
	Id          int64               `json:"id"`
	ProgramName string              `json:"program_name"`
	Description string              `json:"description"`
	Status      string              `json:"status"`
	TrainDays   []internal.TrainDay `json:"train_days"`
}
