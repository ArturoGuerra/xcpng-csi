package xapi

import (
    "errors"
    xenapi "github.com/terra-farm/go-xen-api-client"
)

func (c *xClient) ForceDetachVBD(vbd xenapi.VBDRef, api *xenapi.Client, sess xenapi.SessionRef) error {
    log.Info("VBD.Unplug")
    if err := api.VBD.Unplug(sess, vbd); err != nil {
        log.Info("VBD.UnplugForce")
        if err := api.VBD.UnplugForce(sess, vbd); err != nil {
            return err
        }
    }

    log.Info("VBD.Destory")
    return api.VBD.Destroy(sess, vbd)
}

func (c *xClient) DetachVBD(vbd xenapi.VBDRef, api *xenapi.Client, sess xenapi.SessionRef) error {
    log.Info("VBD.Unplug")
    if err := api.VBD.Unplug(sess, vbd); err != nil {
       return err
    }

    log.Info("VBD.Destory")
    return api.VBD.Destroy(sess, vbd)
}

func (c *xClient) GetVM(api *xenapi.Client, sess xenapi.SessionRef, name string) (xenapi.VMRef, error) {
    log.Info("VM.GetByNameLabel")
    vms, err := api.VM.GetByNameLabel(sess, name)
    if err != nil {
        return "", err
    }

    if len(vms) == 0 {
        return "", errors.New("No VM with this name found")
    }

    if len(vms) > 1 {
        return "", errors.New("More than one VM with this name found")
    }

    vm := vms[0]

    return vm, nil
}
