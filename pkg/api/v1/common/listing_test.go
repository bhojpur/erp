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
	"errors"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var NullSleeper = func(sleepTime time.Duration) {}

type payloadMock struct {
	ID int
}

type DataProviderMock struct {
	countLock           sync.Mutex
	CountContextInput   context.Context
	CountFiltersInput   map[string]interface{}
	CountOutputCount    int
	CountOutputErrorStr string

	readLock         sync.Mutex
	ReadContextInput context.Context
	ProductsToRead   []payloadMock
	ReadErrorStr     string
	ReadBulkFilters  [][]map[string]interface{}
}

func (dpm *DataProviderMock) Count(ctx context.Context, filters map[string]interface{}) (int, error) {
	dpm.countLock.Lock()
	defer dpm.countLock.Unlock()

	dpm.CountContextInput = ctx
	dpm.CountFiltersInput = filters

	var err error
	if dpm.CountOutputErrorStr != "" {
		err = errors.New(dpm.CountOutputErrorStr)
	}

	return dpm.CountOutputCount, err
}

func (dpm *DataProviderMock) Read(ctx context.Context, bulkFilters []map[string]interface{}, callback func(item interface{})) error {
	dpm.readLock.Lock()
	defer dpm.readLock.Unlock()

	dpm.ReadContextInput = ctx
	dpm.ReadBulkFilters = append(dpm.ReadBulkFilters, bulkFilters)

	for _, prod := range dpm.ProductsToRead {
		callback(prod)
	}

	var err error
	if dpm.ReadErrorStr != "" {
		err = errors.New(dpm.ReadErrorStr)
	}

	return err
}

func TestReadingSuccess(t *testing.T) {
	testCases := []struct {
		name                     string
		total                    int
		inputProds               []payloadMock
		listingSettings          ListingSettings
		expectedBulkFilterInputs func() [][]map[string]interface{}
		expectedProdIDs          []int
	}{
		{
			name:  "too small request limit",
			total: 10,
			inputProds: []payloadMock{
				{ID: 1},
				{ID: 2},
			},
			/**
			the DataProviderMock is dummy and gives always 1,2 at every request, so it is expected to be called 5 times
			(total 10 / MaxItemsPerRequest 2 which will generate 5 x {1,2} id responses which are expected here
			*/
			expectedProdIDs: []int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2},
			listingSettings: ListingSettings{
				MaxRequestsCountPerSecond: 0,
				StreamBufferLength:        0,
				MaxItemsPerRequest:        2,
				MaxFetchersCount:          2,
			},
			expectedBulkFilterInputs: func() (res [][]map[string]interface{}) {
				expectedBulkRequestsCount := 5
				res = make([][]map[string]interface{}, 0, expectedBulkRequestsCount)
				expectedBulkInputCount := 1
				for i := 0; i < expectedBulkRequestsCount; i++ {
					bulkItems := make([]map[string]interface{}, 0, expectedBulkInputCount)
					for y := 0; y < expectedBulkInputCount; y++ {
						bulkItems = append(bulkItems, map[string]interface{}{
							"filterKey":     "filterVal",
							"pageNo":        i + 1,
							"recordsOnPage": 2,
						})
					}
					res = append(res, bulkItems)
				}
				return
			},
		},
		{
			name:  "max request limit",
			total: 10001,
			inputProds: []payloadMock{
				{ID: 3},
			},
			expectedProdIDs: []int{3, 3},
			listingSettings: ListingSettings{
				MaxRequestsCountPerSecond: 0,
				StreamBufferLength:        10,
				MaxItemsPerRequest:        10000,
				MaxFetchersCount:          2,
			},
			expectedBulkFilterInputs: func() (res [][]map[string]interface{}) {
				res = make([][]map[string]interface{}, 0, 2)
				bulkItems := make([]map[string]interface{}, 0, 100)
				for y := 0; y < 100; y++ {
					bulkItems = append(bulkItems, map[string]interface{}{
						"filterKey":     "filterVal",
						"pageNo":        y + 1,
						"recordsOnPage": 100,
					})
				}
				res = append(res, bulkItems,
					[]map[string]interface{}{
						{
							"filterKey":     "filterVal",
							"pageNo":        101,
							"recordsOnPage": 100,
						},
					})
				return
			},
		},
		{
			name:  "fetch all in one request",
			total: 1000,
			inputProds: []payloadMock{
				{ID: 4},
			},
			expectedProdIDs: []int{4},
			listingSettings: ListingSettings{
				MaxRequestsCountPerSecond: 0,
				StreamBufferLength:        10,
				MaxItemsPerRequest:        10000,
				MaxFetchersCount:          10,
			},
			expectedBulkFilterInputs: func() (res [][]map[string]interface{}) {
				res = make([][]map[string]interface{}, 0, 2)
				bulkItems := make([]map[string]interface{}, 0, 100)
				for y := 0; y < 10; y++ {
					bulkItems = append(bulkItems, map[string]interface{}{
						"filterKey":     "filterVal",
						"pageNo":        y + 1,
						"recordsOnPage": 100,
					})
				}
				res = append(res, bulkItems)
				return
			},
		},
		{
			name:  "max items per request is impossible",
			total: 100,
			inputProds: []payloadMock{
				{ID: 5},
			},
			expectedProdIDs: []int{5},
			listingSettings: ListingSettings{
				MaxRequestsCountPerSecond: 0,
				StreamBufferLength:        10,
				MaxItemsPerRequest:        10001,
				MaxFetchersCount:          10,
			},
			expectedBulkFilterInputs: func() (res [][]map[string]interface{}) {
				return [][]map[string]interface{}{
					{
						{
							"filterKey":     "filterVal",
							"pageNo":        1,
							"recordsOnPage": 100,
						},
					},
				}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dp := &DataProviderMock{
				CountOutputCount:  testCase.total,
				ProductsToRead:    testCase.inputProds,
				countLock:         sync.Mutex{},
				CountFiltersInput: map[string]interface{}{},
				readLock:          sync.Mutex{},
				ReadBulkFilters:   [][]map[string]interface{}{},
			}
			lister := NewLister(
				testCase.listingSettings,
				dp,
				NullSleeper,
			)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			prodsChan := lister.Get(ctx, map[string]interface{}{"filterKey": "filterVal"})

			actualProds := collectProdsFromChannel(prodsChan)

			assert.Equal(t, map[string]interface{}{"filterKey": "filterVal", "pageNo": 1, "recordsOnPage": 1}, dp.CountFiltersInput)
			assert.Equal(t, ctx, dp.CountContextInput)
			assert.Equal(t, ctx, dp.ReadContextInput)
			assert.ElementsMatch(t, testCase.expectedBulkFilterInputs(), dp.ReadBulkFilters)

			actualProgressCounts := make([]int, 0, len(actualProds))
			actualProdIDs := make([]int, 0, len(actualProds))
			for _, prod := range actualProds {
				assert.NoError(t, prod.Err)
				assert.Equal(t, testCase.total, prod.TotalCount)
				assert.IsType(t, prod.Payload, payloadMock{})
				actualProdIDs = append(actualProdIDs, prod.Payload.(payloadMock).ID)
			}

			sort.Ints(actualProgressCounts)
			sort.Ints(actualProdIDs)

			assert.Equal(t, testCase.expectedProdIDs, actualProdIDs)
		})
	}
}

