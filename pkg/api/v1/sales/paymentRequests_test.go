package sales

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
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	sharedCommon "github.com/bhojpur/erp/pkg/api/v1/common"
	"github.com/bhojpur/erp/pkg/internal/common"
	"github.com/stretchr/testify/assert"
)

func TestGetPaymentsBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"requestName": "getPayments",
				"paymentID":   "1",
			},
			{
				"requestName": "getPayments",
				"paymentID":   "2",
			},
			{
				"requestName": "getPayments",
				"paymentID":   "3",
			},
		})

		bulkResp := GetPaymentsResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []GetPaymentsBulkItem{
				{
					Status: statusBulk,
					PaymentInfos: []PaymentInfo{
						{
							DocumentID: 1,
							Type:       "Some type",
						},
					},
				},
				{
					Status: statusBulk,
					PaymentInfos: []PaymentInfo{
						{
							DocumentID: 2,
							Type:       "Some type",
						},
					},
				},
				{
					Status: statusBulk,
					PaymentInfos: []PaymentInfo{
						{
							DocumentID: 3,
							Type:       "Some type",
						},
					},
				},
			},
		}
		jsonRaw, err := json.Marshal(bulkResp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.GetPaymentsBulk(
		context.Background(),
		[]map[string]interface{}{
			{
				"paymentID": "1",
			},
			{
				"paymentID": "2",
			},
			{
				"paymentID": "3",
			},
		},
		map[string]string{},
	)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 3)

	assert.Equal(t, []PaymentInfo{
		{
			DocumentID: 1,
			Type:       "Some type",
		},
	}, bulkResp.BulkItems[0].PaymentInfos)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)

	assert.Equal(t, []PaymentInfo{
		{
			DocumentID: 2,
			Type:       "Some type",
		},
	}, bulkResp.BulkItems[1].PaymentInfos)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)

	assert.Equal(t, []PaymentInfo{
		{
			DocumentID: 3,
			Type:       "Some type",
		},
	}, bulkResp.BulkItems[2].PaymentInfos)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[2].Status)
}

func TestSavePaymentsBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"requestName": "savePayment",
				"customerID":  "1",
			},
			{
				"requestName": "savePayment",
				"customerID":  "2",
			},
			{
				"requestName": "savePayment",
				"customerID":  "3",
			},
		})

		bulkResp := SavePaymentsResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []SavePaymentsBulkItem{
				{
					Status: statusBulk,
					Records: []SavePaymentID{
						{
							PaymentID: 1,
						},
					},
				},
				{
					Status: statusBulk,
					Records: []SavePaymentID{
						{
							PaymentID: 2,
						},
					},
				},
				{
					Status: statusBulk,
					Records: []SavePaymentID{
						{
							PaymentID: 3,
						},
					},
				},
			},
		}
		jsonRaw, err := json.Marshal(bulkResp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.SavePaymentsBulk(
		context.Background(),
		[]map[string]interface{}{
			{
				"customerID": "1",
			},
			{
				"customerID": "2",
			},
			{
				"customerID": "3",
			},
		},
		map[string]string{},
	)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 3)

	assert.Equal(t, []SavePaymentID{
		{
			PaymentID: 1,
		},
	}, bulkResp.BulkItems[0].Records)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)

	assert.Equal(t, []SavePaymentID{
		{
			PaymentID: 2,
		},
	}, bulkResp.BulkItems[1].Records)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)

	assert.Equal(t, []SavePaymentID{
		{
			PaymentID: 3,
		},
	}, bulkResp.BulkItems[2].Records)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[2].Status)
}
