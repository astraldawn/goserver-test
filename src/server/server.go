package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	PORT := "127.0.0.1:8080"

	http.HandleFunc("/complete/", CompleteTaskFunc)
	http.HandleFunc("/delete/", DeleteTaskFunc)
	http.HandleFunc("/", ShowAllTasksFunc)
	http.HandleFunc("/add/", AddTaskFunc)
	http.HandleFunc("/deleted/", ShowTrashTaskFunc)
	http.HandleFunc("/trash/", TrashTaskFunc)
	http.HandleFunc("/edit/", EditTaskFunc)
	http.HandleFunc("/completed/", ShowCompleteTasksFunc)
	http.HandleFunc("/restore/", RestoreTaskFunc)
	http.HandleFunc("/update/", UpdateTaskFunc)
	http.HandleFunc("/search/", SearchTaskFunc)

	http.HandleFunc("/login", LoginFunc)
	http.HandleFunc("/register", RegisterFunc)
	http.HandleFunc("/change", PwdChangeFunc)
	http.HandleFunc("/logout", LogoutFunc)

	// http.HandleFunc("/admin", HandleAdmin)
	// http.HandleFunc("/add_user", PostAddUser)

	http.Handle("/static/", http.FileServer(http.Dir("public")))
	log.Print("running server on " + PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func CompleteTaskFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Path[len("/complete/"):]
		message := "complete task " + id + " GET"
		w.Write([]byte(message))
	}
}

func DeleteTaskFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.URL.Path[len("/delete/"):]
		w.Write([]byte("Delete the task " + id + " POST"))
	}
}

func AddTaskFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.URL.Path[len("/add/"):]
		w.Write([]byte("Post the task " + id + " POST"))
	}
}

func ShowAllTasksFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		message := "all pending tasks GET"
		w.Write([]byte(message))
	}
}

func ShowTrashTaskFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		message := "show deleted tasks GET"
		w.Write([]byte(message))
	}
}

func TrashTaskFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.URL.Path[len("/trash/"):]
		w.Write([]byte("Trash the task " + id + " POST"))
	}
}

func EditTaskFunc(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/edit/"):]
	if r.Method == "POST" {
		w.Write([]byte("Edit post " + id + " POST"))
	} else {
		w.Write([]byte("Show the edit page " + id + " GET"))
	}
}

func ShowCompleteTasksFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		message := "show completed tasks GET"
		w.Write([]byte(message))
	}
}

func RestoreTaskFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.URL.Path[len("/restore/"):]
		w.Write([]byte("Restore the task " + id + " POST"))
	}
}

func UpdateTaskFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.URL.Path[len("/update/"):]
		w.Write([]byte("Update the task " + id + " POST"))
	}
}

func SearchTaskFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		searchString := r.URL.Path[len("/search/"):]
		w.Write([]byte("Search for " + searchString + " GET"))
	}
}

func LoginFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Write([]byte("Perform login POST"))
	} else {
		w.Write([]byte("Show login page GET"))
	}
}

func RegisterFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Write([]byte("Perform registration (write to DB) POST"))
	} else {
		w.Write([]byte("Show registration page GET"))
	}
}

func LogoutFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.URL.Path[len("/update/"):]
		w.Write([]byte("Perform logout for " + id + " POST"))
	}
}

func PwdChangeFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("Change password GET"))
	}
}

var database Database

//Database encapsulates database
type Database struct {
	db *sql.DB
}

func (db Database) begin() (tx *sql.Tx) {
	tx, err := db.db.Begin()
	if err != nil {
		log.Println(err)
		return nil
	}
	return tx
}

func (db Database) prepare(q string) (stmt *sql.Stmt) {
	stmt, err := db.db.Prepare(q)
	if err != nil {
		log.Println(err)
		return nil
	}
	return stmt
}

func (db Database) query(q string,
	args ...interface{}) (rows *sql.Rows) {
	rows, err := db.db.Query(q, args...)
	if err != nil {
		log.Println(err)
		return nil
	}
	return rows
}

func init() {
	database.db, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}
}

//Close database connection
func Close() {
	database.db.Close()
}

//taskQuery encapsulates Exec()
func taskQuery(sql string, args ...interface{}) error {
	SQL := database.prepare(sql)
	tx := database.begin()
	_, err = tx.Stmt(SQL).Exec(args...)
	if err != nil {
		log.Println("taskQuery: ", err)
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return err
}
