package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"
)

//var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

type Page struct {
	Title   string    `json:"title"`
	Created bool      `json:"created"`
	Timer   time.Time `json:"timer"`
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Pages []Page
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TimerIndex",
		"GET",
		"/timers",
		TimerIndex,
	},
	Route{
		"TimerShow",
		"GET",
		"/timers/{timerId}",
		TimerShow,
	},
	Route{
		"TimerCreate",
		"POST",
		"/timers",
		TimerCreate,
	},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TimerIndex(w http.ResponseWriter, r *http.Request) {
	pages := Pages{
		Page{Title: "Write presentation"},
		Page{Title: "Host meetup"},
	}

	json.NewEncoder(w).Encode(pages)
}

func TimerShow(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// todoId := vars["todoId"]
	fmt.Fprintf(w, "Todo show: %s", r.URL.Path[1:])
}

func TimerCreate(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// todoId := vars["todoId"]
	fmt.Fprintf(w, "Todo show: %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/todos", TimerIndex)
	http.HandleFunc("/todos/", TimerShow)

	http.ListenAndServe(":8080", nil)
	// log.Fatal(http.ListenAndServe(":8080", router))
}

// func main() {
// 	http.HandleFunc("/addtimer/", makeHandler(addHandler))
// ticker := time.NewTicker(time.Second)
// go func() {
// 	for t := range ticker.C {
// 		fmt.Println("Tick at", t.String())
// 	}
// }()
//
// time.Sleep(time.Second * 5)
// ticker.Stop()
// fmt.Println("Ticker stopped")

// enc := json.NewEncoder(os.OpenFile())
// d := map[string]int{"apple": 5, "lettuce": 7}
// enc.Encode(d)

// t0 := time.Now()
// fmt.Println(reflect.TypeOf(t0))
// time.Sleep(time.Second * 5)
// t1 := time.Now()
// fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
// }
