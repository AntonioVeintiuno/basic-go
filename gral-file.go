package main

// this file use operations
import "fmt"

func main() {
  fmt.Println(returnText());
}
// this fuction return two results
func returnText() (string, int) {
  dataone := "Moee"
  datatwo := 21

  return dataone, datatwo
}
