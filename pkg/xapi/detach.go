package xapi

import (
	"errors"

	"github.com/arturoguerra/xcpng-csi/internal/structs"
)

func (c *xClient) detach(volID, nodeID string, zone *structs.Zone) (bool, error) {
	api, session, err := c.Connect(zone)
	if err != nil {
		return false, nil
	}
	defer c.Close(api, session)

	vm, err := c.GetVM(api, session, nodeID)
	if err != nil {
		return false, nil
	}

	log.Info("VDI.GetByUUID")
	vdiUUID, err := api.VDI.GetByUUID(session, volID)
	if err != nil {
		return false, err
	}

	if string(vdiUUID) == "" {
		return false, errors.New("Invalid Volume")
	}

	log.Info("VBD.GetAllRecords")
	vbds, err := api.VBD.GetAllRecords(session)
	if err != nil {
		return false, err
	}

	for ref, vbd := range vbds {
		if vbd.VM == vm && vbd.CurrentlyAttached && vbd.VDI == vdiUUID {
			if err := c.ForceDetachVBD(ref, api, session); err != nil {
				return false, err
			}
		}
	}

	return true, nil
}

// Detach detaches volume from node
func (c *xClient) Detach(volID, nodeID string) error {
	detached := false

	for _, zone := range c.GetZones() {
		if vDetached, err := c.detach(volID, nodeID, zone); err != nil || vDetached {
			if err != nil {
				log.Error(err)
				return err
			}

			detached = true
			break
		}
	}

	if detached == false {
		return errors.New("Unknown error detaching volume")
	}

	return nil
}
