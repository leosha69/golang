package main

import "fmt"

func main() {
    
    var leosha string 
    fmt.Println("Лёша — какой он?")   
    fmt.Scanln(&leosha)
    
    switch leosha {
       case "Потрясный":
        fmt.Println("Да, ты прав!")
        fmt.Println("Да, ты прав!")
        fmt.Println("Да, ты прав!")
        fmt.Println("Да, ты прав!")
        fmt.Println("Да, ты прав!")
        fmt.Println("Да, ты прав!")
        fmt.Println("Да, ты прав!")
        fmt.Println("Да, ты прав!")
        fmt.Println("Да, ты прав!")
        fmt.Println("Да, ты прав!")
       case "Крутой":
        fmt.Println("Он очень крутой!")
       default:
        fmt.Print ("Go fack!")
    }
    
}










