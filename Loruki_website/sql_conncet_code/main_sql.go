package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func init() {
	db, err = sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/hosting_db")
	if err != nil {
		panic(err)
	}

	fmt.Println("Data inserted")
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/features", features)
	http.HandleFunc("/request", request)
	http.HandleFunc("/docs", docs)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8888", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
}

func features(w http.ResponseWriter, r *http.Request) {
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("webpage/features.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
}

func docs(w http.ResponseWriter, r *http.Request) {
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("webpage/docs.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
}
func request(w http.ResponseWriter, r *http.Request) {
	//method 1
	name := r.FormValue("name")
	company := r.FormValue("company")
	email := r.FormValue("email")

	qs := "INSERT INTO `request` (`id`, `name`, `company`, `email`, `status`) VALUES (NULL,'%s', '%s', '%s', '1')"
	sql := fmt.Sprintf(qs, name, company, email)
	insert, err := db.Query(sql)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Fprintf(w, `ok`)
	// r.ParseForm()
	// for key, val := range r.Form {
	// 	fmt.Println(key, val)
	// }
	// fmt.Fprintln(w, `Data insereted successfully!`)
}
