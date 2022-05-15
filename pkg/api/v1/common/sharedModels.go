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
	"encoding/json"

	"github.com/pkg/errors"
)

type (
	//Addresses from getAddresses
	Addresses []Address

	//Address from getAddresses
	Address struct {
		AddressID        int
		OwnerID          int
		TypeID           int
		TypeActivelyUsed int
		Added            int64
		Address2         string
		TypeName         string
		Address          string
		Street           string
		PostalCode       string
		City             string
		State            string
		Country          string
		LastModified
		Attributes
	}

	Attributes struct {
		Attributes []ObjAttribute `json:"attributes"`
	}

	ObjAttribute struct {
		AttributeName  string `json:"attributeName"`
		AttributeType  string `json:"attributeType"`
		AttributeValue string `json:"attributeValue"`
	}

	LongAttribute struct {
		AttributeName  string `json:"attributeName"`
		AttributeValue string `json:"attributeValue"`
	}

	LongAttributes struct {
		LongAttributes []LongAttribute `json:"longAttributes"`
	}

	LastModified struct {
		LastModified           int64  `json:"lastModified"`
		LastModifierEmployeeID int64  `json:"lastModifierEmployeeID"`
		LastModifierUsername   string `json:"lastModifierUsername"`
	}
)

func (u *Address) UnmarshalJSON(data []byte) error {

	raw := struct {
		AddressID        int         `json:"addressID"`
		OwnerID          int         `json:"ownerID"`
		TypeID           json.Number `json:"typeID"`
		TypeActivelyUsed int         `json:"typeActivelyUsed"`
		Added            int64       `json:"added"`
		Address2         string      `json:"address2"`
		TypeName         string      `json:"typeName"`
		Address          string      `json:"address"`
		Street           string      `json:"street"`
		PostalCode       string      `json:"postalCode"`
		City             string      `json:"city"`
		State            string      `json:"state"`
		Country          string      `json:"country"`
		LastModified
		Attributes
	}{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	u.AddressID = raw.AddressID
	u.OwnerID = raw.OwnerID
	if raw.TypeID.String() != "" {
		typeID, err := raw.TypeID.Int64()
		if err != nil {
			return errors.Wrapf(err, "unable to unmarshal address. typeId did not contain an int: %s", raw.TypeID.String())
		}
		u.TypeID = int(typeID)
	}

	u.TypeActivelyUsed = raw.TypeActivelyUsed
	u.Added = raw.Added
	u.Address2 = raw.Address2
	u.TypeName = raw.TypeName
	u.Address = raw.Address
	u.Street = raw.Street
	u.PostalCode = raw.PostalCode
	u.City = raw.City
	u.State = raw.State
	u.Country = raw.Country
	u.LastModified = raw.LastModified
	u.Attributes = raw.Attributes
	return nil
}
