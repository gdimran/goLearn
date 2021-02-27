package main

import "fmt"

type Doctor struct{
	Name string
	Education string
	Age int
	Salary float32
}

//method

func (d Doctor) Speak(){
	fmt.Println(d.Name, "Can speak")
}

func (d Doctor)getName() string{
	return d.Name
}

func main(){

//assign value in struct way 1
var d Doctor
d.Name="Imran"
d.Education="Msc" 
d.Age=28
 d.Salary=20000

//var d=Doctor{Name:"Imran1",Education:"Msc",Age:28,Salary:20000.00,}
//fmt.Println(d.Name, d.Education, d.Age,d.Salary)
d.Speak()
fmt.Println(d.getName())
}