package main

import (
  "fmt"
  // "time"
)

func main() { // time.Sleep(time.Second * 5)
  var suma int = 10 + 11
  var resta int = suma - 11
  var name string = "Antonio"
  var apellido string = "Galvan"
  last_name := "Garcia"
  fmt.Println("Hello " + name +" "+ apellido + " " +last_name)
  fmt.Println(suma)
  fmt.Println(resta)
}
