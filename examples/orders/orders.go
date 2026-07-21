package main

import (
	"database/sql"
	"log"
)

const ORDER_PENDING = 0
const ORDER_CANCELLED = 1

type User struct {
	Id       int     `sql:"id"`
	Username string  `sql:"username"`
	Balance  float64 `sql:"balance"`
}

type Order struct {
	Id          int     `sql:"id"`
	Value       float64 `sql:"value"`
	ReservedFee float64 `sql:"reserved_fee"`
	Status      int     `sql:"status"`
}

func cancelOrder(id int, db *sql.DB) (err error) { _ = "STUB: not implemented"; return nil }

func main() {

	db, err := sql.Open("mysql", "root:@/orders")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = cancelOrder(1, db)
	if err != nil {
		log.Fatal(err)
	}
}
