package main

import (
	"container/ring"
	"time"
)

type proxyStats struct {
	SampleRate int //seconds
	Samples    int //total number of samples
	Ring       *ring.Ring
	InChan     chan int
	DoneChan   chan bool
	counter    int // tracks counts when sampling
}

type statsSummary struct {
	Data      []int     `json:"data"`
	Timestamp time.Time `json:"timestamp"`
	Interval  int       `json:"interval"`
}

func NewProxyStats(samples int, sampleRate int) *proxyStats {
	r := ring.New(samples)

	inChan := make(chan int)
	doneChan := make(chan bool)

	return &proxyStats{
		Samples:    samples,
		SampleRate: sampleRate,
		Ring:       r,
		InChan:     inChan,
		DoneChan:   doneChan,
	}
}

func (pr *proxyStats) Insert(count int) {
	pr.InChan <- count
}

func (pr *proxyStats) Stats() *statsSummary {

	return &statsSummary{}
}

func (pr *proxyStats) Track() {
	ticker := time.NewTicker(time.Second * time.Duration(pr.SampleRate))
	go func(in chan int, done chan bool) {
		for {
			select {
			case <-ticker.C:
				pr.Ring.Value = pr.counter
				pr.Ring.Next()
				pr.counter = 0
			case <-in:
				pr.counter++
			case <-done:
				return
			}
		}
	}(pr.InChan, pr.DoneChan)
}
