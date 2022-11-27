package main

import "fmt"

func main(){
  var table int
  fmt.Print("Enter table number: ")
  fmt.Scanln(&table)
  for i:=1; i<=10; i++{
    fmt.Println(table ,"x", i, "=", (table*i))
    /* alternative ways */
    // fmt.Printf("%d x %d = %d\n", table, i, i*table)
    // fmt.Print(table, " x ", i, " = ", (table*i), "\n")
  }
}
