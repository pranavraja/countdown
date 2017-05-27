package countdown

import (
	"testing"
	"time"
)

func TestCountdown_CountAt(t *testing.T) {
	frozenTime := time.Unix(1234567890, 0)
	halfASecondLater := frozenTime.Add(500 * time.Millisecond)
	oneSecondLater := frozenTime.Add(1 * time.Second)
	oneAndAHalfSecondsLater := frozenTime.Add(1500 * time.Millisecond)
	twoSecondsLater := frozenTime.Add(2 * time.Second)
	tests := []struct {
		name              string
		total             int64
		start             time.Time
		eventTimes        []time.Time
		recencyWeight     float64
		expectedRate      time.Duration
		expectedRemaining int64
	}{
		{"basic countdown 1 per sec", 10, frozenTime, []time.Time{oneSecondLater, twoSecondsLater}, 1, time.Second, 8},
		{"basic countdown 2 per sec", 10, frozenTime, []time.Time{halfASecondLater, oneSecondLater, oneAndAHalfSecondsLater, twoSecondsLater}, 1, 500 * time.Millisecond, 6},
		{"speeding up", 10, frozenTime, []time.Time{oneSecondLater, oneAndAHalfSecondsLater}, 1, 1 * time.Second, 8},
		{"speeding up with recency weight", 10, frozenTime, []time.Time{oneSecondLater, oneAndAHalfSecondsLater}, 0.5, 750 * time.Millisecond, 8},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cd := countdown{
				Total:     test.total,
				Remaining: test.total,
				Weight:    test.recencyWeight,
				last:      test.start,
			}
			for _, evt := range test.eventTimes {
				cd.CountAt(evt)
			}
			if cd.EstimatedRate != test.expectedRate {
				t.Errorf("got rate %s, wanted %s", cd.EstimatedRate, test.expectedRate)
			}
			if cd.Remaining != test.expectedRemaining {
				t.Errorf("got remaining %d, wanted %d", cd.Remaining, test.expectedRemaining)
			}
		})
	}
}
