package main

import (
    "log"
    "net/http"
	"fmt"
	"time"
	"io/ioutil"

	"github.com/gorilla/mux"

)

type EchoServer struct {
    Router  *mux.Router
    Echo []byte
}

func (es *EchoServer) GetHandler(w http.ResponseWriter, r *http.Request) {
	if es.Echo == nil {
		fmt.Fprintln(w, "You can store payloads here")
	} else {
		w.Write(es.Echo)
		sep := "\n===================================================="
		log.Println("Payload retrieved:", sep, "\n", string(es.Echo), sep)
		es.Echo = nil
	}
}

func (es *EchoServer) PostHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
    }
	es.Echo = body
	sep := "\n===================================================="
	log.Println("Payload stored:", sep, "\n", string(body), sep)
	fmt.Fprintln(w, "Stored your payload")
}

func main() {
    r := mux.NewRouter()
	es := EchoServer{r, nil}
    r.HandleFunc("/", es.PostHandler).Methods("POST")
    r.HandleFunc("/", es.GetHandler).Methods("GET")

    srv := &http.Server{
        Handler:      r,
        Addr:         ":8080",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
	log.Println("Starting echo server")
    log.Fatal(srv.ListenAndServe())
}