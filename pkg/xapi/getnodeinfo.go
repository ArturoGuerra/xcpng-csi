package xapi

func (c *xClient) GetNodeInfo(nodeLabel string) *NodeInfo {
	vms, err := c.GetVMByName(nodeLabel)
	if err != nil {
		log.Error(err)
		return nil
	}

    if len(vms) > 1 || len(vms) == 0 {
		return nil
	}

	vm := vms[0]

	if zone := c.GetZoneByUUID(vm.PoolID); zone != nil {
		return &NodeInfo{
			NodeID:   nodeLabel,
			NodeUUID: string(vm.UUID),
			Zone:     zone.Name,
			ZoneUUID: zone.PoolID,
		}
	}

	return nil
}
