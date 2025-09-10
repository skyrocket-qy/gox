package main

import (
	"log"
	"time"
)

func NewWorker(jobs, results chan string) {
	WorkeInit()

	for job := range jobs {
		log.Println("Processing job ", job)
		time.Sleep(1 * time.Second)

		results <- job
	}
}

func WorkeInit() {
	log.Println("initialize worker....")
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

	for range 3 {
		go NewWorker(jobChan, resultsChan)
	}

	for _, job := range jobs {
		jobChan <- job
	}

	for range jobs {
		log.Println(<-resultsChan)
	}
}
