package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type AlertGroup struct {
	id          string
	time        string
	receiver    string
	status      string
	externalurl string
	groupkey    string
}

func main() {
	db, err := sql.Open("postgres", "postgres://alertsnitch:alertsnitch@postgresql.monitoring.svc.cluster.local/alertsnitch?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")

	rows, err := db.Query("SELECT * FROM alertgroup;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	bks := make([]AlertGroup, 0)
	for rows.Next() {
		bk := AlertGroup{}
		err := rows.Scan(&bk.id, &bk.time, &bk.receiver, &bk.status, &bk.externalurl, &bk.groupkey) // order matters
		if err != nil {
			panic(err)
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	for _, bk := range bks {
		// fmt.Println(bk.time, bk.receiver, bk.status, bk.externalurl)
		fmt.Printf("%s, %s, %s, %s, %s, %s\n", bk.id, bk.time, bk.receiver, bk.status, bk.externalurl, bk.groupkey)
	}
}
