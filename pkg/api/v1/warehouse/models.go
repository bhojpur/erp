package warehouse

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
	"encoding/json"

	sharedCommon "github.com/bhojpur/erp/pkg/api/v1/common"
)

type (
	GetWarehousesResponse struct {
		Status     sharedCommon.Status `json:"status"`
		Warehouses Warehouses          `json:"records"`
	}

	Warehouse struct {
		WarehouseID            string      `json:"warehouseID"`
		PricelistID            json.Number `json:"pricelistID"`
		PricelistID2           json.Number `json:"pricelistID2"`
		PricelistID3           json.Number `json:"pricelistID3"`
		PricelistID4           json.Number `json:"pricelistID4"`
		PricelistID5           json.Number `json:"pricelistID5"`
		Name                   string      `json:"name"`
		Code                   string      `json:"code"`
		AddressID              int         `json:"addressID"`
		Address                string      `json:"address"`
		Street                 string      `json:"street"`
		Address2               string      `json:"address2"`
		City                   string      `json:"city"`
		State                  string      `json:"state"`
		Country                string      `json:"country"`
		PINcode                string      `json:"PINcode"`
		StoreGroups            string      `json:"storeGroups"`
		CompanyName            string      `json:"companyName"`
		CompanyCode            string      `json:"companyCode"`
		CompanyVatNumber       string      `json:"companyVatNumber"`
		Phone                  string      `json:"phone"`
		Fax                    string      `json:"fax"`
		Email                  string      `json:"email"`
		Website                string      `json:"website"`
		BankName               string      `json:"bankName"`
		BankAccountNumber      string      `json:"bankAccountNumber"`
		Iban                   string      `json:"iban"`
		Swift                  string      `json:"swift"`
		UsesLocalQuickButtons  int         `json:"usesLocalQuickButtons"`
		DefaultCustomerGroupID int         `json:"defaultCustomerGroupID"`
		IsOfflineInventory     int         `json:"isOfflineInventory"`
		TimeZone               string      `json:"timeZone"`
		sharedCommon.Attributes
	}

	Warehouses []Warehouse

	GetWarehousesBulkItem struct {
		Status     sharedCommon.StatusBulk `json:"status"`
		Warehouses Warehouses              `json:"records"`
	}

	GetWarehousesResponseBulk struct {
		Status    sharedCommon.Status     `json:"status"`
		BulkItems []GetWarehousesBulkItem `json:"requests"`
	}

	SaveWarehouseResult struct {
		WarehouseID int `json:"warehouseID"`
	}

	SaveWarehouseResponse struct {
		Status  sharedCommon.Status   `json:"status"`
		Results []SaveWarehouseResult `json:"records"`
	}

	SaveWarehouseBulkItem struct {
		Status  sharedCommon.StatusBulk `json:"status"`
		Results []SaveWarehouseResult   `json:"records"`
	}

	SaveWarehouseResponseBulk struct {
		Status    sharedCommon.Status     `json:"status"`
		BulkItems []SaveWarehouseBulkItem `json:"requests"`
	}

	SaveInventoryRegistrationResult struct {
		InventoryRegistrationID int `json:"inventoryRegistrationID"`
	}

	SaveInventoryRegistrationResponse struct {
		Status  sharedCommon.Status               `json:"status"`
		Results []SaveInventoryRegistrationResult `json:"records"`
	}

	SaveInventoryWriteOffResult struct {
		InventoryWriteOffID int `json:"inventoryWriteOffID"`
	}
	SaveInventoryWriteOffResponse struct {
		Status  sharedCommon.Status           `json:"status"`
		Results []SaveInventoryWriteOffResult `json:"records"`
	}

	SaveInventoryTransferResult struct {
		InventoryTransferID int `json:"inventoryTransferID"`
	}
	SaveInventoryTransferResponse struct {
		Status  sharedCommon.Status           `json:"status"`
		Results []SaveInventoryTransferResult `json:"records"`
	}

	SaveInventoryRegistrationBulkItem struct {
		Status  sharedCommon.StatusBulk           `json:"status"`
		Results []SaveInventoryRegistrationResult `json:"records"`
	}

	SaveInventoryRegistrationResponseBulk struct {
		Status    sharedCommon.Status                 `json:"status"`
		BulkItems []SaveInventoryRegistrationBulkItem `json:"requests"`
	}

	ReasonCode struct {
		ReasonID                             int    `json:"reasonID"`
		Name                                 string `json:"name"`
		Added                                int    `json:"added"`
		LastModified                         int    `json:"lastModified"`
		Purpose                              string `json:"purpose"`
		Code                                 string `json:"code"`
		ManualDiscountDisablesPromotionTiers []int  `json:"manualDiscountDisablesPromotionTiers"`
	}
	GetReasonCodesResponse struct {
		Status      sharedCommon.Status `json:"status"`
		ReasonCodes []ReasonCode        `json:"records"`
	}
)
