package address

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
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	sharedCommon "github.com/bhojpur/erp/pkg/api/v1/common"
	"github.com/bhojpur/erp/pkg/internal/common"
)

func (cli *Client) GetAddresses(ctx context.Context, filters map[string]string) (addrs []sharedCommon.Address, err error) {
	res := &Response{}

	err = cli.Scan(ctx, "getAddresses", filters, res)
	if err != nil {
		return
	}

	return res.Addresses, nil
}

func (cli *Client) GetAddressTypes(ctx context.Context, filters map[string]string) (addrTypes []Type, err error) {
	res := &TypeResponse{}

	err = cli.Scan(ctx, "getAddressTypes", filters, res)
	if err != nil {
		return
	}

	return res.AddressTypes, nil
}

// GetAddressesBulk will list addresses according to specified filters sending a bulk request to fetch more addresses than the default limit
func (cli *Client) GetAddressesBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetAddressesResponseBulk, error) {
	var addrResp GetAddressesResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getAddresses",
			Filters:    bulkFilterMap,
		})
	}
	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return addrResp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return addrResp, err
	}

	if err := json.Unmarshal(body, &addrResp); err != nil {
		return addrResp, fmt.Errorf("Bhojpur ERP API: failed to unmarshal GetAddressesResponseBulk from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&addrResp.Status) {
		return addrResp, sharedCommon.NewErpError(addrResp.Status.ErrorCode.String(), addrResp.Status.Request+": "+addrResp.Status.ResponseStatus, addrResp.Status.ErrorCode)
	}

	for _, addrBulkItem := range addrResp.BulkItems {
		if !common.IsJSONResponseOK(&addrBulkItem.Status.Status) {
			return addrResp, sharedCommon.NewErpError(addrBulkItem.Status.ErrorCode.String(), addrBulkItem.Status.Request+": "+addrBulkItem.Status.ResponseStatus, addrResp.Status.ErrorCode)
		}
	}

	return addrResp, nil
}

func (cli *Client) SaveAddress(ctx context.Context, filters map[string]string) ([]sharedCommon.Address, error) {
	method := "saveAddress"
	resp, err := cli.SendRequest(ctx, method, filters)
	if err != nil {
		return nil, sharedCommon.NewFromError(method+": request failed", err, 0)
	}
	res := &Response{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError(method+": JSON unmarshal failed", err, 0)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	if len(res.Addresses) == 0 {
		return nil, sharedCommon.NewFromError(method+": no records in response", nil, res.Status.ErrorCode)
	}

	return res.Addresses, nil
}

func (cli *Client) DeleteAddress(ctx context.Context, filters map[string]string) error {
	method := "deleteAddress"
	resp, err := cli.SendRequest(ctx, method, filters)
	if err != nil {
		return sharedCommon.NewFromError(method+": request failed", err, 0)
	}
	res := &DeleteAddressResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return sharedCommon.NewFromError(method+": JSON unmarshal failed", err, 0)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return sharedCommon.NewFromResponseStatus(&res.Status)
	}

	return nil
}

func (cli *Client) DeleteAddressBulk(
	ctx context.Context,
	bulkRequest []map[string]interface{},
	baseFilters map[string]string,
) (DeleteAddressResponseBulk, error) {
	var bulkResp DeleteAddressResponseBulk

	if len(bulkRequest) > sharedCommon.MaxBulkRequestsCount {
		return bulkResp, fmt.Errorf("cannot delete more than %d addresses in one bulk request", sharedCommon.MaxBulkRequestsCount)
	}

	bulkInputs := make([]common.BulkInput, 0, len(bulkRequest))
	for _, bulkInput := range bulkRequest {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "deleteAddress",
			Filters:    bulkInput,
		})
	}

	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return bulkResp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return bulkResp, err
	}

	if err := json.Unmarshal(body, &bulkResp); err != nil {
		return bulkResp, fmt.Errorf("Bhojpur ERP API: failed to unmarshal DeleteAddressResponseBulk from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&bulkResp.Status) {
		return bulkResp, sharedCommon.NewErpError(bulkResp.Status.ErrorCode.String(), bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus, bulkResp.Status.ErrorCode)
	}

	for _, bulkRespItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&bulkRespItem.Status.Status) {
			return bulkResp, sharedCommon.NewErpError(
				bulkRespItem.Status.ErrorCode.String(),
				fmt.Sprintf("%+v", bulkRespItem.Status),
				bulkResp.Status.ErrorCode,
			)
		}
	}

	return bulkResp, nil
}

func (cli *Client) SaveAddressesBulk(ctx context.Context, addrMap []map[string]interface{}, attrs map[string]string) (SaveAddressesResponseBulk, error) {
	var saveAddressesResponseBulk SaveAddressesResponseBulk

	if len(addrMap) > sharedCommon.MaxBulkRequestsCount {
		return saveAddressesResponseBulk, fmt.Errorf("cannot save more than %d addresses in one request", sharedCommon.MaxBulkRequestsCount)
	}

	bulkInputs := make([]common.BulkInput, 0, len(addrMap))
	for _, addr := range addrMap {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "saveAddress",
			Filters:    addr,
		})
	}

	resp, err := cli.SendRequestBulk(ctx, bulkInputs, attrs)
	if err != nil {
		return saveAddressesResponseBulk, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return saveAddressesResponseBulk, err
	}

	if err := json.Unmarshal(body, &saveAddressesResponseBulk); err != nil {
		return saveAddressesResponseBulk, fmt.Errorf("Bhojpur ERP API: failed to unmarshal SaveAddressesResponseBulk from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&saveAddressesResponseBulk.Status) {
		return saveAddressesResponseBulk, sharedCommon.NewErpError(
			saveAddressesResponseBulk.Status.ErrorCode.String(),
			saveAddressesResponseBulk.Status.Request+": "+saveAddressesResponseBulk.Status.ResponseStatus,
			saveAddressesResponseBulk.Status.ErrorCode,
		)
	}

	for _, addrBulkItem := range saveAddressesResponseBulk.BulkItems {
		if !common.IsJSONResponseOK(&addrBulkItem.Status.Status) {
			return saveAddressesResponseBulk, sharedCommon.NewErpError(
				addrBulkItem.Status.ErrorCode.String(),
				fmt.Sprintf("%+v", addrBulkItem.Status),
				saveAddressesResponseBulk.Status.ErrorCode,
			)
		}
	}

	return saveAddressesResponseBulk, nil
}
