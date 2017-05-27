package countdown

import (
	"time"
)

type countdown struct {
	Total         int64
	Remaining     int64
	EstimatedRate time.Duration
	Weight        float64
	Samples       int64

	last time.Time
}

func (c *countdown) Count() {
	c.CountAt(time.Now())
}

func (c *countdown) CountAt(t time.Time) {
	c.Remaining--
	lastEstimatedRate := t.Sub(c.last)
	c.last = t
	if c.EstimatedRate == 0 {
		c.EstimatedRate = lastEstimatedRate
	} else {
		c.EstimatedRate = time.Duration((1-c.Weight)*float64(lastEstimatedRate) + c.Weight*float64(c.EstimatedRate))
	}
}

func (c *countdown) Log() {
	if c.Remaining%(c.Total/c.Samples) == 0 {
		println(100*(c.Total-c.Remaining)/c.Total, "% done -", (c.EstimatedRate * time.Duration(c.Remaining)).String(), "remaining")
	}
}

func NewAt(t time.Time, total int64) *countdown {
	return &countdown{
		Total:     total,
		Remaining: total,
		Weight:    0.2,
		Samples:   100,

		last: t,
	}
}

func New(total int64) *countdown {
	return NewAt(time.Now(), total)
}
