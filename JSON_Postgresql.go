package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type NameList struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func main() {
	//replace user, pass, test with coresponding local postgre db
	db, err := sql.Open("postgres", "postgres://user:pass@localhost/test?sslmode=disable")
	if err != nil {
		log.Fatal("open db:", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT firstname, lastname FROM public.namelist")
	if err != nil {
		log.Fatal("query:", err)
	}
	defer rows.Close()

	var result []NameList
	for rows.Next() {
		var row NameList
		if err := rows.Scan(&row.FirstName, &row.LastName); err != nil {
			log.Fatal("scan:", err)
		}
		result = append(result, row)
	}
	if err := rows.Err(); err != nil {
		log.Fatal("rows err:", err)
	}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("json marshal:", err)
	}

	fmt.Println(string(jsonResult))
}
