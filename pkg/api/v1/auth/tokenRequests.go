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
	"fmt"

	sharedCommon "github.com/bhojpur/erp/pkg/api/v1/common"
	"github.com/bhojpur/erp/pkg/internal/common"
)

//VerifyIdentityToken ...
func (cli *Client) VerifyIdentityToken(ctx context.Context, jwt string) (*SessionInfo, error) {
	method := "verifyIdentityToken"
	params := map[string]string{
		//params.Add("request", method)
		//params.Add("clientCode", cli.clientCode)
		//params.Add("setContentType", "1")
		"jwt": jwt,
	}
	resp, err := cli.SendRequest(ctx, method, params)
	if err != nil {
		return nil, err
	}
	res := &verifyIdentityTokenResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError(fmt.Sprintf("unmarshaling %s response failed", method), err, 0)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	return &res.Result, nil
}

//GetIdentityToken ...
func (cli *Client) GetIdentityToken(ctx context.Context) (*IdentityToken, error) {
	method := "getIdentityToken"

	resp, err := cli.SendRequest(ctx, method, map[string]string{})
	if err != nil {
		return nil, sharedCommon.NewFromError(fmt.Sprintf("%s request failed", method), err, 0)
	}
	res := &getIdentityTokenResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError(fmt.Sprintf("unmarshaling %s response failed", method), err, 0)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	return &res.Result, nil
}

//GetJWTToken executes the getJWTToken query (https://learn-api.bhojpur.net/requests/getjwttoken).
func (cli *Client) GetJWTToken(ctx context.Context) (*JwtToken, error) {

	resp, err := cli.SendRequest(ctx, "getJwtToken", map[string]string{})
	if err != nil {
		return nil, err
	}
	var res JwtTokenResponse

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, sharedCommon.NewFromError("error decoding GetJWTToken response", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	return &res.Records, nil
}

//only for partnerClient
func (cli *PartnerClient) GetJWTToken(ctx context.Context) (*JwtToken, error) {

	resp, err := cli.SendRequest(ctx, "getJwtToken", map[string]string{})
	if err != nil {
		return nil, err
	}
	var res JwtTokenResponse

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, sharedCommon.NewFromError("error decoding GetJWTToken response", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	return &res.Records, nil
}
