package main

import "fmt"

func main(){
  var year int

  fmt.Print("Enter year: ")
  fmt.Scanln(&year)

  if year % 4 == 0{
    if year % 100 == 0{
      if year % 400 == 0{
        fmt.Println("Leap year.")
      }else{
        fmt.Println("Not a leap year.")
      }
    }else{
      fmt.Println("Leap year.")
    }
  }else{
    fmt.Println("Not a leap year.")
  }
}