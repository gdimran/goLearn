package main

import "fmt"


func main(){
//fmt.Println("Enter a number: ")
//var input float64

//fmt.Scanf("%f", &input)

//output := input * 2
//fmt.Println(output)

//x := 5
//x += 1
//fmt.Println(x)

//fmt.Println("Enter temperature in Farenhit: ")
//var F float64
//fmt.Scanf("%f", &F)
//c :=(F - 32) * 5/9
//fmt.Println("Temperature in Celsius: ", c)

const perftInMeter = 0.3048
var ft float64
fmt.Println("Enter area size in Footer: ")
fmt.Scanf("%f", &ft)
meter := (ft * perftInMeter)
fmt.Println("Your area in meter: ", meter)

for i := 1; i <= 10; i++ {
    fmt.Println(i)
  }



}
