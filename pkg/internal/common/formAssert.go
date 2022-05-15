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
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertFormValues(t *testing.T, r *http.Request, expectedValues map[string]interface{}) {
	for expectedKey, expectedValue := range expectedValues {
		actualValue := r.FormValue(expectedKey)
		assert.Equal(t, expectedValue, actualValue)
	}
}

func AssertRequestBulk(t *testing.T, r *http.Request, expectedBulkRequest []map[string]interface{}) {
	requestsRaw := r.FormValue("requests")
	decodedValue, err := url.QueryUnescape(requestsRaw)
	assert.NoError(t, err)

	var actualBulkRequest []map[string]interface{}
	err = json.Unmarshal([]byte(decodedValue), &actualBulkRequest)
	assert.NoError(t, err)

	assert.Equal(t, expectedBulkRequest, actualBulkRequest)
}
