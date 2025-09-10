package main

import (
	"sync"
	"time"
)

func main() {
	teacher := NewTeacher()
	students := []*Student{}
	for _, name := range []string{"A", "B", "C", "D", "E"} {
		students = append(students, &Student{
			Name: name,
		})
	}
	StartGame(teacher, students)
}

func StartGame(teacher *Teacher, students []*Student) {
	numOfStudents := len(students)
	round := 1
	questionChan := make(chan Question, numOfStudents)
	waitAnswer := &sync.WaitGroup{}
	whoAnswer := WhoAnswer{
		Name:  "",
		Mutex: sync.Mutex{},
	}

	go teacher.RunBehavior(questionChan, waitAnswer, &whoAnswer, &round, numOfStudents)
	for _, student := range students {
		go student.RunBehavior(questionChan, waitAnswer, &whoAnswer, &round)
	}

	// Prevent main process interrupt
	time.Sleep(time.Second * 60)
}
