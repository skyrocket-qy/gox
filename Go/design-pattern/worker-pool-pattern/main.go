package main

import (
	"fmt"
	"time"
)

func NewWorker(jobs chan string, results chan string) {
	WorkeInit()
	for job := range jobs {
		fmt.Println("Processing job ", job)
		time.Sleep(1 * time.Second)
		results <- job
	}
}

func WorkeInit() {
	fmt.Println("initialize worker....")
	time.Sleep(3 * time.Second)
}

func main() {
	jobs := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
	}

	jobChan := make(chan string, len(jobs))
	resultsChan := make(chan string, len(jobs))

	for i := 0; i < 3; i++ {
		go NewWorker(jobChan, resultsChan)
	}

	for _, job := range jobs {
		jobChan <- job
	}

	for i := 0; i < len(jobs); i++ {
		fmt.Println(<-resultsChan)
	}
}
