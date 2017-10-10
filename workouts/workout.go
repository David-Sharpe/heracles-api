package workouts

import (
    "fmt"
    "github.com/go-pg/pg"
    // "github.com/go-pg/pg/orm"
)

// DataObject defines an interface for objects stored in the database
type DataObject interface{
    Create(db *pg.DB)
    Read(db *pg.DB)
    Update(db *pg.DB)
    Delte(db *pg.DB)
}

type Exercise struct {
    Name string `json:"name"`
    Weight float32 `json:"weight"`
    Units string `json:"units"`
    Sets int32 `json:"sets"`
    Reps int32 `json:"reps"`
} 

type Workout struct {
    ID int64 `json:"id"`
    Name string `json:"name"`
    Exercises []Exercise `json:"exercises"`
}

func(workout *Workout) Create(db *pg.DB) error {
    var err error
    fmt.Println("Saving the workout!")
    fmt.Printf("%+v\n", *workout)
    err = db.Insert(workout)
    fmt.Printf("Error after Insert, %v\n", err)
    return err
}

func (workout *Workout) Read(db *pg.DB) error {
    fmt.Println("Get the workout!")
    fmt.Println(*workout)
    return db.Select(workout)
}

func (workout *Workout) Update(db *pg.DB) error {
    var err error
    fmt.Println("Updating.")
    err = db.Update(workout)
    fmt.Printf("Error after Update, %v\n", err)
    return err
}

func (workout *Workout) Delete(db *pg.DB) error {
    var err error
    fmt.Println("Deleting.")
    err = db.Delete(workout)
    fmt.Printf("Error after Delete, %v\n", err)
    return err
}
