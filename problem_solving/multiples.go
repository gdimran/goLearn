package main  
  
import (  
 "fmt"  
)  
  
func sum(array []int) int {  
 result := 0  
 for _, v := range array {  
  result += v  
 }  
 return result  
}  
func main() {  
var num []int
for i:=0; i<1000; i++{
if i%3==0 || i%5==0{
	num = append(num , i)
	}
}

 fmt.Print("Result :", sum(num))  
}  
