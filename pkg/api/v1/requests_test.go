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
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	sharedCommon "github.com/bhojpur/erp/pkg/api/v1/common"
	"github.com/bhojpur/erp/pkg/internal/common"
	"github.com/stretchr/testify/assert"
)

func TestGetEmployeesBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"requestName": "getEmployees",
				"employeeID":  "1",
			},
			{
				"requestName": "getEmployees",
				"employeeID":  "2",
			},
			{
				"requestName": "getEmployees",
				"employeeID":  "3",
			},
		})

		bulkResp := GetEmployeesResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []GetEmployeesResponseBulkItem{
				{
					Status: statusBulk,
					Employees: []Employee{
						{
							EmployeeID: "1",
							FullName:   "Name 1",
						},
					},
				},
				{
					Status: statusBulk,
					Employees: []Employee{
						{
							EmployeeID: "2",
							FullName:   "Name 2",
						},
					},
				},
				{
					Status: statusBulk,
					Employees: []Employee{
						{
							EmployeeID: "3",
							FullName:   "Name 3",
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

	c := newErpClient(cli)

	bulkResp, err := c.GetEmployeesBulk(
		context.Background(),
		[]map[string]interface{}{
			{
				"employeeID": "1",
			},
			{
				"employeeID": "2",
			},
			{
				"employeeID": "3",
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

	assert.Equal(t, []Employee{
		{
			EmployeeID: "1",
			FullName:   "Name 1",
		},
	}, bulkResp.BulkItems[0].Employees)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)

	assert.Equal(t, []Employee{
		{
			EmployeeID: "2",
			FullName:   "Name 2",
		},
	}, bulkResp.BulkItems[1].Employees)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)

	assert.Equal(t, []Employee{
		{
			EmployeeID: "3",
			FullName:   "Name 3",
		},
	}, bulkResp.BulkItems[2].Employees)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[2].Status)
}
