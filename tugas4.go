package main

import "fmt"

func main(){
  var Aldo,Yosep = tinggi(182,178)
  fmt.Println("Aldo :",Aldo," cm")
  fmt.Println("Yosep :",Yosep," cm")
}

func tinggi(x int,y int)(int, int){
  var a = x ;
  var b = y;
  return a,b
}
