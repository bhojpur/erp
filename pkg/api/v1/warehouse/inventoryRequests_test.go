package warehouse

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

func TestSaveInventoryRegistration(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode":              "someclient",
			"sessionKey":              "somesess",
			"request":                 "saveInventoryRegistration",
			"inventoryRegistrationID": "12345",
			"creatorID":               "2234",
		})

		resp := SaveInventoryRegistrationResponse{
			Status:  sharedCommon.Status{ResponseStatus: "ok"},
			Results: []SaveInventoryRegistrationResult{{InventoryRegistrationID: 999}},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"inventoryRegistrationID": "12345",
		"creatorID":               "2234",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	regisrationID, err := cl.SaveInventoryRegistration(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, 999, regisrationID)
}

func TestSaveInventoryRegistrationBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"inventoryRegistrationID": "123",
				"creatorID":               "2",
				"requestName":             "saveInventoryRegistration",
			},
			{
				"warehouseID":   "334",
				"stocktakingID": "233",
				"supplierID":    "455",
				"requestName":   "saveInventoryRegistration",
			},
		})

		bulkResp := SaveInventoryRegistrationResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []SaveInventoryRegistrationBulkItem{
				{
					Status:  statusBulk,
					Results: []SaveInventoryRegistrationResult{{InventoryRegistrationID: 3456}},
				},
				{
					Status:  statusBulk,
					Results: []SaveInventoryRegistrationResult{{InventoryRegistrationID: 3457}},
				},
			},
		}
		jsonRaw, err := json.Marshal(bulkResp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := []map[string]interface{}{
		{
			"inventoryRegistrationID": "123",
			"creatorID":               "2",
		},
		{
			"warehouseID":   "334",
			"stocktakingID": "233",
			"supplierID":    "455",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.SaveInventoryRegistrationBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 2)
	assert.Equal(t, []SaveInventoryRegistrationResult{{InventoryRegistrationID: 3456}}, bulkResp.BulkItems[0].Results)
	assert.Equal(t, []SaveInventoryRegistrationResult{{InventoryRegistrationID: 3457}}, bulkResp.BulkItems[1].Results)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
}

func TestSaveInventoryWriteOff(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode":              "someclient",
			"sessionKey":              "somesess",
			"request":                 "saveInventoryWriteOff",
			"inventoryRegistrationID": "12345",
			"creatorID":               "2234",
		})

		resp := SaveInventoryWriteOffResponse{
			Status:  sharedCommon.Status{ResponseStatus: "ok"},
			Results: []SaveInventoryWriteOffResult{{InventoryWriteOffID: 999}},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"inventoryRegistrationID": "12345",
		"creatorID":               "2234",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	writeOffID, err := cl.SaveInventoryWriteOff(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, 999, writeOffID)
}

func TestSaveInventoryTransfer(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode":              "someclient",
			"sessionKey":              "somesess",
			"request":                 "saveInventoryTransfer",
			"inventoryRegistrationID": "12345",
			"creatorID":               "2234",
		})

		resp := SaveInventoryTransferResponse{
			Status:  sharedCommon.Status{ResponseStatus: "ok"},
			Results: []SaveInventoryTransferResult{{InventoryTransferID: 999}},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"inventoryRegistrationID": "12345",
		"creatorID":               "2234",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	TransferID, err := cl.SaveInventoryTransfer(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, 999, TransferID)
}
