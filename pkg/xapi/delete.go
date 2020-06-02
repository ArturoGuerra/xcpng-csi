package xapi

import (
	"errors"
	"fmt"

	"github.com/arturoguerra/xcpng-csi/internal/structs"
)

func (c *xClient) deleteVolume(volID string, zone *structs.Zone) (bool, error) {
	api, session, err := c.Connect(zone)
	if err != nil {
		log.Error(err)
		// Needs to return nil as error even if it's the correct node cuz we don't know that
		// IDK who to blame if the CSI Team for overlooking this or xcpng for not having a centralized api :/ looking at you xen orchestra lol
		return false, nil
	}

	defer c.Close(api, session)

	vdi, err := api.VDI.GetByUUID(session, volID)
	if err != nil {
		log.Error(err)
		return false, nil
		//return fmt.Errorf("Could not get VDI by UUID: %s error: %s", volID, err.Error())
	}

	vbds, err := api.VBD.GetAllRecords(session)
	if err != nil {
		return false, fmt.Errorf("Error getting all VBDs error: %s", err.Error())
	}

	for ref, vbd := range vbds {
		if vbd.VDI == vdi && vbd.CurrentlyAttached {
			if err = api.VBD.Unplug(session, ref); err != nil {
				return false, fmt.Errorf("Error unpluging VBD error: %s", err.Error())
			}
			if err = api.VBD.Destroy(session, ref); err != nil {
				return false, fmt.Errorf("Error destroying VBD error: %s", err.Error())
			}
		}
	}

	if err = api.VDI.Destroy(session, vdi); err != nil {
		return false, fmt.Errorf("Could not destory VDI by UUID: %s, error: %s", volID, err.Error())
	}

	return true, nil
}

func (c *xClient) DeleteVolume(volID string) error {
	deleted := false
	for _, zone := range c.GetZones() {
		if vDeleted, err := c.deleteVolume(volID, zone); err != nil || vDeleted {
			if err != nil {
				log.Error(err)
				return err
			}

			deleted = true
			break
		}
	}

	if deleted == false {
		return errors.New("Unknown error deleting volume")
	}

	return nil
}
