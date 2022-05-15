package customer

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

type Manager interface {
	SaveCustomer(ctx context.Context, filters map[string]string) (*CustomerImportReport, error)
	SaveCustomerBulk(ctx context.Context, customerMap []map[string]interface{}, attrs map[string]string) (SaveCustomerResponseBulk, error)
	GetCustomers(ctx context.Context, filters map[string]string) ([]Customer, error)
	GetCustomersWithStatus(ctx context.Context, filters map[string]string) (*GetCustomersResponse, error)
	GetCustomersBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetCustomersResponseBulk, error)
	DeleteCustomer(ctx context.Context, filters map[string]string) error
	DeleteCustomerBulk(ctx context.Context, customerMap []map[string]interface{}, attrs map[string]string) (DeleteCustomersResponseBulk, error)
	VerifyCustomerUser(ctx context.Context, username, password string) (*WebshopClient, error)
	ValidateCustomerUsername(ctx context.Context, username string) (bool, error)
	GetCustomerGroups(ctx context.Context, filters map[string]string) ([]CustomerGroup, error)
	// GetCustomerBalance will retrieve current balance (store credit) for requested customers.
	GetCustomerBalance(ctx context.Context, filters map[string]string) ([]CustomerBalance, error)
	GetSuppliers(ctx context.Context, filters map[string]string) ([]Supplier, error)
	GetSuppliersBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetSuppliersResponseBulk, error)
	SaveSupplier(ctx context.Context, filters map[string]string) (*CustomerImportReport, error)
	SaveSupplierBulk(ctx context.Context, suppliers []map[string]interface{}, attrs map[string]string) (SaveSuppliersResponseBulk, error)
	DeleteSupplier(ctx context.Context, filters map[string]string) error
	DeleteSupplierBulk(ctx context.Context, supplierMap []map[string]interface{}, attrs map[string]string) (DeleteSuppliersResponseBulk, error)
	AddCustomerRewardPoints(ctx context.Context, filters map[string]string) (AddCustomerRewardPointsResult, error)
	AddCustomerRewardPointsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (AddCustomerRewardPointsResponseBulk, error)
	GetCompanyTypes(ctx context.Context, filters map[string]string) ([]CompanyType, error)
	SaveCompanyType(ctx context.Context, filters map[string]string) (*SaveCompanyTypeResponse, error)
	SaveSupplierGroup(ctx context.Context, filters map[string]string) (*SaveSupplierGroupResponse, error)
}
