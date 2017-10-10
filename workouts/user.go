package workouts

type User struct {
    ID int64
    Name string
    Email string
    Password string
    Workouts []Workout
}
