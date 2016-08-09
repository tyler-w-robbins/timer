package main

import (
	"html/template"
	"net/http"
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

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := Status{Name: "foo", Running: true, Seconds: time.Now()}
	err := templates.ExecuteTemplate(w, "index.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ajaxHandler(w http.ResponseWriter, r *http.Request) {}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ajax", ajaxHandler)
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