func TestReadCountError(t *testing.T) {
	dp := &DataProviderMock{
		CountOutputCount:    1,
		CountOutputErrorStr: "some count error",
	}

	lister := NewLister(ListingSettings{}, dp, NullSleeper)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	prodsChan := lister.Get(ctx, map[string]interface{}{"filterKey": "filterVal"})

	actualProds := collectProdsFromChannel(prodsChan)

	assert.Len(t, actualProds, 1)
	assert.EqualError(t, actualProds[0].Err, "some count error")
}

func TestCancelReading(t *testing.T) {
	dp := &DataProviderMock{
		CountOutputCount: 10,
		ProductsToRead: []payloadMock{
			{ID: 4},
			{ID: 5},
			{ID: 6},
			{ID: 7},
			{ID: 8},
		},
		countLock:         sync.Mutex{},
		CountFiltersInput: map[string]interface{}{},
		readLock:          sync.Mutex{},
		ReadBulkFilters:   [][]map[string]interface{}{},
	}
	lister := NewLister(
		ListingSettings{
			MaxRequestsCountPerSecond: 0,
			StreamBufferLength:        10,
			MaxItemsPerRequest:        100,
			MaxFetchersCount:          10,
		},
		dp,
		NullSleeper,
	)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	prodsChan := lister.Get(ctx, map[string]interface{}{"filterKey": "filterVal"})

	actualProds := collectProdsFromChannel(prodsChan)
	assert.NotEqual(t, 5, actualProds)
	assert.NotEqual(t, 5, dp.ReadBulkFilters)
}

func TestReadItemsError(t *testing.T) {
	dp := &DataProviderMock{
		CountOutputCount: 1,
		ReadErrorStr:     "some read items error",
	}

	lister := NewLister(ListingSettings{}, dp, NullSleeper)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	prodsChan := lister.Get(ctx, map[string]interface{}{})

	actualProds := collectProdsFromChannel(prodsChan)

	assert.Len(t, actualProds, 1)
	assert.EqualError(t, actualProds[0].Err, "some read items error")
}

