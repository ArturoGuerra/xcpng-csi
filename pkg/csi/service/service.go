package service

import (
    //"fmt"
    //"strings"
    //"sync"
    //"sync/atomic"
    //"//github.com/golang/protobuf/ptypes"
    "github.com/container-storage-interface/spec/lib/go/csi"
    "github.com/arturoguerra/xcpng-csi/pkg/xapi"
)

const (
    Name = "csi.xcpng.arturonet.com"
    VendorVersion = "0.69"
)

var Manifest = map[string]string{
    "url": "https://github.com/arturoguerra/kube-xcpng-csi",
}

type Service interface {
    csi.ControllerServer
    csi.IdentityServer
    csi.NodeServer
}

type service struct {
    XClient  xapi.XClient
    NodeID   string
}

func New(xclient xapi.Client, nodeid string) Service {
    return &service{
        XClient: xclient,
        NodeID:  nodeid,
    }
}
