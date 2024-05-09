package model

import "github.com/kirillmc/gin_test_server/internal"

type TrainDay struct {
	Id          int64               `json:"id"`
	DayName     string              `json:"day_name"`
	Description string              `json:"description"`
	Exercises   []internal.Exercise `json:"exercises"`
}
