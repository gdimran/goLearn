package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/couchbase/gocb/v2"
)

var (
	db  *gocb.Cluster
	err error
)

func init() {
	// //couchbase connection block
	db, err = gocb.Connect(
		"localhost",
		gocb.ClusterOptions{
			Username: "imran",
			Password: "123456",
		})
	checkErr(err)

	// We wait until the bucket is definitely connected and setup.
	err = db.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connection established successfully.")
}

func main() {
	//serving perspective request
	http.HandleFunc("/", home)
	http.HandleFunc("/request", request)
	http.HandleFunc("/features", features)
	http.HandleFunc("/docs", docs)

	//serving file from server to client
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("../assets"))))

	//localhost running on port 8888
	http.ListenAndServe(":8888", nil)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../template/base.gohtml")
	checkErr(err)

	tmpl.Execute(w, nil)
}

func features(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../template/base.gohtml")
	checkErr(err)

	tmpl, err = tmpl.ParseFiles("../wpage/features.gohtml")
	checkErr(err)

	tmpl.Execute(w, nil)
}

func docs(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../template/base.gohtml")
	checkErr(err)

	tmpl, err = tmpl.ParseFiles("../webpage/docs.gohtml")
	checkErr(err)

	// tmpl, err = tmpl.ParseFiles("wpage/test.gohtml")
	// checkErr(err)

	tmpl.Execute(w, nil)
}

func request(w http.ResponseWriter, r *http.Request) {
	//gettint form data
	name := r.FormValue("name")
	company := r.FormValue("company")
	email := r.FormValue("email")

	//Insert Document Start //
	q := `INSERT INTO golang_project (KEY, VALUE) VALUES ( "request::3", { "aid": "%s", "name": "%s", "company": "%s", "email": "%s", "type": "%s", "status": %d })`
	query := fmt.Sprintf(q, "request::3", name, company, email, "request", 1)

	_, err := db.Query(query, &gocb.QueryOptions{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Data inserted successfully.")
	}
	// Insert Document End //
}
