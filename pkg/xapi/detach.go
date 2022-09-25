package xapi

import "github.com/arturoguerra/go-xolib/pkg/xoclient"

// Detach detaches volume from node
func (c *xClient) Detach(volID, nodeID string) error {
	vdiRef := xoclient.VDIRef(volID)
	vmRef := xoclient.VMRef(nodeID)

	// gets VBD from volID and nodeID
	vbds, err := c.GetVBDsFromVDI(vdiRef)
	if err != nil {
		return err
	}

	log.Infof("Detach: deleting VBDs (%d)", len(vbds))

	// deletes said VBD
	for _, vbd := range vbds {
		if vbd.VM == vmRef {
			if err = c.DeleteVBD(vbd.UUID); err != nil {
				return err
			}
		}
	}

	return nil
}
