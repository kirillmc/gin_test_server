package model

import "github.com/kirillmc/gin_test_server/internal"

type TrainPrograms struct {
	TrainPrograms []internal.TrainProgram `json:"train_programs"`
}
