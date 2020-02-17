package provider

import (
    "context"
    "net"

    "github.com/rexray/gocsi"
    "github.com/arturoguerra/gologging"
    "github.com/arturoguerra/xcpng-csi/pkg/xapi"
    "github.com/arturoguerra/xcpng-csi/pkg/csi/service"
)

const (
    Version = "0.69"
)

var Manifest = map[string]string{
    "url": "https://github.com/arturoguerra/xcpng-csi"
}

func New(xclient xapi.XClient, nodeid) gocsi.StoragePluginProvider {
    svc := service.New(xclient, nodeid)
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
