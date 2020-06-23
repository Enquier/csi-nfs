//go:generate go generate ./core

package main

import (
	"context"

	"github.com/rexray/gocsi"

	"github.com/enquier/csi-nfs/provider"
	"github.com/enquier/csi-nfs/service"
)

// main is ignored when this package is built as a go plug-in.
func main() {
	gocsi.Run(
		context.Background(),
		service.Name,
		"An NFS Container Storage Interface (CSI) Plugin",
		usage,
		provider.New())
}

const usage = ``
