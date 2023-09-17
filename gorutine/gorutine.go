package gorutine

import (
    "fmt"
    "sync"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "processing job", j)
        results <- j * 2
    }
}

func Gorutine() {
    numJobs := 10
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    // Create three worker goroutines
    numWorkers := 3
    var wg sync.WaitGroup
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go func(workerId int) {
            defer wg.Done()
            worker(workerId, jobs, results)
        }(i)
    }

    // Send jobs to the workers
    for j := 1; j <= numJobs; j++ {
        fmt.Println("assign job --> ", j)

        jobs <- j
    }
    close(jobs)

    // Collect results from the workers
    wg.Wait()
    close(results)

    // Print the results
    for r := range results {
        fmt.Println(r)
    }
}