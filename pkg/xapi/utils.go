package xapi

import (
	"github.com/arturoguerra/go-xolib/pkg/xoclient"
	"github.com/arturoguerra/xcpng-csi/internal/structs"
)

// GetZoneByLabel get zone based on region and zone labels
func (c *xClient) GetZoneByLabel(zoneLabel string) *structs.Zone {
	for _, zone := range c.GetZones() {
		if zone.Name == zoneLabel {
			return zone
		}
	}

	return nil
}

func (c *xClient) GetZoneByUUID(uuid string) *structs.Zone {
	for _, zone := range c.GetZones() {
		if zone.PoolID == uuid {
			return zone
		}
	}

	return nil
}

// ValidTopology verifies region and zone are valid
func (c *xClient) ValidTopology(zoneLabel string) bool {
	if zone := c.GetZoneByLabel(zoneLabel); zone != nil {
		return true
	}

	return false
}

// GetStorageRepo gets storageRepository from zone
func (c *xClient) GetStorageRepo(zone *structs.Zone, datastore string) *xoclient.SRRef {
	if datastore == "" {
		datastore = zone.Default
	}

	for _, storage := range zone.Storage {
		if storage.Name == datastore {
			return storage.SR
		}
	}

	return nil
}

// GetZones returns list of zones
func (c *xClient) GetZones() []*structs.Zone {
	return c.Zones
}
