package main
import "fmt"


func find (a float64, b float64) (float64, float64){
  
  p := a*b
  s :=(a+b)*2
  fmt.Println ("Периметр: ", p)
  fmt.Println ("Площадь", s)
  return a, b
}

func main() {
  var a float64
  var b float64
  fmt.Println ("Введите сторону а")
  fmt.Scanln (&a)
  fmt.Println ("Введите сторону b")
  fmt.Scanln (&b)
  find (a, b)
  
  
}