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
	"strconv"
	"sync"
	"time"

	addresses "github.com/bhojpur/erp/pkg/api/v1/address"
	auth "github.com/bhojpur/erp/pkg/api/v1/auth"
	sharedCommon "github.com/bhojpur/erp/pkg/api/v1/common"
	company "github.com/bhojpur/erp/pkg/api/v1/company"
	customers "github.com/bhojpur/erp/pkg/api/v1/customer"
	log "github.com/bhojpur/erp/pkg/api/v1/log"
	pos "github.com/bhojpur/erp/pkg/api/v1/pos"
	prices "github.com/bhojpur/erp/pkg/api/v1/price"
	products "github.com/bhojpur/erp/pkg/api/v1/product"
	documents "github.com/bhojpur/erp/pkg/api/v1/purchase"
	sales "github.com/bhojpur/erp/pkg/api/v1/sales"
	servicediscovery "github.com/bhojpur/erp/pkg/api/v1/service"
	"github.com/bhojpur/erp/pkg/api/v1/warehouse"
	"github.com/bhojpur/erp/pkg/internal/common"
)

type Client struct {
	commonClient *common.Client
	//Address requests
	AddressProvider addresses.Manager
	//Token requests
	AuthProvider auth.Provider
	//Company and Conf parameter requests
	CompanyManager company.Manager
	//Customers and suppliers requests
	CustomerManager customers.Manager
	//POS related requests
	PosManager pos.Manager
	//ListingDataProvider related requests
	ProductManager products.Manager
	//SalesDocuments, Payments, Projects, ShoppingCart, VatRates
	SalesManager sales.Manager
	//Warehouse requests
	WarehouseManager warehouse.Manager
	//Prices requests
	PricesManager prices.Manager
	//Documents requests
	DocumentsManager documents.Manager
	//Service Discovery
	ServiceDiscoverer servicediscovery.ServiceDiscoverer
}

func (c *Client) InvalidateSession() {
	c.commonClient.InvalidateSession()
}

func (c *Client) GetSession() (sessionKey string, err error) {
	return c.commonClient.GetSession()
}

//SendParametersInRequestBody indicates to the client that the request should add the data payload in the
//request body instead of using the query parameters. Using the request body eliminates the query size
//limitations imposed by the maximum URL length
func (c *Client) SendParametersInRequestBody() {
	c.commonClient.SendParametersInRequestBody()
}

//NewUnvalidatedClient returns a new Client without validating any of the incoming parameters giving the
//developer more flexibility
func NewUnvalidatedClient(sk, cc, partnerKey string, httpCli *http.Client) *Client {
	comCli := common.NewClient(sk, cc, partnerKey, httpCli, nil)
	return newErpClient(comCli)
}

//NewClientFromCredentials makes a verifyUser Bhojpur ERP API call and initializes the client struct
func NewClientFromCredentials(username, password, clientCode string, customCli *http.Client) (*Client, error) {
	if customCli == nil {
		customCli = common.GetDefaultHTTPClient()
	}
	sessionKey, err := auth.VerifyUser(username, password, clientCode, customCli)
	if err != nil {
		return nil, err
	}

	return NewClient(sessionKey, clientCode, customCli)
}

// NewClient Takes three params:
// sessionKey string obtained from credentials or jwt
// clientCode Bhojpur ERP customer identification number
// and a custom http Client if needs to be overwritten. if nil will use default http client provided by the SDK
//The headersSetToEveryRequest function will be executed on every request and supplied with the request name. There is an example in the /examples of you to use it
func NewClient(sessionKey string, clientCode string, customCli *http.Client) (*Client, error) {
	if sessionKey == "" || clientCode == "" {
		return nil, errors.New("sessionKey and clientCode are required")
	}
	comCli := common.NewClient(sessionKey, clientCode, "", customCli, nil)
	return newErpClient(comCli), nil
}

//NewClientWithCustomHeaders enables defining the function that will set headers to every request by your own
func NewClientWithCustomHeaders(customHTTPCli *http.Client, headersSetToEveryRequest func(requestName string) url.Values) (*Client, error) {
	if headersSetToEveryRequest == nil {
		return nil, errors.New("the function that will set headers to every request is a required argument")
	}
	return newErpClient(common.NewClient("", "", "", customHTTPCli, headersSetToEveryRequest)), nil
}

//NewClientWithURL creates a new Client which can have a static URL which is not affected by clientCode
// nor the headersSetToEveryRequest function if set. If the url parameter is set to an empty string, the URL
// is still resolved normally. This allows creating clients which have a static url in your unit tests but function
// normally in the rest of your code
func NewClientWithURL(sessionKey, clientCode, partnerKey, url string, httpCli *http.Client, headersSetToEveryRequest func(requestName string) url.Values) (*Client, error) {
	if (sessionKey == "" || clientCode == "") && headersSetToEveryRequest == nil {
		return nil, errors.New("Either sessionKey and clientCode or a function for header generation is required")
	}
	comCli := common.NewClientWithURL(sessionKey, clientCode, partnerKey, url, httpCli, headersSetToEveryRequest)
	return newErpClient(comCli), nil
}

