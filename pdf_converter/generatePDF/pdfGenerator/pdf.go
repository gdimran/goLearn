package pdfGenerator

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"

)

//pdf requestpdf struct
type RequestPdf struct {
	body string
}

//new request to pdf function
func NewRequestPdf(body string) *RequestPdf {
	return &RequestPdf{
		body: body,
	}
}

//parsing template function

func (r *RequestPdf) ParseTemplate(templateName string, data interface{}) error{
	t := time.Now().Unix()
	//write whole the body

	if _, err := os.Stat("cloneTemplate/"); os.IsNotExist(err) {
		errDir:=os.Mkdir("cloneTemplate/",0777)
		if errDir !nil{
			log.Fatal(errDir)
		}
	}

	err1:=ioutil.WriteFile("cloneTemplate/"+strconv.FormatInt(int64(t),10)+".html",[]byte(r.body),0644)

	if err1 != nil{
		panic(err1)
	}

	f,err:=os.Open("cloneTemplate/"+strconv.FormatInt(int64(t),10)+ ".html")
	if f !=nil {
		defer f.Close()
	}
	if err !=nil {
		log.Fatal(err)
	}
	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfg.WriteFile(pdfPath)
	if err != nil {
		log.Fatal(err)
	}

	dir, err := os.Getwd()
	if err != nil{
		panic(err)
	}

	defer os.RemoveAll(dir + "/cloneTemplate")

	return true, nil

}
