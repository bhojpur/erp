package common

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"sync"
	"time"
)

//Throttler abstracts limiting of API requests
type Throttler interface {
	Throttle()
}

type Sleeper func(sleepTime time.Duration)

var sleepThrottler *SleepThrottler

//SleepThrottler implements sleeping logic for requests throttling
type SleepThrottler struct {
	LimitPerSecond int
	LastTimestamp  int64
	Count          int
	sl             Sleeper
	lock           sync.Mutex
}

//NewSleepThrottler creates SleepThrottler
func NewSleepThrottler(limitPerSecond int, sl Sleeper) *SleepThrottler {
	if sleepThrottler == nil {
		sleepThrottler = &SleepThrottler{
			LimitPerSecond: limitPerSecond,
			LastTimestamp:  time.Now().Unix(),
			Count:          0,
			sl:             sl,
			lock:           sync.Mutex{},
		}
	}

	return sleepThrottler
}

//Throttle implements throttling method
func (rt *SleepThrottler) Throttle() {
	rt.lock.Lock()
	defer rt.lock.Unlock()

	if rt.LimitPerSecond <= 0 {
		return
	}

	rt.Count++
	now := time.Now().Unix()
	if now != rt.LastTimestamp {
		rt.LastTimestamp = now
		rt.Count = 1
		return
	}

	if rt.Count >= rt.LimitPerSecond {
		rt.sl(time.Second)
	}
}

type ThrottlerMock struct {
	WasTriggered bool
}

func (tm *ThrottlerMock) Throttle() {
	tm.WasTriggered = true
}
