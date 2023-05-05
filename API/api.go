package API

import (
	"awesomeProject/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

const port = "localhost:8080"

type API struct {
	router *mux.Router
	db     *sql.DB
}

func NewAPI() (*API, error) {
	res := new(API)
	res.router = mux.NewRouter()

	return res, nil
}

func (a *API) Start() error {
	db, err := database.Connect("root:Sukalol78!@tcp(localhost:3306)/universities")
	if err != nil {
		fmt.Println(err)
		return err
	}
	a.db = db

	a.router.HandleFunc("/platform", a.handleAddPlatform())

	return http.ListenAndServe(port, a.router)
}

func (a *API) Stop() {
	fmt.Println("Stopping API...")
	a.db.Close()
	fmt.Println("API stopped...")
}

func (a *API) handleAddPlatform() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "can't read body", http.StatusBadRequest)
			return
		}
		err = request.Body.Close()
		if err != nil {
			http.Error(writer, "can't close body", http.StatusInternalServerError)
			return
		}

		var msg AddPlatform
		err = json.Unmarshal(body, &msg)
		if err != nil {
			http.Error(writer, "error during unmarshal", http.StatusBadRequest)
			return
		}

		switch request.Method {
		case "POST":
			_, err = a.db.Exec(database.AddPlatformQuery, msg.PlatformName)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
		case "DELETE":
			ins, err := a.db.Prepare(database.DeletePlatformQuery1)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			ins.Exec(msg.PlatformName)
			ins, err = a.db.Prepare(database.DeletePlatformQuery2)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			ins.Exec(msg.PlatformName)
			ins, err = a.db.Prepare(database.DeletePlatformQuery3)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			ins.Exec(msg.PlatformName)
		}

		writer.WriteHeader(http.StatusOK)
	}
}
