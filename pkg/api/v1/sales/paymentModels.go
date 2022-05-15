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
	PaymentStatus string
	PaymentType   string

	PaymentInfo struct {
		DocumentID             int                         `json:"documentID"` // Invoice ID
		PaymentID              int                         `json:"paymentID"`
		CustomerID             int                         `json:"customerID"`
		TypeID                 string                      `json:"typeID"`
		BankTransactionID      int                         `json:"bankTransactionID"`
		Type                   string                      `json:"type"` // CASH, TRANSFER, CARD, CREDIT, GIFTCARD, CHECK, TIP
		Date                   string                      `json:"date"`
		Sum                    string                      `json:"sum"`
		CardHolder             string                      `json:"cardHolder"`
		CardType               string                      `json:"cardType"`
		CardNumber             string                      `json:"cardNumber"`
		AuthorizationCode      string                      `json:"authorizationCode"`
		ReferenceNumber        string                      `json:"referenceNumber"`
		CurrencyRate           string                      `json:"currencyRate"`
		CashPaid               string                      `json:"cashPaid"`
		CashChange             string                      `json:"cashChange"`
		CurrencyCode           string                      `json:"currencyCode"` // INR, EUR, USD
		Info                   string                      `json:"info"`         // Information about the payer or payment transaction
		Added                  uint64                      `json:"added"`
		IsPrepayment           uint64                      `json:"isPrepayment"`
		StoreCredit            uint64                      `json:"storeCredit"`
		BankAccount            string                      `json:"bankAccount"`
		BankDocumentNumber     string                      `json:"bankDocumentNumber"`
		BankDate               string                      `json:"bankDate"`
		BankPayerAccount       string                      `json:"bankPayerAccount"`
		BankPayerName          string                      `json:"bankPayerName"`
		BankPayerCode          string                      `json:"bankPayerCode"`
		BankSum                string                      `json:"bankSum"`
		BankReferenceNumber    string                      `json:"bankReferenceNumber"`
		BankDescription        string                      `json:"bankDescription"`
		BankCurrency           string                      `json:"bankCurrency"`
		ArchivalNumber         string                      `json:"archivalNumber"`
		PaymentServiceProvider string                      `json:"paymentServiceProvider"`
		Aid                    string                      `json:"aid"`
		ApplicationLabel       string                      `json:"applicationLabel"`
		PinStatement           string                      `json:"pinStatement"`
		CryptogramType         string                      `json:"cryptogramType"`
		Cryptogram             string                      `json:"cryptogram"`
		ExpirationDate         string                      `json:"expirationDate"`
		EntryMethod            string                      `json:"entryMethod"`
		TransactionNumber      string                      `json:"transactionNumber"`
		TransactionId          string                      `json:"transactionId"`
		TransactionType        string                      `json:"transactionType"`
		TransactionTime        int64                       `json:"transactionTime"`
		UpiPaymentID           string                      `json:"upiPaymentID"`
		CertificateBalance     string                      `json:"certificateBalance"`
		StatusCode             string                      `json:"statusCode"`
		StatusMessage          string                      `json:"statusMessage"`
		GiftCardVatRateID      int                         `json:"giftCardVatRateID"`
		LastModified           uint64                      `json:"lastModified"`
		Attributes             []sharedCommon.ObjAttribute `json:"attributes"`
	}

	GetPaymentsBulkItem struct {
		Status       sharedCommon.StatusBulk `json:"status"`
		PaymentInfos []PaymentInfo           `json:"records"`
	}

	GetPaymentsResponseBulk struct {
		Status    sharedCommon.Status   `json:"status"`
		BulkItems []GetPaymentsBulkItem `json:"requests"`
	}

	SavePaymentID struct {
		PaymentID int `json:"paymentID"`
	}

	SavePaymentsBulkItem struct {
		Status  sharedCommon.StatusBulk `json:"status"`
		Records []SavePaymentID         `json:"records"`
	}

	SavePaymentsResponseBulk struct {
		Status    sharedCommon.Status    `json:"status"`
		BulkItems []SavePaymentsBulkItem `json:"requests"`
	}
)
