package service

import "errors"

// GetTopologyFromLabels gets the Topology labels from map[string]string
func (s *service) GetTopologyFromLabels(labels map[string]string) (zone string, err error) {
	found := false

	zone, found = labels[ZoneLabel]
	if !found {
		return "", errors.New("Invalid Zone")
	}

	if valid := s.XClient.ValidTopology(zone); !valid {
		return "", errors.New("Invalid Topology")
	}

	return zone, err
}
