package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type PgTest struct {
	Id   string
	Name string
}

func QueryPGTEST(db *sql.DB) []PgTest {
	rs := []PgTest{}
	rows, selectErr := db.Query("SELECT * FROM PG_TEST")
	if nil != selectErr {
		fmt.Printf("Select fail %s \n", selectErr.Error())
		return rs
	}

	for rows.Next() {
		t := PgTest{}
		rows.Scan(&t.Id, &t.Name)
		rs = append(rs, t)
	}

	return rs
}

func main() {
	db, err := sql.Open("postgres", "postgres://ocsen:ocsen-hoc-code@localhost:5432/ocsenDB?sslmode=disable")
	defer db.Close()
	if nil != err {
		fmt.Printf("Connect fail %s \n", err.Error())
		return
	}
	uuidStr := uuid.New().String()
	insertReuslt, insertErr := db.Exec("INSERT INTO PG_TEST(ID,NAME) VALUES($1, $2)", uuidStr, "Binh Minh")
	if nil != insertErr {
		fmt.Printf("Select fail %s \n", insertErr.Error())
		return
	} else {
		id, _ := insertReuslt.RowsAffected()
		fmt.Printf("Id %d \n", id)
	}

	var count int = 0
	db.QueryRow("SELECT COUNT(*) AS C FROM PG_TEST").Scan(&count)
	fmt.Printf("Count: %d \n", count)

	pgTestList := QueryPGTEST(db)
	fmt.Println(pgTestList)

	deleteResult, deleteErr := db.Exec("DELETE FROM PG_TEST WHERE ID = $1", uuidStr)
	if nil != deleteErr {
		fmt.Printf("Select fail %s \n", insertErr.Error())
		return
	} else {
		id, _ := deleteResult.RowsAffected()
		fmt.Printf("Id %d \n", id)
	}
	pgTestList = QueryPGTEST(db)
	fmt.Println(pgTestList)
}
