package main

// this file use operations
import "fmt"

func main() {
  // fmt.Println(returnText());
  fmt.Print("Order 1 -->");
  fmt.Println(hats("$", 45));

  fmt.Println("~~~~~~~~~~~~~~~~");

  fmt.Print("Order 2 -->");
  fmt.Println(hats("$", 24 ));
}

func hats(coin string, order float32) (string, string, float32) {
  price := func() float32 {
    return order * 7
  }
  return "Hat price is: ", coin, price()
}

// this fuction return two results
func returnText() (string, int) {
  dataone := "Moee"
  datatwo := 21

  return dataone, datatwo
}
