package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Student struct {
	Name string
}

func NewStudent(name string) *Student {
	return &Student{
		Name: name,
	}
}

func (s *Student) RunBehavior(questionChan chan Question, waitAnswer *sync.WaitGroup, whoAnswer *WhoAnswer, round *int) {
	curRound := 0
	for {
		if *round <= curRound {
			continue
		}
		curRound = *round
		question := <-questionChan
		s.SeeAndThink(question)

		for {
			if whoAnswer.Mutex.TryLock() {
				if whoAnswer.Name == "" {
					whoAnswer.Name = s.Name
					s.Answer(question)
				} else {
					fmt.Printf("Student %s: %s, you win.\n", s.Name, whoAnswer.Name)
				}
				whoAnswer.Mutex.Unlock()
				break
			}
		}

		waitAnswer.Done()
	}
}

func (s *Student) SeeAndThink(question Question) {
	thinkTime := time.Duration(rand.Intn(3)+1) * time.Second
	time.Sleep(thinkTime)
}

func (s *Student) Answer(q Question) {
	fmt.Printf("Student %s: ", s.Name)
	var ans int
	switch q.C {
	case "+":
		ans = q.A + q.B
	case "-":
		ans = q.A - q.B
	case "*":
		ans = q.A * q.B
	case "/":
		if q.B == 0 {
			fmt.Println("Cannot divide by zero!")
			return
		}
		ans = q.A / q.B
	}
	fmt.Printf("%d %s %d = %d!\n", q.A, q.C, q.B, ans)
}
