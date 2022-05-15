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

type (
	ShoppingAppliedPromotions struct {
		Count        int `json:"count"`
		PromotionID  int `json:"promotionID"`
		RewardPoints int `json:"rewardPoints"`
	}
	ShoppingCartTotals struct {
		Rows              []ShoppingCartProduct       `json:"rows"`
		NetTotal          float64                     `json:"netTotal"`
		VATTotal          float64                     `json:"vatTotal"`
		Total             float64                     `json:"total"`
		AppliedPromotions []ShoppingAppliedPromotions `json:"appliedPromotions"`
	}
	ShoppingCartTotalsWithFullRows struct {
		Rows              []map[string]interface{}    `json:"rows"`
		NetTotal          float64                     `json:"netTotal"`
		VATTotal          float64                     `json:"vatTotal"`
		Total             float64                     `json:"total"`
		AppliedPromotions []ShoppingAppliedPromotions `json:"appliedPromotions"`
	}
	ShoppingCartProduct struct {
		RowNumber            int     `json:"rowNumber"`
		ProductID            string  `json:"productID"`
		Amount               string  `json:"amount"`
		VatRateID            int     `json:"vatrateID"`
		VatRate              string  `json:"vatRate"`
		OriginalPrice        float64 `json:"originalPrice"`
		OriginalPriceWithVAT float64 `json:"originalPriceWithVAT"`
		PromotionDiscount    float64 `json:"promotionDiscount"`
		ManualDiscount       float64 `json:"manualDiscount"`
		Discount             float64 `json:"discount"`
		FinalPrice           float64 `json:"finalPrice"`
		FinalPriceWithVAT    float64 `json:"finalPriceWithVAT"`
		RowNetTotal          float64 `json:"rowNetTotal"`
		RowVAT               float64 `json:"rowVAT"`
		RowTotal             float64 `json:"rowTotal"`
	}
)
