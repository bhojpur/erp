package api

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
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	sharedCommon "github.com/bhojpur/erp/pkg/api/v1/common"
	"github.com/bhojpur/erp/pkg/internal/common"
)

//this interface sums up the general requests here
type Manager interface {
	GetCountries(ctx context.Context, filters map[string]string) ([]Country, error)
	GetUserRights(ctx context.Context, filters map[string]string) ([]UserRights, error)
	GetEmployees(ctx context.Context, filters map[string]string) ([]Employee, error)
	GetEmployeesBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetEmployeesResponseBulk, error)
	GetBusinessAreas(ctx context.Context, filters map[string]string) ([]BusinessArea, error)
	GetCurrencies(ctx context.Context, filters map[string]string) ([]Currency, error)
	SaveEvent(ctx context.Context, filters map[string]string) (int, error)
	GetEvents(ctx context.Context, filters map[string]string) ([]Event, error)
	LogProcessingOfCustomerData(ctx context.Context, filters map[string]string) error
	GetUserOperationsLog(ctx context.Context, filters map[string]string) (*GetUserOperationsLogResponse, error)
}

// GetCountries will list countries according to specified filters.
func (c *Client) GetCountries(ctx context.Context, filters map[string]string) ([]Country, error) {
	resp, err := c.commonClient.SendRequest(ctx, GetCountriesMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetCountriesResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("failed to unmarshal GetCountriesResponse", err, 0)
	}
	if !common.IsJSONResponseOK((*sharedCommon.Status)(&res.Status)) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return res.Countries, nil
}

//GetUserName from GetUserRights Bhojpur ERP API request
func (c *Client) GetUserRights(ctx context.Context, filters map[string]string) ([]UserRights, error) {

	resp, err := c.commonClient.SendRequest(ctx, GetUserRightsMethod, filters)
	if err != nil {
		return nil, sharedCommon.NewFromError(GetUserRightsMethod+" request failed", err, 0)
	}
	res := &GetUserRightsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("unmarshaling GetUserRightsResponse failed", err, 0)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	if len(res.Records) == 0 {
		return nil, errors.New("no records found")
	}

	return res.Records, nil
}

// GetEmployees will list employees according to specified filters.
func (c *Client) GetEmployees(ctx context.Context, filters map[string]string) ([]Employee, error) {
	resp, err := c.commonClient.SendRequest(ctx, GetEmployeesMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetEmployeesResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("failed to unmarshal GetEmployeesResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return res.Employees, nil
}

func (c *Client) GetEmployeesBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetEmployeesResponseBulk, error) {
	var bulkResp GetEmployeesResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getEmployees",
			Filters:    bulkFilterMap,
		})
	}
	resp, err := c.commonClient.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return bulkResp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return bulkResp, err
	}

	if err := json.Unmarshal(body, &bulkResp); err != nil {
		return bulkResp, fmt.Errorf("Bhojpur ERP API: failed to unmarshal GetEmployeesResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&bulkResp.Status) {
		return bulkResp, sharedCommon.NewErpError(bulkResp.Status.ErrorCode.String(), bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus, bulkResp.Status.ErrorCode)
	}

	for _, bulkItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&bulkItem.Status.Status) {
			return bulkResp, sharedCommon.NewErpError(bulkItem.Status.ErrorCode.String(), bulkItem.Status.Request+": "+bulkItem.Status.ResponseStatus, bulkResp.Status.ErrorCode)
		}
	}

	return bulkResp, nil
}

// GetBusinessAreas will list business areas according to specified filters.
func (c *Client) GetBusinessAreas(ctx context.Context, filters map[string]string) ([]BusinessArea, error) {
	resp, err := c.commonClient.SendRequest(ctx, GetBusinessAreasMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetBusinessAreasResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("failed to unmarshal GetBusinessAreasResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return res.BusinessAreas, nil
}

// GetCurrencies will list currencies according to specified filters.
func (c *Client) GetCurrencies(ctx context.Context, filters map[string]string) ([]Currency, error) {
	resp, err := c.commonClient.SendRequest(ctx, GetCurrenciesMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetCurrenciesResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("failed to unmarshal GetCurrenciesResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return res.Currencies, nil
}

func (c *Client) LogProcessingOfCustomerData(ctx context.Context, filters map[string]string) error {
	resp, err := c.commonClient.SendRequest(ctx, logProcessingOfCustomerDataMethod, filters)
	if err != nil {
		return sharedCommon.NewFromError("logProcessingOfCustomerData request failed", err, 0)
	}

	if resp.StatusCode != http.StatusOK {
		return sharedCommon.NewFromError(fmt.Sprintf("Logging response HTTP status is %d", resp.StatusCode), nil, 0)
	}

	return nil
}

func (c *Client) GetUserOperationsLog(ctx context.Context, filters map[string]string) (*GetUserOperationsLogResponse, error) {
	resp, err := c.commonClient.SendRequest(ctx, GetUserOperationsLog, filters)
	if err != nil {
		return nil, err
	}
	var res GetUserOperationsLogResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("failed to unmarshal getUserOperationsLog", err, 0)
	}
	if !strings.EqualFold(res.Status.ResponseStatus, "ok") {
		return nil, sharedCommon.NewFromResponseStatus(&sharedCommon.Status{
			Request:        res.Status.Request,
			ResponseStatus: res.Status.ResponseStatus,
			ErrorCode:      res.Status.ErrorCode,
			ErrorField:     res.Status.ErrorField,
		})
	}
	return &res, nil
}

func (c *Client) SaveEvent(ctx context.Context, filters map[string]string) (int, error) {
	resp, err := c.commonClient.SendRequest(ctx, SaveEventMethod, filters)
	if err != nil {
		return 0, err
	}
	var res SaveEventResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return 0, sharedCommon.NewFromError(fmt.Sprintf("failed to unmarshal %s response", SaveEventMethod), err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return 0, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return res.Records[0].EventID, nil
}

func (c *Client) GetEvents(ctx context.Context, filters map[string]string) ([]Event, error) {
	resp, err := c.commonClient.SendRequest(ctx, GetEvents, filters)
	if err != nil {
		return nil, err
	}
	var res GetEventsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("failed to unmarshal GetEmployeesResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return res.Events, nil
}
