package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mickelsonm/demo-api/models/tickets"
)

func main() {
	router := mux.NewRouter()
	//just the default handler
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	//GET tickets
	router.HandleFunc("/tickets", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ticks, err := tickets.GetAll()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			toJSON(w, ticks)
		case http.MethodPost, http.MethodPut, http.MethodDelete:
			var ticket tickets.Ticket

			err := json.NewDecoder(r.Body).Decode(&ticket)
			if err != nil {
				http.Error(w, "Invalid input provided when creating ticket", http.StatusBadRequest)
				return
			}

			if r.Method == http.MethodPost {
				err = ticket.Add()
			} else if r.Method == http.MethodPut {
				err = ticket.Update()
			} else {
				err = ticket.Delete()
			}

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			toJSON(w, ticket)
		default:
			http.Error(w, "Unsupported method", http.StatusBadRequest)
			return
		}
	})

	//GET a ticket by ID
	router.HandleFunc("/tickets/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if id < 1 {
			err = errors.New("Invalid ticket ID")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		t := tickets.Ticket{ID: id}

		if err = t.Get(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		toJSON(w, t)
	})

	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(loggedRouter)))
}

func toJSON(w http.ResponseWriter, data interface{}) {
	js, err := json.MarshalIndent(data, " ", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(js)
}
