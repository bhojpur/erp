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

import "context"

type (
	ProjectManager interface {
		GetProjects(ctx context.Context, filters map[string]string) ([]Project, error)
		GetProjectStatus(ctx context.Context, filters map[string]string) ([]ProjectStatus, error)
	}
	DocumentManager interface {
		SaveSalesDocument(ctx context.Context, filters map[string]string) (SaleDocImportReports, error)
		SaveSalesDocumentBulk(
			ctx context.Context,
			bulkFilters []map[string]interface{},
			baseFilters map[string]string,
		) (respBulk SaveSalesDocumentResponseBulk, err error)
		GetSalesDocuments(ctx context.Context, filters map[string]string) ([]SaleDocument, error)
		GetSalesDocumentsWithStatus(ctx context.Context, filters map[string]string) (*GetSalesDocumentResponse, error)
		GetSalesDocumentsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetSaleDocumentResponseBulk, error)
		DeleteDocument(ctx context.Context, filters map[string]string) error
		SavePurchaseDocument(ctx context.Context, filters map[string]string) (PurchaseDocImportReports, error)
		SavePurchaseDocumentBulk(
			ctx context.Context,
			bulkFilters []map[string]interface{},
			baseFilters map[string]string,
		) (respBulk SavePurchaseDocumentResponseBulk, err error)
		DeleteDocumentsBulk(
			ctx context.Context,
			bulkFilters []map[string]interface{},
			baseFilters map[string]string,
		) (respBulk DeleteResponseBulk, err error)
	}

	VatRateManager interface {
		GetVatRates(ctx context.Context, filters map[string]string) (VatRates, error)
		GetVatRatesBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetVatRatesResponseBulk, error)
		SaveVatRate(ctx context.Context, filters map[string]string) (*SaveVatRateResult, error)
		SaveVatRateBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (SaveVatRateResponseBulk, error)
		SaveVatRateComponent(ctx context.Context, filters map[string]string) (*SaveVatRateComponentResult, error)
		SaveVatRateComponentBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (SaveVatRateComponentResponseBulk, error)
	}

	AssignmentsManger interface {
		SaveAssignment(ctx context.Context, filters map[string]string) (int64, error)
	}

	ReportsManager interface {
		GetSalesReport(ctx context.Context, filters map[string]string) (*GetSalesReport, error)
	}

	Manager interface {
		ProjectManager
		DocumentManager
		VatRateManager
		AssignmentsManger
		ReportsManager
		GetCoupons(ctx context.Context, filters map[string]string) (*GetCouponsResponse, error)
		//payment requests
		SavePayment(ctx context.Context, filters map[string]string) (int64, error)
		SavePaymentsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (SavePaymentsResponseBulk, error)
		GetPayments(ctx context.Context, filters map[string]string) ([]PaymentInfo, error)
		GetPaymentsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetPaymentsResponseBulk, error)
		DeletePayment(ctx context.Context, filters map[string]string) error
		DeletePaymentsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (DeleteResponseBulk, error)

		//shopping cart
		CalculateShoppingCart(ctx context.Context, filters map[string]string) (*ShoppingCartTotals, error)
		CalculateShoppingCartWithFullRowsResponse(ctx context.Context, filters map[string]string) (*ShoppingCartTotalsWithFullRows, error)
	}
)
