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
	common2 "github.com/bhojpur/erp/pkg/api/v1/common"
)

type (
	GetProjectsResponse struct {
		Status   common2.Status `json:"status"`
		Projects []Project      `json:"records"`
	}

	GetProjectStatusesResponse struct {
		Status          common2.Status  `json:"status"`
		ProjectStatuses []ProjectStatus `json:"records"`
	}

	Project struct {
		ProjectID    uint   `json:"projectID"`
		Name         string `json:"name"`
		CustomerID   uint   `json:"customerID"`
		CustomerName string `json:"customerName"`
		EmployeeID   uint   `json:"employeeID"`
		EmployeeName string `json:"employeeName"`
		TypeID       uint   `json:"typeID"`
		TypeName     string `json:"typeName"`
		StatusID     uint   `json:"statusID"`
		StatusName   string `json:"statusName"`
		StartDate    string `json:"startDate"`
		EndDate      string `json:"endDate"`
		Notes        string `json:"notes"`
		LastModified uint64 `json:"lastModified"`
	}

	ProjectStatus struct {
		ProjectStatusID uint   `json:"projectStatusID"`
		Name            string `json:"name"`
		Finished        byte   `json:"finished"`
		Added           uint64 `json:"added"`
		LastModified    uint64 `json:"lastModified"`
	}
)
