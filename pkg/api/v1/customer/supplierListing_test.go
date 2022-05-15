package customer

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
	"fmt"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"
	"time"

	sharedCommon "github.com/bhojpur/erp/pkg/api/v1/common"
	"github.com/bhojpur/erp/pkg/internal/common"
	"github.com/stretchr/testify/assert"
)

func sendRequest(w http.ResponseWriter, errStatus sharedCommon.ApiError, totalCount int, supplierIDsBulk [][]int) error {
	bulkResp := GetSuppliersResponseBulk{
		Status: sharedCommon.Status{ResponseStatus: "ok"},
	}

	bulkItems := make([]GetSuppliersResponseBulkItem, 0, len(supplierIDsBulk))
	for _, supplierIDs := range supplierIDsBulk {
		suppliers := make([]Supplier, 0, len(supplierIDs))
		for _, id := range supplierIDs {
			suppliers = append(suppliers, Supplier{
				SupplierId: uint(id),
				FullName:   fmt.Sprintf("Some Supplier %d", id),
			})
		}
		statusBulk := sharedCommon.StatusBulk{}
		if errStatus == 0 {
			statusBulk.ResponseStatus = "ok"
		} else {
			statusBulk.ResponseStatus = "not ok"
		}
		statusBulk.RecordsTotal = totalCount
		statusBulk.ErrorCode = errStatus
		statusBulk.RecordsInResponse = len(supplierIDs)

		bulkItems = append(bulkItems, GetSuppliersResponseBulkItem{
			Status:    statusBulk,
			Suppliers: suppliers,
		})
	}
	bulkResp.BulkItems = bulkItems

	jsonRaw, err := json.Marshal(bulkResp)
	if err != nil {
		return err
	}

	_, err = w.Write(jsonRaw)
	if err != nil {
		return err
	}
	return nil
}

func TestListingCountSuccess(t *testing.T) {
	const totalCount = 10
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parsedRequest, err := common.ExtractBulkFiltersFromRequest(r)
		assert.NoError(t, err)
		if err != nil {
			return
		}

		assert.Equal(t, "someclient", parsedRequest["clientCode"])
		assert.Equal(t, "somesess", parsedRequest["sessionKey"])
		requests := parsedRequest["requests"].([]map[string]interface{})
		assert.Equal(t, float64(1), requests[0]["pageNo"])
		assert.Equal(t, float64(1), requests[0]["recordsOnPage"])
		assert.Equal(t, "getSuppliers", requests[0]["requestName"])
		assert.Equal(t, "smeval", requests[0]["somekey"])

		err = sendRequest(w, 0, totalCount, [][]int{{1}})
		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	defer srv.Close()

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	supplierClient := NewClient(baseClient)
	supplierDataProvider := NewSupplierListingDataProvider(supplierClient)

	actualCount, err := supplierDataProvider.Count(context.Background(), map[string]interface{}{"somekey": "smeval", "pageNo": 1, "recordsOnPage": 1})
	assert.NoError(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, totalCount, actualCount)
}

func TestListingCountError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := sendRequest(w, sharedCommon.MalformedRequest, 0, [][]int{{1}})
		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	defer srv.Close()

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	suppliersClient := NewClient(baseClient)
	suppliersDataProvider := NewSupplierListingDataProvider(suppliersClient)

	actualCount, err := suppliersDataProvider.Count(context.Background(), map[string]interface{}{"somekey": "smeval"})
	assert.Error(t, err)
	if err == nil {
		return
	}
	assert.Contains(t, err.Error(), sharedCommon.MalformedRequest.String())
	assert.Equal(t, 0, actualCount)
}

func TestListingCountWithNoBulkItems(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := sendRequest(w, 0, 0, [][]int{})
		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	defer srv.Close()

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	suppliersClient := NewClient(baseClient)
	suppliersDataProvider := NewSupplierListingDataProvider(suppliersClient)

	actualCount, err := suppliersDataProvider.Count(context.Background(), map[string]interface{}{"somekey": "smeval"})
	assert.NoError(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, 0, actualCount)
}

