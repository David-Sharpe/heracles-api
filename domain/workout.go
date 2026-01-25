package domain

type Workout struct {
	Name   string  `json:"name"`
	Reps   int     `json:"reps"`
	Sets   int     `json:"sets"`
	Weight float64 `json:"weight"`
	Unit   string  `json:"unit"`
}
