package interval

import "time"

type Interval interface {
	Sleep()
}

type sleepInterval struct {
	duration time.Duration
}

func NewSleepInterval(duration time.Duration) *sleepInterval {
	return &sleepInterval{
		duration: duration,
	}
}

func (s *sleepInterval) Sleep() {
	time.Sleep(s.duration)
}

type noOpInterval struct{}

func (n *noOpInterval) Sleep() {}

func NewNoOpInterval() *noOpInterval {
	return &noOpInterval{}
}
