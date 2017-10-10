package main

import (
    "fmt"
    "github.com/David-Sharpe/heracles/workouts"
    "github.com/go-pg/pg"
    "github.com/go-pg/pg/orm"
    "github.com/joho/godotenv"
    "os"
)

func main() {
    godotenv.Load("../.env")
    fmt.Println("Building database")
    db := pg.Connect(&pg.Options{
        User: "postgres",
        Password: os.Getenv("DB_PASSWORD"),
        Database: "heracles",
    })
    
    err := db.CreateTable(&workouts.User{}, &orm.CreateTableOptions{Temp: false,})
    fmt.Println(err)
}
