package xapi

import (
	"errors"

	"github.com/arturoguerra/xcpng-csi/internal/structs"
	xenapi "github.com/terra-farm/go-xen-api-client"
)

func (c *xClient) ForceDetachVBD(vbd xenapi.VBDRef, api *xenapi.Client, sess xenapi.SessionRef) error {
	log.Info("VBD.Unplug")
	if err := api.VBD.Unplug(sess, vbd); err != nil {
		log.Info("VBD.UnplugForce")
		if err := api.VBD.UnplugForce(sess, vbd); err != nil {
			return err
		}
	}

	log.Info("VBD.Destory")
	return api.VBD.Destroy(sess, vbd)
}

func (c *xClient) DetachVBD(vbd xenapi.VBDRef, api *xenapi.Client, sess xenapi.SessionRef) error {
	log.Info("VBD.Unplug")
	if err := api.VBD.Unplug(sess, vbd); err != nil {
		return err
	}

	log.Info("VBD.Destory")
	return api.VBD.Destroy(sess, vbd)
}

func (c *xClient) GetVM(api *xenapi.Client, sess xenapi.SessionRef, name string) (xenapi.VMRef, error) {
	log.Info("VM.GetByNameLabel")
	vms, err := api.VM.GetByNameLabel(sess, name)
	if err != nil {
		return "", err
	}

	if len(vms) == 0 {
		return "", errors.New("No VM with this name found")
	}

	if len(vms) > 1 {
		return "", errors.New("More than one VM with this name found")
	}

	vm := vms[0]

	return vm, nil
}

// GetRegion returns region based on region label
func (c *xClient) GetRegion(label string) *structs.Region {
	for _, region := range c.Regions {
		if region.Name == label {
			return region
		}
	}

	return nil
}

// GetZone returns zone based on region and zone label
func (c *xClient) GetZone(region *structs.Region, zoneLabel string) *structs.Zone {
	for _, zone := range region.Zones {
		if zone.Name == zoneLabel {
			return zone
		}
	}

	return nil
}

// GetZoneFromLabel get zone based on region and zone labels
func (c *xClient) GetZoneFromLabel(regionLabel, zoneLabel string) *structs.Zone {
	if region := c.GetRegion(regionLabel); region != nil {
		if zone := c.GetZone(region, zoneLabel); zone != nil {
			return zone
		}
	}

	return nil

}

// ValidTopology verifies region and zone are valid
func (c *xClient) ValidTopology(regionLabel, zoneLabel string) bool {
	region := c.GetRegion(regionLabel)
	if region == nil {
		return false
	}

	zone := c.GetZone(region, zoneLabel)
	if zone == nil {
		return false
	}

	return true
}

// GetStorageRepo gets storageRepository from zone
func (c *xClient) GetStorageRepo(zone *structs.Zone, datastore string) string {
	if datastore == "" {
		datastore = zone.Default
	}

	for _, storage := range zone.Storage {
		if storage.Name == datastore {
			return storage.SR
		}
	}

	return ""
}

// GetRegions returns list of regions
func (c *xClient) GetRegions() []*structs.Region {
	return c.Regions
}

// GetZones returns list of zones
func (c *xClient) GetZones() []*structs.Zone {
	zones := make([]*structs.Zone, 0)

	for _, region := range c.Regions {
		for _, zone := range region.Zones {
			zones = append(zones, zone)
		}
	}

	return zones
}
