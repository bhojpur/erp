package price

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
	GetSupplierPriceLists(ctx context.Context, filters map[string]string) ([]PriceList, error)
	AddProductToSupplierPriceList(ctx context.Context, filters map[string]string) (*ChangeProductToSupplierPriceListResult, error)
	EditProductToSupplierPriceList(ctx context.Context, filters map[string]string) (*ChangeProductToSupplierPriceListResult, error)
	ChangeProductToSupplierPriceListBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (ChangeProductToSupplierPriceListResponseBulk, error)
	GetSupplierPriceListsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetPriceListsResponseBulk, error)
	GetProductsInPriceList(ctx context.Context, filters map[string]string) ([]ProductsInPriceList, error)
	GetProductsInPriceListWithStatus(ctx context.Context, filters map[string]string) (GetProductsInPriceListResponse, error)
	GetProductsInPriceListBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetProductsInPriceListResponseBulk, error)
	GetProductsInSupplierPriceList(ctx context.Context, filters map[string]string) ([]ProductsInSupplierPriceList, error)
	GetProductsInSupplierPriceListBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (ProductsInSupplierPriceListResponseBulk, error)
	DeleteProductsFromSupplierPriceList(ctx context.Context, filters map[string]string) (*DeleteProductsFromSupplierPriceListResult, error)
	DeleteProductsFromSupplierPriceListBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (DeleteProductsFromSupplierPriceListResponseBulk, error)
	SaveSupplierPriceList(ctx context.Context, filters map[string]string) (*SaveSupplierPriceListResult, error)
	SaveSupplierPriceListBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (SaveSupplierPriceListResponseBulk, error)
	GetPriceLists(ctx context.Context, filters map[string]string) (*GetRegularPriceListResult, error)
	GetPriceListsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetRegularPriceListResponseBulk, error)
	SavePriceList(ctx context.Context, filters map[string]string) (*SavePriceListResult, error)
	SavePriceListBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (SavePriceListResponseBulk, error)
	AddProductToPriceList(ctx context.Context, filters map[string]string) (*ChangeProductToPriceListResult, error)
	EditProductToPriceList(ctx context.Context, filters map[string]string) (*ChangeProductToPriceListResult, error)
	ChangeProductToPriceListBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (ChangeProductToPriceListResponseBulk, error)
	DeleteProductsFromPriceList(ctx context.Context, filters map[string]string) (*DeleteProductsFromPriceListResult, error)
	DeleteProductsFromPriceListBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (DeleteProductsFromPriceListResponseBulk, error)
}
