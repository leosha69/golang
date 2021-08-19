package main
import "fmt"



func main() {
  
  mylabel:
  for i := 0; i < 10;  i++ {
    fmt.Println("Hello World")
    if i > 5 {
      break mylabel
    }
  }
  
}