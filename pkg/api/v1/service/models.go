package service

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

type getServiceEndpointsResponse struct {
	Status  common2.Status
	Records []ServiceEndpoints `json:"records"`
}
type Endpoint struct {
	IsSandbox     bool   `json:"isSandbox"`
	Url           string `json:"url"`
	Documentation string `json:"documentation"`
}

type ServiceEndpoints struct {
	Chain        Endpoint `json:"chain"`
	PIM          Endpoint `json:"pim"`
	WMS          Endpoint `json:"wms"`
	Promotion    Endpoint `json:"promotion"`
	Reports      Endpoint `json:"reports"`
	JSON         Endpoint `json:"json"`
	Assignments  Endpoint `json:"assignments"`
	AccountAdmin Endpoint `json:"account-admin"`
	VisitorQueue Endpoint `json:"visitor-queue"`
	Loyalty      Endpoint `json:"loyalty"`
	CDN          Endpoint `json:"cdn"`
	Tasks        Endpoint `json:"tasks"`
	Webhook      Endpoint `json:"webhook"`
	User         Endpoint `json:"user"`
	Import       Endpoint `json:"import"`
	EMS          Endpoint `json:"ems"`
	ClockIn      Endpoint `json:"clockin"`
	Ledger       Endpoint `json:"ledger"`
	Auth         Endpoint `json:"auth"`
	CRM          Endpoint `json:"crm"`
	DCP          Endpoint `json:"dcp"`
	Sales        Endpoint `json:"sales"`
	Pricing      Endpoint `json:"pricing"`
	Inventory    Endpoint `json:"inventory"`
	Chair        Endpoint `json:"chair"`
	PosAPI       Endpoint `json:"pos-api"`
	ERP          Endpoint `json:"erp"`
}
