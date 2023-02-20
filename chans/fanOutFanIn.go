package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        time.Sleep(time.Second)
        results <- j * j
    }
}


func fanOutFanIn() {
    numJobs := 10
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)

    for a := 1; a <= numJobs; a++ {
        fmt.Println(<-results)
    }
}

func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        fanOutFanIn()
        wg.Done()
    }()
    wg.Wait()
}
