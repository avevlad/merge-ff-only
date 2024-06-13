package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int) {
    fmt.Printf("Worker %sss starting\n", id)

    time.Sleep(time.Second)
    fmt.Printf("Worker %d done", id)
}

func main() {

    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(100)

        go func() {
            defer wg.Done()
            worker(i)
        }()
    }

    wg.Wait()

}
