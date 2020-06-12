package xapi

import (
	"errors"

	"github.com/arturoguerra/go-xolib/pkg/xoclient"
	"github.com/arturoguerra/xcpng-csi/internal/structs"
)

// CreateVolume creates in specific region/zone and storageRepo
func (c *xClient) CreateVolume(name, fsType, datastore string, size int64, zone *structs.Zone) (*xoclient.VDIRef, error) {
	// gets SR uuid from Name (using zone config)
	srRef := c.GetStorageRepo(zone, datastore)
	if srRef == nil {
		return nil, errors.New("Missing SR")
	}

	// gets SR using SRUUID
	sr, err := c.GetSRByUUID(*srRef)
	if err != nil {
		return nil, err
	}

	// creates volume in SR
	vdiRef, err := c.CreateVDI(name, size, sr.UUID)
	if err != nil {
		return nil, err
	}

	// returns VDIRef
	return vdiRef, nil
}
