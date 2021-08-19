//go 1.10.4

package main
import "fmt"


func main() {
  var arr [18] int
  for i := 0; i < len(arr); i++ {
    fmt.Println(i)
  }
}