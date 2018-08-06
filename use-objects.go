package main
// on this document use method type to create new type or objec to use on golang
import "fmt"

type Car struct {
  brand string
  model string
  year int
  price float32
  coupe bool
}

func main() {
  var red_car = Car{"Honda", "Civic", 2018, 300.00, false}

  fmt.Println(red_car.brand)
  fmt.Println(red_car.model)
}