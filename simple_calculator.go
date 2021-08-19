//go 1.10.4

package main
import "fmt"


func main() {
  
  var a, b int
  var c string
  fmt.Println("Введите пример")
  fmt.Scanln (&a, &c, &b)
  switch c {
    case "+":
      fmt.Println (a+b)
    case "-":
      fmt.Println (a-b)
    case "*":
      fmt.Println (a*b)
    case "/":
      fmt.Println (a/b)
  }
  
}