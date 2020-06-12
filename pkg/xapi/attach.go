package xapi

import (
	"fmt"

	"github.com/arturoguerra/go-xolib/pkg/xoclient"
)

// Attach attaches volume to node
func (c *xClient) Attach(volID, nodeID, fstype string) (string, error) {
	vdiRef := xoclient.VDIRef(volID)
	vm, err := c.GetVMByName(nodeID)
	if err != nil {
		return "", err
	}

	vdi, err := c.GetVDIByUUID(vdiRef)
	if err != nil {
		return "", err
	}

	vbds, err := c.GetVBDsFromVDI(vdi.UUID)
	if err != nil {
		return "", err
	}

	var vbd *xoclient.VBD

	for _, vbd := range vbds {
		if vbd.VM != vm.UUID {
			if err := c.DeleteVBD(vbd.UUID); err != nil {
				log.Error(err)
			}
		} else {
			vbd = vbd
		}
	}

	if vbd == nil {
		if err = c.AttachVBD(vdi.UUID, vm.UUID); err != nil {
			return "", err
		}

		vbds, err := c.GetVBDsFromVDI(vdi.UUID)
		if err != nil {
			return "", err
		}

		for _, vbd := range vbds {
			if vbd.VM == vm.UUID {
				vbd = vbd
			}
		}
	}

	return fmt.Sprintf("/dev/%s", vbd.Device), nil
}
