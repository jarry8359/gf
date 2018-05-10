package main

import (
    "gitee.com/johng/gf/g/os/gflock"
    "fmt"
    "time"
)

func main() {
    l := gflock.New("1.lock")
    fmt.Println(l.Path())
    fmt.Println(l.Lock())
    fmt.Println("lock 1")
    fmt.Println(l.Lock())
    fmt.Println("lock 1")
    time.Sleep(time.Hour)
}
