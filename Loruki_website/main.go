package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/mateors/mcb"
)

//var db *sql.DB
var db *mcb.DB
var err error

func init() {
	db = mcb.Connect("localhost", "imran", "123456")
	res, err := db.Ping()
	if err != nil {

		fmt.Println(res)
		os.Exit(1)
	}
	fmt.Println(res, err)
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

// func request(w http.ResponseWriter, r *http.Request) {
// 	//method 1
// 	name := r.FormValue("name")
// 	company := r.FormValue("company")
// 	email := r.FormValue("email")

// 	qs := "INSERT INTO `request` (`id`, `name`, `company`, `email`, `status`) VALUES (NULL,'%s', '%s', '%s', '1')"
// 	sql := fmt.Sprintf(qs, name, company, email)
// 	insert, err := db.Query(sql)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer insert.Close()
// 	fmt.Fprintf(w, `ok`)
// }

type RequestTable struct {
	ID      string `json:"aid"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Company string `json:"company"`
	Type    string `json:"type"`
	Status  int    `json:"status"`
}

func request(w http.ResponseWriter, r *http.Request) {

	//method-1
	// name := r.FormValue("name")
	// company := r.FormValue("company")
	// email := r.FormValue("email")

	// fmt.Println(name, company, email)
	// fmt.Fprintf(w, `received %s %s %s`, name, company, email) //response

	//method-2
	r.ParseForm()
	for key, val := range r.Form {
		fmt.Println(key, val)
	}
	uniqueNumber := time.Now().UnixNano() / (1 << 22)
	uid := strconv.FormatInt(uniqueNumber, 10)
	var reqTable RequestTable

	r.Form.Add("bucket", "golang_project")
	r.Form.Add("aid", "request::"+uid) //we will update later
	r.Form.Add("type", "request")
	r.Form.Add("status", "1")
	pRes := db.Insert(r.Form, &reqTable)
	fmt.Println(pRes.Status, pRes.Errors)

	fmt.Fprintf(w, `OK`)
}
