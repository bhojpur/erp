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
	"net/http"
	"net/url"
)

type AuthFunc func(string) url.Values

//NewClientWithURL allows creating a new Client with a hardcoded URL. Useful for testing purposes
func NewClientWithURL(sk, cc, partnerKey, url string, httpCli *http.Client, headersForEveryRequestFunc AuthFunc) *Client {
	constr := &ClientConstructor{}
	constr.WithSessionKey(sk)
	constr.WithClientCode(cc)
	constr.WithURL(url)
	constr.WithHttpClient(httpCli)
	constr.WithPartnerKey(partnerKey)
	constr.WithHeaderFunc(headersForEveryRequestFunc)

	return constr.Build()
}

func NewClient(sk, cc, partnerKey string, httpCli *http.Client, headersForEveryRequestFunc AuthFunc) *Client {
	constr := &ClientConstructor{}
	constr.WithSessionKey(sk)
	constr.WithClientCode(cc)
	constr.WithHttpClient(httpCli)
	constr.WithPartnerKey(partnerKey)
	constr.WithHeaderFunc(headersForEveryRequestFunc)

	return constr.Build()
}

type ClientConstructor struct {
	sk                         string
	url                        string
	partnerKey                 string
	clientCode                 string
	httpCli                    *http.Client
	headersForEveryRequestFunc AuthFunc
	sessionProvider            SessionProvider
}

func (cc *ClientConstructor) Build() *Client {
	var sessionProvider SessionProvider
	if cc.sessionProvider != nil {
		sessionProvider = cc.sessionProvider
	} else {
		sessionProvider = &DefaultSessionProvider{SessionKey: cc.sk}
	}

	if cc.httpCli == nil {
		cc.httpCli = GetDefaultHTTPClient()
	}

	cli := &Client{
		httpClient:      cc.httpCli,
		sessionProvider: sessionProvider,
		clientCode:      cc.clientCode,
		partnerKey:      cc.partnerKey,
		headersFunc:     cc.headersForEveryRequestFunc,
	}

	if cli.headersFunc == nil {
		cli.headersFunc = cli.getDefaultMandatoryHeaders
	}
	if cc.url == "" {
		if cc.clientCode != "" {
			cli.Url = GetBaseURL(cc.clientCode)
		} else {
			cli.Url = GetBaseURLFromAuthFunc(cli.headersFunc)
		}
	} else {
		cli.Url = cc.url
	}
	return cli
}

func (cc *ClientConstructor) WithSessionKey(sk string) {
	cc.sk = sk
}

func (cc *ClientConstructor) WithURL(url string) {
	cc.url = url
}

func (cc *ClientConstructor) WithPartnerKey(partnerKey string) {
	cc.partnerKey = partnerKey
}

func (cc *ClientConstructor) WithClientCode(clientCode string) {
	cc.clientCode = clientCode
}

func (cc *ClientConstructor) WithHttpClient(httpCli *http.Client) {
	cc.httpCli = httpCli
}

func (cc *ClientConstructor) WithHeaderFunc(headersForEveryRequestFunc AuthFunc) {
	cc.headersForEveryRequestFunc = headersForEveryRequestFunc
}

func (cc *ClientConstructor) WithSessionProvider(sessProv SessionProvider) {
	cc.sessionProvider = sessProv
}

type SessionProvider interface {
	GetSession() (sessionKey string, err error)
	Invalidate()
}

type DefaultSessionProvider struct {
	SessionKey string
}

func (dsp *DefaultSessionProvider) GetSession() (sessionKey string, err error) {
	return dsp.SessionKey, nil
}

func (dsp *DefaultSessionProvider) Invalidate() {
	dsp.SessionKey = ""
}

type Client struct {
	Url                         string
	httpClient                  *http.Client
	clientCode                  string
	partnerKey                  string
	headersFunc                 AuthFunc
	sessionProvider             SessionProvider
	sendParametersInRequestBody bool
}

//SendParametersInRequestBody indicates to the client that the request should add the data payload in the
//request body instead of using the query parameters. Using the request body eliminates the query size
//limitations imposed by the maximum URL length
func (cli *Client) SendParametersInRequestBody() {
	cli.sendParametersInRequestBody = true
}

func (cli *Client) Close() {
	cli.httpClient.CloseIdleConnections()
}
