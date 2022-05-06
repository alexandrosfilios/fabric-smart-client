/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package db

import (
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/db"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/db/driver"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/db/driver/unversioned"
)

type Driver struct {
}

func (o *Driver) NewVersioned(dataSourceName string) (driver.VersionedPersistence, error) {
	return OpenDB(dataSourceName)
}

func (o *Driver) New(dataSourceName string) (driver.Persistence, error) {
	db, err := OpenDB(dataSourceName)
	if err != nil {
		return nil, err
	}
	return &unversioned.Unversioned{Versioned: db}, nil
}

func init() {
	db.Register("orion", &Driver{})
}
