package xapi

import (
	"github.com/arturoguerra/go-logging"
	"github.com/arturoguerra/go-xolib/pkg/xoclient"
	"github.com/arturoguerra/go-xolib/pkg/xolib"
	"github.com/arturoguerra/xcpng-csi/internal/structs"
)

var log = logging.New()

type (
	// Config is the xolib configuration
	Config xolib.Config

	// XClient interface
	XClient interface {
		Attach(string, string, string) (string, error)
		Detach(string, string) error
		IsAttached(string, string) (bool, error)
		CreateVolume(string, string, string, int64, *structs.Zone) (*xoclient.VDIRef, error)
		DeleteVolume(string) error
		ValidTopology(string) bool
		GetZoneByLabel(string) *structs.Zone
		GetZones() []*structs.Zone
		GetNodeInfo(string) *NodeInfo
	}

	xClient struct {
		xoclient.Client
		ClusterID string
		NodeID    string
		PoolID    string
		Zones     []*structs.Zone
	}

	// NodeInfo contains information indentifing a node
	NodeInfo struct {
		NodeID   string
		NodeUUID string
		Zone     string
		ZoneUUID string
	}
)

// New creates new XCP-ng client
func New(clusterID, nodeID string, zones []*structs.Zone) XClient {
	return &xClient{
		ClusterID: clusterID,
		NodeID:    nodeID,
		Zones:     zones,
	}
}
