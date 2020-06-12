package service

import (
	"os"
	"strings"
	"sync"

	"github.com/arturoguerra/go-logging"
	"github.com/arturoguerra/xcpng-csi/internal/structs"
	"github.com/arturoguerra/xcpng-csi/pkg/xapi"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/mitchellh/mapstructure"
	"github.com/rexray/gocsi"
	"gopkg.in/go-playground/validator.v8"
)

const (
	// Name driver base
	Name = "csi.xcpng.ar2ro.io"
	// VendorVersion is DriverVersion
	VendorVersion = "1.1.0"
	// UnixSocketPrefix is the CSI Socket path
	UnixSocketPrefix = "unix://"

	// ZoneLabel is the kubernetes Zone node label
	ZoneLabel = "topology.kubernetes.io/zone"
	// ZoneUUID is the pool UUID
	ZoneUUID = Name + "/pool-uuid"
	// NodeUUID is the node UUID
	NodeUUID = Name + "/node-uuid"
	// NodeName is the vm name
	NodeName = Name + "/node"
)

// Manifest contains information about the CSI Driver
var Manifest = map[string]string{
	"url": "https://github.com/arturoguerra/kube-xcpng-csi",
}

var (
	log      = logging.New()
	gigabyte = 1024 * 1024 * 1024
	minSize  = 1 * gigabyte
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
	// Service interface that contains all the required CSI Methods
	Service interface {
		csi.ControllerServer
		csi.IdentityServer
		csi.NodeServer
	}

	service struct {
		XClient   xapi.XClient
		NodeID    string
		ClusterID string
		Zones     []*structs.Zone
		/* CreateVolume */
		CVMux sync.Mutex
		/* ControllerPublishVolume */
		PVMux sync.Mutex
	}

	// Params represent the StorageClass Parameters fields
	Params struct {
		FSType    string `json:"FSType"`
		Datastore string `json:"Datastore"`
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
		params.FSType = "ext4"
	}

	return &params, nil
}

// New Creates a new CSI Driver Client
func New(xclient xapi.XClient, nodeID string) Service {
	return &service{
		XClient: xclient,
		NodeID:  nodeID,
	}
}
