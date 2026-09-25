package main

import (
	"fmt"
	"go-netscan/checker"
	"sync"
	"time"
)

var wg sync.WaitGroup

const numWorkers = 3

func main() {
	checkers := []checker.Checker{
		checker.HTTPChecker{URL: "https://golang.org", Timeout: 2 * time.Second},
		checker.TCPChecker{Address: "127.0.0.1:8080", Timeout: 1 * time.Second},
	}

	jobs := make(chan checker.Checker, len(checkers))
	results := make(chan checker.Result, len(checkers))

	for range numWorkers {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	for _, c := range checkers {
		jobs <- c
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		if res.Err != nil {
			fmt.Printf("[%s] %s | latency: %v | error: %v\n", res.Status, res.Target, res.Latency, res.Err)
		} else {
			fmt.Printf("[%s] %s | latency: %v\n", res.Status, res.Target, res.Latency)
		}
	}
}

func worker(jobs <-chan checker.Checker, results chan<- checker.Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for c := range jobs {
		res := c.Check()
		results <- res
	}
}
