package main 

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
)

type Option struct {
	ID int `json: "id"`
	Value string `json: "value"`
}

type RequestOption struct {
	Value string `json: "value"`
}

func CreateApiServer(db *sql.DB) {

	port := ":8000"

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		res, err := db.Query("SELECT * FROM options")
		var values []Option


		for res.Next() {
			var id int
			var value string 

			err = res.Scan(&id, &value)

			CheckError(err)

			values = append(values, Option{ ID: id, Value: value})

		}


		fmt.Println("Welcome to API")
		
		data, _ := json.Marshal(values)

		w.Write(data)

	}).Methods("GET")


	router.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		
        vars := r.PostFormValue("Value")

        w.WriteHeader(http.StatusOK)

        insert, err := db.Prepare("INSERT options SET id=?,value=?")
        CheckError(err)

        res,err := insert.Exec(nil, vars);
        CheckError(err)

        id, err := res.LastInsertId()
        CheckError(err)

        var idInt int

        idInt = int(id)

        var opt = Option{ID: idInt, Value: vars}

        fmt.Println(id)

		data, _ := json.Marshal(opt)

		w.Write(data)

	}).Methods("POST")

	log.Fatal(http.ListenAndServe(port, router))


}
