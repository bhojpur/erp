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
	"errors"
	"net/http"
	"sync"
	"testing"

	"github.com/bhojpur/erp/pkg/internal/common"
	"github.com/stretchr/testify/assert"
)

func TestGetSessionKeyUserSuccess(t *testing.T) {
	sessKeyUserToGive := SessionKeyUser{
		UserID:             "123",
		UserName:           "someName",
		EmployeeName:       "someEmplName",
		EmployeeID:         "someEmplID",
		GroupID:            "someGroupID",
		GroupName:          "someGroupName",
		IPAddress:          "1.1.1.1",
		SessionKey:         "someSess",
		SessionLength:      10,
		LoginUrl:           "https://login.yunica.net",
		YunicaPOSVersion:   "123",
		YunicaPOSAssetsURL: "https://asset.pos.yunica.net",
		BhojpurURL:         "https://erp.bhojpur.net",
		IdentityToken:      "identityToken",
		Token:              "token",
	}
	payload := SessionKeyUserResponse{
		Records: []SessionKeyUser{sessKeyUserToGive},
	}

	bodyMock := common.NewMockFromStruct(payload)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       bodyMock,
	}

	cl := &common.ClientMock{
		ErrToGive:      nil,
		ResponseToGive: resp,
		Lock:           sync.Mutex{},
	}

	sessKeyUserActual, err := GetSessionKeyUser("pramila", "welcome1234", cl)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sessKeyUserToGive, *sessKeyUserActual)
	assert.Len(t, cl.Requests, 1)
	if len(cl.Requests) != 1 {
		return
	}

	req := cl.Requests[0]
	assert.Equal(
		t,
		"https://erp.bhojpur.net/api/?clientCode=welcome1234&doNotGenerateIdentityToken=1&request=getSessionKeyUser&sessionKey=pramila",
		req.URL.String(),
	)
	assert.Equal(t, "application/json", req.Header.Get("Accept"))
	assert.True(t, bodyMock.WasClosed)
}

func TestGetSessionKeyUserInvalidBody(t *testing.T) {
	cl := &common.ClientMock{
		ErrToGive: nil,
		ResponseToGive: &http.Response{
			StatusCode: http.StatusOK,
			Body: &common.BodyMock{
				Body:       common.NewMockFromStr("lala"),
				WasClosed:  false,
				CloseError: nil,
			},
		},
		Lock: sync.Mutex{},
	}

	_, err := GetSessionKeyUser("sess124", "code124", cl)
	assert.Error(t, err)
	if err == nil {
		return
	}
	assert.Contains(t, err.Error(), "Bhojpur ERP API: failed to decode SessionKeyUserResponse")
}

func TestGetSessionKeyUserZeroRecords(t *testing.T) {
	payload := SessionKeyUserResponse{
		Records: []SessionKeyUser{},
	}

	bodyMock := common.NewMockFromStruct(payload)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       bodyMock,
	}

	cl := &common.ClientMock{
		ErrToGive:      nil,
		ResponseToGive: resp,
		Lock:           sync.Mutex{},
	}

	_, err := GetSessionKeyUser("sess125", "code125", cl)
	assert.Error(t, err)
	if err == nil {
		return
	}
	assert.Contains(t, err.Error(), "Bhojpur ERP API: getSessionKeyUser: no records in response, status: Error, code: 0")
}

func TestGetSessionKeyUserError(t *testing.T) {
	payload := SessionKeyUserResponse{
		Records: []SessionKeyUser{},
	}

	bodyMock := common.NewMockFromStruct(payload)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       bodyMock,
	}

	cl := &common.ClientMock{
		ErrToGive:      errors.New("some bad error"),
		ResponseToGive: resp,
		Lock:           sync.Mutex{},
	}

	_, err := GetSessionKeyUser("sess126", "code126", cl)
	assert.Error(t, err)
	if err == nil {
		return
	}
	assert.EqualError(t, err, "Bhojpur ERP API: failed to call getSessionKeyUser request: some bad error, status: Error, code: 0")
}

func TestGetSessionKeyUserWrongRespCode(t *testing.T) {
	sessKeyUserToGive := SessionKeyUser{UserID: "123"}
	payload := SessionKeyUserResponse{
		Records: []SessionKeyUser{sessKeyUserToGive},
	}

	bodyMock := common.NewMockFromStruct(payload)
	resp := &http.Response{
		StatusCode: http.StatusBadRequest,
		Body:       bodyMock,
	}

	cl := &common.ClientMock{
		ErrToGive:      nil,
		ResponseToGive: resp,
		Lock:           sync.Mutex{},
	}

	_, err := GetSessionKeyUser("sess127", "code127", cl)
	assert.Error(t, err)
	if err == nil {
		return
	}
	assert.Contains(t, err.Error(), "wrong response status code: 400")
}
