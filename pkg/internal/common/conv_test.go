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
	"testing"

	"github.com/bhojpur/erp/pkg/api/v1/common"
	"github.com/stretchr/testify/assert"
)

type structToConvert struct {
	ID              uint                  `json:"id"`
	SupplierType    string                `json:"supplierType"`
	FullName        string                `json:"fullName"`
	CompanyName     string                `json:"companyName"`
	FirstName       string                `json:"firstName"`
	LstName         string                `json:"lastName"`
	GroupId         uint                  `json:"groupID"`
	GroupName       string                `json:"groupName"`
	Phone           string                `json:"phone"`
	Mobile          string                `json:"mobile"`
	Email           string                `json:"email"`
	Fax             string                `json:"fax"`
	Code            string                `json:"code"`
	IntegrationCode string                `json:"integrationCode"`
	VatrateID       uint                  `json:"vatrateID"`
	CurrencyCode    string                `json:"currencyCode"`
	DeliveryTermsID uint                  `json:"deliveryTermsID"`
	CountryId       uint                  `json:"countryID"`
	CountryName     string                `json:"countryName"`
	CountryCode     string                `json:"countryCode"`
	Address         string                `json:"address"`
	Gln             string                `json:"GLN"`
	Attributes      []common.ObjAttribute `json:"attributes"`

	// Detail fields
	VatNumber           string `json:"vatNumber"`
	Skype               string `json:"skype"`
	Website             string `json:"website"`
	BankName            string `json:"bankName"`
	BankAccountNumber   string `json:"bankAccountNumber"`
	BankIBAN            string `json:"bankIBAN"`
	BankSWIFT           string `json:"bankSWIFT"`
	Birthday            string `json:"birthday"`
	CompanyID           uint   `json:"companyID"`
	ParentCompanyName   string `json:"parentCompanyName"`
	SupplierManagerID   uint   `json:"supplierManagerID"`
	SupplierManagerName string `json:"supplierManagerName"`
	PaymentDays         uint   `json:"paymentDays"`
	Notes               string `json:"notes"`
	LastModified        string `json:"lastModified"`
	Added               uint64 `json:"added"`
}

func TestConvertingStructToMap(t *testing.T) {
	s := structToConvert{
		ID:                  1,
		SupplierType:        "some type",
		FullName:            "Some full name",
		CompanyName:         "Bhojpur Consulting",
		FirstName:           "some first",
		LstName:             "some last",
		GroupId:             3,
		GroupName:           "some group",
		Phone:               "3334444444",
		Mobile:              "8788682735",
		Email:               "no@mail.me",
		Fax:                 "341234343241",
		Code:                "32413",
		IntegrationCode:     "341324",
		VatrateID:           5,
		CurrencyCode:        "inr",
		DeliveryTermsID:     7,
		CountryId:           6,
		CountryName:         "India",
		CountryCode:         "IN",
		Address:             "Nirbhaya Dihra, Basauri, Piro",
		Gln:                 "gln222",
		Attributes:          []common.ObjAttribute{},
		VatNumber:           "3431241",
		Skype:               "nono",
		Website:             "bhojpur.net",
		BankName:            "some swiss bank",
		BankAccountNumber:   "3413412434",
		BankIBAN:            "341t45243535",
		BankSWIFT:           "some swift",
		Birthday:            "26.03.2018",
		CompanyID:           9,
		ParentCompanyName:   "some parent",
		SupplierManagerID:   10,
		SupplierManagerName: "Some manager",
		PaymentDays:         11,
		Notes:               "some notes",
		LastModified:        "2018-03-26 00:00:00",
		Added:               1,
	}

	actualMap, err := ConvertStructToMap(s)
	assert.NoError(t, err)
	if err != nil {
		return
	}
	expectedMap := map[string]interface{}{
		"GLN":                 "gln222",
		"added":               float64(1),
		"address":             "Bagar, Arrah",
		"attributes":          []interface{}{},
		"bankAccountNumber":   "3413412434",
		"bankIBAN":            "341t45243535",
		"bankName":            "some swiss bank",
		"bankSWIFT":           "some swift",
		"birthday":            "26.03.2018",
		"code":                "32413",
		"companyID":           float64(9),
		"companyName":         "Yunica Retail",
		"countryCode":         "IN",
		"countryID":           float64(6),
		"countryName":         "India",
		"currencyCode":        "inr",
		"deliveryTermsID":     float64(7),
		"email":               "no@mail.me",
		"fax":                 "341234343241",
		"firstName":           "some first",
		"fullName":            "Some full name",
		"groupID":             float64(3),
		"groupName":           "some group",
		"integrationCode":     "341324",
		"lastModified":        "2018-11-28 00:00:00",
		"lastName":            "some last",
		"mobile":              "341431434",
		"notes":               "some notes",
		"parentCompanyName":   "some parent",
		"paymentDays":         float64(11),
		"phone":               "3334444444",
		"skype":               "nono",
		"supplierID":          float64(1),
		"supplierManagerID":   float64(10),
		"supplierManagerName": "Some manager",
		"supplierType":        "some type",
		"vatNumber":           "3431241",
		"vatrateID":           float64(5),
		"website":             "bhojpur.net",
	}
	assert.Equal(t, expectedMap, actualMap)
}
