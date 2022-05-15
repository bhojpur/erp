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
	"time"

	"github.com/bhojpur/erp/pkg/api/v1/log"
)

type Waiter interface {
	Wait(dur time.Duration)
}

type SleepWaiter struct{}

func (sw SleepWaiter) Wait(dur time.Duration) {
	time.Sleep(dur)
}

type Connector struct {
	SessionCleaner        func() error
	Connect               func() (err error)
	AttemptsCount         uint
	Waiter                Waiter
	WaitingInterval       time.Duration
	WaitingIncrementCoeff uint
}

func (c Connector) Run() error {
	var i uint
	var err error
	if c.Waiter == nil {
		c.Waiter = SleepWaiter{}
	}
	for i = 0; i < c.AttemptsCount; i++ {
		err = c.Connect()
		if err == nil {
			return nil
		}

		switch e := err.(type) {
		case *ErpError:
			if e.Code == APISessionExpired {
				log.Log.Log(log.Error, "failed to connect because auth session is expired: %v", err)
				log.Log.Log(log.Debug, "will invalidate session")
				err := c.SessionCleaner()
				if err != nil {
					return err
				}
			}
		default:
			log.Log.Log(log.Error, "failed to connect: %v", err)
		}

		log.Log.Log(log.Debug, "will retry the connection attempt after %v sleep interval, connections attempts left: %d", c.WaitingInterval, c.AttemptsCount-i)
		c.Waiter.Wait(c.WaitingInterval * time.Duration(c.WaitingIncrementCoeff*i+1))
	}

	return err
}
