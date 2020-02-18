package xapi

import (
    xenapi "github.com/terra-farm/go-xen-api-client"
)

func (c *xClient) Detach(volname, nodename string) error {
    api, session, err := c.Connect()
    if err != nil {
        return err
    }
    defer c.Close(api, session)

    vm, err := c.GetVM(api, session, nodename)
    if err != nil {
        return err
    }

    log.Info("VDI.GetAllRecords")
    vdis, err := api.VDI.GetAllRecords(session)
    if err != nil {
        return err
    }

    var vdiUUID xenapi.VDIRef
    for ref, vdi := range vdis {
        if vdi.NameLabel == volname && !vdi.IsASnapshot {
            vdiUUID = ref
        }
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
