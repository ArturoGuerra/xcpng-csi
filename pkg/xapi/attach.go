package xapi

import (
	"errors"
	"fmt"

	"github.com/arturoguerra/go-xolib/pkg/xoclient"
)

// Attach attaches volume to node
func (c *xClient) Attach(volID, nodeID, fstype string) (string, error) {
	vdiRef := xoclient.VDIRef(volID)

	// retrieve requested VDI
	vdi, err := c.GetVDIByUUID(vdiRef)
	if err != nil {
		return "", err
	}

	// retrieve requested node
	vm, err := c.GetVMFromK8sNode(nodeID)
	if err != nil {
		return "", err
	}

	// retrieve VBDs for VDI
	vbds, err := c.GetVBDsFromVDI(vdi.UUID)
	if err != nil {
		return "", err
	}

	// we have more than one VBD for VDI so something is wrong
	if len(vbds) > 1 {
		return "", errors.New("Too much VBDs, something is wrong")
	}

	// if we have no VBDs for VDI, we should attach
	if len(vbds) == 0 {
		// lets attach one
		log.Infof("Attaching VDI: %s to VM: %s", vdi.UUID, vm.UUID)
		if err = c.AttachVBD(vdi.UUID, vm.UUID); err != nil {
			// if attach faild return error
			return "", err
		}

		// TODO: maybe give some time for VBD to become available
		// time.Sleep(5)

		// restart attachment procedure
		return "", errors.New("VDI attached successfully, waiting for VBD to become available")
	}

	var vbd *xoclient.VBD

	if vbds[0].VM == vm.UUID {
		// found attachment on proper node
		vbd = vbds[0]
	}

	if vbd == nil {
		return "", errors.New("Failed to attach volume (no VBD)")
	}

	if vbd.Device == "" {
		return "", errors.New("Failed to attach volume (no device)")
	}

	// return device to mount
	dev := fmt.Sprintf("/dev/%s", vbd.Device)
	log.Infof("Attached dev: %s, VDI: %s to VM: %s", dev, vdi.UUID, vm.UUID)
	return dev, nil
}
