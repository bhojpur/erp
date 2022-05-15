package common

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
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/bhojpur/erp/pkg/api/v1/common"
	"github.com/bhojpur/erp/pkg/api/v1/log"
)

type BulkInput struct {
	MethodName string
	Filters    map[string]interface{}
}

func IsJSONResponseOK(responseStatus *common.Status) bool {
	return strings.EqualFold(responseStatus.ResponseStatus, "ok")
}

func getHTTPRequest(cli *Client, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest("POST", cli.Url, body)
	if err != nil {
		return nil, common.NewFromError("failed to build HTTP request", err, 0)

	}
	return req, err
}

const (
	clientCode = "clientCode"
	sessionKey = "sessionKey"
)

func (cli *Client) getDefaultMandatoryHeaders(request string) url.Values {
	params := url.Values{}
	params.Add("request", request)
	params.Add("setContentType", "1")
	params.Add(clientCode, cli.clientCode)
	if cli.partnerKey != "" {
		params.Add("partnerKey", cli.partnerKey)
	}

	return params
}

func setParams(params url.Values, filters map[string]string) {
	for k, v := range filters {
		params.Add(k, v)
	}
}

func (cli *Client) SendRequest(ctx context.Context, apiMethod string, filters map[string]string) (*http.Response, error) {
	log.Log.Log(log.Debug, "will call %s with filters %+v", apiMethod, filters)
	params := cli.headersFunc(apiMethod)
	log.Log.Log(log.Debug, "extracted headers %+v", params)

	params, err := cli.addSessionParams(params)
	if err != nil {
		return nil, err
	}

	setParams(params, filters)

	var req *http.Request
	if cli.sendParametersInRequestBody {
		req, err = getHTTPRequest(cli, strings.NewReader(params.Encode()))
		if err != nil {
			return nil, common.NewFromError("failed to build HTTP request", err, 0)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, err = getHTTPRequest(cli, nil)
		if err != nil {
			return nil, common.NewFromError("failed to build HTTP request", err, 0)
		}
		req.URL.RawQuery = params.Encode()
	}
	req = req.WithContext(ctx)

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, common.NewFromError(fmt.Sprintf("%v request failed", apiMethod), err, 0)
	}
	log.Log.Log(log.Debug, "got response with code: %d", resp.StatusCode)
	return resp, nil
}

func (cli *Client) addSessionParams(params url.Values) (url.Values, error) {
	sk, err := cli.sessionProvider.GetSession()
	params.Add(sessionKey, sk)

	return params, err
}

func (cli *Client) InvalidateSession() {
	cli.sessionProvider.Invalidate()
}

func (cli *Client) GetSession() (sessionKey string, err error) {
	return cli.sessionProvider.GetSession()
}

type DestRespWithStatus interface {
	GetStatus() *common.Status
}

func (cli *Client) Scan(ctx context.Context, apiMethod string, filters map[string]string, dest DestRespWithStatus) error {
	resp, err := cli.SendRequest(ctx, apiMethod, filters)
	if err != nil {
		return common.NewFromError(apiMethod+" request failed", err, 0)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return common.NewFromError("unmarshalling of response has failed", err, 0)
	}

	if err := json.Unmarshal(body, dest); err != nil {
		return common.NewFromError("unmarshalling of response failed", err, 0)
	}

	status := dest.GetStatus()
	if !IsJSONResponseOK(dest.GetStatus()) {
		return common.NewErpErrorf(
			status.ErrorCode.String(),
			"request name: %s, error field: %s, response status: %s, body: %s",
			status.ErrorCode,
			status.Request,
			status.ErrorField,
			status.ResponseStatus,
			string(body),
		)
	}

	return nil
}

func (cli *Client) SendRequestBulk(ctx context.Context, inputs []BulkInput, filters map[string]string) (*http.Response, error) {
	log.Log.Log(log.Debug, "will call Bulk request with inputs %+v and filters %+v", inputs, filters)
	bulkRequest := make([]map[string]interface{}, 0, len(inputs))
	for _, input := range inputs {
		bulkItemFilters := input.Filters
		bulkItemFilters["requestName"] = input.MethodName

		bulkRequest = append(bulkRequest, bulkItemFilters)
	}

	jsonRequests, err := json.Marshal(bulkRequest)
	if err != nil {
		return nil, common.NewFromError("failed to build requests payload", err, 0)
	}

	filters["requests"] = string(jsonRequests)

	var params url.Values
	if cli.headersFunc != nil {
		params = cli.headersFunc("")
		params.Del("request")
		params, err = cli.addSessionParams(params)
		if err != nil {
			return nil, err
		}

	} else {
		params = make(url.Values)
	}

	setParams(params, filters)

	req, err := getHTTPRequest(cli, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, common.NewFromError("failed to build HTTP request", err, 0)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req = req.WithContext(ctx)

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, common.NewFromError("Bulk request failed", err, 0)
	}
	log.Log.Log(log.Debug, "got response from Bulk API with status %d", resp.StatusCode)
	return resp, nil
}

func doRequest(req *http.Request, cli *Client) (*http.Response, error) {
	resp, err := cli.httpClient.Do(req)
	return resp, err
}

func ConvertSourceToJsonStrIfPossible(source interface{}) string {
	data, err := json.Marshal(source)
	if err != nil {
		return fmt.Sprintf("%+v", source)
	}

	return string(data)
}

func Die(err error) {
	if err != nil {
		panic(err)
	}
}
