package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Teacher struct {
	MathSymbols []string
}

func NewTeacher() *Teacher {
	return &Teacher{
		MathSymbols: []string{"+", "-", "*", "/"},
	}
}

func (t *Teacher) WarmUp() {
	fmt.Println("Teacher: Guys, are you ready?")
	time.Sleep(time.Second * 3)
}

func (t *Teacher) AskMathQuestion() Question {
	return Question{
		A: rand.Intn(101),
		B: rand.Intn(101),
		C: t.MathSymbols[rand.Intn(len(t.MathSymbols))],
	}
}

func (t *Teacher) RunBehavior(questionChan chan Question, waitAnswer *sync.WaitGroup, whoAnswer *WhoAnswer, round *int, numOfStudents int) {
	for ; ; *round++ {
		fmt.Printf("# Round %d start\n", *round)

		t.WarmUp()
		question := t.AskMathQuestion()
		for i := 0; i < numOfStudents; i++ {
			questionChan <- question
		}

		waitAnswer.Add(numOfStudents)
		waitAnswer.Wait()

		// Clear status
		whoAnswer.Name = ""
		fmt.Printf("# Round %d end\n", *round)
		time.Sleep(time.Second * 3)
	}
}
