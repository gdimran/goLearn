package main

import "fmt"

func sum(array []int) int {  
 result := 0  
 for _, v := range array {  
  result += v  
 }  
 return result  
}  

func main(){
    var n int
    t1:=0
    t2:=1
    nextTerm:=0
    var num []int
    fmt.Print("Enter the number of terms : ")
    fmt.Scan(&n)
    fmt.Print("Fibonacci Series :")
    for i:=1;i<=n;i++ {
        if(i==1){
		if (t1/2==0){
num = append(num , t1)
            fmt.Print(" ",t1)
            continue
}
	    
}
        if(i==2){
if (t1/2==0){
num = append(num , t2)
            fmt.Print(" ",t2)
            continue
}
		
        }
        nextTerm = t1 + t2
        t1=t2
        t2=nextTerm
	if (nextTerm/2==0){
	num = append(num , nextTerm)
        fmt.Print(" ",nextTerm)
	}
         
    }
fmt.Print("Result :", sum(num))  
}