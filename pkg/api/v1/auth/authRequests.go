package auth

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
	"net/url"

	sharedCommon "github.com/bhojpur/erp/pkg/api/v1/common"

	"github.com/bhojpur/erp/pkg/internal/common"
)

//VerifyUser will give you session key
func VerifyUser(username, password, clientCode string, client *http.Client) (string, error) {
	requestUrl := fmt.Sprintf(common.BaseUrl, clientCode)
	params := url.Values{}
	params.Add("username", username)
	params.Add("clientCode", clientCode)
	params.Add("password", password)
	params.Add("request", "verifyUser")

	req, err := http.NewRequest("POST", requestUrl, nil)
	if err != nil {
		return "", sharedCommon.NewFromError("failed to build HTTP request", err, 0)
	}

	req.URL.RawQuery = params.Encode()
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return "", sharedCommon.NewFromError("failed to build VerifyUser request", err, 0)
	}

	res := &VerifyUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", sharedCommon.NewFromError("failed to decode VerifyUserResponse", err, 0)
	}
	if len(res.Records) < 1 {
		return "", sharedCommon.NewFromError("VerifyUser: no records in response", nil, res.Status.ErrorCode)
	}
	return res.Records[0].SessionKey, nil
}

//pass filters (including clientCode and sessionKey), pass client code, context and http client
func VerifyUserV2(ctx context.Context, filters map[string]string, clientCode string, cli *http.Client) (string, error) {
	requestUrl := fmt.Sprintf(common.BaseUrl, clientCode)
	params := url.Values{}
	for k, v := range filters {
		params.Add(k, v)
	}
	params.Add("request", "verifyUser")
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestUrl, nil)
	if err != nil {
		return "", sharedCommon.NewFromError("failed to build HTTP request", err, 0)
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Add("Accept", "application/json")
	resp, err := cli.Do(req)

	if err != nil {
		return "", sharedCommon.NewFromError("failed to build VerifyUser request", err, 0)
	}
	res := &VerifyUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", sharedCommon.NewFromError("failed to decode VerifyUserResponse", err, 0)
	}

	if res.Status.ErrorCode != 0 {
		return "", sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return res.Records[0].SessionKey, nil
}

func VerifyUserV3(ctx context.Context, filters map[string]string, clientCode string, cli *http.Client) (*VerifyUserResponse, error) {
	requestUrl := fmt.Sprintf(common.BaseUrl, clientCode)
	params := url.Values{}
	for k, v := range filters {
		params.Add(k, v)
	}
	params.Add("request", "verifyUser")
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestUrl, nil)
	if err != nil {
		return nil, sharedCommon.NewFromError("failed to build HTTP request", err, 0)
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Add("Accept", "application/json")
	resp, err := cli.Do(req)

	if err != nil {
		return nil, sharedCommon.NewFromError("failed to build VerifyUser request", err, 0)
	}
	res := &VerifyUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("failed to decode VerifyUserResponse", err, 0)
	}

	if res.Status.ErrorCode != 0 {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return res, nil
}

//VerifyUserFull executes the Bhojpur ERP API VerifyUser call and returns an object containing most of the resulting data.
//If it is necessary to specify the length of the created session or pass some other additional parameters
//to the underlying Bhojpur ERP API call, this can be done using the inputParams map.
func VerifyUserFull(ctx context.Context, username, password, clientCode string, inputParams map[string]string, cli *http.Client) (*SessionKeyUser, error) {
	requestUrl := fmt.Sprintf(common.BaseUrl, clientCode)
	params := url.Values{}
	for k, v := range inputParams {
		params.Add(k, v)
	}
	params.Add("username", username)
	params.Add("clientCode", clientCode)
	params.Add("password", password)
	params.Add("request", "verifyUser")
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestUrl, nil)
	if err != nil {
		return nil, sharedCommon.NewFromError("failed to build HTTP request", err, 0)
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Add("Accept", "application/json")
	resp, err := cli.Do(req)

	if err != nil {
		return nil, sharedCommon.NewFromError("failed to build VerifyUser request", err, 0)
	}
	res := &VerifyUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("failed to decode VerifyUserResponse", err, 0)
	}

	if res.Status.ErrorCode != 0 {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	if len(res.Records) < 1 {
		return nil, errors.New("verifyUser: no records in response")
	}
	return &res.Records[0], nil
}

