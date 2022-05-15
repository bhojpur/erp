package product

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
	GetProducts(ctx context.Context, filters map[string]string) ([]Product, error)
	GetProductsCount(ctx context.Context, filters map[string]string) (int, error)
	GetProductsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetProductsResponseBulk, error)
	GetProductUnits(ctx context.Context, filters map[string]string) ([]ProductUnit, error)
	GetProductCategories(ctx context.Context, filters map[string]string) ([]ProductCategory, error)
	GetProductCategoriesBulk(
		ctx context.Context,
		bulkFilters []map[string]interface{},
		baseFilters map[string]string,
	) (respBulk GetProductCategoryResponseBulk, err error)
	GetProductBrands(ctx context.Context, filters map[string]string) ([]ProductBrand, error)
	GetBrands(ctx context.Context, filters map[string]string) ([]ProductBrand, error)
	GetProductPriorityGroups(ctx context.Context, filters map[string]string) (GetProductPriorityGroups, error)
	GetProductPriorityGroupBulk(
		ctx context.Context,
		bulkFilters []map[string]interface{},
		baseFilters map[string]string,
	) (respBulk GetProductPriorityGroupResponseBulk, err error)
	GetProductGroups(ctx context.Context, filters map[string]string) ([]ProductGroup, error)
	GetProductGroupsBulk(
		ctx context.Context,
		bulkFilters []map[string]interface{},
		baseFilters map[string]string,
	) (respBulk GetProductGroupResponseBulk, err error)
	GetProductStock(ctx context.Context, filters map[string]string) ([]GetProductStock, error)
	GetProductStockFile(ctx context.Context, filters map[string]string) ([]GetProductStockFile, error)
	GetProductStockFileBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetProductStockFileResponseBulk, error)
	GetProductStockBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetProductStockResponseBulk, error)
	SaveProduct(ctx context.Context, filters map[string]string) (SaveProductResult, error)
	SaveProductBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (SaveProductResponseBulk, error)
	DeleteProduct(ctx context.Context, filters map[string]string) error
	DeleteProductBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (DeleteProductResponseBulk, error)
	SaveAssortment(ctx context.Context, filters map[string]string) (SaveAssortmentResult, error)
	SaveAssortmentBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (SaveAssortmentResponseBulk, error)
	AddAssortmentProducts(ctx context.Context, filters map[string]string) (AddAssortmentProductsResult, error)
	AddAssortmentProductsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (AddAssortmentProductsResponseBulk, error)
	EditAssortmentProducts(ctx context.Context, filters map[string]string) (EditAssortmentProductsResult, error)
	EditAssortmentProductsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (EditAssortmentProductsResponseBulk, error)
	RemoveAssortmentProducts(ctx context.Context, filters map[string]string) (RemoveAssortmentProductResult, error)
	RemoveAssortmentProductsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (RemoveAssortmentProductResponseBulk, error)
	SaveProductCategory(ctx context.Context, filters map[string]string) (result SaveProductCategoryResult, err error)
	SaveProductCategoryBulk(
		ctx context.Context,
		bulkFilters []map[string]interface{},
		baseFilters map[string]string,
	) (respBulk SaveProductCategoryResponseBulk, err error)
	SaveBrand(ctx context.Context, filters map[string]string) (result SaveBrandResult, err error)
	SaveBrandBulk(
		ctx context.Context,
		bulkFilters []map[string]interface{},
		baseFilters map[string]string,
	) (respBulk SaveBrandResponseBulk, err error)
	SaveProductPriorityGroup(ctx context.Context, filters map[string]string) (result SaveProductPriorityGroupResult, err error)
	SaveProductPriorityGroupBulk(
		ctx context.Context,
		bulkFilters []map[string]interface{},
		baseFilters map[string]string,
	) (respBulk SaveProductPriorityGroupResponseBulk, err error)
	SaveProductGroup(ctx context.Context, filters map[string]string) (result SaveProductGroupResult, err error)
	SaveProductGroupBulk(
		ctx context.Context,
		bulkFilters []map[string]interface{},
		baseFilters map[string]string,
	) (respBulk SaveProductGroupResponseBulk, err error)
	DeleteProductGroup(ctx context.Context, filters map[string]string) error
	DeleteProductGroupBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (DeleteProductGroupResponseBulk, error)
	GetProductFiles(ctx context.Context, filters map[string]string) (GetProductFilesResponse, error)
	GetProductPictures(ctx context.Context, filters map[string]string) ([]Image, error)
	GetProductPicturesBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetProductPicturesResponseBulk, error)
}
