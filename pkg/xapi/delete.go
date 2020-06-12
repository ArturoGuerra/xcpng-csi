package xapi

import "github.com/arturoguerra/go-xolib/pkg/xoclient"

func (c *xClient) DeleteVolume(volID string) error {
	vdiRef := xoclient.VDIRef(volID)
	vdi, err := c.GetVDIByUUID(vdiRef)
	if err != nil {
		return err
	}

	// deletes all VBDs in the vdi
	vbds, err := c.GetVBDsFromVDI(vdi.UUID)
	if err != nil {
		return err
	}

	for _, vbd := range vbds {
		if err := c.DeleteVBD(vbd.UUID); err != nil {
			log.Error(err)
		}
	}

	// deletes vd
	if err = c.DeleteVDI(vdi.UUID); err != nil {
		return err
	}

	return nil
}
