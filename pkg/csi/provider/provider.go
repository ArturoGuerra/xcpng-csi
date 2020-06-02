package provider

import (
	"github.com/arturoguerra/xcpng-csi/pkg/csi/service"
	"github.com/arturoguerra/xcpng-csi/pkg/xapi"
	"github.com/rexray/gocsi"
)

// New Creates the CSI Driver
func New(xclient xapi.XClient, nodeID string) gocsi.StoragePluginProvider {
	svc := service.New(xclient, nodeID)
	return &gocsi.StoragePlugin{
		Controller: svc,
		Identity:   svc,
		Node:       svc,
		EnvVars: []string{
			gocsi.EnvVarSerialVolAccess + "=true",
			gocsi.EnvVarSpecValidation + "=true",
			gocsi.EnvVarRequirePubContext + "=true",
		},
	}
}
