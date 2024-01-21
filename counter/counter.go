package counter

import (
	"fmt"
	"io"
	"os"

	"github.com/isaporiti/tpg-tools/interval"
)

type counter struct {
	count    uint
	writer   io.Writer
	interval interval.Interval
}

func NewCounter(options ...option) (*counter, error) {
	c := &counter{
		writer:   os.Stdout,
		interval: interval.NewNoOpInterval(),
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
