package main

import "fmt"


func update(a *int){
	fmt.Println(a)
	*a=*a+10
}


func main(){
      
	var x int
	var y *int
	fmt.Println(x, &x)
	fmt.Println(y, &y)

	x=15
	y=&x  //referencing value
	fmt.Println(x, y)
	fmt.Println("y dereferencing value is", *y)
	
	update(&x)
	fmt.Println(" Update value is", x)
}