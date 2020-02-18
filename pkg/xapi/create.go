package xapi

import (
    "fmt"
    xenapi "github.com/terra-farm/go-xen-api-client"
)

func (c *xClient) CreateVolume(name, sr, fstype string, size int) error {
    api, session, err := c.Connect()
    if err != nil {
        return err
    }
    defer c.Close(api, session)

    srs, err := api.SR.GetByNameLabel(session, sr)
    if err != nil {
        return fmt.Errorf("Could not list SRs for name label: %s, error: %s", sr, err.Error())
    }

    if len(srs) > 1 {
        return fmt.Errorf("Too many SRs where found for thr name label: %s", sr)
    }

    if len(srs) < 1 {
        return fmt.Errorf("No SR was found for name label: %s", sr)
    }

    if _, err := api.VDI.Create(session, xenapi.VDIRecord{
        NameDescription: "XCP-ng CSI Driver for Kubernetes",
        NameLabel:       name,
        SR:              srs[0],
        Type:            xenapi.VdiTypeUser,
        VirtualSize:     size,
    }); err != nil {
        return err
    }

    return nil
}
