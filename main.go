package main

import (
	"log"

	"github.com/David-Sharpe/heracles-api/workouts"
	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	// "github.com/go-pg/pg/v9"
	// "github.com/go-pg/pg/v9/orm"
	"fmt"
	// "text/template"
	// "io/ioutil"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"goji.io"
	"goji.io/pat"
)

// var db *pg.DB
var cache redis.Conn
var notice []byte

func getIDFromRequest(request *http.Request) int64 {
	id, err := strconv.ParseInt(pat.Param(request, "id"), 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("id is %#v\n", id)
	return id
}

func getWorkouts(writer http.ResponseWriter, request *http.Request) {
	var res string
	fmt.Printf("getting workout")
	res, err := redis.String(cache.Do("GET", getIDFromRequest(request)))
	if err != nil {
		workout := workouts.Workout{
			ID:        getIDFromRequest(request),
			Exercises: []workouts.Exercise{},
		}
		// workout.Read(db)
		temp, _ := json.Marshal(workout)
		res = string(temp)
		//     cache.Do("SET", getIDFromRequest(request), res)
	} else {
		fmt.Printf("err is %#v", err)
	}

	fmt.Printf("%#v\n", res)
	fmt.Fprintf(writer, res)
}

func postWorkout(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var content workouts.Workout

	decoder.Decode(&content)
	// content.Create(db)

	response, _ := json.Marshal(content)
	cache.Do("SET", content.ID, response)
	fmt.Fprintf(writer, "%s\n", response)
	defer request.Body.Close()
}

func putWorkout(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("Entering put bitches\n")
	id, _ := strconv.ParseInt(pat.Param(request, "id"), 10, 64)
	fmt.Printf("id: %d\n", id)
	fmt.Printf("request: %+v\n", request)
	decoder := json.NewDecoder(request.Body)
	var content workouts.Workout

	err := decoder.Decode(&content)
	if err != nil {
		panic(err)
	}
	fmt.Printf("putting workout: %+v\n", content)
	content.ID = id
	// content.Update(db)
	response, _ := json.Marshal(content)
	cache.Do("SET", id, response)
	fmt.Fprintf(writer, "%s\n", response)
}

func deleteWorkout(writer http.ResponseWriter, request *http.Request) {
	id, _ := strconv.ParseInt(pat.Param(request, "id"), 10, 64)
	// workout := workouts.Workout{
	//     ID: id,
	// }
	// workout.Delete(db)
	cache.Do("DEL", id)
	fmt.Fprintf(writer, "deleted")
}

func home(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "OK\n")
}

// func buildDB(writer http.ResponseWriter, request *http.Request) {
//     err := db.CreateTable(&workouts.Workout{}, &orm.CreateTableOptions{Temp: false,})
//     fmt.Fprintf(writer, err.Error())
// }

// func getNotified(writer http.ResponseWriter, request *http.Request) {
//     body, err := ioutil.ReadAll(request.Body)
//     notice = body
//     if err != nil {
//         fmt.Fprintf(writer, "OK")
//     } else {
//         fmt.Fprintf(writer, "Fail")
//     }
// }

func retrieveNotifications(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, string(notice))
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("middleware\n")
		fmt.Printf("Method: %p\n", r)
		next.ServeHTTP(w, r)
		fmt.Printf("reponse: %p\n", r)
	})
}
func main() {
	godotenv.Load()
	println("starting up\n")
	// config := newrelic.NewConfig("heracles-api", os.Getenv("NEW_RELIC_KEY"))
	// app, err := newrelic.NewApplication(config)
	// dbOptions, _ := pg.ParseURL(os.Getenv("DATABASE_URL"))
	// fmt.Printf("%+v\n", *dbOptions)

	// messages := make(chan *workouts.DataObject)
	// if os.Getenv("ENVIRONMENT") == "dev" {
	//     dbOptions.TLSConfig = nil
	// }
	// db = pg.Connect(dbOptions)

	var connectionError error
	cache, connectionError = redis.DialURL(os.Getenv("REDIS_URL"))
	if connectionError != nil {
		// Handle error
	}
	defer cache.Close()

	mux := goji.NewMux()
	// workoutMux := goji.SubMux()

	mux.HandleFunc(pat.Get("/"), home)
	// mux.HandleFunc(pat.Get("/notifications"), retrieveNotifications)
	// mux.HandleFunc(pat.Post("/notifications"), getNotified)

	// mux.Handle(pat.Get("/workouts/:id"), middleware(http.HandlerFunc(getWorkouts)))
	// mux.Handle(pat.New("/workouts/*"), middleware(http.HandlerFunc(getWorkouts)))

	// workoutMux.HandleFunc(pat.Get(""), getWorkouts)
	mux.HandleFunc(pat.Get("/workouts/:id"), getWorkouts)
	mux.HandleFunc(pat.Post("/workouts/"), postWorkout)
	mux.HandleFunc(pat.Put("/workouts/:id"), putWorkout)
	// mux.HandleFunc(pat.Put("/workouts/:id"), putWorkout)
	mux.HandleFunc(pat.Delete("/workouts/:id"), deleteWorkout)
	localCors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost"},
		AllowCredentials: true,
		AllowedMethods:   []string{"PUT", "POST", "GET", "DELETE", "HEAD"},
		Debug:            true,
	})
	// handler := cors.Default().Handler(mux)
	handler := localCors.Handler(mux)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handler))
	// log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), mux))
}
