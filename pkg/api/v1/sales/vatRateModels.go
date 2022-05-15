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
	sharedCommon "github.com/bhojpur/erp/pkg/api/v1/common"
)

type (
	VatRate struct {
		ID         string                      `json:"id"`
		Name       string                      `json:"name"`
		Rate       string                      `json:"rate"`
		Code       string                      `json:"code"`
		Active     string                      `json:"active"`
		Attributes []sharedCommon.ObjAttribute `json:"attributes"`
		//Added        string `json:"added"`
		LastModified string `json:"lastModified"`
		//IsReverseVat int    `json:"isReverseVat"`
		//ReverseRate int `json:"reverseRate"`
	}

	VatRates []VatRate

	NetTotalsByTaxRate struct {
		VatrateID int     `json:"vatrateID"`
		Total     float64 `json:"total"`
	}

	//GetVatRatesResponse ...
	GetVatRatesResponse struct {
		Status   sharedCommon.Status `json:"status"`
		VatRates []VatRate           `json:"records"`
	}

	GetVatRatesBulkItem struct {
		Status   sharedCommon.StatusBulk `json:"status"`
		VatRates []VatRate               `json:"records"`
	}

	GetVatRatesResponseBulk struct {
		Status    sharedCommon.Status   `json:"status"`
		BulkItems []GetVatRatesBulkItem `json:"requests"`
	}

	SaveVatRateResult struct {
		VatRateID int `json:"vatRateID"`
	}

	SaveVatRateResultResponse struct {
		Status            sharedCommon.Status `json:"status"`
		SaveVatRateResult []SaveVatRateResult `json:"records"`
	}

	SaveVatRateBulkItem struct {
		Status  sharedCommon.StatusBulk `json:"status"`
		Records []SaveVatRateResult     `json:"records"`
	}
	SaveVatRateResponseBulk struct {
		Status    sharedCommon.Status   `json:"status"`
		BulkItems []SaveVatRateBulkItem `json:"requests"`
	}

	SaveVatRateComponentResult struct {
		VatRateComponentID int `json:"vatRateComponentID"`
	}

	SaveVatRateComponentResultResponse struct {
		Status                     sharedCommon.Status          `json:"status"`
		SaveVatRateComponentResult []SaveVatRateComponentResult `json:"records"`
	}

	SaveVatRateComponentBulkItem struct {
		Status  sharedCommon.StatusBulk      `json:"status"`
		Records []SaveVatRateComponentResult `json:"records"`
	}
	SaveVatRateComponentResponseBulk struct {
		Status    sharedCommon.Status            `json:"status"`
		BulkItems []SaveVatRateComponentBulkItem `json:"requests"`
	}
)
