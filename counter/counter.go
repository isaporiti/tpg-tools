package counter

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type counter struct {
	count  int
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

type option func(*counter) error

func WithInitialCount(value int) option {
	return func(c *counter) error {
		if value < 0 {
			return ErrNoNegativeValues
		}

		c.count = value
		return nil
	}
}

func WithWriter(w io.Writer) option {
	return func(c *counter) error {
		c.writer = w
		return nil
	}
}

var ErrNoNegativeValues = errors.New("no negative values allowed")

func (c *counter) Next() (next int) {
	next = c.count
	c.count++
	return next
}

func (c *counter) Run(times uint) {
	for i := 0; i < int(times); i++ {
		fmt.Fprintln(c.writer, c.Next())
	}
}
