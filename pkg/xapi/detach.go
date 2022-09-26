package xapi

import (
	"github.com/arturoguerra/go-xolib/pkg/xoclient"
)

// Detach detaches volume from node
func (c *xClient) Detach(volID, nodeID string) error {
	vdiRef := xoclient.VDIRef(volID)
	vm, err := c.GetVMFromK8sNode(nodeID)
	if err != nil {
		return err
	}

	log.Infof("Detach: Got %s vmref %s", nodeID, vm.UUID)

	// gets VBDs from volID
	vbds, err := c.GetVBDsFromVDI(vdiRef)
	if err != nil {
		return err
	}

	log.Infof("Detach: deleting VBDs (%d)", len(vbds))

	// deletes said VBD
	for _, vbd := range vbds {
		// FIXME: should we actually limit this to single VM?
		// why just not delete all VBDs?
		if vbd.VM == vm.UUID {
			// should we disconnect before removing VBD?
			// log.Infof("Disconnecting VBD uuid (%s)", vbd.UUID)
			// if err = c.DisconnectVBD(vbd.UUID); err != nil {
			// 	log.Error(err)
			// }

			log.Infof("Detach: removing VBD %s from VM: %s", vbd.UUID, vm.UUID)
			if err = c.DeleteVBD(vbd.UUID); err != nil {
				return err
			}
		}
	}

	return nil
}
