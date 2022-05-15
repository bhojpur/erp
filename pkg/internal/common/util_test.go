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
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSendRequestInBody(t *testing.T) {
	calledTimes := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calledTimes++
		assert.Equal(t, "", r.URL.Query().Get("clientCode"))
		assert.Equal(t, "", r.URL.Query().Get("sessionKey"))
		assert.Equal(t, "", r.URL.Query().Get("someKey"))
		AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
			"someKey":    "someValue",
		})
	}))
	defer srv.Close()

	cli := NewClientWithURL(
		"somesess",
		"someclient",
		"",
		srv.URL,
		&http.Client{
			Timeout: 5 * time.Second,
		},
		nil,
	)
	cli.SendParametersInRequestBody()

	resp, err := cli.SendRequest(
		context.Background(),
		"getSuppliers",
		map[string]string{"someKey": "someValue"},
	)
	assert.NoError(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, calledTimes)
}

func TestSendRequestInQuery(t *testing.T) {
	calledTimes := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calledTimes++
		assert.Equal(t, "someclient", r.URL.Query().Get("clientCode"))
		assert.Equal(t, "somesess", r.URL.Query().Get("sessionKey"))
		assert.Equal(t, "someValue", r.URL.Query().Get("someKey"))
		AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
			"someKey":    "someValue",
		})
	}))
	defer srv.Close()

	cli := NewClientWithURL(
		"somesess",
		"someclient",
		"",
		srv.URL,
		&http.Client{
			Timeout: 5 * time.Second,
		},
		nil,
	)

	resp, err := cli.SendRequest(
		context.Background(),
		"getSuppliers",
		map[string]string{"someKey": "someValue"},
	)
	assert.NoError(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, calledTimes)
}

func TestSendRequestBulk(t *testing.T) {
	calledTimes := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calledTimes++

		AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
			"someKey":    "someValue",
		})

		AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"requestName":   "getSuppliers",
				"recordsOnPage": "10",
				"pageNo":        "1",
			},
			{
				"requestName":   "getSuppliers",
				"recordsOnPage": "10",
				"pageNo":        "2",
			},
		})
	}))

	defer srv.Close()

	cli := NewClientWithURL(
		"somesess",
		"someclient",
		"",
		srv.URL,
		&http.Client{
			Timeout: 5 * time.Second,
		},
		nil,
	)

	resp, err := cli.SendRequestBulk(
		context.Background(),
		[]BulkInput{
			{
				MethodName: "getSuppliers",
				Filters: map[string]interface{}{
					"recordsOnPage": "10",
					"pageNo":        "1",
				},
			},
			{
				MethodName: "getSuppliers",
				Filters: map[string]interface{}{
					"recordsOnPage": "10",
					"pageNo":        "2",
				},
			},
		},
		map[string]string{"someKey": "someValue"},
	)
	assert.NoError(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, calledTimes)
}
