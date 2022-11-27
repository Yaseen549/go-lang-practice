package main

import "fmt"

func main(){
  var bmi float32
  var weight float32
  var height float32

  fmt.Print("Enter Height in cm: ")
  fmt.Scanln(&height)
  fmt.Print("Enter Weight in kg: ")
  fmt.Scanln(&weight)

  height = height/100

  bmi = weight/(height*height)
  fmt.Println(bmi)
}