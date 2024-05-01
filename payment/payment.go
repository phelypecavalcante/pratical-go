package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	p := Payment{
		From:   "Wile. E. Coyote",
		To:     "ACME",
		Amount: 123.34,
	}

	p.Process()
	p.Process()
}

type Payment struct {
	From   string
	To     string
	Amount float64 // USD

	once sync.Once
}

func (p *Payment) Process() {
	p.once.Do(func() {
		p.process(time.Now())
	})
}

func (p *Payment) process(t time.Time) {
	ts := t.Format(time.RFC3339)
	fmt.Printf("[%s] %s -> $%.2f -> %s\n", ts, p.From, p.Amount, p.To)
}
