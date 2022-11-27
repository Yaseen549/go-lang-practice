package main

import "fmt"

func main(){
  var val1 int
  var val2 int
  var op string
  // op = "-"

  fmt.Print("Enter Value 1: ")
  fmt.Scanln(&val1)
  fmt.Print("Enter Operand: ")
  fmt.Scanln(&op)
  fmt.Print("Enter Value 2: ")
  fmt.Scanln(&val2)

  if op == "+"{
    fmt.Println(val1, op, val2, "=", val1+val2)
  }else if op == "-"{
    fmt.Println(val1, op, val2, "=", val1-val2)
  }else if op == "*"{
    fmt.Println(val1, op, val2, "=", val1*val2)
  }else if op == "/"{
    fmt.Println(val1,op,val2, "=", val1/val2)
  }

}