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
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type WaiterMock struct {
	WaitingDurations []time.Duration
}

func (wm *WaiterMock) Wait(dur time.Duration) {
	wm.WaitingDurations = append(wm.WaitingDurations, dur)
}

func TestSingleConnectSuccess(t *testing.T) {
	sessionCleanerWasCalled := false
	connectAmount := 0
	w := &WaiterMock{}
	const connectsCount = 5
	c := Connector{
		SessionCleaner: func() error {
			sessionCleanerWasCalled = true
			return nil
		},
		Connect: func() (err error) {
			connectAmount++
			return
		},
		AttemptsCount: connectsCount,
		Waiter:        w,
	}

	err := c.Run()
	assert.NoError(t, err)

	assert.False(t, sessionCleanerWasCalled)

	assert.Equal(t, 1, connectAmount)

	assert.Len(t, w.WaitingDurations, 0)
}

func TestSessionExpiration(t *testing.T) {
	sessionCleanerWasCalled := false
	connectAmount := 0
	w := &WaiterMock{}
	const connectsCount = 5
	c := Connector{
		SessionCleaner: func() error {
			sessionCleanerWasCalled = true
			return nil
		},
		Connect: func() (err error) {
			if connectAmount == 0 {
				err = &ErpError{
					error:   errors.New("conn failure"),
					Status:  "Some status",
					Message: "Some message",
					Code:    APISessionExpired,
				}
			}
			connectAmount++
			return
		},
		AttemptsCount:         connectsCount,
		Waiter:                w,
		WaitingInterval:       time.Second,
		WaitingIncrementCoeff: 10,
	}

	err := c.Run()
	if err != nil {
		t.Error(err)
	}

	assert.True(t, sessionCleanerWasCalled)
	assert.Equal(t, 2, connectAmount)
	assert.Len(t, w.WaitingDurations, 1)
	assert.Equal(t, time.Second, w.WaitingDurations[0])
}

func TestMultipleConnectFailure(t *testing.T) {
	sessionCleanerWasCalled := false
	connectAmount := 0
	w := &WaiterMock{}
	const connectsCount = 5
	c := Connector{
		SessionCleaner: func() error {
			sessionCleanerWasCalled = true
			return nil
		},
		Connect: func() (err error) {
			connectAmount++
			err = errors.New("conn failure")
			return
		},
		AttemptsCount:         connectsCount,
		Waiter:                w,
		WaitingInterval:       time.Second,
		WaitingIncrementCoeff: 2,
	}

	err := c.Run()
	assert.Error(t, err, "conn failure")
	assert.False(t, sessionCleanerWasCalled)
	assert.Equal(t, 5, connectAmount)
	assert.Len(t, w.WaitingDurations, 5)
	assert.Equal(t, time.Second, w.WaitingDurations[0])
	assert.Equal(t, time.Second*3, w.WaitingDurations[1])
	assert.Equal(t, time.Second*5, w.WaitingDurations[2])
	assert.Equal(t, time.Second*7, w.WaitingDurations[3])
	assert.Equal(t, time.Second*9, w.WaitingDurations[4])
}

func TestSessionCleanFailure(t *testing.T) {
	connectAmount := 0
	w := &WaiterMock{}
	const connectsCount = 5
	c := Connector{
		SessionCleaner: func() error {
			return errors.New("some session failure")
		},
		Connect: func() (err error) {
			if connectAmount == 0 {
				err = &ErpError{
					error:   errors.New("conn failure"),
					Status:  "Some status",
					Message: "Some message",
					Code:    APISessionExpired,
				}
			}
			connectAmount++
			return
		},
		AttemptsCount: connectsCount,
		Waiter:        w,
	}

	err := c.Run()
	assert.Error(t, err, "some session failure")
	assert.Equal(t, 1, connectAmount)
	assert.Len(t, w.WaitingDurations, 0)
}
