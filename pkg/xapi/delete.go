package xapi

import (
    "fmt"
)

func (c *xClient) DeleteVolume(id string) error {
    api, session, err := c.Connect()
    if err != nil {
        return err
    }

    defer c.Close(api, session)

    vdi, err := api.VDI.GetByUUID(session, id)
    if err != nil {
        return fmt.Errorf("Could not get VDI by UUID: %s, error: %s", id, err.Error())
    }


    vbds, err := api.VBD.GetAllRecords(session)
    if err != nil {
        return fmt.Errorf("Error getting all VBDs error: %s", err.Error())
    }

    for ref, vbd := range vbds {
        if vbd.VDI == vdi && vbd.CurrentlyAttached {
            if err = api.VBD.Unplug(session, ref); err != nil {
                return fmt.Errorf("Error unpluging VBD error: %s", err.Error())
            }

            if err = api.VBD.Destroy(session, ref); err != nil {
                return fmt.Errorf("Error destroying VBD error: %s", err.Error())
            }
        }
    }

    if err = api.VDI.Destroy(session, vdi); err != nil {
        return fmt.Errorf("Could not destory VDI by UUID: %s, error: %s", id, err.Error())
    }

    return nil
}
