package service

import "errors"

// GetTopologyLabels gets the Topology labels from map[string]string
func (s *service) GetTopologyLabels(labels map[string]string) (region string, zone string, err error) {
	found := false

	region, found = labels[RegionLabel]
	if !found {
		return "", "", errors.New("Invalid Region")
	}

	zone, found = labels[ZoneLabel]
	if !found {
		return "", "", errors.New("Invalid Zone")
	}

	if valid := s.XClient.ValidTopology(region, zone); !valid {
		return "", "", errors.New("Invalid Topology")
	}

	return region, zone, err
}
