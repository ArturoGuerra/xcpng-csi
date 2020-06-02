package xapi

import (
	"fmt"

	"github.com/arturoguerra/xcpng-csi/internal/structs"
	xenapi "github.com/terra-farm/go-xen-api-client"
)

func (c *xClient) Connect(zone *structs.Zone) (*xenapi.Client, xenapi.SessionRef, error) {
	api, err := xenapi.NewClient(fmt.Sprintf("https://%s", zone.Credentials.Host), nil)
	if err != nil {
		return nil, "", err
	}

	sess, err := api.Session.LoginWithPassword(zone.Credentials.User, zone.Credentials.Password, "1.0", "Kubernetes-driver")
	if err != nil {
		return nil, "", err
	}

	return api, sess, nil
}

func (c *xClient) Close(api *xenapi.Client, sess xenapi.SessionRef) {
	api.Session.Logout(sess)
}
