package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type NameList struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Define a global database handle
var db *sql.DB

func main() {
	//replace user, pass, test with coresponding local postgre db
	var err error
	db, err = sql.Open("postgres", "postgres://user:pass@localhost/test?sslmode=disable")
	if err != nil {
		log.Fatal("open db:", err)
	}
	defer db.Close()

	http.HandleFunc("/namelist", handleTableA)
	log.Println("Starting server on :8888")
	err = http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleTableA(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT firstname, lastname FROM public.namelist")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var result []NameList
	for rows.Next() {
		var row NameList
		if err := rows.Scan(&row.FirstName, &row.LastName); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result = append(result, row)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResult)
}
