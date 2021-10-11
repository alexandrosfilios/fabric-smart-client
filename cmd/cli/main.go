/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"os"

	"github.com/hyperledger-labs/fabric-smart-client/integration/nwo/cmd/artifactgen"
	"github.com/hyperledger-labs/fabric-smart-client/integration/nwo/cmd/cryptogen"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/cli"
	view "github.com/hyperledger-labs/fabric-smart-client/platform/view/services/client/view/cmd"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [view | cryptogen | artifactsgen]\n", os.Args[0])
	os.Exit(1)
}

func reArrangeArgs() {
	var args []string
	args = append(args, os.Args[0])
	args = append(args, os.Args[2:]...)
	os.Args = args
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	switch os.Args[1] {
	case "cryptogen":
		reArrangeArgs()
		cryptogen.Gen()
		return
	case "artifactsgen":
		reArrangeArgs()
		artifactgen.Gen()
		return
	case "view":
		cli := cli.NewCLI("sc", "Command line client for Fabric Smart Client")
		view.RegisterViewCommand(cli)
		cli.Run(os.Args[1:])
		return
	default:
		usage()
	}

}
