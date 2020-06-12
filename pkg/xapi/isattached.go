package xapi

import "github.com/arturoguerra/go-xolib/pkg/xoclient"

func (c *xClient) IsAttached(volID, nodeID string) (bool, error) {
	vdiRef := xoclient.VDIRef(volID)
	vm, err := c.GetVMByName(nodeID)
	if err != nil {
		return false, err
	}

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
