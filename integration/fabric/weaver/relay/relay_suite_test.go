/*
Copyright IBM Corp All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package relay_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hyperledger-labs/fabric-smart-client/integration"
)

func TestEndToEnd(t *testing.T) {
	t.Skip("Re-enable this when we figure out how to work with weaver")
	RegisterFailHandler(Fail)
	RunSpecs(t, "Two Fabric Networks Suite with Weaver Relay")
}

func StartPort() int {
	return integration.TwoFabricNetworksWithWeaverRelayPort.StartPortForNode()
}
