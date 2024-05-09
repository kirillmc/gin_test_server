package model

//type TrainPrograms struct {
//	TrainPrograms []TrainProgram `json:"train_programs"`
//}
//
//type TrainProgram struct {
//	Id          int64      `json:"id"`
//	ProgramName string     `json:"program_name"`
//	Description string     `json:"description"`
//	Status      string     `json:"status"`
//	TrainDays   []TrainDay `json:"train_days"`
//}
//
//type TrainDay struct {
//	Id          int64      `json:"id"`
//	DayName     string     `json:"day_name"`
//	Description string     `json:"description"`
//	Exercises   []Exercise `json:"exercises"`
//}
//
//type Exercise struct {
//	Id           int64    `json:"id"`
//	ExerciseName string   `json:"exercise_name"`
//	Pictures     []string `json:"pictures"`
//	Description  string `json:"description"`
//	Sets         []Set  `json:"sets"`
//}

type Set struct {
	Id       int64   `json:"id"`
	Quantity int64   `json:"quantity"`
	Weight   float64 `json:"weight"`
}
