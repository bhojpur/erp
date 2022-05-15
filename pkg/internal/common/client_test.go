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
	"net/url"
	"strings"
	"testing"
)

func TestClientCodeInsideClientsURL(t *testing.T) {
	authFunc := func(requestName string) url.Values {
		params := url.Values{}
		params.Add("request", requestName)
		params.Add(clientCode, "1234")
		return params
	}
	t.Run("auth func case", func(t *testing.T) {
		c := NewClient("", "", "", nil, authFunc)
		if !strings.Contains(c.Url, "1234") {
			t.Error(c.clientCode)
			return
		}
	})

	t.Run("client code provided case", func(t *testing.T) {
		c := NewClient("", "1234", "", nil, nil)
		if !strings.Contains(c.Url, "1234") {
			t.Error(c.clientCode)
			return
		}
	})
}
