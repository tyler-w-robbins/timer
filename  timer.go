package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"
)

var templates = template.Must(template.ParseFiles("index.html"))

// type Page struct {
// 	Title   string    `json:"title"`
// 	Created bool      `json:"created"`
// 	Timer   time.Time `json:"timer"`
// }

type Status struct {
	Name    string    `json:"name"`
	Running bool      `json:"running"`
	Seconds time.Time `json:"seconds"` // make integer
}

var statusArray = []Status{}

// var enc = json.NewEncoder(os.OpenFile("status.json"))
var t0 = time.Now()

func indexHandler(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	htmlString := string(file)
	fmt.Fprint(w, htmlString)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	fmt.Fprint(w, t1.Sub(t0))
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	r.ParseForm()
	log.Println(r.Form)
	fmt.Println(r.PostFormValue("name"))
	fmt.Fprint(w, t1.Sub(t0))
}

func stopHandler(w http.ResponseWriter, r *http.Request) {}

func main() {
	fmt.Println(reflect.TypeOf(statusArray))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/start", startHandler)
	http.HandleFunc("/stop", stopHandler)
	http.ListenAndServe(":8080", nil)
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
