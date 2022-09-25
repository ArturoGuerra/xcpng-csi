package xapi

func (c *xClient) GetNodeInfo(nodeLabel string) *NodeInfo {
	vms, err := c.GetVMByName(nodeLabel)
	if err != nil {
		log.Error(err)
		return nil
	}

	if len(vms) == 0 {
		log.Info("No xen nodes found")
		return nil
	}

	if len(vms) > 1 {
		log.Infof("Multiple xen nodes found (%d)", len(vms))
		return nil
	}

	vm := vms[0]

	log.Infof("Getting zone by uuid %s", vm.PoolID)
	if zone := c.GetZoneByUUID(vm.PoolID); zone != nil {
		return &NodeInfo{
			NodeID:   nodeLabel,
			NodeUUID: string(vm.UUID),
			Zone:     zone.Name,
			ZoneUUID: zone.PoolID,
		}
	}

	log.Info("No Zone found")

	return nil
}