//SwitchUser executes the Bhojpur ERP API SwitchUser call and returns an object containing most of the resulting data.
//If it is necessary to specify the length of the created session or pass some other additional parameters
//to the underlying Bhojpur ERP API call, this can be done using the inputParams map.
func SwitchUser(ctx context.Context, sessionKey, pin, clientCode string, inputParams map[string]string, cli *http.Client) (*SessionKeyUser, error) {
	requestUrl := fmt.Sprintf(common.BaseUrl, clientCode)
	params := url.Values{}
	for k, v := range inputParams {
		params.Add(k, v)
	}
	params.Add("sessionKey", sessionKey)
	params.Add("cardCode", pin)
	params.Add("clientCode", clientCode)
	params.Add("request", "switchUser")
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestUrl, nil)
	if err != nil {
		return nil, sharedCommon.NewFromError("failed to build HTTP request", err, 0)
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Add("Accept", "application/json")
	resp, err := cli.Do(req)

	if err != nil {
		return nil, sharedCommon.NewFromError("failed to build SwitchUser request", err, 0)
	}
	res := &SwitchUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("failed to decode SwitchUserResponse", err, 0)
	}

	if res.Status.ErrorCode != 0 {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	if len(res.Records) < 1 {
		return nil, errors.New("switchUser: no records in response")
	}
	return &res.Records[0], nil
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

//GetSessionKeyUser returns user information for the used session key
func GetSessionKeyUser(sessionKey string, clientCode string, client HttpClient) (*SessionKeyUser, error) {
	requestUrl := fmt.Sprintf(common.BaseUrl, clientCode)
	params := url.Values{}
	params.Add("sessionKey", sessionKey)
	params.Add("doNotGenerateIdentityToken", "1")
	params.Add("request", "getSessionKeyUser")
	params.Add("clientCode", clientCode)

	req, err := http.NewRequest("POST", requestUrl, nil)
	if err != nil {
		return nil, sharedCommon.NewFromError("failed to build HTTP request", err, 0)
	}

	req.URL.RawQuery = params.Encode()
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, sharedCommon.NewFromError("failed to call getSessionKeyUser request", err, 0)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		var body []byte
		if resp.Body != nil {
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				body = []byte{}
			}
		}

		return nil, fmt.Errorf("wrong response status code: %d, body: %s", resp.StatusCode, string(body))
	}

	res := &SessionKeyUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("failed to decode SessionKeyUserResponse", err, 0)
	}
	if len(res.Records) < 1 {
		return nil, sharedCommon.NewFromError("getSessionKeyUser: no records in response", nil, 0)
	}
	return &res.Records[0], nil
}

//GetSessionKeyInfo returns session key expiration info
func GetSessionKeyInfo(sessionKey string, clientCode string, client HttpClient) (*SessionKeyInfo, error) {
	requestUrl := fmt.Sprintf(common.BaseUrl, clientCode)
	params := url.Values{}
	params.Add("sessionKey", sessionKey)
	params.Add("request", "getSessionKeyInfo")
	params.Add("clientCode", clientCode)

	req, err := http.NewRequest("POST", requestUrl, nil)
	if err != nil {
		return nil, sharedCommon.NewFromError("failed to build HTTP request", err, 0)
	}

	req.URL.RawQuery = params.Encode()
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, sharedCommon.NewFromError("failed to call getSessionKeyInfo request", err, 0)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		var body []byte
		if resp.Body != nil {
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				body = []byte{}
			}
		}

		return nil, fmt.Errorf("wrong response status code: %d, body: %s", resp.StatusCode, string(body))
	}

	res := &SessionKeyInfoResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("failed to decode SessionKeyInfoResponse", err, 0)
	}
	if len(res.Records) < 1 {
		return nil, sharedCommon.NewFromError("getSessionKeyUser: no records in response", nil, res.Status.ErrorCode)
	}
	return &res.Records[0], nil
}
