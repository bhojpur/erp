package address

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
	sharedCommon "github.com/bhojpur/erp/pkg/api/v1/common"
)

type (
	//GetAddressesResponse ..
	Response struct {
		Status    sharedCommon.Status    `json:"status"`
		Addresses sharedCommon.Addresses `json:"records"`
	}

	TypeResponse struct {
		Status       sharedCommon.Status `json:"status"`
		AddressTypes []Type              `json:"records"`
	}

	Type struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Added        string `json:"added"`
		LastModified string `json:"lastModified"`
	}

	GetAddressesResponseBulkItem struct {
		Status    sharedCommon.StatusBulk `json:"status"`
		Addresses sharedCommon.Addresses  `json:"records"`
	}

	GetAddressesResponseBulk struct {
		Status    sharedCommon.Status            `json:"status"`
		BulkItems []GetAddressesResponseBulkItem `json:"requests"`
	}

	SaveAddressResp struct {
		AddressID int `json:"addressID"`
	}

	SaveAddressesResponseBulkItem struct {
		Status  sharedCommon.StatusBulk `json:"status"`
		Records []SaveAddressResp       `json:"records"`
	}

	SaveAddressesResponseBulk struct {
		Status    sharedCommon.Status             `json:"status"`
		BulkItems []SaveAddressesResponseBulkItem `json:"requests"`
	}

	DeleteAddressResponse struct {
		Status sharedCommon.Status `json:"status"`
	}

	DeleteAddressBulkItem struct {
		Status sharedCommon.StatusBulk `json:"status"`
	}

	DeleteAddressResponseBulk struct {
		Status    sharedCommon.Status     `json:"status"`
		BulkItems []DeleteAddressBulkItem `json:"requests"`
	}
)

func (r Response) GetStatus() *sharedCommon.Status {
	return &r.Status
}

func (r TypeResponse) GetStatus() *sharedCommon.Status {
	return &r.Status
}
