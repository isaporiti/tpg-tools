package counter

import (
	"fmt"
	"io"
	"os"
	"time"
)

type counter struct {
	count    uint
	writer   io.Writer
	interval Interval
}

type Interval interface {
	Sleep()
}

type interval struct {
	duration time.Duration
}

func NewInterval(duration time.Duration) *interval {
	return &interval{
		duration: duration,
	}
}

func (s *interval) Sleep() {
	time.Sleep(s.duration)
}

func NewCounter(options ...option) (*counter, error) {
	c := &counter{
		writer:   os.Stdout,
		interval: NewInterval(0 * time.Nanosecond),
	}

	for _, opt := range options {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *counter) Next() (next uint) {
	next = c.count
	c.count++
	return next
}

func (c *counter) Run(times uint) {
	for i := 0; i < int(times); i++ {
		c.interval.Sleep()
		fmt.Fprintln(c.writer, c.Next())
	}
}
