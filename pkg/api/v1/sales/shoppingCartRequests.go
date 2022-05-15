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
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	sharedCommon "github.com/bhojpur/erp/pkg/api/v1/common"
	"github.com/bhojpur/erp/pkg/internal/common"
)

func (cli *Client) CalculateShoppingCart(ctx context.Context, filters map[string]string) (*ShoppingCartTotals, error) {

	resp, err := cli.SendRequest(ctx, "calculateShoppingCart", filters)
	if err != nil {
		return nil, sharedCommon.NewFromError("CalculateShoppingCart: error sending request", err, 0)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, sharedCommon.NewFromError(fmt.Sprintf("CalculateShoppingCart: bad response status code: %d", resp.StatusCode), nil, 0)
	}

	var respData struct {
		Status  sharedCommon.Status
		Records []*ShoppingCartTotals
	}

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, sharedCommon.NewFromError("CalculateShoppingCart: unmarshaling response failed", err, 0)
	}
	if !common.IsJSONResponseOK(&respData.Status) {
		return nil, sharedCommon.NewErpError(respData.Status.ErrorCode.String(), respData.Status.Request+": "+respData.Status.ResponseStatus, respData.Status.ErrorCode)
	}
	if len(respData.Records) < 1 {
		return nil, sharedCommon.NewFromError("CalculateShoppingCart: no records in response", nil, respData.Status.ErrorCode)
	}

	return respData.Records[0], nil
}

func (cli *Client) CalculateShoppingCartWithFullRowsResponse(ctx context.Context, filters map[string]string) (*ShoppingCartTotalsWithFullRows, error) {
	resp, err := cli.SendRequest(ctx, "calculateShoppingCart", filters)
	if err != nil {
		return nil, sharedCommon.NewFromError("CalculateShoppingCart: error sending request", err, 0)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, sharedCommon.NewFromError(fmt.Sprintf("CalculateShoppingCart: bad response status code: %d", resp.StatusCode), nil, 0)
	}

	var respData struct {
		Status  sharedCommon.Status
		Records []*ShoppingCartTotalsWithFullRows
	}

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, sharedCommon.NewFromError("CalculateShoppingCart: unmarshaling response failed", err, 0)
	}
	if !common.IsJSONResponseOK(&respData.Status) {
		return nil, sharedCommon.NewErpError(respData.Status.ErrorCode.String(), respData.Status.Request+": "+respData.Status.ResponseStatus, respData.Status.ErrorCode)
	}
	if len(respData.Records) < 1 {
		return nil, sharedCommon.NewFromError("CalculateShoppingCart: no records in response", nil, respData.Status.ErrorCode)
	}

	return respData.Records[0], nil
}
