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

func getWorkout(writer http.ResponseWriter, request *http.Request) {
    workout := workouts.Workout{
        ID: getIDFromRequest(request),
    }
    workout.Read(db)
    res, _ := json.Marshal(workout)
    fmt.Fprintf(writer, string(res))
}

func postWorkout(writer http.ResponseWriter, request *http.Request) {
    decoder := json.NewDecoder(request.Body)
    var content workouts.Workout

    decoder.Decode(&content)
    content.Create(db)

    fmt.Fprintf(writer, "%+v\n", content)
    defer request.Body.Close()
}

func putWorkout(writer http.ResponseWriter, request *http.Request) {
    id, _ := strconv.ParseInt(pat.Param(request, "id"), 10, 64)
    decoder := json.NewDecoder(request.Body)
    var content workouts.Workout

    decoder.Decode(&content)
    fmt.Printf("%v\n", content)
    content.ID = id
    content.Update(db)
    fmt.Fprintf(writer, "%+v\n", content)
}

func deleteWorkout(writer http.ResponseWriter, request *http.Request) {
    id, _ := strconv.ParseInt(pat.Param(request, "id"), 10, 64)
    workout := workouts.Workout{
        ID: id,
    }
    workout.Delete(db)
    fmt.Fprintf(writer, "deleted")
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

    // messages := make(chan *workouts.DataObject)
    db = pg.Connect(&pg.Options{
        User: "postgres",
        Password: os.Getenv("DB_PASSWORD"),
        Database: "heracles"
    })

    mux := goji.NewMux()
    // handler, _ := gohaml.NewHamlHandler("./")
    // http.HandleFunc("/", handler)
    mux.HandleFunc(pat.Get("/"), home)
    mux.HandleFunc(pat.Get("/workouts/:id"), getWorkout)
    mux.HandleFunc(pat.Post("/workouts/"), postWorkout)
    mux.HandleFunc(pat.Put("/workouts/:id"), putWorkout)
    mux.HandleFunc(pat.Delete("/workouts/:id"), deleteWorkout)
    // mux.Handle("/", handler);
    // http.ListenAndServe("localhost:8000", handler)
    http.ListenAndServe("localhost:8000", mux)
}
