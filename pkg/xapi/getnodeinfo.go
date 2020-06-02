package xapi

import "github.com/arturoguerra/xcpng-csi/internal/structs"

func (c *xClient) getNodeInfo(nodeLabel string, zone *structs.Zone) (*NodeInfo, error) {
	api, session, err := c.Connect(zone)
	if err != nil {
		return nil, err
	}

	defer c.Close(api, session)

	vm, err := c.GetVM(api, session, nodeLabel)
	if err != nil {
		return nil, err
	}

	return &NodeInfo{
		NodeID: string(vm),
	}, nil

}

func (c *xClient) GetNodeInfo(nodeLabel string) *NodeInfo {
	for _, region := range c.Regions {
		for _, zone := range region.Zones {
			if node, err := c.getNodeInfo(nodeLabel, zone); err == nil && node != nil {
				node.Region = region.Name
				node.Zone = zone.Name
				return node
			}
		}
	}

	return nil
}
