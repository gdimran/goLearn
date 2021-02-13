package main

type weight float32

var name string ="Imran"


type Book struct{
	Title string
	Author string
	ISBN	string
	Price float32
	Pages int
}

var b1 Book
/*b1.Title ="An Introduction to programming Go"
b1.Author="CALEB DOXY"
b1.ISBN="978-1478355823"
b1.Price=10.50
b1.Pages=165*/

//fmt.Println(b1)

/*
car := struct{
	make string
	model string
	mileage int
}{
	make: "Ford",
	model: "Taurus",
	mileage: 20000,
}*/

//fmt.Println(car)
type car struct{
	make string
	model string
	mileage int
}
var c= car{
	make: "Ford",
	model: "Taurus",
	mileage: 20000,
}

