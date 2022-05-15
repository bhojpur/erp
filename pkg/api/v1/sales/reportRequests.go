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

func (cli *Client) GetSalesReport(ctx context.Context, filters map[string]string) (*GetSalesReport, error) {
	var salesReportResp GetSalesReport
	resp, err := cli.SendRequest(ctx, "getSalesReport", filters)
	if err != nil {
		return nil, sharedCommon.NewFromError("getSalesReport: error sending request", err, 0)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, sharedCommon.NewFromError(fmt.Sprintf("getSalesReport: bad response status code: %d", resp.StatusCode), nil, 0)
	}

	if err := json.NewDecoder(resp.Body).Decode(&salesReportResp); err != nil {
		return nil, sharedCommon.NewFromError("getSalesReport: unmarshaling response failed", err, 0)
	}
	if !common.IsJSONResponseOK(&salesReportResp.Status) {
		return &salesReportResp, sharedCommon.NewErpError(
			salesReportResp.Status.ErrorCode.String(),
			salesReportResp.Status.Request+": "+salesReportResp.Status.ResponseStatus,
			salesReportResp.Status.ErrorCode,
		)
	}
	if len(salesReportResp.Records) < 1 {
		return &salesReportResp, sharedCommon.NewFromError("getSalesReport: no records in response", nil, salesReportResp.Status.ErrorCode)
	}

	return &salesReportResp, nil
}
