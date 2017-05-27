# Prerequisites

- [go 1.8](https://golang.org/doc/install)

# Setup

	go get github.com/pranavraja/countdown

# Usage

	cd := countdown.New(1000)
	for i := 0; i < 1000; i++ {
		cd.Count()
		cd.Log()
	}

By default it will actually log every 1% progress, i.e. 100 samples. You can change that by setting `cd.Samples`.

If you want a custom log format, you can just read `cd.Remaining`, `cd.Total` and `cd.EstimatedRate` yourself and have at it.

# Background

This uses an [exponential moving
average](https://en.wikipedia.org/wiki/Moving_average#Exponential_moving_average)
to account for fluctuations. The default recency weight is set to 0.2, which
means that the most recent value is not as significant in calculating the
moving average. You can override this by setting `cd.Weight`.