func TestReadingGroupedSuccess(t *testing.T) {
	testCases := []struct {
		name                string
		total               int
		inputProds          []payloadMock
		listingSettings     ListingSettings
		expectedProdIdsFlat []int
		itemsCountPerGroup  int
		expectedGroupCounts []int
	}{
		{
			name:  "4 groups per 5 items with total 20 elements and 1 consumer",
			total: 20,
			inputProds: []payloadMock{
				{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}, {ID: 6}, {ID: 7}, {ID: 8}, {ID: 9}, {ID: 10},
				{ID: 11}, {ID: 12}, {ID: 13}, {ID: 14}, {ID: 15}, {ID: 16}, {ID: 17}, {ID: 18}, {ID: 19}, {ID: 20},
			},
			expectedProdIdsFlat: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			listingSettings: ListingSettings{
				MaxRequestsCountPerSecond: 0,
				StreamBufferLength:        0,
				MaxItemsPerRequest:        100,
				MaxFetchersCount:          1,
			},
			itemsCountPerGroup:  5,
			expectedGroupCounts: []int{5, 5, 5, 5},
		},
		{
			name:       "4 groups per 3 items with total 10 elements and 10 consumer",
			total:      12,
			inputProds: []payloadMock{{ID: 1}, {ID: 2}, {ID: 3}},
			//will make ceil(10/3)=4 requests in total, so {1,2,3} * 4
			expectedProdIdsFlat: []int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3},
			listingSettings: ListingSettings{
				StreamBufferLength: 3,
				MaxItemsPerRequest: 3,
				MaxFetchersCount:   10,
			},
			itemsCountPerGroup:  3,
			expectedGroupCounts: []int{3, 3, 3, 3},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dp := &DataProviderMock{
				CountOutputCount:  testCase.total,
				ProductsToRead:    testCase.inputProds,
				countLock:         sync.Mutex{},
				CountFiltersInput: map[string]interface{}{},
				readLock:          sync.Mutex{},
				ReadBulkFilters:   [][]map[string]interface{}{},
			}
			lister := NewLister(
				testCase.listingSettings,
				dp,
				NullSleeper,
			)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			prodsChanGrouped := lister.GetGrouped(ctx, map[string]interface{}{"filterKey": "filterVal"}, testCase.itemsCountPerGroup)

			groups, flatItems, flatIDs, actualGroupCounts := collectProdsGroupsFromChannel(prodsChanGrouped)

			assert.Equal(t, map[string]interface{}{"filterKey": "filterVal", "pageNo": 1, "recordsOnPage": 1}, dp.CountFiltersInput)
			assert.Equal(t, ctx, dp.CountContextInput)
			assert.Equal(t, ctx, dp.ReadContextInput)
			assert.ElementsMatch(t, testCase.expectedProdIdsFlat, flatIDs)
			assert.Len(t, groups, len(testCase.expectedGroupCounts))

			for _, prod := range flatItems {
				assert.NoError(t, prod.Err)
				assert.Equal(t, testCase.total, prod.TotalCount)
				assert.IsType(t, prod.Payload, payloadMock{})
			}
			assert.ElementsMatch(t, testCase.expectedGroupCounts, actualGroupCounts)
		})
	}
}

func collectProdsFromChannel(prodsChan ItemsStream) []Item {
	actualProds := make([]Item, 0)
	doneChan := make(chan struct{}, 1)
	go func() {
		defer close(doneChan)
		for prod := range prodsChan {
			actualProds = append(actualProds, prod)
		}
	}()

mainLoop:
	for {
		select {
		case <-doneChan:
			break mainLoop
		case <-time.After(time.Second):
			break mainLoop
		}
	}

	return actualProds
}

func collectProdsGroupsFromChannel(
	prodsChanGrouped ItemsStreamGrouped,
) (groups [][]Item, flatItems []Item, flatIDs []int, actualGroupCounts []int) {
	groups = make([][]Item, 0)
	flatItems = make([]Item, 0)
	flatIDs = make([]int, 0)
	actualGroupCounts = make([]int, 0)
	doneChan := make(chan struct{}, 1)
	go func() {
		defer close(doneChan)
		for prodGroup := range prodsChanGrouped {
			groups = append(groups, prodGroup)
			for _, prod := range prodGroup {
				flatItems = append(flatItems, prod)
				id := prod.Payload.(payloadMock).ID
				flatIDs = append(flatIDs, id)
			}
			actualGroupCounts = append(actualGroupCounts, len(prodGroup))
		}
	}()

mainLoop:
	for {
		select {
		case <-doneChan:
			break mainLoop
		case <-time.After(time.Second):
			break mainLoop
		}
	}

	return
}
