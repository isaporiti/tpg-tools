package counter

import (
	"bytes"
	"testing"
	"time"

	"github.com/isaporiti/tpg-tools/interval"
)

func TestCounter_NewCounter(t *testing.T) {
	t.Run("NewCounter returns a pointer to a counter", func(t *testing.T) {
		t.Parallel()

		c, _ := NewCounter()

		if c == nil {
			t.Errorf("Expected a new counter, got %v", c)
		}
	})

	t.Run("NewCounter accepts an initial value and sets it to the to the counter", func(t *testing.T) {
		t.Parallel()
		c, _ := NewCounter(WithInitialCount(5))

		got := c.Next()
		if got != 5 {
			t.Errorf("Expected 5, got %v", c.count)
		}
	})

	t.Run("NewCounter accepts an optional io.Writer", func(t *testing.T) {
		t.Parallel()
		writer := &bytes.Buffer{}
		c, _ := NewCounter(WithWriter(writer))

		c.Run(1)

		if writer.String() == "" {
			t.Errorf("Expected writer to be written to.")
		}
	})
}

func TestCounter_Next(t *testing.T) {
	t.Run("Next should return 0 on first call", func(t *testing.T) {
		t.Parallel()
		c, _ := NewCounter()

		got := c.Next()

		if got != 0 {
			t.Errorf("Expected 0, got %v", got)
		}
	})

	t.Run("Next should return 0, 1, 2, 3 and so on on subsequent calls", func(t *testing.T) {
		t.Parallel()
		c, _ := NewCounter()

		for want := uint(0); want < 10; want++ {
			got := c.Next()
			if got != want {
				t.Errorf("Expected %v, got %v", want, got)
				return
			}
		}
	})
}

func TestCounter_Run(t *testing.T) {
	t.Run("Run accpets a number of iterations and prints the result of Next that many times", func(t *testing.T) {
		t.Parallel()
		writer := &bytes.Buffer{}
		c, _ := NewCounter(WithWriter(writer))

		c.Run(5)

		got := writer.String()
		want := "0\n1\n2\n3\n4\n"
		if got != want {
			t.Errorf("Expected '%v', got '%v'", want, got)
		}
	})

	t.Run("Run prints the succession of numbers with time intervals", func(t *testing.T) {
		t.Parallel()

		writer := &bytes.Buffer{}
		interval := newStubInterval(1 * time.Millisecond)
		c, _ := NewCounter(WithWriter(writer), WithInterval(interval))

		c.Run(3)

		got := interval.(*stubInterval).calls
		if got != 3 {
			t.Errorf("Expected 3 calls to interval.Sleep, got %v", got)
		}
	})

}

type stubInterval struct {
	duration time.Duration
	calls    int
}

func (s *stubInterval) Sleep() {
	s.calls++
}

func newStubInterval(d time.Duration) interval.Interval {
	return &stubInterval{duration: d}
}
