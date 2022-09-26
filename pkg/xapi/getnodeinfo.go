package xapi

func (c *xClient) GetNodeInfo(nodeLabel string) *NodeInfo {
	vm, err := c.GetVMFromK8sNode(nodeLabel)
	if err != nil {
		log.Error(err)
		return nil
	}

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
