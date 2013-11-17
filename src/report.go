package gorinju

import (
	"fmt"
	"time"
)

type Report struct {
	url string
	time time.Duration
}

func (rep Report) PrintURL () {
	fmt.Printf("Target URL :: %s\n", rep.url)
}

func (rep Report) PrintTime () {
	fmt.Printf("Time       :: %s\n", rep.time.String())
}

func (rep Report) Print() {
	rep.PrintURL()
	rep.PrintTime()
}
