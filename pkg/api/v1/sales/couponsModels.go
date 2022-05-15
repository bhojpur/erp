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

import sharedCommon "github.com/bhojpur/erp/pkg/api/v1/common"

type (
	Coupon struct {
		Added                      int    `json:"added"`
		CampaignID                 string `json:"campaignID"`
		Code                       string `json:"code"`
		CouponID                   int    `json:"couponID"`
		Description                string `json:"description"`
		IssuedFromDate             string `json:"issuedFromDate"`
		IssuedUntilDate            string `json:"issuedUntilDate"`
		LastModified               int    `json:"lastModified"`
		Measure                    string `json:"measure"`
		Name                       string `json:"name"`
		PrintedAutomaticallyInPOS  int    `json:"printedAutomaticallyInPOS"`
		PrintingCostInRewardPoints int    `json:"printingCostInRewardPoints"`
		PromptCashier              int    `json:"promptCashier"`
		Threshold                  string `json:"threshold"`
		ThresholdType              string `json:"thresholdType"`
		Treshold                   int    `json:"treshold"`
		TresholdType               string `json:"tresholdType"`
		WarehouseID                string `json:"warehouseID"`
	}

	GetCouponsResponse struct {
		Status  sharedCommon.Status `json:"status"`
		Coupons []Coupon            `json:"records"`
	}
)
