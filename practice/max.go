package main

import "fmt"

func main(){
x := []int{
  48,96,86,68,
  57,82,63,70,
  37,34,83,27,
  19,97, 9,17,
}

var min=x[0]
var max=x[0]
for i:=1; i<len(x); i++{
	if min>x[i]{
		min=x[i]
	}

if max<x[i]{
		max=x[i]
	}
	
}
fmt.Println(min)
fmt.Println(max)


xs := []float64{98,93,77,82,83}

  total := 0.0
  for _, v := range xs {
    total += v
  }
  fmt.Println(total / float64(len(xs)))

}