func TestReadSuccess(t *testing.T) {
	const limit = 2
	const offset = 1
	const totalCount = 10
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parsedRequest, err := common.ExtractBulkFiltersFromRequest(r)
		assert.NoError(t, err)
		if err != nil {
			return
		}

		assert.Equal(t, "someclient", parsedRequest["clientCode"])
		assert.Equal(t, "somesess", parsedRequest["sessionKey"])

		requests := parsedRequest["requests"].([]map[string]interface{})
		assert.Len(t, requests, 1)

		assert.Equal(t, float64(offset), requests[0]["pageNo"])
		assert.Equal(t, float64(limit), requests[0]["recordsOnPage"])

		assert.Equal(t, "getSuppliers", requests[0]["requestName"])
		assert.Equal(t, "smeval", requests[0]["somekey"])

		err = sendRequest(w, 0, totalCount, [][]int{{1, 2}, {3, 4}, {5}})
		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	defer srv.Close()

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	suppliersClient := NewClient(baseClient)
	suppliersDataProvider := NewSupplierListingDataProvider(suppliersClient)

	actualProdIDs := make([]int, 0, 5)
	err := suppliersDataProvider.Read(
		context.Background(),
		[]map[string]interface{}{
			{
				"somekey":       "smeval",
				"pageNo":        1,
				"recordsOnPage": 2,
			},
		},
		func(item interface{}) {
			assert.IsType(t, item, Supplier{})
			actualProdIDs = append(actualProdIDs, int(item.(Supplier).SupplierId))
		},
	)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, []int{1, 2, 3, 4, 5}, actualProdIDs)
}

func TestReadError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := sendRequest(w, sharedCommon.MalformedRequest, 10, [][]int{{1}})
		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	defer srv.Close()

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	suppliersClient := NewClient(baseClient)
	suppliersDataProvider := NewSupplierListingDataProvider(suppliersClient)

	err := suppliersDataProvider.Read(
		context.Background(),
		[]map[string]interface{}{{"somekey": "smeval"}},
		func(item interface{}) {},
	)
	assert.Error(t, err)
	if err == nil {
		return
	}

	assert.Contains(t, err.Error(), sharedCommon.MalformedRequest.String())
}

func TestReadSuccessIntegration(t *testing.T) {
	const totalCount = 11
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parsedRequest, err := common.ExtractBulkFiltersFromRequest(r)
		assert.NoError(t, err)
		if err != nil {
			return
		}

		requests := parsedRequest["requests"].([]map[string]interface{})
		assert.Len(t, requests, 1)

		if requests[0]["pageNo"] == float64(1) {
			err = sendRequest(w, 0, totalCount, [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}})
		} else {
			err = sendRequest(w, 0, totalCount, [][]int{{11}})
		}

		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	defer srv.Close()

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	suppliersClient := NewClient(baseClient)
	suppliersDataProvider := NewSupplierListingDataProvider(suppliersClient)

	lister := sharedCommon.NewLister(
		sharedCommon.ListingSettings{
			MaxRequestsCountPerSecond: 0,
			StreamBufferLength:        10,
			MaxItemsPerRequest:        10,
			MaxFetchersCount:          10,
		},
		suppliersDataProvider,
		func(sleepTime time.Duration) {},
	)

	prodsChan := lister.Get(context.Background(), map[string]interface{}{})

	actualProdIDs := collectProdIDsFromChannel(prodsChan)
	sort.Ints(actualProdIDs)

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, actualProdIDs)
}

func collectProdIDsFromChannel(prodsChan sharedCommon.ItemsStream) []int {
	actualProdIDs := make([]int, 0)
	doneChan := make(chan struct{}, 1)
	go func() {
		defer close(doneChan)
		for prod := range prodsChan {
			actualProdIDs = append(actualProdIDs, int(prod.Payload.(Supplier).SupplierId))
		}
	}()

mainLoop:
	for {
		select {
		case <-doneChan:
			break mainLoop
		case <-time.After(time.Second * 5):
			break mainLoop
		}
	}

	return actualProdIDs
}
