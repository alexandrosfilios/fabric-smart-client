/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package badger

import (
	"github.com/hyperledger-labs/fabric-smart-client/platform/view"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/db"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/db/driver"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/db/driver/unversioned"
)

type Driver struct{}

func (v *Driver) NewVersioned(sp view.ServiceProvider, dataSourceName string) (driver.VersionedPersistence, error) {
	return OpenDB(dataSourceName)
}

func (v *Driver) New(sp view.ServiceProvider, dataSourceName string) (driver.Persistence, error) {
	db, err := OpenDB(dataSourceName)
	if err != nil {
		return nil, err
	}
	return &unversioned.Unversioned{Versioned: db}, nil
}

func init() {
	db.Register("badger", &Driver{})
}
