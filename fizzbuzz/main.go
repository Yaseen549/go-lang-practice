package main

import "fmt"

func main(){
  var n int
  fmt.Print("How many numbers: ")
  fmt.Scanln(&n)
  a := 1
  b := 1
  fmt.Println(a)
  fmt.Println(b)
  for i:=3;i<=n;i++{
    if i%3 ==0 && i%5 ==0{
      fmt.Println("FizzBuzz")
    }else if i%3==0{
      fmt.Println("Fizz")
    }else if i%5==0{
      fmt.Println("Buzz")
    }else{
      fmt.Println(i)
    }
  }
}