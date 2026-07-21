package main

import "database/sql"

func recordStats(db *sql.DB, userID, productID int64) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func main() {

	db, err := sql.Open("mysql", "root@/blog")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = recordStats(db, 1, 5); err != nil {
		panic(err)
	}
}
