package xapi

import (
    "fmt"
    xenapi "github.com/terra-farm/go-xen-api-client"
)

func (c *xClient) Connect() (*xenapi.Client, xenapi.SessionRef, error) {
    api, err := xenapi.NewClient(fmt.Sprintf("https://%s", c.Host), nil)
    if err != nil {
        return nil, "", err
    }

    sess, err := api.Session.LoginWithPassword(c.Username, c.Password, "1.0", "Kubernetes-driver")
    if err != nil {
        return nil, "", err
    }

    return api, sess, nil
}

func (c *xClient) Close(api * xenapi.Client, sess xenapi.SessionRef) {
    api.Session.Logout(sess)
}
