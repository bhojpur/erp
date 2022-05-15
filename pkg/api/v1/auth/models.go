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
	common2 "github.com/bhojpur/erp/pkg/api/v1/common"
)

type (
	VerifyUserResponse struct {
		Status  common2.Status   `json:"status"`
		Records []SessionKeyUser `json:"records"`
	}
	SwitchUserResponse struct {
		Status  common2.Status   `json:"status"`
		Records []SessionKeyUser `json:"records"`
	}

	verifyIdentityTokenResponse struct {
		Status common2.Status `json:"status"`
		Result SessionInfo    `json:"records"`
	}

	SessionInfo struct {
		SessionKey string `json:"sessionKey"`
	}

	getIdentityTokenResponse struct {
		Status common2.Status `json:"status"`
		Result IdentityToken  `json:"records"`
	}
	IdentityToken struct {
		Jwt string `json:"identityToken"`
	}
	JwtTokenResponse struct {
		Status  common2.Status `json:"status"`
		Records JwtToken       `json:"records"`
	}
	JwtToken struct {
		Token string `json:"token"`
	}

	SessionKeyUserResponse struct {
		Records []SessionKeyUser `json:"records"`
	}

	SessionKeyUser struct {
		UserID               string `json:"userID"`
		UserName             string `json:"userName"`
		EmployeeName         string `json:"employeeName"`
		EmployeeID           string `json:"employeeID"`
		GroupID              string `json:"groupID"`
		GroupName            string `json:"groupName"`
		IPAddress            string `json:"ipAddress"`
		SessionKey           string `json:"sessionKey"`
		SessionLength        int    `json:"sessionLength"`
		LoginUrl             string `json:"loginUrl"`
		YunicaPOSVersion     string `json:"yunicaPOSVersion"`
		YunicaPOSAssetsURL   string `json:"yunicaPOSAssetsURL"`
		BhojpurURL           string `json:"bhojpurURL"`
		IdentityToken        string `json:"identityToken"`
		Token                string `json:"token"`
		CustomerRegistryURLs []struct {
			Priority int64  `json:"priority"`
			Token    string `json:"token"`
			URL      string `json:"url"`
			Weight   int64  `json:"weight"`
		} `json:"customerRegistryURLs"`
	}

	SessionKeyInfoResponse struct {
		Status  common2.Status   `json:"status"`
		Records []SessionKeyInfo `json:"records"`
	}
	SessionKeyInfo struct {
		CreationUnixTime string `json:"creationUnixTime"`
		ExpireUnixTime   string `json:"expireUnixTime"`
	}
)
