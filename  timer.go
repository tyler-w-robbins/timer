package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var templates = template.Must(template.ParseFiles("index.html"))

type Status struct {
	Name    string        `json:"name"`
	Running bool          `json:"running"`
	Seconds time.Time     `json:"seconds"` // make integer
	StopDur time.Duration `json:"seconds"`
}

var statii []*Status
var statusMap = make(map[string]Status)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	htmlString := string(file)
	fmt.Fprint(w, htmlString)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println(r.Form)
	tempName := r.Form.Get("name")
	if _, ok := statusMap[tempName]; ok {
		if statusMap[tempName].Running == true {
			t1 := time.Now()
			fmt.Fprint(w, (statusMap[tempName].StopDur+t1.Sub(statusMap[tempName].Seconds)).String()+"start")
		} else if statusMap[tempName].StopDur > 0 {
			fmt.Fprint(w, statusMap[tempName].StopDur.String()+" - pausedstop")
		} else {
			fmt.Fprint(w, statusMap[tempName].StopDur.String()+"stop")
		}

	} else {
		fmt.Fprint(w, "0m0s")
	}

}

func startHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println(r.Form)
	tempName := r.Form.Get("name")
	t1 := time.Now()
	if statusMap[tempName].StopDur > 0 {
		statusMap[tempName] = Status{Running: true}
		log.Println(statusMap[tempName].StopDur)
	} else {
		statusMap[tempName] = Status{Running: true, Seconds: t1}
	}
	fmt.Fprint(w, statusMap[tempName].Seconds.String()+"start")
}

func stopHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println(r.Form)
	tempName := r.Form.Get("name")
	t1 := time.Now()
	statusMap[tempName] = Status{Running: false, Seconds: t1, StopDur: t1.Sub(statusMap[tempName].Seconds)}
	fmt.Fprint(w, t1.Sub(statusMap[tempName].Seconds).String()+" - pausedstop")
}

func main() {
	statusMap["systest"] = Status{Running: true, Seconds: time.Now()}
	// fmt.Println(reflect.TypeOf(statusArray))
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
