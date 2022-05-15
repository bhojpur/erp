package api

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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

//ClientMock mocks HttpClient interface
type ClientMock struct {
	ErrToGive      error
	ResponseToGive *http.Response
	Requests       []*http.Request
	Lock           sync.Mutex
}

//Do HttpClient interface implementation
func (cm *ClientMock) Do(req *http.Request) (*http.Response, error) {
	cm.Lock.Lock()
	defer cm.Lock.Unlock()

	cm.Requests = append(cm.Requests, req)

	return cm.ResponseToGive, cm.ErrToGive
}

//BodyMock implements resp.Body
type BodyMock struct {
	Body       io.Reader
	WasClosed  bool
	CloseError error
}

//NewFromStr creates BodyMock from a string
func NewFromStr(body string) *BodyMock {
	return &BodyMock{
		Body:       strings.NewReader(body),
		WasClosed:  false,
		CloseError: nil,
	}
}

//NewFromStruct creates BodyMock from a struct converted to json or string
func NewFromStruct(input interface{}) *BodyMock {
	c, err := json.Marshal(input)
	if err != nil {
		return &BodyMock{
			Body:       strings.NewReader(fmt.Sprintf("%+v", input)),
			WasClosed:  false,
			CloseError: nil,
		}
	}
	buf := bytes.NewBuffer(c)
	return &BodyMock{
		Body:       buf,
		WasClosed:  false,
		CloseError: nil,
	}
}

//Read io.Reader implementation
func (bm *BodyMock) Read(p []byte) (n int, err error) {
	return bm.Body.Read(p)
}

//Close io.Closer implementation
func (bm *BodyMock) Close() error {
	bm.WasClosed = true
	return bm.CloseError
}
