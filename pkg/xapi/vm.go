package xapi

import (
	"errors"
	"github.com/arturoguerra/go-xolib/pkg/xoclient"
)

// return Xen VM reference for specified node
func (c *xClient) GetVMFromK8sNode(nodeID string) (*xoclient.VM, error) {
	vms, err := c.GetVMByName(nodeID)
	if err != nil {
		return nil, err
	}

	if len(vms) == 0 {
		log.Infof("No Xen nodes found matching: %s", nodeID)
		return nil, errors.New("Unable to find Xen VM")
	}

	if len(vms) > 1 {
		log.Infof("Multiple xen nodes found (%d)", len(vms))
		return nil, errors.New("Unable to find proper Xen VM")
	}

	return vms[0], nil
}
