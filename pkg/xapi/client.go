package xapi

import (
	"github.com/arturoguerra/go-logging"
	"github.com/arturoguerra/xcpng-csi/internal/structs"
)

var log = logging.New()

type (
	// XClient interface
	XClient interface {
		Attach(string, string, string, string, *structs.Zone) (string, error)
		Detach(string, string) error
		IsAttached(string, string, *structs.Zone) (bool, error)
		CreateVolume(string, string, string, int, *structs.Zone) (string, error)
		DeleteVolume(string) error
		ValidTopology(string, string) bool
		GetZoneFromLabel(string, string) *structs.Zone
		GetRegions() []*structs.Region
		GetZones() []*structs.Zone
		GetNodeInfo(string) *NodeInfo
	}

	xClient struct {
		Regions []*structs.Region
	}

	// NodeInfo contains information indentifing a node
	NodeInfo struct {
		NodeID string
		Region string
		Zone   string
	}
)

// New creates new XCP-ng client
func New(regions []*structs.Region) XClient {
	return &xClient{
		Regions: regions,
	}
}
