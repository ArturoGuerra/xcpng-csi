package xapi

func (c *xClient) GetNodeInfo(nodeLabel string) *NodeInfo {
	vm, err := c.GetVMByName(nodeLabel)
	if err != nil {
		log.Error(err)
		return nil
	}

	if zone := c.GetZoneByUUID(vm.PoolID); zone == nil {
		return &NodeInfo{
			NodeID:   nodeLabel,
			NodeUUID: string(vm.UUID),
			Zone:     zone.Name,
			ZoneUUID: zone.PoolID,
		}
	}

	return nil
}
