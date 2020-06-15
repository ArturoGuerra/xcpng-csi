package xapi

import (
	"errors"

	"github.com/arturoguerra/go-xolib/pkg/xoclient"
)

func (c *xClient) IsAttached(volID, nodeID string) (bool, error) {
	vdiRef := xoclient.VDIRef(volID)
	vms, err := c.GetVMByName(nodeID)
	if err != nil {
		return false, err
	}

	if len(vms) > 1 || len(vms) == 0 {
		return false, errors.New("Issue fetching VM")
	}

	vm := vms[0]

	vbds, err := c.GetVBDsFromVM(vm.UUID)
	if err != nil {
		return false, err
	}

	for _, vbd := range vbds {
		if vbd.VDI == vdiRef {
			return true, nil
		}
	}

	return false, nil
}
