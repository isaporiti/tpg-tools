package counter

import (
	"fmt"
	"io"
	"os"
)

type counter struct {
	count  uint
	writer io.Writer
}

func NewCounter(options ...option) (*counter, error) {
	c := &counter{
		writer: os.Stdout,
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
		fmt.Fprintln(c.writer, c.Next())
	}
}
