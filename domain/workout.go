package domain

import "go/types"

type Workout struct {
	Name   string  `json:"name"`
	Reps   int     `json:"reps"`
	Sets   int     `json:"sets"`
	Weight float64 `json:"weight"`
	Unit   string  `json:"unit"`
}

type Plan struct {
	Name   string  `json:"name"`
	types.Slice `json:"slice"`
}
