package main

// this file use operations
import "fmt"

func main() {
  var numero1 float32 = 10
  var numero2 float32 = 6

  fmt.Println("Operación 1")
  operations(numero1, numero2)

  fmt.Println("~~~~~~~~~~~~~~~~~~~~")
  var numero3 float32 = 20
  var numero4 float32 = 12
  fmt.Println("Operación 2")
  operations(numero3, numero4)
}

// this function return a number type float32 as seen in the function params
func operation(n1 float32, n2 float32, op string) float32 {
  var result float32
  if op == "+" {
    result = n1 + n2
  } else if op == "-" {
    result = n1 - n2
  } else if op == "*" {
    result = n1 * n2
  } else if op == "/" {
    result = n1 / n2
  }
  return result
}

func operations(numero1 float32, numero2 float32) {
  //Suma
  fmt.Print("La suma es: ")
  // fmt.Println(numero1 + numero2)
  fmt.Println(operation(numero1, numero2, "+"))

  //Resta
  fmt.Print("La resta es: ")
  // fmt.Println(numero1 - numero2)
  fmt.Println(operation(numero1, numero2, "-"))

  //Multiplicaión
  fmt.Print("La multiplicación es: ")
  // fmt.Println(numero1 * numero2)
  fmt.Println(operation(numero1, numero2, "*"))

  //División
  fmt.Print("La división es: ")
  // fmt.Println(numero1 / numero2)
  fmt.Println(operation(numero1, numero2, "/"))
}