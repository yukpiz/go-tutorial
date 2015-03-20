package main

import (
    "fmt"
    "net"
    "os"
    "time"
)

func main() {
    fmt.Println("The time is", time.Now())
    fmt.Println(os.Open("PlayGround.go"))
    fmt.Println(net.Dial("tcp", "http://google.com"))
}
