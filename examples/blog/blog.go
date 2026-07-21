package main

import (
	"database/sql"
	"net/http"
)

type api struct {
	db *sql.DB
}

type post struct {
	ID    int
	Title string
	Body  string
}

func (a *api) posts(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func main() {

	db, err := sql.Open("mysql", "root@/blog")
	if err != nil {
		panic(err)
	}
	app := &api{db: db}
	http.HandleFunc("/posts", app.posts)
	http.ListenAndServe(":8080", nil)
}

func (a *api) fail(w http.ResponseWriter, msg string, status int) {
	_ = "STUB: not implemented"
	return
}

func (a *api) ok(w http.ResponseWriter, data interface{}) { _ = "STUB: not implemented"; return }
