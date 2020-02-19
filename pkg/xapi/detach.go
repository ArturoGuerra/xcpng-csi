package xapi

import (
    "errors"
)

func (c *xClient) Detach(volId, nodeID string) error {
    api, session, err := c.Connect()
    if err != nil {
        return err
    }
    defer c.Close(api, session)

    vm, err := c.GetVM(api, session, nodeID)
    if err != nil {
        return err
    }

    log.Info("VDI.GetByUUID")
    vdiUUID, err := api.VDI.GetByUUID(session, volId)
    if err != nil {
        return err
    }

    if string(vdiUUID) == "" {
        return errors.New("Invalid Volume")
    }

    log.Info("VBD.GetAllRecords")
    vbds, err := api.VBD.GetAllRecords(session)
    if err != nil {
        return err
    }

    for ref, vbd := range vbds {
        if vbd.VM == vm && vbd.CurrentlyAttached && vbd.VDI == vdiUUID {
            if err := c.ForceDetachVBD(ref, api, session); err != nil {
                return err
            }
        }
    }

    return nil
}
