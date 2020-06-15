package xapi

import (
	"fmt"

	"github.com/arturoguerra/go-xolib/pkg/xoclient"
)

// Attach attaches volume to node
func (c *xClient) Attach(volID, nodeID, fstype string) (string, error) {
	vdiRef := xoclient.VDIRef(volID)
	vms, err := c.GetVMByName(nodeID)
	if err != nil {
		return "", err
	}

	if len(vms) > 1 || len(vms) == 0 {
		return "", errors.New("Error fetching VMs")
	}

	vm := vms[0]

	vdi, err := c.GetVDIByUUID(vdiRef)
	if err != nil {
		return "", err
	}

	vbds, err := c.GetVBDsFromVDI(vdi.UUID)
	if err != nil {
		return "", err
	}

	for _, VBD := range vbds {
		if err := c.DeleteVBD(VBD.UUID); err != nil {
			log.Error(err)
		}
	}

	if err = c.AttachVBD(vdi.UUID, vm.UUID); err != nil {
		return "", err
	}

	vbds, err = c.GetVBDsFromVDI(vdi.UUID)
	if err != nil {
		return "", err
	}

	if len(vbds) != 1 {
		return "", fmt.Errorf("Found %d VBDs", len(vbds))
	}

	var vbd *xoclient.VBD

	for _, VBD := range vbds {
		if VBD.VM == vm.UUID {
			vbd = VBD
		}
	}

	return fmt.Sprintf("/dev/%s", vbd.Device), nil
}
