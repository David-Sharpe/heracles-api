package main

import (
    "github.com/realistschuckle/gohaml"
    "github.com/joho/godotenv"
    "github.com/David-Sharpe/heracles-api/workouts"
    "github.com/go-pg/pg"
    "fmt"
    "text/template"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "os"
    "strconv"

    "goji.io"
    "goji.io/pat"
)

var db *pg.DB

func getIDFromRequest(request *http.Request) int64 {
    id, err := strconv.ParseInt(pat.Param(request, "id"), 10, 64)
    if err != nil {
        panic(err)
    }
    return id
}

func postWorkout(writer http.ResponseWriter, request *http.Request) {
    decoder := json.NewDecoder(request.Body)
    var content workouts.Workout
    decoder.Decode(&content)
    content.Create(db)
    fmt.Fprintf(writer, "%+v\n", content)
    defer request.Body.Close()
}

func getWorkout(writer http.ResponseWriter, request *http.Request) {
    var content workouts.Workout
    content.ID = getIDFromRequest(request)
    content.Read(db)
    fmt.Fprintf(writer, "%+v\n", content)
}

func putWorkout(writer http.ResponseWriter, request *http.Request) {
    var content workouts.Workout
    decoder := json.NewDecoder(request.Body)
    decoder.Decode(&content)
    content.ID = getIDFromRequest(request)
    content.Update(db)
    fmt.Fprintf(writer, "%+v\n", content)
}

func deleteWorkout(writer http.ResponseWriter, request *http.Request) {
    var content workouts.Workout
    content.ID = getIDFromRequest(request)
    content.Delete(db)
    fmt.Fprintf(writer, "%+v\n", content)
}

func home(writer http.ResponseWriter, request *http.Request) {
    var scope = make(map[string]interface{})
    scope["lang"] = "HAML"
    content, _ := ioutil.ReadFile("sample.haml")
    engine, _ := gohaml.NewEngine(string(content))
    output := engine.Render(scope)
    homeTemplate := template.Must(template.New("").Parse(output))
    homeTemplate.Execute(writer, workouts.Workout { Name: "test"})
}

func main() {
    godotenv.Load()
    db = pg.Connect(&pg.Options{
        User: "postgres",
        Password: os.Getenv("DB_PASSWORD"),
        Database: "heracles",
    })

    // hamlHandler, _ := gohaml.NewHamlHandler("./haml")
    mux := goji.NewMux()
    mux.HandleFunc(pat.Get("/workouts/:id"), getWorkout)
    mux.HandleFunc(pat.Put("/workouts/:id"), putWorkout)
    mux.HandleFunc(pat.Post("/workouts/"), postWorkout)
    mux.HandleFunc(pat.Delete("/workouts/:id"), deleteWorkout)
    http.ListenAndServe("localhost:8000", mux)
}
