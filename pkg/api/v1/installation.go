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
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	sharedCommon "github.com/bhojpur/erp/pkg/api/v1/common"
)

type InstallationRequest struct {
	CompanyName string `json:"companyName"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	SendEmail   int    `json:"sendEmail"`
}
type InstallationResponse struct {
	ClientCode int    `json:"clientCode"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
}

func CreateInstallation(baseUrl, partnerKey string, filters map[string]string, httpCli *http.Client) (*InstallationResponse, error) {

	if httpCli == nil {
		return nil, errors.New("no http cli provided")
	}

	params := url.Values{}
	for k, v := range filters {
		params.Add(k, v)
	}
	params.Add("request", createInstallationMethod)
	params.Add("partnerKey", partnerKey)

	req, err := http.NewRequest("POST", baseUrl, nil)
	if err != nil {
		return nil, sharedCommon.NewFromError("failed to build HTTP request", err, 0)

	}
	req.URL.RawQuery = params.Encode()
	resp, err := httpCli.Do(req)
	if err != nil {
		return nil, sharedCommon.NewFromError("CreateInstallation: error sending POST request", err, 0)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, sharedCommon.NewFromError(fmt.Sprintf("CreateInstallation: bad response status code: %d", resp.StatusCode), nil, 0)
	}

	var respData struct {
		Status  sharedCommon.Status
		Records []InstallationResponse
	}

	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, sharedCommon.NewFromError("CreateInstallation: error decoding JSON response body", err, 0)
	}
	if respData.Status.ErrorCode != 0 {
		return nil, sharedCommon.NewFromError(fmt.Sprintf("CreateInstallation: API error %s", respData.Status.ErrorCode), nil, respData.Status.ErrorCode)
	}
	if len(respData.Records) < 1 {
		return nil, sharedCommon.NewFromError("CreateInstallation: no records in response", nil, respData.Status.ErrorCode)
	}

	return &respData.Records[0], nil
}