func newErpClient(c *common.Client) *Client {
	return &Client{
		commonClient:      c,
		AddressProvider:   addresses.NewClient(c),
		AuthProvider:      auth.NewClient(c),
		CompanyManager:    company.NewClient(c),
		CustomerManager:   customers.NewClient(c),
		PosManager:        pos.NewClient(c),
		ProductManager:    products.NewClient(c),
		SalesManager:      sales.NewClient(c),
		WarehouseManager:  warehouse.NewClient(c),
		ServiceDiscoverer: servicediscovery.NewClient(c),
		PricesManager:     prices.NewClient(c),
		DocumentsManager:  documents.NewClient(c),
	}
}

type ClientBuilder struct {
	UserName                   string                 //if set this will be used to fetch session key every time when session gets outdated
	Password                   string                 //if set this will be used to fetch session key every time when session gets outdated
	ClientCode                 string                 //required value for all requests
	SessionKey                 string                 //if you don't set SessionProvider this key will be used to auth all requests
	DefaultSessionLenSeconds   int                    //set the length of dynamically created sessions
	URL                        string                 //change the base API url
	PartnerKey                 string                 //set the partner key
	HttpCli                    *http.Client           //you can adjust the http client transport options here
	HeadersForEveryRequestFunc common.AuthFunc        //this will set headers for all outgoing requests except for the session key
	SessionProvider            common.SessionProvider //custom session establishing logic, if not set DynamicSessionProvider is used which requires UserName and Password
}

type DynamicSessionProvider struct {
	ClientCode               string
	UserName                 string
	Pass                     string
	SessionKey               string
	SessionValidTill         *time.Time
	DefaultSessionLenSeconds int
	Lock                     sync.Mutex
	HTTPClient               *http.Client
}

func (dsp *DynamicSessionProvider) Invalidate() {
	dsp.Lock.Lock()
	defer dsp.Lock.Unlock()
	dsp.SessionKey = ""
}

func (dsp *DynamicSessionProvider) GetSession() (sessionKey string, err error) {
	dsp.Lock.Lock()
	defer dsp.Lock.Unlock()

	if dsp.isSessionValid() {
		log.Log.Log(log.Debug, "will use the cached key which is valid till %v", dsp.SessionValidTill)
		return dsp.SessionKey, nil
	}

	log.Log.Log(log.Debug, "will request new session key since the old one is not valid %v", dsp.SessionValidTill)
	sessionKey, validTill, err := dsp.getAuthUserFromAPI()
	if err != nil {
		return "", err
	}

	log.Log.Log(log.Debug, "got new session key with validity till %v", validTill)

	dsp.SessionKey = sessionKey
	dsp.SessionValidTill = validTill

	return dsp.SessionKey, nil
}

func (dsp *DynamicSessionProvider) isSessionValid() bool {
	if dsp.SessionKey == "" {
		return false
	}
	if dsp.SessionValidTill == nil {
		return true
	}

	return dsp.SessionValidTill.After(time.Now().UTC())
}

func (dsp *DynamicSessionProvider) getAuthUserFromAPI() (sessionKey string, validTill *time.Time, err error) {
	requestUrl := fmt.Sprintf(common.BaseUrl, dsp.ClientCode)
	params := url.Values{}
	params.Add("username", dsp.UserName)
	params.Add("clientCode", dsp.ClientCode)
	params.Add("password", dsp.Pass)
	params.Add("sessionLength", strconv.Itoa(dsp.DefaultSessionLenSeconds))
	params.Add("request", "verifyUser")

	log.Log.Log(log.Debug,
		"will call verifyUser with client code %s, user name %s and session length %d seconds",
		dsp.ClientCode,
		dsp.UserName,
		dsp.DefaultSessionLenSeconds,
	)

	req, err := http.NewRequest("POST", requestUrl, nil)
	if err != nil {
		return "", nil, err
	}

	client := dsp.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}

	req.URL.RawQuery = params.Encode()
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return "", nil, err
	}

	res := &auth.VerifyUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", nil, fmt.Errorf("failed to decode VerifyUserResponse %w", err)
	}

	if len(res.Records) < 1 {
		return "", nil, &sharedCommon.ErpError{
			Status:  res.Status.ResponseStatus,
			Message: "No records in the VerifyUserResponse",
			Code:    res.Status.ErrorCode,
		}
	}

	sessionKey = res.Records[0].SessionKey
	sessionLength := res.Records[0].SessionLength

	sessionValidTill := time.Now().UTC().Add(time.Second * time.Duration(sessionLength))
	validTill = &sessionValidTill
	return
}

func (cb ClientBuilder) Build() *Client {
	constr := &common.ClientConstructor{}
	constr.WithClientCode(cb.ClientCode)

	if cb.SessionProvider != nil {
		constr.WithSessionProvider(cb.SessionProvider)
	} else {
		sessProvider := &DynamicSessionProvider{
			ClientCode:               cb.ClientCode,
			UserName:                 cb.UserName,
			Pass:                     cb.Password,
			DefaultSessionLenSeconds: cb.DefaultSessionLenSeconds,
			Lock:                     sync.Mutex{},
		}

		constr.WithSessionProvider(sessProvider)
	}

	constr.WithPartnerKey(cb.PartnerKey)
	constr.WithURL(cb.URL)
	constr.WithHeaderFunc(cb.HeadersForEveryRequestFunc)
	constr.WithHttpClient(cb.HttpCli)
	constr.WithSessionKey(cb.SessionKey)

	baseClient := constr.Build()

	return newErpClient(baseClient)
}
