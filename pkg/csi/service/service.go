package service

import (
    "os"
    "sync"
    "strings"
    "github.com/rexray/gocsi"
    "github.com/arturoguerra/go-logging"
    "gopkg.in/go-playground/validator.v8"
    "github.com/container-storage-interface/spec/lib/go/csi"
    "github.com/arturoguerra/xcpng-csi/pkg/xapi"
    "github.com/mitchellh/mapstructure"
)

const (
    Name = "csi.xcpng.arturonet.com"
    VendorVersion = "1.0.0"
    UnixSocketPrefix = "unix://"
)

var Manifest = map[string]string{
    "url": "https://github.com/arturoguerra/kube-xcpng-csi",
}

var (
    log = logging.New()
    gigabyte = 1024 * 1024 * 1024
    minSize  = 5 * gigabyte
)

// Work around if node dies and old csi.sock is left behind.
// NOTE: This can cause issues if two instances of a node are scheduled in the same node but that would be an extreme edge case.
func init() {
    sockPath := os.Getenv(gocsi.EnvVarEndpoint)
    sockPath = strings.TrimPrefix(sockPath, UnixSocketPrefix)
    if len(sockPath) > 1 {
        os.Remove(sockPath)
    }
}

type (
    Service interface {
        csi.ControllerServer
        csi.IdentityServer
        csi.NodeServer
    }

    service struct {
        XClient  xapi.XClient
        NodeID   string
        /* CreateVolume */
        CVMux    sync.Mutex
        /* ControllerPublishVolume */
        PVMux    sync.Mutex
    }

    Params struct {
        SR     string `json:"SR" validate:"required"`
        FSType string `json:"FSType"`
    }
)

func (s *service) ParseParams(items map[string]string) (*Params, error) {
    var params Params
    mapstructure.Decode(items, &params)
    v := validator.New(&validator.Config{TagName: "validate"})
    if err := v.Struct(&params); err != nil {
        return nil, err
    }

    if params.FSType == "" {
        log.Info("Missing fstype assuming EXT4")
        params.FSType = "ext4"
    }

    return &params, nil
}

func New(xclient xapi.XClient, nodeid string) Service {
    return &service{
        XClient: xclient,
        NodeID:  nodeid,
    }
}
