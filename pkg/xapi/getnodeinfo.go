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
		log.Infof("Getting region info for: %s", region.Name)
		for _, zone := range region.Zones {
			log.Infof("Getting Zone info for: %s in %s region", zone.Name, region.Name)
			node, err := c.getNodeInfo(nodeLabel, zone)
			if err != nil {
				log.Error(err)
			} else if node != nil {
				log.Infof("Found NodeInfo for %s in zone: %s region: %s", nodeLabel, zone.Name, region.Name)
				node.Region = region.Name
				node.Zone = zone.Name
				return node
			}
		}
	}

	return nil
}
