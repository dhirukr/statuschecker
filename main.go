package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Status string

const (
	CHECKING = "CHECKING"
	UP       = "UP"
	DOWN     = "DOWN"
)

var m = map[string]Status{}

// {"http://www.google.com": false, }

type Urls struct {
	Websites []string `json:"websites"`
}

func Check(name string) (status bool) {
	resp, err := http.Get(name)
	if err == nil && resp.StatusCode == http.StatusOK {

		return true
	} else if err != nil {

		return false
	}
	return
}

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", ReqHandler)
	http.ListenAndServe(":8080", nil)
}

func ReqHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		updateMap(r, &m)
	case "GET":
		name := r.URL.Query().Get("name")
		if name != "" {
			if _, ok := m[name]; ok {
				fmt.Println("Checking individually ", string(name), " is ", m[name])
			} else {
				fmt.Println("Website url not in map.Add the url first")
			}

			return
		} else {
			getStatus(r)

		}
		go updateStatus(&m)

	default:
		fmt.Println("Unexpected command")
	}

}

func updateStatus(m *map[string]Status) {

	for {
		for key := range *m {

			updateStatusUtil(key)

		}
		//fmt.Println("Status is checked and updated")
		//fmt.Printf("%+v\n", m)
		for k, v := range *m {
			fmt.Println(k, "is", v)
		}
		fmt.Println("\n")
		time.Sleep(60 * time.Second)
	}
}

func updateStatusUtil(key string) {

	status := Check(key)
	if status {

		m[key] = UP
	} else {

		m[key] = DOWN
	}
}

func updateMap(r *http.Request, m *map[string]Status) {

	urls := Urls{}
	err := json.NewDecoder(r.Body).Decode(&urls)
	if err != nil {
		log.Fatal("Unable to decode JSON request body:", err)
	}
	for _, val := range urls.Websites {
		if _, ok := (*m)[val]; !ok {
			(*m)[val] = CHECKING
		}
	}
	fmt.Println("Map is updated")
}

func getStatus(r *http.Request) {

	if len(m) == 0 {
		fmt.Println("Website list is empty")
		return
	}

	for k, v := range m {
		fmt.Println(k, "is", v)
	}
	fmt.Println("\n")

}
