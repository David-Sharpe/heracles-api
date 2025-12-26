package main

import (
	"encoding/json"
	"fmt"
	"heracles-api/domain"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello World!\n")
		if err != nil {
			return
		}
	})

	http.HandleFunc("/workouts", func(w http.ResponseWriter, r *http.Request) {
		workout := domain.Workout{Name: "Bench press"}
		err := json.NewEncoder(w).Encode(workout)
		if err != nil {
			return
		}
	})

	http.HandleFunc("/plans", func(w http.ResponseWriter, r *http.Request) {
		workout := domain.Workout{Name: "Squat"}
		if r.Method == "POST" {
			err := json.NewEncoder(w).Encode(workout)
			if err != nil {
				return
			}
		} else {
			workout.Name = "Bench"
			err := json.NewEncoder(w).Encode(workout)
			if err != nil {
				return
			}
		}

	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
