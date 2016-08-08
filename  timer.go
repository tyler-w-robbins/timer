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
		"/status",
		TimerIndex,
	},
	Route{
		"TimerStart",
		"POST",
		"/start",
		TimerStart,
	},
	Route{
		"TimerStop",
		"POST",
		"/stop",
		TimerStop,
	},
}

//
// func loadPage(title string) (*Page, error) {
// 	filename := "start.json"
// 	body, err := OPioutil.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Page{Title: title, Timer: body}, nil
// }

func Index(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>",
		p.Title, p.Title, p.Timer)
}

func TimerIndex(w http.ResponseWriter, r *http.Request) {
	pages := Pages{
		Page{Title: "Write presentation", Timer: time.Now()},
		Page{Title: "Host meetup"},
	}

	json.NewEncoder(w).Encode(pages)
}

// func TimerShow(w http.ResponseWriter, r *http.Request) {
// 	// vars := mux.Vars(r)
// 	// todoId := vars["todoId"]
// 	fmt.Fprintf(w, "Todo show: %s", r.URL.Path[1:])
// }

func TimerStart(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// todoId := vars["todoId"]
	fmt.Fprintf(w, "Todo show: %s", r.URL.Path[1:])
}

func TimerStop(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// todoId := vars["todoId"]
	fmt.Fprintf(w, "Todo show: %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/todos", TimerIndex)
	http.HandleFunc("/todos/", TimerStart)

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
