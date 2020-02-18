package xapi

import (
    "fmt"
)

func (c *xClient) DeleteVolume(name string) error {
    api, session, err := c.Connect()
    if err != nil {
        return err
    }

    defer c.Close(api, session)

    vdis, err := api.VDI.GetByNameLabel(session, name)
    if err != nil {
        return fmt.Errorf("Could not list VDIs for name label %s, error: %s", name, err.Error())
    }

    if len(vdis) > 1 {
        return fmt.Errorf("Too many VDIs where found for name label: %s", name)
    }

    if len(vdis) > 0 {
        vdi := vdis[0]

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
            return fmt.Errorf("Could not destory VDI for name label: %s, error: %s", name, err.Error())
        }
    }

    return nil
}
