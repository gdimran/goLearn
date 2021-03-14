package main

import "strings"

func Capitalize(st string) []string {
  
	var newString[] string
  for i:=0; i<len(st); i++{
    if i % 2 !=0 {
      newString[i] = strings[i].ToUpper(st);
    }else{
      newString[i] = strings[i].TpLower(st);
    }
  }
  
  return newString.Join('')
}

func main(){
  Capitalize("Imran")
}