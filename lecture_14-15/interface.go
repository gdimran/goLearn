package main

import "fmt"


type Email interface{

	Write(string,string,string)
	Send() string
	Read()
}

type Test struct{
	To string
	From string
	Subject string
	Message string
}

func (t Test) Write(to string, from string, subject string){
	fmt.Println(to,from,subject)
}

func (t Test) Send() string{
	return t.To
}

func (t Test) Read(){
fmt.Println( "Email received from",t.From,"Message:", t.Message)
}                                                                                                                                                    

func main(){
	
//var e Email

//fmt.Println(e)

var t Test
t.To="ujjal@gmail.com"
t.From ="gdimran@gmail.com"
t.Subject="Test email"
t.Message="Hello this is test message"

t.Write(t.To,t.From,t.Subject)

t.Read()

}