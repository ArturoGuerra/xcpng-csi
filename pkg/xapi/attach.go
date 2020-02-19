package xapi

import (
    "fmt"
    "errors"
    "github.com/arturoguerra/xcpng-csi/pkg/errs"
    xenapi "github.com/terra-farm/go-xen-api-client"
)

func (c *xClient) Attach(volId, NodeID, rawMode, fstype string) (string, error) {
    xmode := xenapi.VbdModeRW

    api, session, err := c.Connect()
    if err != nil {
        return "", err
    }

    defer c.Close(api, session)

    vm, err := c.GetVM(api, session, NodeID)
    if err != nil {
        log.Error(err)
        return "", errs.New(errs.InvalidNode)
    }

    log.Info("VM.GetAllAllowedVBDDevices")
    vbdDevices, err := api.VM.GetAllowedVBDDevices(session, vm)
    if err != nil {
        return "", err
    }

    if len(vbdDevices) < 0 {
        return "", errors.New("No VBD Devices are available")
    }

    log.Info("VDI.GetByUUID")
    vdiUUID, err := api.VDI.GetByUUID(session, volId)
    if err != nil {
        return "", err
    }

    if string(vdiUUID) == "" {
        return "", errors.New(errs.InvalidVolume)
    }

    log.Info("VBD.GetAllRecords")
    vbds, err := api.VBD.GetAllRecords(session)
    if err != nil {
        return "", err
    }

    for ref, vbd := range vbds {
        if vbd.VDI == vdiUUID && vbd.CurrentlyAttached {
            log.Info("Attempting to safely detach VDI")
            if err := c.DetachVBD(ref, api, session); err != nil {
                return "", err
            }
        }
    }

    log.Info("VBD.Create")
    vbdUUID, err := api.VBD.Create(session, xenapi.VBDRecord{
        Bootable:    false,
        Mode:        xmode,
        Type:        xenapi.VbdTypeDisk,
        Unpluggable: true,
        Userdevice:  vbdDevices[0],
        VDI:         vdiUUID,
        VM:          vm,
    })
    if err != nil {
        return "", err
    }

    log.Info("VBD.Plug")
    if err = api.VBD.Plug(session, vbdUUID); err != nil {
        return "", err
    }

    log.Info("VBD.GetDevice")
    device, err := api.VBD.GetDevice(session, vbdUUID)
    if err != nil {
        return "", err
    }

    return fmt.Sprintf("/dev/%s", device), nil
}
