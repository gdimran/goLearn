package main

import "fmt"



func add(x int, y int) int{
result := x+y
return result
}
func addMethode2(a, b int)int{
r := a+b
return r
}

func minus(c, d int) (outPut int){
outPut = c+d
return outPut
}

func rectangle(length int, width int) (area int, parameter int) {
	parameter = 2 * (length + width)
	area = length * width
	return 
}
//passing address to a function

func update( ag *int, n *string){
	*ag = *ag + 10
	*n = *n + " Hossain"
return
}

func main(){
        m:= minus(40,30)
	x2:= addMethode2(20,30)
	x:= add(10,30)
	fmt.Println(x,x2,m)

	var ar, p int
	ar, p = rectangle(20, 30)
	fmt.Println("Area:", ar)
	fmt.Println("Parameter:", p)

	var age = 25
	var name = "Imran"
	fmt.Println("Before: ", name, age)

	update(&age, &name)
	fmt.Println("After: ", name,age)
	


}