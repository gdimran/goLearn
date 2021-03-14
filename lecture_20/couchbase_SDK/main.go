package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/mateors/mcb"
)

var db *mcb.DB

type RequestTable struct {
	ID      string `json:"aid"`
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Type    string `json:"type"`
	Status  int    `json:"type"`
}

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

	//How to insert into couchbase bucket
	var myData RequestTable

	form := make(url.Values, 0)
	form.Add("bucket", "golang_project") //bucket Name
	form.Add("aid", "request::2")        //document ID
	form.Add("name", "Mostain Billah")
	form.Add("company", "Mators")
	form.Add("email", "mostain.billah@gmail.com")
	form.Add("type", "request")
	form.Add("status", "1") //what type of data or table name in general (SQL)

	p := db.Insert(form, &myData)    //pass by reference (&myData)
	fmt.Println("Status:", p.Status) //p.Status == Success means data successfully inserted to bucket.

	//How to retrieve from couchbase bucket (selected fields only)

	pres := db.Query("SELECT aid,name,company,email FROM golang_project WHERE type='request'")
	rows := pres.GetRows()

	fmt.Println("Total Rows:", len(rows))
	fmt.Println(rows)

	//How to retrieve from couchbase bucket (All fields using *)

	pres := db.Query("SELECT * FROM golang_project WHERE type='request'")
	rows := pres.GetBucketRows("golang_project") //bucketName as argument

	fmt.Println("Total Rows:", len(rows))
	fmt.Println(rows)

}
