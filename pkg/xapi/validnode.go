package xapi

func (c *xClient) ValidNode(NodeID string) (bool, error) {
	api, session, err := c.Connect();
    if err != nil {
		return false, err
	}

	defer c.Close(api, session)

	if _, err := c.GetVM(api, session, NodeID); err != nil {
		return false, nil
	}

	return true, nil
}